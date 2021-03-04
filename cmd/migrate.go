package cmd

import (
	"board-games/config"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate"
	"os"
)

// TODO: configから取得して環境を分けられるようにする
const (
	Source   = "file://cmd/migration/sql"
	Database = "mysql://%s%s@tcp(%s%s)/%s"
)

// declare command line options
var (
	Command = flag.String("exec", "", "set up or down as a argument")
	Force   = flag.Bool("f", false, "force exec fixed sql")
	Step    = flag.Int("step", 99999999, "step migrate versions")
)

// available command lists
var AvailableExecCommands = map[string]string{
	"up":      "Execute up sqls",
	"steps":   "Execute readDown sqls",
	"version": "Just check current migrate version",
}

func main() {
	config.LoadEnv() // TODO: 後で作る
	flag.Parse()
	if len(*Command) < 1 {
		fmt.Println("\nError: no argument")
		showUsageMessage() // TODO: 後で作る
		os.Exit(1)
		return
	}

	config, _ := config.GetConfig() // TODO: 後で作る
	user := config.Db.User
	pass := config.Db.Pass
	host := config.Db.Host
	port := config.Db.Port
	database := config.Db.Database

	m, err := migrate.New(Source, fmt.Sprintf(Database, user, pass, host, port, database))
	if err != nil {
		fmt.Println("err", err)
	}

	version, dirty, err := m.Version()
	showVersionInfo(version, dirty, err) // TODO: 後で作る

	fmt.Println("Command: exec", *Command)
	fmt.Println("Step: version", *Step)
	applyQuery(m, version, dirty) // TODO: 後で作る
}

// exec up or down sqls
func applyQuery(m *migrate.Migrate, version uint, dirty bool) {
	if dirty && *Force {
		fmt.Println("force=true: force execute current version sql")
		_ = m.Force(int(version))
	}

	var err error
	switch *Command {
	case "up":
		err = m.Up()
	case "steps":
		err = m.Steps(*Step)
	case "version":
		// do nothing
		return
	default:
		fmt.Println("\nError: invalid command '" + *Command + "'\n")
		showUsageMessage()
		os.Exit(1)
	}

	if err != nil && err.Error() != "no change" {
		fmt.Println("err", err)
		os.Exit(1)
	} else {
		fmt.Println("Success:", *Command+"\n")
		fmt.Println("updated version info")
		version, dirty, err := m.Version()
		showVersionInfo(version, dirty, err) // TODO: 後で作る
	}

}

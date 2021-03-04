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

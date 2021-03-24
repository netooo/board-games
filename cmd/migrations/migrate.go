package main

import (
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate"
	"os"
)

const (
	Source   = "file://cmd/migrations/sql"
	Database = "mysql://%s:%s@tcp(%s:%s)/%s"
)

var (
	Command = flag.String("exec", "", "set up or down as a argument")
	Force   = flag.Bool("f", false, "force exec fixed sql")
	Step    = flag.Int("step", 999999999, "step migrate version")
)

var AvailableCommands = map[string]string{
	"up":      "Execute up sqls",
	"step":    "Execute readDown sqls",
	"version": "Just check current migrate version",
}

func main() {
	flag.Parse()
	if len(*Command) < 1 {
		fmt.Println("\nError: no argument")
		showUsageMessage()
		os.Exit(1)
		return
	}

	DBName := os.Getenv("DB_NAME")
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")

	m, err := migrate.New(Source, fmt.Sprintf(Database, DBUser, DBPass, DBHost, DBPort, DBName))
	if err != nil {
		fmt.Println("Error: ", err)
	}

	version, dirty, err := m.Version()
	showVersionInfo(version, dirty, err)

	fmt.Println("Command: exec", *Command)
	fmt.Println("Step: version:", *Step)
	// applyQuery(m, version, dirty) TODO: 後で追加する
}

func showUsageMessage() {
	fmt.Println("--------------------------------")
	fmt.Println("Usage")
	fmt.Println("	go run migrate.go -exec <Command>")
	fmt.Println("Available exec commands: ")
	for availableCommand, detail := range AvailableCommands {
		fmt.Println("  " + availableCommand + " : " + detail)
	}
	fmt.Println("--------------------------------")
}

func showVersionInfo(version uint, dirty bool, err error) {
	fmt.Println("--------------------------------")
	fmt.Println("Version : ", version)
	fmt.Println("Dirty   : ", dirty)
	fmt.Println("Error   : ", err)
	fmt.Println("--------------------------------")
}

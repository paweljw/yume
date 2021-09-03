package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
	"yume/models"

	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "server",
				Aliases: []string{"s"},
				Usage:   "Serves Yume",
				Action:  server,
			},
			{
				Name:    "migrate",
				Aliases: []string{"m"},
				Usage:   "Runs data migrations",
				Action: func(c *cli.Context) error {
					migrations, err := filepath.Glob("./resources/*.sql")

					if err != nil {
						return err
					}

					sort.Strings(migrations)

					err = models.EstablishConnection()

					if err != nil {
						return err
					}

					latestMigration := 0
					latestMigrationStr, err := os.ReadFile("./resources/migration.lock")

					if err == nil {
						latestMigration, _ = strconv.Atoi(string(latestMigrationStr))
					}

					for _, migration := range migrations {
						migrationParts := strings.Split(migration, "/")
						migrationIdStr := strings.Split(migrationParts[len(migrationParts)-1], "_")[0]
						migrationId, _ := strconv.Atoi(migrationIdStr)

						if migrationId <= latestMigration {
							log.Printf("Skipping old migration %s", migration)
						} else {
							log.Printf("Applying migration %s", migration)
							body, err := os.ReadFile("./" + migration)
							if err != nil {
								log.Fatal(err)
							}

							err = models.Db.Exec(string(body)).Error

							if err != nil {
								log.Fatal(err)
							}

							err = os.WriteFile("./resources/migration.lock", []byte(strconv.Itoa(migrationId)), 0666)

							if err != nil {
								log.Fatal(err)
							}
						}
					}
					return nil
				},
			},
			{
				Name:    "gen_migration",
				Aliases: []string{"gm"},
				Usage:   "Generates a new data migration file",
				Action: func(c *cli.Context) error {
					date := time.Now().Unix()
					name := c.Args().Get(0)
					filename := fmt.Sprintf("%d_%s.sql", date, name)

					if filename == "" {
						log.Fatal("Empty name is not allowed.")
						return nil
					}

					_, err := os.Create("./resources/" + filename)
					log.Printf("Created ./resources/%s", filename)

					return err
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

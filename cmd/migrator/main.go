package main

import (
	"broadcaster/internal/modules"
	"broadcaster/internal/variables"
	"embed"
	"fmt"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/seletskiy/go-log"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

type MigrateLogger struct {
	Logger *zap.Logger
}

func (l MigrateLogger) Printf(format string, v ...interface{}) {
	l.Logger.Info(strings.ReplaceAll(fmt.Sprintf(format, v...), "\n", ""))
}
func (l MigrateLogger) Verbose() bool {
	return true
}

func main() {
	modules.Init()
	logger := modules.Logger()
	logger.Info(variables.Banner("Migrator"))

	defer func() {
		logger.Info("shutdown")
		_ = logger.Sync()
	}()

	log.Info("Open source and database")
	sourceDriver, err := iofs.New(migrationFiles, "migrations")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	dbDriver, err := postgres.WithInstance(modules.Database().DB, &postgres.Config{})
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	migrateInstance, err := migrate.NewWithInstance("iofs", sourceDriver, "_", dbDriver)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer func() {
		if err == nil {
			if _, err := migrateInstance.Close(); err != nil {
				logger.Error(err.Error())
			}
		}
	}()

	migrateInstance.Log = &MigrateLogger{Logger: logger}
	migrateInstance.PrefetchMigrations = 10
	migrateInstance.LockTimeout = 15 * time.Second

	app := &cli.App{
		Name:     variables.Banner("Migrator"),
		Version:  variables.Version,
		Usage:    "apply new migrations",
		HideHelp: false,
		Commands: []cli.Command{
			{
				Name: "version",
				Action: func(c *cli.Context) error {
					version, dirty, err := migrateInstance.Version()
					if err != nil {
						return err
					}
					if dirty {
						log.Info(fmt.Sprintf("%v (dirty)", version))
					} else {
						log.Info(fmt.Sprintf("%v", version))
					}
					return nil
				},
			},
			{
				Name: "generate",
				Action: func(c *cli.Context) error {
					return GenerateMigration(c.Args().First(), logger)
				},
			},
			{
				Name: "up",
				Action: func(c *cli.Context) error {
					stepsStr := c.Args().First()
					if len(stepsStr) == 0 {
						return migrateInstance.Up()
					} else if steps, err := strconv.Atoi(stepsStr); err != nil {
						return err
					} else {
						return migrateInstance.Steps(steps)
					}
				},
			}, {
				Name: "redo",
				Action: func(c *cli.Context) error {
					if err := migrateInstance.Steps(-1); err != nil {
						return err
					}
					return migrateInstance.Steps(1)
				},
			}, {
				Name: "down",
				Action: func(c *cli.Context) error {
					stepsStr := c.Args().First()
					if len(stepsStr) == 0 {
						return fmt.Errorf("required number of steps or all")
					} else if stepsStr == "all" {
						return migrateInstance.Down()
					} else if steps, err := strconv.Atoi(stepsStr); err != nil {
						return err
					} else {
						return migrateInstance.Steps(-steps)
					}
				},
			}, {
				Name: "force",
				Action: func(c *cli.Context) error {
					versionStr := c.Args().First()
					if len(versionStr) == 0 {
						return fmt.Errorf("required number of steps or all")
					} else if version, err := strconv.Atoi(versionStr); err != nil {
						return err
					} else {
						return migrateInstance.Force(version)
					}
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logger.Warn(err.Error())
	}
}

func createFile(filename string) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return err
	}
	return f.Close()
}

func GenerateMigration(name string, log *zap.Logger) error {
	dir := filepath.Dir("cmd/migrator/migrations/")
	version := strconv.FormatInt(time.Now().Unix(), 10)
	for _, direction := range []string{"up", "down"} {
		basename := fmt.Sprintf("%s_%s.%s.%s", version, name, direction, "sql")
		filename := filepath.Join(dir, basename)
		if err := createFile(filename); err != nil {
			return err
		}
		absPath, _ := filepath.Abs(filename)
		log.Info(absPath)
	}
	return nil
}

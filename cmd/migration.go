package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"path/filepath"
	"stories/cmd/config"
	"stories/pkg/stories/store"
	"strings"
)

const (
	rollBackStep = -1
	cutSet       = "file://"
	databaseName = "postgres"
)

func runMigrations() {
	newMigrate, err := newMigrate()
	if err != nil {
		fmt.Println(err)
	}

	if err := newMigrate.Up(); err != nil {
		if err == migrate.ErrNoChange {
			return
		}
		fmt.Println(err)
	}
}

func rollBackMigrations() {
	newMigrate, err := newMigrate()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := newMigrate.Steps(rollBackStep); err != nil {
		if err == migrate.ErrNoChange {
			return
		}
		fmt.Println(err)
	}
}

func newMigrate() (*migrate.Migrate, error) {
	cfg := config.NewConfig()

	dbHandler := store.NewDBHandler(cfg.GetDatabaseConfig(), zap.NewExample())

	db, err := dbHandler.GetDB()
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	sourcePath, err := getSourcePath(cfg.GetMigrationPath())
	if err != nil {
		return nil, err
	}

	return migrate.NewWithDatabaseInstance(sourcePath, databaseName, driver)
}

func getSourcePath(directory string) (string, error) {
	directory = strings.TrimLeft(directory, cutSet)
	absPath, err := filepath.Abs(directory)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", cutSet, absPath), nil
}

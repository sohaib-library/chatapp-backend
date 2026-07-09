package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/pressly/goose/v3"
)

const defaultMigrationsDir = "database/migration"

func RunMigrations(db *sql.DB) error {
	return runGoose(db, goose.Up)
}

func RollbackMigration(db *sql.DB) error {
	return runGoose(db, goose.Down)
}

func runGoose(db *sql.DB, fn func(*sql.DB, string, ...goose.OptionsFunc) error) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("set goose dialect: %w", err)
	}

	if _, err := os.Stat(defaultMigrationsDir); err != nil {
		return fmt.Errorf("migrations directory %q: %w", defaultMigrationsDir, err)
	}

	if err := fn(db, defaultMigrationsDir); err != nil {
		if err == goose.ErrNoNextVersion || err == goose.ErrNoCurrentVersion {
			return nil
		}
		return fmt.Errorf("run goose migrations: %w", err)
	}

	return nil
}

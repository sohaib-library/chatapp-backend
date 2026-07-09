package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

type Config struct {
	Host       string
	Port       string
	User       string
	Password   string
	DBName     string
	SSLMode    string
	ServerPort string
}

func Database(envFile string) *sql.DB {
	if err := godotenv.Load(envFile); err != nil {
		log.Fatal(err)
	}

	cfg, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	if err := ensureDatabase(ctx, cfg); err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("pgx", cfg.databaseURL())
	if err != nil {
		log.Fatal(err)
	}

	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	return db
}

func Migertions(db *sql.DB) {
	if err := RunMigrations(db); err != nil {
		log.Fatal(err)
	}
}

func loadConfig() (*Config, error) {
	cfg := &Config{
		Host:       firstNonEmpty(os.Getenv("POSTGRES_HOST"), os.Getenv("DB_HOST")),
		Port:       firstNonEmpty(os.Getenv("POSTGRES_PORT"), os.Getenv("DB_PORT"), "5432"),
		User:       firstNonEmpty(os.Getenv("POSTGRES_USER"), os.Getenv("DB_USER")),
		Password:   firstNonEmpty(os.Getenv("POSTGRES_PASSWORD"), os.Getenv("DB_PASSWORD")),
		DBName:     firstNonEmpty(os.Getenv("POSTGRES_DB"), os.Getenv("DB_NAME"), "chatapp"),
		SSLMode:    firstNonEmpty(os.Getenv("POSTGRES_SSLMODE"), os.Getenv("DB_SSLMODE"), "disable"),
		ServerPort: firstNonEmpty(os.Getenv("SERVER_PORT"), "8000"),
	}

	if os.Getenv("DATABASE_URL") != "" {
		return cfg, nil
	}

	if cfg.Host == "" || cfg.User == "" {
		return nil, fmt.Errorf("database is not configured: set DATABASE_URL or POSTGRES_*/DB_* environment variables")
	}

	return cfg, nil
}

func (c *Config) databaseURL() string {
	if databaseURL := os.Getenv("DATABASE_URL"); databaseURL != "" {
		return databaseURL
	}
	return buildDatabaseURL(c.User, c.Password, c.Host, c.Port, c.DBName, c.SSLMode)
}

func (c *Config) adminDatabaseURL() string {
	return buildDatabaseURL(c.User, c.Password, c.Host, c.Port, "postgres", c.SSLMode)
}

func ensureDatabase(ctx context.Context, cfg *Config) error {
	if canConnect(ctx, cfg.databaseURL()) {
		return nil
	}

	conn, err := pgx.Connect(ctx, cfg.adminDatabaseURL())
	if err != nil {
		return fmt.Errorf("connect to postgres: %w", err)
	}
	defer conn.Close(ctx)

	var exists bool
	if err := conn.QueryRow(ctx,
		"SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)",
		cfg.DBName,
	).Scan(&exists); err != nil {
		return fmt.Errorf("check database existence: %w", err)
	}

	if exists {
		return nil
	}

	createSQL := fmt.Sprintf(`CREATE DATABASE %s`, pgx.Identifier{cfg.DBName}.Sanitize())
	if _, err := conn.Exec(ctx, createSQL); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "42501" {
			return fmt.Errorf(
				`permission denied to create database %q: ask a PostgreSQL admin to run:
  CREATE DATABASE %s OWNER %s;`,
				cfg.DBName, cfg.DBName, conn.Config().User,
			)
		}
		return fmt.Errorf("create database %q: %w", cfg.DBName, err)
	}

	return nil
}

func canConnect(ctx context.Context, databaseURL string) bool {
	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		return false
	}
	defer conn.Close(ctx)

	return conn.Ping(ctx) == nil
}

func buildDatabaseURL(user, password, host, port, dbName, sslMode string) string {
	if password != "" {
		return fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=%s",
			user, password, host, port, dbName, sslMode,
		)
	}

	return fmt.Sprintf(
		"postgres://%s@%s:%s/%s?sslmode=%s",
		user, host, port, dbName, sslMode,
	)
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

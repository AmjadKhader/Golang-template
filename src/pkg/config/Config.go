package repos

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"path/filepath"

	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
)

// DB holds the database connection pool.
var DB *sql.DB

func Setup() {
	// Load configuration from config.toml
	viper.SetConfigFile("../../config.toml") //
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Database configuration
	dbUser := viper.GetString("database.user")
	dbPassword := viper.GetString("database.password")
	dbName := viper.GetString("database.dbname")
	migrations := viper.GetString("migrations.dir")

	connStr := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Run database migrations.
	runMigrations(DB, migrations)

	// Verify the connection.
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}

}

func runMigrations(db *sql.DB, migrations string) {
	absPath, err := filepath.Abs("../../" + migrations)

	if err != nil {
		log.Fatal(err)
	}

	goose.SetDialect("postgres")
	if err := goose.Up(db, absPath); err != nil {
		log.Fatalf("could not run migrations: %v", err)
	}
	fmt.Println("Migrations ran successfully")
}

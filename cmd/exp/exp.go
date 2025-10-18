package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) ToString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode,
	)
}

func main() {
	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "paul",
		Password: "dune",
		Database: "lensvault",
		SSLMode:  "disable",
	}

	db, err := sql.Open("pgx", cfg.ToString())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to DB")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT NOT NULL
	  );
	
	  CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		amount INT,
		description TEXT
	  );`)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tables created.")

	name := "',''); DROP TABLE users; --"
	email := "new@smith.io"
	// query := fmt.Sprintf(`
	//   INSERT INTO users (name, email)
	//   VALUES ('%s', '%s');`, name, email)
	// fmt.Printf("Executing: %s\n", query)
	// _, err = db.Exec(query)
	row := db.QueryRow(`
		INSERT INTO users (name, email)
		VALUES ($1, $2) RETURNING id;`, name, email)
	var id int
	err = row.Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("User created. id =", id)
}

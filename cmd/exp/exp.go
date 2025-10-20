package main

import (
	"fmt"

	"github.com/Vedjw/lensvault/models"
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
	config := models.DefaultPostgresConfig()
	db, err := models.Open(config)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to DB")

	us := models.UserService{
		DB: db,
	}
	nu := &models.NewUser{
		FirstName: "ved",
		LastName:  "W",
		Age:       23,
		Email:     "xyz@io.com",
		Password:  "asdf",
	}
	user, err := us.Create(nu)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	// _, err = db.Exec(`
	// CREATE TABLE IF NOT EXISTS users (
	// 	pk SERIAL PRIMARY KEY,
	// 	id UUID DEFAULT uuidv7() NOT NULL,
	// 	name TEXT NOT NULL,
	// 	email TEXT NOT NULL UNIQUE,
	// 	password_hash TEXT NOT NULL
	// );`)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Tables created.")

	// name := "',''); DROP TABLE users; --"
	// email := "new@smith.io"
	// // query := fmt.Sprintf(`
	// //   INSERT INTO users (name, email)
	// //   VALUES ('%s', '%s');`, name, email)
	// // fmt.Printf("Executing: %s\n", query)
	// // _, err = db.Exec(query)
	// row := db.QueryRow(`
	// 	INSERT INTO users (name, email, password_hash)
	// 	VALUES ($1, $2, $3) RETURNING pk, id;`, name, email, "yolo")
	// var pk int
	// var uuid string
	// err = row.Scan(&pk, &uuid)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("User created. at = %d with id = %s", pk, uuid)
}

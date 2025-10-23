package models

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	PK           uint
	ID           string
	FirstName    string
	LastName     string
	Age          int
	Email        string
	PasswordHash string
}

type UserService struct {
	DB *sql.DB
}

type NewUser struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
	Password  string
}

func (us UserService) Create(nu *NewUser) (*User, error) {
	nu.Email = strings.ToLower(nu.Email)
	//TODO: replace hardcoded pepper with os.GetEnv()
	password := nu.Password + "-" + "dvorak"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("models: create user: %w", err)
	}
	passwordHash := string(hash)

	user := &User{
		FirstName:    nu.FirstName,
		LastName:     nu.LastName,
		Age:          nu.Age,
		Email:        nu.Email,
		PasswordHash: passwordHash,
	}
	row := us.DB.QueryRow(`
		INSERT INTO users (first_name, last_name, age, email, password_hash)
		VALUES ($1, $2, $3, $4, $5) RETURNING pk, id;`,
		nu.FirstName, nu.LastName, nu.Age, nu.Email, passwordHash)
	err = row.Scan(&user.PK, &user.ID)
	if err != nil {
		return nil, fmt.Errorf("models: create user: %w", err)
	}
	return user, nil
}

func (us UserService) Authenticate(email, password string) (*User, error) {
	email = strings.ToLower(email)
	user := &User{
		Email: email,
	}
	row := us.DB.QueryRow(`
  		SELECT pk, id, first_name, last_name, age, password_hash
  		FROM users WHERE email=$1`,
		email,
	)
	err := row.Scan(&user.PK, &user.ID, &user.FirstName, &user.LastName, &user.Age, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("authenticate: %w", err)
	}

	password = password + "-" + "dvorak"
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("authenticate: %w", err)
	}
	return user, nil
}

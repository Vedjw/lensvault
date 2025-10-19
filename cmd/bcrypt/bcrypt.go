// Mini bcrypt CLI for Learning

package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type currHash struct {
	password []byte
	hash     []byte
}

func loadEnvfile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, err
}

func getEnvField(fd *[]byte, field string) string {
	lines := strings.Split(string(*fd), "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 && parts[0] == field {
			return parts[1]
		}
	}
	return "nf"
}

func hashNewPassword(pep string) *currHash {
	var passwordPlain string
	fmt.Println("Enter Your Password:")
	fmt.Scan(&passwordPlain)

	password := passwordPlain + "-" + pep
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return &currHash{
		password: []byte(password),
		hash:     hash,
	}
}

func comparepassword(chd *currHash, pep string) {
	var passwordPlain string
	fmt.Println("Enter Your Password:")
	fmt.Scan(&passwordPlain)

	password := passwordPlain + "-" + pep
	err := bcrypt.CompareHashAndPassword(chd.hash, []byte(password))
	if err != nil {
		fmt.Println("Invaild Password")
		return
	}
	fmt.Println("Correct Password")
}

func main() {
	envdata, err := loadEnvfile("../../.env")
	if err != nil {
		panic(err)
	}
	pepper := getEnvField(&envdata, "PASSWORD_PEPPER")
	if pepper == "nf" {
		fmt.Println("Environment variable not found")
	}

	var currHashData *currHash
	var cmd string
	for {
		fmt.Println("Enter command, hash, compare or quit:")
		fmt.Scan(&cmd)

		switch cmd {
		case "hash":
			currHashData = hashNewPassword(pepper)
			if currHashData != nil {
				fmt.Printf("Your Password: %s, Hash generated: %s\n", currHashData.password, currHashData.hash)
			}
		case "compare":
			comparepassword(currHashData, pepper)
		case "quit":
			fmt.Println("Come back again")
			return
		default:
			fmt.Println("Invalid Command")
		}
	}
}

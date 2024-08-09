package main

import (
	"fmt"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running tern command: %s\n", string(out))
		panic(err)
	}

	fmt.Println("Migrations applied successfully")
}

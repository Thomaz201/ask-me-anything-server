package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	fmt.Println("DATABASE_USERNAME:", os.Getenv("DATABASE_USERNAME"))
	fmt.Println("DATABASE_PASSWORD:", os.Getenv("DATABASE_PASSWORD"))
	fmt.Println("DATABASE_NAME:", os.Getenv("DATABASE_NAME"))

	cmd := exec.Command("tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

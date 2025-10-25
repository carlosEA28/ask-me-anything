package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		os.Exit(1)
	}

	// Define the tern command
	cmd := exec.Command("tern", "migrate",
		"--migrations", "./internal/store/pgstore/migrations",
		"--config", "./internal/store/pgstore/migrations/tern.conf")

	// Capture command output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	fmt.Println("Running migrations...")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running migrations: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Migrations completed successfully!")
}

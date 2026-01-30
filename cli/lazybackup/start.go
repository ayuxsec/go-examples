package main

import (
	"fmt"
)

func start(mode string) error {
	// Handle general flags
	if *debugEnabled {
		fmt.Println("[DEBUG] Debug mode enabled")
	}

	fmt.Println("Using config file:", *configFile)

	// Mode-specific logic
	switch mode {
	case "local":
		if sourcePath == "" {
			return fmt.Errorf("you must provide --source for local backup")
		}
		fmt.Println("Backing up locally from:", sourcePath)

	case "s3":
		if sourcePath == "" {
			return fmt.Errorf("you must provide --source for S3 backup")
		}
		fmt.Println("Backing up to S3 from:", sourcePath)

	case "restore":
		fmt.Println("Restoring backup...")

	default:
		return fmt.Errorf("unknown mode: %s", mode)
	}

	return nil
}

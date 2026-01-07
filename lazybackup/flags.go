package main

import "github.com/spf13/pflag"

var (
	// general flags
	debugEnabled *bool
	configFile   *string

	// backup flags
	sourcePath string
)

func createGeneralFlags() *pflag.FlagSet {
	generalFlags := pflag.NewFlagSet("general", pflag.ExitOnError)

	debugEnabled = generalFlags.Bool("debug", false, "Enable debug logging")
	configFile = generalFlags.String("config", "config.yaml", "Path to config file")

	return generalFlags
}

func createBackupFlags() *pflag.FlagSet {
	backupFlags := pflag.NewFlagSet("backup", pflag.ExitOnError)
	backupFlags.StringVar(&sourcePath, "source", "", "Source directory to backup")
	return backupFlags
}

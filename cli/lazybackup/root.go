package main

import "github.com/spf13/cobra"

var localCommand = &cobra.Command{
	Use:   "local",
	Short: "Backup files to local storage",
	RunE: func(cmd *cobra.Command, args []string) error {
		return start("local")
	},
}

var s3Command = &cobra.Command{
	Use:   "s3",
	Short: "Backup files to S3",
	RunE: func(cmd *cobra.Command, args []string) error {
		return start("s3")
	},
}

var restoreCommand = &cobra.Command{
	Use:   "restore",
	Short: "Restore files from backup",
	RunE: func(cmd *cobra.Command, args []string) error {
		return start("restore")
	},
}

func GetRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "lazybackup",
	}

	generalFlags := createGeneralFlags()
	backupFlags := createBackupFlags()

	rootCmd.AddCommand(localCommand)
	rootCmd.AddCommand(s3Command)
	rootCmd.AddCommand(restoreCommand)

	// attach flags
	localCommand.Flags().AddFlagSet(generalFlags)
	localCommand.Flags().AddFlagSet(backupFlags)

	s3Command.Flags().AddFlagSet(generalFlags)
	s3Command.Flags().AddFlagSet(backupFlags)

	restoreCommand.Flags().AddFlagSet(generalFlags)

	rootCmd.Flags().AddFlagSet(generalFlags)

	return rootCmd
}

/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
	"github.com/spf13/cobra"
)

type RunEFunc func(cmd *cobra.Command, args []string) error

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "16-cli_with_cobra",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func connectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}

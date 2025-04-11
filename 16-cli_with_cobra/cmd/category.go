/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/rafaelpapastamatiou/goexpert/16-cli_with_cobra/cmd/category"
	"github.com/rafaelpapastamatiou/goexpert/16-cli_with_cobra/internal/database"
	"github.com/spf13/cobra"
)

// CategoryCmd represents the category command
var CategoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(CategoryCmd)

	db, err := connectToDatabase()
	if err != nil {
		panic(err)
	}

	categoryDb := database.NewCategory(db)

	CategoryCmd.AddCommand(category.NewCreateCmd(categoryDb))
	CategoryCmd.AddCommand(category.NewListCmd(categoryDb))
}

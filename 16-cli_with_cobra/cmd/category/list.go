/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package category

import (
	"fmt"

	"github.com/rafaelpapastamatiou/goexpert/16-cli_with_cobra/internal/database"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
func NewListCmd(categoryDb *database.Category) *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all categories",
		Long:  "List all categories in the database.",
		RunE:  runList(categoryDb),
	}

	return listCmd
}

func runList(categoryDb *database.Category) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		categories, err := categoryDb.List()
		if err != nil {
			panic(err)
		}

		if len(categories) == 0 {
			fmt.Println("No categories found.")
			return nil
		}

		fmt.Println("Categories:")
		for _, category := range categories {
			fmt.Printf("ID: %s, Name: %s, Description: %s\n", category.ID, category.Name, category.Description)
		}

		return nil
	}
}

/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package category

import (
	"fmt"

	"github.com/rafaelpapastamatiou/goexpert/16-cli_with_cobra/internal/database"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
func NewCreateCmd(categoryDb *database.Category) *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new category",
		Long:  `Create a new category with the specified name and description.`,
		RunE:  runCreate(categoryDb),
	}

	createCmd.Flags().StringP("name", "n", "", "Name of the category")
	createCmd.MarkFlagRequired("name")

	createCmd.Flags().StringP("description", "d", "", "Description of the category")
	createCmd.MarkFlagRequired("description")

	return createCmd
}

func runCreate(categoryDb *database.Category) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		fmt.Printf("Creating category with name: %s and description: %s\n", name, description)

		category, err := categoryDb.Create(name, description)
		if err != nil {
			return err
		}

		fmt.Printf("Created category with ID: %s\n", category.ID)
		return nil
	}
}

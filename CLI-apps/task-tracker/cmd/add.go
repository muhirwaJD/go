package cmd

import (
	"fmt"
	"music/crud"
	"music/database"
	"music/models"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adding a new task",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		artist := args[1]
		year, err := strconv.Atoi(args[2])
		if err != nil {
			return fmt.Errorf("year must be a number, got %q", args[2])
		}
		err = crud.Create(database.DB, models.Song{Name: name, Artist: artist, Year: year})

		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

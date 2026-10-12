/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"music/crud"
	"music/database"
	"music/models"
	"strconv"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update <name> <artist> <year> <id>",
	Short: "Update a song by ID",
	Args:  cobra.ExactArgs(4),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		artist := args[1]
		year, err := strconv.Atoi(args[2])
		if err != nil {
			return fmt.Errorf("year must be a number, got %q", args[2])
		}
		id, err := strconv.Atoi(args[3])
		if err != nil {
			return fmt.Errorf("id must be a number, got %q", args[3])
		}
		updated, err := crud.Update(database.DB, models.Song{Name: name, Artist: artist, Year: year, ID: id})
		if err != nil {
			return err
		}
		fmt.Printf("Updated song #%d: %s — %s (%d)\n", updated.ID, updated.Name, updated.Artist, updated.Year)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

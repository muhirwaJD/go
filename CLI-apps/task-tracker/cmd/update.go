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
	Use:   "update",
	Short: "Updating a Song",
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
		err = crud.Update(database.DB, models.Song{id, name, artist, year})
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

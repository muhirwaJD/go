package cmd

import (
	"fmt"
	"music/crud"
	"music/database"
	"music/models"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a song by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("id must be a number, got %q", args[0])
		}

		deletedSong, err := crud.DeleteSong(database.DB, models.Song{ID: id})
		if err != nil {
			return err
		}

		fmt.Printf("Deleted song #%d: %s — %s (%d)\n", deletedSong.ID, deletedSong.Name, deletedSong.Artist, deletedSong.Year)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

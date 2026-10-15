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
	Use:   "delete",
	Short: "Deleting a Song",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("id must be a number, got %q", args[3])
		}

		deletedSong, err := crud.DeleteSong(database.DB, models.Song{
			ID: id,
		})
		fmt.Printf("Deleted Song: %s by %s, year %d\n", deletedSong.Name, deletedSong.Artist, deletedSong.Year)
		return err
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

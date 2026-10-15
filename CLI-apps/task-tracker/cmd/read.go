package cmd

import (
	"fmt"
	"music/crud"
	"music/database"

	"github.com/spf13/cobra"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "List all songs",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		songs, err := crud.ReadSongs(database.DB)

		if err != nil {
			return err
		}
		if len(songs) == 0 {
			fmt.Println(`No songs yet. Add one with: music add "<name>" "<artist>" <year>`)
			return nil
		}
		for _, s := range songs {
			fmt.Printf("#%d: %s — %s (%d)\n", s.ID, s.Name, s.Artist, s.Year)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}

package main

import (
	"fmt"
	"log"
	"mymodule/database"
	"mymodule/models"
)

var songs = []models.Song{
	{Name: "Blinding Lights", Artist: "The Weeknd", Year: 2019},
	{Name: "Starboy", Artist: "The Weeknd", Year: 2016},
	{Name: "Essence", Artist: "Wizkid ft. Tems", Year: 2020},
	{Name: "Money Constant", Artist: "Davido ft. Tems", Year: 2020},
}

var updateSongs = []models.Song{
	{Name: "Degaine", Artist: "Aya Nakamura and Damso", Year: 2018, ID: 73},
}

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect: ", err)
	}
	fmt.Println("Connected toc database ✓")
	defer db.Close()

	if err := database.SetupConstraints(db); err != nil {
		log.Fatal("Failed to setup constraints: ", err)
	}
	fmt.Println("Constraints set up ✓")

	// Insert songs (skip if duplicate)
	//fmt.Println("\nInserting songs...")
	//for _, song := range songs {
	//	err := database.InsertSong(db, song)
	//	if err != nil {
	//		// Ignore duplicate key error (constraint working as intended)
	//		if strings.Contains(err.Error(), "unique constraint") {
	//			fmt.Printf("⚠ Skipped duplicate: %s by %s\n", song.Name, song.Artist)
	//		} else {
	//			log.Fatal("Insert error:", err)
	//		}
	//	} else {
	//		fmt.Printf("✓ Inserted: %s by %s (%d)\n", song.Name, song.Artist, song.Year)
	//	}
	//}

	// Pull and display all songs
	fmt.Println("\n--- Songs in Database ---")
	data, err := database.ReadSongs(db)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, song := range data {
		fmt.Printf("[ID: %d] %s by %s (%d)\n", song.ID, song.Name, song.Artist, song.Year)
	}

	// Update a song
	//fmt.Println("\n--- Updating a song in the Database ---")
	////fmt.Println(updateSongs[0])
	//for _, newSong := range updateSongs {
	//	err := database.UpdateSong(db, newSong)
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println("Tres Bien Song Updated😆")
	//}

	fmt.Println("\n--- Deleting a song in the Database ---")
	var deleteSong models.Song
	deleteSong.ID = 3
	var id = deleteSong.ID

	//fmt.Println("This is what we are SENDING", id)
	err = database.DeleteSong(db, id)
	if err != nil {
		log.Fatal(err)
	}
}

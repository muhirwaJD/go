package main

import (
	"fmt"
	"net/http"
)

type Songs struct {
	Name   string
	Artist string
}

func helloJd(w http.ResponseWriter, r *http.Request) {
	const name = "David"
	_, err := fmt.Fprintln(w, "Hello", name)
	if err != nil {
		return
	}
}

func printSong(b Songs, w http.ResponseWriter) {
	_, err := fmt.Fprintf(w, "Name: %s\n", b.Name)
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(w, "Artist: %s\n\n", b.Artist)
	if err != nil {
		return
	}
}

func songDetails(w http.ResponseWriter, r *http.Request) {
	var song1 Songs
	var song2 Songs

	song1.Name = "Mami Wata"
	song1.Artist = "Gazo et Tiakola"

	song2.Name = "Jolie Bebe"
	song2.Artist = "Naza"

	printSong(song1, w)
	printSong(song2, w)

	c := map[int]string{1: "Davinah", 2: "Davinah", 3: "Davinah"}
	fmt.Println(c)
}


// Learning GO
// Simple M3U player
// Author: Tales Luna <tales.ferreira.luna@gmail.com>

// Floyd Player (with mplayer)

package main

import "os"

func main() {

	// Set m3u file ... OH GOD!! Why you do not create a "Open File Dialog" ?

	if len(os.Args) > 1 {
		FILE := os.Args[1]
		// Parse to list
		playlist := readM3U(FILE)

		// OK, open GUI
		if len(playlist.Medias) > 0 {
			setPlaylist(playlist)
		} else {
			println("\n [!] " + FILE + " is empty or not a valid M3U list \n")
		}
	} else {
		println("\n [!] Use: ./floyd <path/you/list.m3u> \n")
	}
}

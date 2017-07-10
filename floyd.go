
// Learning GO
// Simple M3U player
// Author: Tales Luna <tales.ferreira.luna@gmail.com>

// Floyd Player (with mplayer)

package main

func main() {

	// Set m3u file ... OH GOD!! Why you do not create a "Open File Dialog" ?
	FILE := "path/to/you/file.m3u"

	// Parse to list
	playlist := readM3U(FILE)

	// OK, open GUI
	if len(playlist.Medias) > 0 {
		setPlaylist(playlist)
	}


}

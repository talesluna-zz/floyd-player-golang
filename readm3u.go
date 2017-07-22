
// Learning GO
// Simple M3U player
// Author: Tales Luna <tales.ferreira.luna@gmail.com>

// Floyd Player (with mplayer)

package main

import (
	"strings"
	"os"
	"bufio"
	"strconv"
)

// This function receive a m3u file path, read the file and parse file lines in Playlist
func readM3U(path string) (Playlist) {

	// Try open file and read buffer
	file, err := os.Open(path)
	if err != nil {
		println("\n [!] " + err.Error() + " \n")
		os.Exit(0)
	}

	defer file.Close()

	// Get file and create a empty playlist
	buf 		:= bufio.NewScanner(file)
	playlist 	:= Playlist{}

	// M3U prefixes
	logoPrefix	:= "tvg-logo="
	groupPrefix	:= "group-title="

	// Media struct
	mediaInfo	:= &MediaInfo{}

	// To control new medias
	newMedia	:= false
	mediaNumber	:= 1

	// Read file lines to populate playlist
	for buf.Scan() {

		extLine := buf.Text()

		// If is a new media (media is composed of two lines) create new "mediaInfo" for newMedia
		if newMedia {
			mediaInfo = &MediaInfo{}
			newMedia = false
		}

		// If exist '#EXTINF' in line, this line have a media information
		if strings.HasPrefix(extLine, "#EXTINF") {

			// Split line to get all prefix of info
			extLine 		:= strings.SplitN(extLine, ",", -1)
			mediaInfo.Name 	= strings.TrimSpace(extLine[1])
			extLine 		= strings.Split(extLine[0], " ")

			// Looping for info prefixes
			for i := 0; i < len(extLine); i++ {
				if strings.HasPrefix(extLine[i], logoPrefix) {

					mediaInfo.LogoURI = getPrefixText(extLine[i], logoPrefix)

				} else if strings.HasPrefix(extLine[i], groupPrefix) {

					mediaInfo.Group = getPrefixText(extLine[i], groupPrefix)

				}
			}

		} else if !strings.HasPrefix(extLine, "#") {
			// If in line no have nothing #, this line is a literal URI of media
			// Too set a newMedia true, and go to next
			mediaInfo.MediaURI	= extLine
			newMedia = true
			mediaNumber++
		} else {
			// Else ... continue scan
			continue
		}

		// If Media have a Name and MediaURI, insert there in playlist
		if len(mediaInfo.Name) > 0 && len(mediaInfo.MediaURI) > 0 {

			playlist.MediasInfo	= append(playlist.MediasInfo, *mediaInfo)
			playlist.Medias		= append(playlist.Medias, strconv.Itoa(mediaNumber) + " - " + mediaInfo.Name)

		}
	}

	// After scan, return a playlist
	return playlist
}

// Get value of prefix "tag"
func getPrefixText(text string, prefix string) (string)  {
	text = strings.Replace(text, prefix, "", 1)
	text = strings.Replace(text, "\"", "", -1)
	return strings.TrimSpace(text)
}
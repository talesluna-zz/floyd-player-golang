
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

/*
	Parse M3U List to array of Media
 */
func readM3U(path string) (Playlist) {

	// Try open file and read buffer
	file, err := os.Open(path)
	if err != nil {
		println("\n [!] " + err.Error() + " \n")
		os.Exit(0)
	}

	defer file.Close()

	buf 		:= bufio.NewScanner(file)
	playlist 	:= Playlist{}

	// M3U prefixes
	logoPrefix 	:= "tvg-logo="
	groupPrefix := "group-title="

	// Media struct
	mediaInfo 	:= &MediaInfo{}

	// To control new medias
	newMedia	:= false
	mediaNumber	:= 1

	for buf.Scan() {

		extLine := buf.Text()

		if newMedia {
			mediaInfo = &MediaInfo{}
			newMedia = false
		}

		if strings.HasPrefix(extLine, "#EXTINF") {
			extLine := strings.SplitN(extLine, ",", -1)

			mediaInfo.Name = strings.TrimSpace(extLine[1])

			extLine = strings.Split(extLine[0], " ")

			for i := 0; i < len(extLine); i++ {
				if strings.HasPrefix(extLine[i], logoPrefix) {

					mediaInfo.LogoURI = getPrefixText(extLine[i], logoPrefix)

				} else if strings.HasPrefix(extLine[i], groupPrefix) {

					mediaInfo.Group = getPrefixText(extLine[i], groupPrefix)

				}
			}

		} else if !strings.HasPrefix(extLine, "#") {
			mediaInfo.MediaURI	= extLine
			newMedia = true
			mediaNumber++
		} else {
			continue
		}

		if len(mediaInfo.Name) > 0 && len(mediaInfo.MediaURI) > 0 {
			playlist.MediasInfo = append(playlist.MediasInfo, *mediaInfo)
			playlist.Medias		= append(playlist.Medias, strconv.Itoa(mediaNumber) + " - " + mediaInfo.Name)
		}
	}
	return playlist
}

/*
	Get a tag data by prefix
 */
func getPrefixText(text string, prefix string) (string)  {
	text = strings.Replace(text, prefix, "", 1)
	text = strings.Replace(text, "\"", "", -1)
	return strings.TrimSpace(text)
}
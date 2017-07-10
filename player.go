
// Learning GO
// Simple M3U player
// Author: Tales Luna <tales.ferreira.luna@gmail.com>

// Floyd Player (with mplayer)

package main

import (
	"os"
	"os/exec"
	"github.com/google/gxui"
	"github.com/google/gxui/themes/dark"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/math"
	"net/http"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	_ "image/gif"
	_ "golang.org/x/image/bmp"
)

type Playlist struct {
	Medias		[]string
	MediasInfo 	[]MediaInfo
}

type MediaInfo struct {
	Name		string
	Group		string
	LogoURI		string
	MediaURI	string
}

var playlist 		= Playlist{}
var defaultTitle	= "Floyd Player"

/*
	After loaded, set playlist to populate GUI list
 */
func setPlaylist(list Playlist)  {
	playlist = list
	// Start UI
	gl.StartDriver(initPlayerWindow)
}

/*
	Start a main window with list of medias
 */
func initPlayerWindow(driver gxui.Driver)  {

	// Window init settings
	width 	:= 1424
	height 	:= 768
	title	:= defaultTitle

	// Create Window
	theme	:= dark.CreateTheme(driver)
	window	:= theme.CreateWindow(width, height, title)

	// Window pos-properties
	window.SetBorderPen(gxui.Pen{5, gxui.Gray15})
	window.SetPadding(math.Spacing{10,0,10,0})
	window.OnClose(closePlayer)

	// Media info panel
	layout := theme.CreateSplitterLayout()
	layout.SetOrientation(gxui.Horizontal)

	// Panels, right and left
	panelLeft 	:= theme.CreatePanelHolder()
	panelRight 	:= theme.CreatePanelHolder()
	panelLeft.SetSize(math.Size{400, 600})

	// Info text
	infoLayout 	:= theme.CreateSplitterLayout()
	infoLabel 	:= theme.CreateLabel()
	infoLabel.SetVerticalAlignment(gxui.AlignTop)
	infoLabel.SetText("Selecione algo na playlist")
	infoLayout.AddChild(infoLabel)

	// List Adapter
	listAdapter := gxui.CreateDefaultAdapter()
	listAdapter.SetItems(playlist.Medias)
	listAdapter.SetSize(math.Size{400, 40})

	// List for playlist medias
	playlistTree := theme.CreateList()
	playlistTree.SetAdapter(listAdapter)
	playlistTree.OnSelectionChanged(func(item gxui.AdapterItem) {

		// Get media info
		infoIndex := listAdapter.ItemIndex(item)
		media := playlist.MediasInfo[infoIndex]

		// Change window title
		window.SetTitle(defaultTitle + " | " + media.Name)
		panelRight.RemovePanel(infoLayout)
		infoLayout = getInfoLayout(media, theme)
		panelRight.AddPanel(infoLayout, "Assistir")
	})

	// Add elements to panels
	panelLeft.AddPanel(playlistTree, "Minha Playlist")
	panelRight.AddPanel(infoLayout, "Assistir")

	// Add panels to layout
	layout.AddChild(panelLeft)
	layout.AddChild(panelRight)
	layout.SetChildWeight(panelLeft, 3)

	// Add layout to window
	window.AddChild(layout)
}

/*
	Make a layout with selected media info
 */
func getInfoLayout(media MediaInfo, theme gxui.Theme) (gxui.SplitterLayout) {

	infoLayout 	:= theme.CreateSplitterLayout()
	titleLabel 	:= theme.CreateLabel()

	// Show media name and group
	titleLabel.SetText(media.Group + " - " + media.Name)
	titleLabel.SetVerticalAlignment(gxui.AlignMiddle)
	titleLabel.SetHorizontalAlignment(gxui.AlignCenter)
	titleLabel.SetColor(gxui.Gray40)
	titleLabel.SetMargin(math.Spacing{5,5,5,5})

	// Load RGBA of log to make texture
	logoRgba := loadLogo(media.LogoURI)
	if logoRgba != nil {

		texture := theme.Driver().CreateTexture(logoRgba, 1)

		imageLogo := theme.CreateImage()
		imageLogo.SetMargin(math.Spacing{10,10,0,0})
		imageLogo.SetTexture(texture)

		infoLayout.AddChild(imageLogo)
		infoLayout.SetChildWeight(imageLogo, 1)
	}

	// Play button
	playButton := theme.CreateButton()
	playButton.SetText("Assistir")
	playButton.SetVerticalAlignment(gxui.AlignMiddle)
	playButton.SetHorizontalAlignment(gxui.AlignCenter)
	playButton.SetBackgroundBrush(gxui.Brush{gxui.Green20})

	// Exec mplayer and play media link
	playButton.OnClick(func(event gxui.MouseEvent) {
		exec.Command("mplayer","-zoom","-fs",media.MediaURI).Run()
	})

	// Info layout childs
	infoLayout.AddChild(titleLabel)
	infoLayout.AddChild(playButton)
	infoLayout.SetChildWeight(titleLabel, 0.8)
	infoLayout.SetChildWeight(playButton, 0.2)

	return infoLayout
}

/*
	Only to close appliection
 */
func closePlayer()  {
	os.Exit(0)
}

/*
	This function get a HTTP image and return to set on info layout
 */
func loadLogo(url string) (image.Image) {

	// Get a web image
	resp, err := http.Get(url)

	if err != nil {
		return nil
	}

	// Decode to image
	source, _, err := image.Decode(resp.Body)

	if err != nil {
		return  nil
	}

	// Set a image size
	pointMax 	:= image.Point{source.Bounds().Max.X,source.Bounds().Max.Y}
	pointMin 	:= image.Point{source.Bounds().Min.X,source.Bounds().Min.Y}
	rectangle 	:= image.Rectangle{pointMin,pointMax}

	// Draw image rgba
	rgba := image.NewRGBA(rectangle)
	draw.Draw(rgba, rectangle, source, image.ZP, draw.Src)

	return rgba
}
package main

import "encoding/xml"

type Gamelist struct {
	XMLName xml.Name `xml:"gameList"`
	Games   []Game   `xml:"game"`
	Folders []Folder `xml:"folder"`
}

type Game struct {
	XMLName     xml.Name `xml:"game"`
	Path        string   `xml:"path"`
	Name        string   `xml:"name"`
	Description string   `xml:"desc"`
	ImagePath   string   `xml:"image"`
	VideoPath   string   `xml:"video"`
	MarqueePath string   `xml:"marquee"`
	ReleaseDate string   `xml:"releasedate"`
	Thumbnail   string   `xml:"thumbnail"`
	Developer   string   `xml:"developer"`
	Publisher   string   `xml:"publisher"`
	Genre       string   `xml:"genre"`
	Players     string   `xml:"players"`
	Rating      string   `xml:"rating"`
}

type Folder struct {
	XMLName     xml.Name `xml:"folder"`
	Path        string   `xml:"path"`
	Name        string   `xml:"name"`
	Description string   `xml:"desc"`
	ImagePath   string   `xml:"image"`
	VideoPath   string   `xml:"video"`
	MarqueePath string   `xml:"marquee"`
	ReleaseDate string   `xml:"releasedate"`
	Thumbnail   string   `xml:"thumbnail"`
	Developer   string   `xml:"developer"`
	Publisher   string   `xml:"publisher"`
	Genre       string   `xml:"genre"`
	Players     string   `xml:"players"`
	Rating      string   `xml:"rating"`
}

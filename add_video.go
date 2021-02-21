package main

import (
	"encoding/csv"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	gamelistInputPath := flag.String("gamelist", "", "path to input gamelist.xml file.")
	fileMapPath := flag.String("filemap", "", "path to video file map CSV.")
	videoFileDir := flag.String("videodir", "", "path to directory with video files.")
	gamelistOutPath := flag.String("out", "", "path to output gamelist.xml file.")

	flag.Parse()

	if gamelistInputPath == nil || fileMapPath == nil || videoFileDir == nil || gamelistOutPath == nil {
		flag.PrintDefaults()
		return
	}

	videoFileMap, err := readCSVAsMap(*fileMapPath, 0, 1)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Couldn't open %v", fileMapPath), err)
		return
	}

	gamelist, err := readGamelist(*gamelistInputPath)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Couldn't open %v", gamelistInputPath), err)
		return
	}

	out := &Gamelist{}

	for _, game := range gamelist.Games {
		gameFilename := filepath.Base(game.Path)
		//fmt.Println(gameFilename)
		videoFilename, ok := videoFileMap[gameFilename]
		if !ok {
			out.Games = append(out.Games, game)
			continue // this is fine - we just don't have a video for this one at the moment!
		}

		videoPath := *videoFileDir + videoFilename

		game.VideoPath = videoPath

		out.Games = append(out.Games, game)
	}

	//for _, folder := range gamelist.Folders {
	//	fmt.Printf("Folder: %s\n", folder.Name)
	//}

	if err := writeGamelist(out, *gamelistOutPath); err != nil {
		log.Fatalln("Couldn't write output gamelist: ", err)
		return
	}
}

func readCSVAsMap(filename string, keyCol, valCol int) (map[string]string, error) {
	fileHandle, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	outMap := make(map[string]string)

	csvReader := csv.NewReader(fileHandle)
	for {
		csvRow, err := csvReader.Read()
		if err == io.EOF {
			return outMap, nil
		}
		if err != nil {
			return nil, err
		}

		key := csvRow[keyCol]
		val := csvRow[valCol]

		outMap[key] = val
	}
}

func writeGamelist(gamelist *Gamelist, path string) error {
	file, err := xml.MarshalIndent(gamelist, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, file, 0644)
}

func readGamelist(filename string) (*Gamelist, error) {
	fileHandle, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	byteValue, err := ioutil.ReadAll(fileHandle)
	if err != nil {
		return nil, err
	}

	gamelist := &Gamelist{}
	err = xml.Unmarshal(byteValue, gamelist)
	if err != nil {
		return nil, err
	}

	if err := fileHandle.Close(); err != nil {
		return nil, err
	}

	return gamelist, nil
}

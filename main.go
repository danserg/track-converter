package main

import (
	"encoding/xml"
	"fmt"
	"github.com/danserg/track-converter/gpstypes"
	"golang.org/x/net/html/charset"
	"io"
	"os"
)

// ParseFile returns a Gpx gpstypes from gpx file
func ParseFile(path string) (*Gpx, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %s", err)
	}
	return parse(file)
}

func parse(r io.Reader) (*Gpx, error) {
	var gpx Gpx
	d := xml.NewDecoder(r)
	d.CharsetReader = charset.NewReaderLabel
	err := d.Decode(&gpx)
	if err != nil {
		return nil, fmt.Errorf("could not decode gpx to gpstypes: %s", err)
	}
	return &gpx, nil
}

func main() {
	g, err := ParseFile("./sample/garmin-tracks.gpx")
	if err != nil {
		panic(err)
	}
	fmt.Println(g.Version)
	for _, track := range g.Trk {
		for _, segment := range track.Trkseg {
			for _, pt := range segment.Trkpt {
				// Do something with pt.Lat, pt.Lon, etc...
				fmt.Println("pt: ", pt.Lat, pt.Lon, pt.Ele)

			}
			//fmt.Println("segment:", segment)
		}
		fmt.Println("track:", track.Trkseg)
	}

	for _, wpt := range g.Wpt {
		fmt.Println("waypoints:", wpt.Name, wpt.Cmt, wpt.Lat, wpt.Lon, wpt.Ele)

	}

}

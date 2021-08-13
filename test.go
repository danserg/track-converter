package main

import (
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"io"
	"os"
	"track-converter/gpstypes"
)

// ParseFile returns a Gpx gpstypes from gpx file
func ParseFile(path string) (*gpstypes.Kml, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %s", err)
	}
	return parse(file)
}

func parse(r io.Reader) (*gpstypes.Kml, error) {
	var kml gpstypes.Kml
	d := xml.NewDecoder(r)
	d.CharsetReader = charset.NewReaderLabel
	err := d.Decode(&kml)
	if err != nil {
		return nil, fmt.Errorf("could not decode gpx to gpstypes: %s", err)
	}
	return &kml, nil
}

func main() {
	g, err := ParseFile("./sample/desna_splav.kml")
	if err != nil {
		panic(err)
	}
	//fmt.Println(g.Document.Placemark.Name)
	for _, name := range g.Document.Folder.Folder.Placemark {
		fmt.Println("name: ", name.Name)
	}
	//for _, track := range g.Trk {
	//	for _, segment := range track.Trkseg {
	//		for _, pt := range segment.Trkpt {
	//			// Do something with pt.Lat, pt.Lon, etc...
	//			fmt.Println("pt: ", pt.Lat, pt.Lon, pt.Ele)
	//		}
	//	}
	//}
	//for _, wpt := range g.Wpt {
	//	fmt.Println("waypoints:", wpt.Name, wpt.Lat, wpt.Lon)
	//}
}

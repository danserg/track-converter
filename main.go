package main

import (
	"encoding/xml"
	"time"
	"os"
	"fmt"
	"io"
	"golang.org/x/net/html/charset"
)

// Gpx is a root node in .gpx document.
type Gpx struct {
	XMLName  xml.Name  `xml:"gpx"`
	Version  string    `xml:"version,attr"`
	Creator  string    `xml:"creator,attr"`
	Wpt      []*Wpt    `xml:"wpt"`
	Rte      []*Rte    `xml:"rte"`
	Trk      []*Trk    `xml:"trk"`

}
// Rte represents route - an ordered list of waypoints representing a series of turn points leading to a destination.
type Rte struct {
	Name   string  `xml:"name"`
	Cmt    string  `xml:"cmt"`
	Desc   string  `xml:"desc"`
	Src    string  `xml:"src"`
	Number uint    `xml:"number"`
	Type   string  `xml:"type"`
	Rtept  []*Wpt  `xml:"rtept"`
}

// Trk represents a track - an ordered list of points describing a path.
type Trk struct {
	Name   string    `xml:"name"`
	Cmt    string    `xml:"cmt"`
	Desc   string    `xml:"desc"`
	Src    string    `xml:"src"`
	Number uint      `xml:"number"`
	Type   string    `xml:"type"`
	Trkseg []*Trkseg `xml:"trkseg"`
}

// Trkseg holds a list of Track Points which are logically connected in order.
type Trkseg struct {
	Trkpt []*Wpt `xml:"trkpt"`
}

// Wpt represents a waypoint, point of interest, or named feature on a map.
type Wpt struct {
	Lat          float64   `xml:"lat,attr"`
	Lon          float64   `xml:"lon,attr"`
	Ele          float64   `xml:"ele"`
	Time         time.Time `xml:"time"`
	Magvar       float64   `xml:"magvar"`
	GeoIDHeight  float64   `xml:"geoidheight"`
	Name         string    `xml:"name"`
	Cmt          string    `xml:"cmt"`
	Desc         string    `xml:"desc"`
	Src          string    `xml:"src"`
	Sym          string    `xml:"sym"`
	Type         string    `xml:"type"`
	Fix          string    `xml:"fix"` // http://www.topografix.com/gpx/1/1/#type_fixType
	Sat          uint      `xml:"sat"`
	Hdop         float64   `xml:"hdop"`
	Vdop         float64   `xml:"vdop"`
	Pdop         float64   `xml:"pdop"`
	AgeOfGPSData float64   `xml:"ageofgpsdata"`
	DGPSID       uint      `xml:"dgpsid"`
}

// ParseFile returns a Gpx struct from gpx file
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
		return nil, fmt.Errorf("could not decode gpx to struct: %s", err)
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

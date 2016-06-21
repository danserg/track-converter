package main

import (
	"encoding/xml"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Gpx struct {
	XMLName      xml.Name `xml:"http://www.topografix.com/GPX/1/1 gpx"`
	XmlNsXsi     string   `xml:"xmlns:xsi,attr,omitempty"`
	XmlSchemaLoc string   `xml:"xsi:schemaLocation,attr,omitempty"`
	Version      string   `xml:"version,attr"`
	Creator      string   `xml:"creator,attr"`
	Tracks       []GpxTrk `xml:"trk"`
}

type Trk struct {
	XMLName string `xml:"trkseg"`
}

type GpxWpt struct {
	Lat          float64   `xml:"lat,attr"`
	Lon          float64   `xml:"lon,attr"`
	Ele          float64   `xml:"ele,omitempty"`
	Timestamp    time.Time `xml:"time,omitempty"`
	MagVar       string    `xml:"magvar,omitempty"`
	GeoIdHeight  string    `xml:"geoidheight,omitempty"`
	Name         string    `xml:"name,omitempty"`
	Cmt          string    `xml:"cmt,omitempty"`
	Desc         string    `xml:"desc,omitempty"`
	Src          string    `xml:"src,omitempty"`
	Sym          string    `xml:"sym,omitempty"`
	Type         string    `xml:"type,omitempty"`
	Fix          string    `xml:"fix,omitempty"`
	Sat          int       `xml:"sat,omitempty"`
	Hdop         float64   `xml:"hdop,omitempty"`
	Vdop         float64   `xml:"vdop,omitempty"`
	Pdop         float64   `xml:"pdop,omitempty"`
	AgeOfGpsData float64   `xml:"ageofgpsdata,omitempty"`
	DGpsId       int       `xml:"dgpsid,omitempty"`
}

type GpxTrkseg struct {
	XMLName xml.Name `xml:"trkseg"`
	Points  []GpxWpt `xml:"trkpt"`
}

type GpxTrk struct {
	XMLName  xml.Name    `xml:"trk"`
	Name     string      `xml:"name,omitempty"`
	Cmt      string      `xml:"cmt,omitempty"`
	Desc     string      `xml:"desc,omitempty"`
	Src      string      `xml:"src,omitempty"`
	Number   int         `xml:"number,omitempty"`
	Type     string      `xml:"type,omitempty"`
	Segments []GpxTrkseg `xml:"trkseg"`
}

func main() {

}

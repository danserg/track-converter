package gpstypes

import (
	"encoding/xml"
	"time"
)

// Gpx is a root node in .gpx document.
type Gpx struct {
	XMLName xml.Name `xml:"gpx"`
	Version string   `xml:"version,attr"`
	Creator string   `xml:"creator,attr"`
	Wpt     []*Wpt   `xml:"wpt"`
	Rte     []*Rte   `xml:"rte"`
	Trk     []*Trk   `xml:"trk"`
}

// Rte represents route - an ordered list of waypoints representing a series of turn points leading to a destination.
type Rte struct {
	Name   string `xml:"name"`
	Cmt    string `xml:"cmt"`
	Desc   string `xml:"desc"`
	Src    string `xml:"src"`
	Number uint   `xml:"number"`
	Type   string `xml:"type"`
	Rtept  []*Wpt `xml:"rtept"`
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

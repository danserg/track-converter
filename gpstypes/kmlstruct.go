package gpstypes

import "encoding/xml"

type Kml struct {
	XMLName  xml.Name `xml:"kml"`
	Text     string   `xml:",chardata"`
	Xmlns    string   `xml:"xmlns,attr"`
	Document struct {
		Text   string `xml:",chardata"`
		Folder struct {
			Text  string `xml:",chardata"`
			Name  string `xml:"name"`
			Open  string `xml:"open"`
			Style struct {
				Text      string `xml:",chardata"`
				ListStyle struct {
					Text         string `xml:",chardata"`
					ListItemType string `xml:"listItemType"`
					BgColor      string `xml:"bgColor"`
				} `xml:"ListStyle"`
			} `xml:"Style"`
			Folder struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name"`
				Open  string `xml:"open"`
				Style struct {
					Text      string `xml:",chardata"`
					ListStyle struct {
						Text         string `xml:",chardata"`
						ListItemType string `xml:"listItemType"`
						BgColor      string `xml:"bgColor"`
					} `xml:"ListStyle"`
				} `xml:"Style"`
				Placemark []struct {
					Text        string `xml:",chardata"`
					Name        string `xml:"name"`
					Description string `xml:"description"`
					Style       struct {
						Text       string `xml:",chardata"`
						LabelStyle struct {
							Text  string `xml:",chardata"`
							Color string `xml:"color"`
							Scale string `xml:"scale"`
						} `xml:"LabelStyle"`
						IconStyle struct {
							Text  string `xml:",chardata"`
							Scale string `xml:"scale"`
							Icon  struct {
								Text string `xml:",chardata"`
								Href string `xml:"href"`
							} `xml:"Icon"`
							HotSpot struct {
								Text   string `xml:",chardata"`
								X      string `xml:"x,attr"`
								Y      string `xml:"y,attr"`
								Xunits string `xml:"xunits,attr"`
								Yunits string `xml:"yunits,attr"`
							} `xml:"hotSpot"`
						} `xml:"IconStyle"`
						LineStyle struct {
							Text  string `xml:",chardata"`
							Color string `xml:"color"`
							Width string `xml:"width"`
						} `xml:"LineStyle"`
					} `xml:"Style"`
					Point struct {
						Text        string `xml:",chardata"`
						Extrude     string `xml:"extrude"`
						Coordinates string `xml:"coordinates"`
					} `xml:"Point"`
					LineString struct {
						Text        string `xml:",chardata"`
						Extrude     string `xml:"extrude"`
						Coordinates string `xml:"coordinates"`
					} `xml:"LineString"`
				} `xml:"Placemark"`
			} `xml:"Folder"`
		} `xml:"Folder"`
	} `xml:"Document"`
}


package entities

import "encoding/xml"

// PropStat is metadata about a property
type PropStat struct {
	XMLName xml.Name `xml:"propstat"`
	Status  string   `xml:"status"`
	Prop    *Prop    `xml:",omitempty"`
}

// Response is a multistatus response entity
type Response struct {
	XMLName   xml.Name    `xml:"response"`
	Href      string      `xml:"href"`
	PropStats []*PropStat `xml:"propstat,omitempty"`
}

// Multistatus is a request to find properties on an an entity or collection
type Multistatus struct {
	XMLName   xml.Name    `xml:"DAV: multistatus"`
	Responses []*Response `xml:"response,omitempty"`
}

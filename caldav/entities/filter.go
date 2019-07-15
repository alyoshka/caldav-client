package entities

import (
	"encoding/xml"

	"github.com/alyoshka/caldav-go/caldav/values"
	"github.com/alyoshka/caldav-go/icalendar/properties"
)

// Filter is a CalDAV query filter entity
type Filter struct {
	XMLName         xml.Name         `xml:"urn:ietf:params:xml:ns:caldav filter"`
	ComponentFilter *ComponentFilter `xml:",omitempty"`
}

// ComponentFilter used to filter down calendar components, such as VCALENDAR > VEVENT
type ComponentFilter struct {
	XMLName         xml.Name             `xml:"urn:ietf:params:xml:ns:caldav comp-filter"`
	Name            values.ComponentName `xml:"name,attr"`
	ComponentFilter *ComponentFilter     `xml:",omitempty"`
	TimeRange       *TimeRange           `xml:",omitempty"`
	PropertyFilter  *PropertyFilter      `xml:",omitempty"`
	ParameterFilter *ParameterFilter     `xml:",omitempty"`
}

// TimeRange used to restrict component filters to a particular time range
type TimeRange struct {
	XMLName   xml.Name         `xml:"urn:ietf:params:xml:ns:caldav time-range"`
	StartTime *values.DateTime `xml:"start,attr"`
	EndTime   *values.DateTime `xml:"end,attr"`
}

// PropertyFilter used to restrict component filters to a property value
type PropertyFilter struct {
	XMLName         xml.Name                `xml:"urn:ietf:params:xml:ns:caldav prop-filter"`
	Name            properties.PropertyName `xml:"name,attr"`
	TextMatch       *TextMatch              `xml:",omitempty"`
	ParameterFilter *ParameterFilter        `xml:",omitempty"`
}

// ParameterFilter used to restrict component filters to a parameter value
type ParameterFilter struct {
	XMLName   xml.Name                 `xml:"urn:ietf:params:xml:ns:caldav param-filter"`
	Name      properties.ParameterName `xml:"name,attr"`
	TextMatch *TextMatch               `xml:",omitempty"`
}

// TextMatch used to match properties by text value
type TextMatch struct {
	XMLName         xml.Name             `xml:"urn:ietf:params:xml:ns:caldav text-match"`
	Collation       values.TextCollation `xml:"collation,attr,omitempty"`
	NegateCondition values.HumanBoolean  `xml:"attr,negate-condition,omitempty"`
	Content         string               `xml:",innerxml"`
}

// NewPropertyMatcher creates a new CalDAV property value matcher
func NewPropertyMatcher(name properties.PropertyName, content string) *PropertyFilter {
	pf := new(PropertyFilter)
	pf.Name = name
	pf.TextMatch = new(TextMatch)
	pf.TextMatch.Content = content
	return pf
}

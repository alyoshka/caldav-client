package properties

import "strings"

type PropertyName string

const (
	UIDPropertyName                 PropertyName = "UID"
	CommentPropertyName                          = "COMMENT"
	OrganizerPropertyName                        = "ORGANIZER"
	AttendeePropertyName                         = "ATTENDEE"
	ExceptionDateTimesPropertyName               = "EXDATE"
	RecurrenceDateTimesPropertyName              = "RDATE"
)

type ParameterName string

const (
	CanonicalNameParameterName ParameterName = "CN"
	TimeZoneIdPropertyName                   = "TZID"
)

func (p PropertyName) Equals(test string) bool {
	return strings.EqualFold(string(p), test)
}

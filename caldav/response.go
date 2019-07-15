package caldav

import (
	"io/ioutil"
	"log"

	"github.com/alyoshka/caldav-go/icalendar"
	"github.com/alyoshka/caldav-go/utils"
	"github.com/alyoshka/caldav-go/webdav"
)

var _ = log.Print

// Response is a WebDAV response object
type Response webdav.Response

// WebDAV downcasts the response to the WebDAV interface
func (r *Response) WebDAV() *webdav.Response {
	return (*webdav.Response)(r)
}

// Decode decodes a CalDAV iCalendar response into the provided interface
func (r *Response) Decode(into interface{}) error {
	if body := r.Body; body == nil {
		return nil
	} else if encoded, err := ioutil.ReadAll(body); err != nil {
		return utils.NewError(r.Decode, "unable to read response body", r, err)
	} else {
		// log.Printf("IN: %+v", string(encoded))
		if err := icalendar.Unmarshal(string(encoded), into); err != nil {
			return utils.NewError(r.Decode, "unable to decode response body", r, err)
		} else {
			return nil
		}
	}
}

// NewResponse creates a new WebDAV response object
func NewResponse(response *webdav.Response) *Response {
	return (*Response)(response)
}

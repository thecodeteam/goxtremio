package goxtremio

import (
	"regexp"

	xms "github.com/emccode/goxtremio/api/v3"
)

type Event *xms.Event

//GetEvents returns a list or a specific events filtered by severity,
//eventCode, or description
func (c *Client) GetEvents(
	severity, eventCode, descRxPatt string) ([]Event, error) {

	events, err := c.api.GetEvents(severity)
	if err != nil {
		return nil, err
	}

	var filtered []Event
	rx, _ := regexp.Compile(descRxPatt)

	for _, e := range events.Events {
		if (eventCode == "" || e.EventCode == eventCode) &&
			(e.Description == "" || rx.MatchString(e.Description)) {
			filtered = append(filtered, e)
		}
	}
	return filtered, nil
}

package goxtremio

import (
	"regexp"

	xmsv3 "github.com/emccode/goxtremio/api/v3"
)

func getEvents(severity string) ([]*xmsv3.Event, error) {
	events, err := xms.GetEvents(severity)
	if err != nil {
		return nil, err
	}
	return events.Events, nil
}

//GetEvents returns a list or a specific events filtered by severity, eventcode, or description
func GetEvents(severity string, eventCode string, descriptionRegex string) ([]*xmsv3.Event, error) {
	events, err := getEvents(severity)
	if err != nil {
		return nil, err
	}

	var eventsFiltered []*xmsv3.Event
	r, _ := regexp.Compile(descriptionRegex)

	for _, event := range events {
		if (eventCode == "" || event.EventCode == eventCode) && (event.Description == "" || r.MatchString(event.Description)) {
			eventsFiltered = append(eventsFiltered, event)
		}
	}
	return eventsFiltered, nil
}

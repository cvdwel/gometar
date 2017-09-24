package parser

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/eugecm/gometar/wind"
)

var groupRegexps = []string{
	`(?P<source>VRB|[0-9]{3})`,
	`(?P<gust>[0-9]{2}G)?`,
	`(?P<speed>[0-9]{2})`,
	`(?P<unit>KT|MPS)`,
	`( `,
	`(?P<varifrom>[0-9]{3})V(?P<varito>[0-9]{3})`,
	`)?`,
}

// WParser is an implementation of WindParser
type WParser struct {
	groupRegexp *regexp.Regexp
}

// New creates an instance of WParser
func New() wind.WindParser {
	groupRegexpString := strings.Join(groupRegexps, "")
	groupRegexp := regexp.MustCompile(groupRegexpString)

	return &WParser{groupRegexp: groupRegexp}
}

// Parse takes a string representing the Wind component of a METAR report
// and builds a corresponding wind.Group (or an error if the wind information
// could not be parsed
func (w *WParser) Parse(input string) (wind.Group, error) {
	matches := w.groupRegexp.FindStringSubmatch(input)

	// get variable component
	variable := false
	if matches[1] == "VRB" {
		variable = true
	}

	// get source component
	source := 0
	if !variable { // source is only given in non-variable reports
		sourceInt, err := strconv.Atoi(matches[1])
		if err != nil {
			return wind.Group{}, err
		}
		source = sourceInt
	}

	return wind.Group{
		Variable: variable,
		Source:   source,
	}, nil
}
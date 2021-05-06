package lada

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const idPattern = `[_a-z][_a-z0-9]*`
const placeHolderPattern = `\s*` +
	`(?:(?P<name>` + idPattern + `))` +
	`(?P<filters>(\s*\\|\s*(` + idPattern + `))*)?` +
	`(?P<invalid>.*?)` +
`\s*`

var charactersToEscape = []string{"[", "]", "{", "}", "*", "+", "?", "|", "^", "$", ".", "\\"}

func escapeSequence(sequence string) string  {
	escaped := sequence
	for _, ch := range charactersToEscape {
		strings.Replace(escaped, ch, "\\" + ch, -1)
	}
	return escaped
}

type TemplateFilters map[string]func(string) string

type TemplateString string

func (t TemplateString) SubstituteWithTag(items map[string]string, open string, close string) (string, error) {
	return t.SubstituteWithTagAndFilters(items, open, close, make(TemplateFilters))
}

func (t TemplateString) Substitute(items map[string]string) (string, error) {
	return t.SubstituteWithTagAndFilters(items, "{", "}", make(TemplateFilters))
}

func (t TemplateString) SubstituteWithFilters(items map[string]string, filters TemplateFilters) (string, error) {
	return t.SubstituteWithTagAndFilters(items, "{", "}", filters)
}

func (t TemplateString) SubstituteWithTagAndFilters(
	items map[string]string,
	open string,
	close string,
	filters TemplateFilters,
) (string, error) {

	var err error = nil

	result := string(t)
	// match escapes and save them
	result = strings.Replace(result, "`" + open + "`", "&#open;", -1)
	result = strings.Replace(result, "`" + close + "`", "&#close;", -1)
	pattern := regexp.MustCompile("(?m)" +
		escapeSequence(open) +
		placeHolderPattern +
		escapeSequence(close),
	)

	groupNames := pattern.SubexpNames()
	print(groupNames)
	for _, match := range pattern.FindAllStringSubmatch(result, -1) {
		stringToReplace := match[0]
		parameterName := match[1]
		rawFilters := match[2]
		invalid := match[5]
		if invalid != "" {
			err = errors.New(fmt.Sprintf("invalid expression `%s` in `%s`", invalid, stringToReplace))
		}

		if rawFilters != "" {
			filters := strings.Split(rawFilters, "|")
			for _, filter := range filters {

			}
		}

		if value, ok := items[parameterName]; ok {
			result = strings.Replace(result, stringToReplace, value, -1)
		} else {
			result = strings.Replace(result, stringToReplace, "", -1)
		}
	}

	// revert back escapes
	result = strings.Replace(result, "&#open;", open, -1)
	result = strings.Replace(result, "&#close;", close, -1)
	return result, err
}
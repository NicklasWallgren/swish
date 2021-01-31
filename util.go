package swish

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
)

func getGroupsFromRegExp(subject string, expression string) map[string]string {
	compRegEx := regexp.MustCompile(expression)
	match := compRegEx.FindStringSubmatch(subject)

	paramsMap := make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}

	return paramsMap
}

func getIDFromLocation(location string) (string, error) {
	expression := "(paymentrequests|refunds)/(?P<ID>.*)"

	groups := getGroupsFromRegExp(location, expression)

	if value, ok := groups["ID"]; ok {
		return value, nil
	}

	return "", fmt.Errorf("could not derive order id from location header. Location: %s", location)
}

func tryReadCloserToString(readCloser io.ReadCloser) string {
	buf := new(bytes.Buffer)

	if _, err := buf.ReadFrom(readCloser); err != nil {
		return ""
	}

	return buf.String()
}

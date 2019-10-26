package swish

import (
	"fmt"
	"net/http"
	"regexp"
)

func isValidHttpResponse(statusCode int, httpStatusCodes []int) bool {
	for _, validStatusCode := range httpStatusCodes {
		if statusCode == validStatusCode {
			return true
		}
	}
	return false
}

func isHttpStatusCodeWithinRange(statusCode int, statusCodeRange statusCodeRange) bool {
	return statusCode >= statusCodeRange.start && statusCode <= statusCodeRange.end
}

func getHttpHeaderValue(header string, response *http.Response) string {
	return response.Header.Get(header)
}

func getGroupsFromRegExp(subject string, expression string) map[string]string {
	var compRegEx = regexp.MustCompile(expression)
	match := compRegEx.FindStringSubmatch(subject)

	paramsMap := make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}

	return paramsMap
}

func getIdFromLocation(location string) (string, error) {
	expression := "(paymentrequests|refunds)/(?P<Id>.*)"

	groups := getGroupsFromRegExp(location, expression)

	if value, ok := groups["Id"]; ok {
		return value, nil
	}

	return "", fmt.Errorf("could not derive order id from location header. Location: %s", location)
}

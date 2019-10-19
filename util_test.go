package swish

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestGetGroupsFromRegExp(t *testing.T) {
	subject := "https://mss.cpc.getswish.net/swish-cpcapi/api/v1/paymentrequests/98BF074EE6CA42F7BCBE175182D59659"
	expression := "paymentrequests/(?P<Id>.*)"

	result := getGroupsFromRegExp(subject, expression)
	assert.Equal(t, result["Id"], "98BF074EE6CA42F7BCBE175182D59659")
}

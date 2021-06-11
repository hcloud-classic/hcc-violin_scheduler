package http

import (
	"net/url"
	"strings"
)

func compatibleRFC3986Encode(str string) string {
	resultStr := str
	resultStr = strings.Replace(resultStr, "+", "%20", -1)
	return resultStr
}

func queryURLEncoder(queryString string) string {
	params := url.Values{
		"query_string": {queryString},
	}

	urlEncode := compatibleRFC3986Encode(params.Encode())
	urlEncode = urlEncode[len("query_string="):]

	return urlEncode
}

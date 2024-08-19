package hash

import (
	"net/url"
	"sort"
	"strings"
)

func Sign(secret string, param map[string]string) string {
	stringB := StringB(secret, param)
	return strings.ToUpper(Md5(stringB))
}

func StringB(secret string, param map[string]string) string {
	stringA := StringA(param)
	stringB := stringA + "key=" + secret
	return stringB
}

func StringA(param map[string]string) string {
	var keys []string
	delete(param, "sign")
	for k := range param {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var stringA string
	for _, k := range keys {
		stringA = stringA + k + "=" + url.QueryEscape(param[k]) + "&"
	}
	return stringA
}

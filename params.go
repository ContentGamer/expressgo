package expressgo

import "regexp"

func parsePathParams(urlTemplate string, value string) map[string]string {
	result := make(map[string]string)

	regexPattern := regexp.MustCompile(`:([a-zA-Z0-9]+)`).ReplaceAllString(urlTemplate, `(?P<$1>[a-zA-Z0-9]+)`)
	re := regexp.MustCompile("^" + regexPattern + "$")

	match := re.FindStringSubmatch(value)
	if match == nil {
		return result
	}

	names := re.SubexpNames()
	for i, name := range names {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	return result
}

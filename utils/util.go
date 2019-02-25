package utils

import (
	"log"
	"net/url"
)

// GetRelativeURL will return the full URL
func GetRelativeURL(href string, baseURL string) string {
	parsedBase, err := url.Parse(baseURL)
	if err != nil {
		log.Println(err)
		return ""
	}

	parsed, err := url.Parse(href)
	if err != nil {
		log.Println(err)
		return ""
	}

	var relativeURL url.URL
	relativeURL.Scheme = parsedBase.Scheme
	relativeURL.Host = parsedBase.Host
	relativeURL.Path = parsed.Path

	return relativeURL.String()
}

// IsExternalDomain will return a boolean indicating whether the passed href
// is part of an external domain
func IsExternalDomain(href string) bool {
	parsed, err := url.Parse(href)
	if err != nil {
		log.Println(err)
		return false
	}

	if parsed.Scheme == "" && parsed.Host == "" {
		return false
	}

	return true
}

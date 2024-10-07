package main

import (
	"net/url"
	"strings"
)

/*

Types of Registrant current

1) Github
2) Artifact Hub
3) Meshery

*/

func FindRegistrant(source_url *url.URL) string {
	var registrant string
	protocol := source_url.Scheme
	domain := source_url.Host
	urlPath := source_url.Path

	if protocol == "git" {
		if isMeshery(urlPath) {
			registrant = "meshery"
		} else {
			registrant = "GitHub"
		}
	} else {
		if domain == "github.com" && !strings.HasSuffix(strings.ToLower(urlPath), ".tgz") {
			if isMeshery(urlPath) {
				registrant = "meshery"
			} else {
				registrant = "GitHub"
			}
		} else {
			registrant = "Artifact Hub"
		}
	}

	return registrant
}

func isMeshery(urlPath string) bool {
	segments := strings.Split(strings.Trim(urlPath, "/"), "/")

	for _, segment := range segments {
		if strings.EqualFold(segment, "meshery") {
			return true
		}
	}

	return false
}

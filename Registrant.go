package main

import (
	"net/url"
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
	path := source_url.Path
	//query := source_url.RawQuery

	if protocol == "git" {
		if isMeshery(path) {
			registrant = "meshery"
		} else {
			registrant = "GitHub"
		}
	} else {
		if domain == "github.com" && checktgzformat(path) == false {
			if isMeshery(path) {
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

func checktgzformat(s string) bool {
	if len(s) <= 3 {
		return false
	}
	if s[len(s)-3:] == "tgz" {
		return true
	}
	return false
}

func isMeshery(s string) bool {
	if len(s) <= 8 {
		return false
	}
	if s[:8] == "/meshery" {
		return true
	}
	return false
}

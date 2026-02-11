package main

import "strings"

func cleanInput(text string) []string {
	var s []string
	text = strings.TrimSpace(text)
	lower := strings.ToLower(text)
	s = strings.Fields(lower)
	return s
}

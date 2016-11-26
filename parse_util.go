// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "regexp"

// ------------ SCRIPT TAGS

func startOfScriptTag(line string) bool {
	re := `<script(.+(<\/script>))?`
	match := regexp.MustCompile(re).FindStringSubmatch(line)

	return match != nil && len(match) == 3 && match[2] != "</script>"
}

func endOfScriptTag(line string) bool {
	re := `((<script).+)?<\/script>`
	match := regexp.MustCompile(re).FindStringSubmatch(line)

	return match != nil && len(match) == 3 && match[2] != "<script"
}

// ------------ LEFT BRACKET
func isLeftBracket(line string) bool {
	re := `(\{)(.*(\}))?`
	matches := regexp.MustCompile(re).FindAllStringSubmatch(line, -1)

	if matches == nil {
		return false
	}

	match := matches[len(matches)-1]

	return match != nil && len(match) == 4 && match[3] != "}"
}

func parseLeftBracket(line string) string {
	re := `(.+)?\{(.+)?`
	match := regexp.MustCompile(re).FindStringSubmatch(line)

	return match[1] + "{ldelim}" + match[2]
}

// ------------ RIGHT BRACKET
func isRightBracket(line string) bool {
	re := `((\{).+)?\}`
	match := regexp.MustCompile(re).FindStringSubmatch(line)

	return match != nil && len(match) == 3 && match[2] != "{"
}

func parseRightBracket(line string) string {
	re := `(.+[^delim])?\}(.+)?`
	match := regexp.MustCompile(re).FindStringSubmatch(line)

	return match[1] + "{rdelim}" + match[2]
}

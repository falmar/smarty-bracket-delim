// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// ----------------------- BRACES
func parseBraces(inputFile io.Reader, outputFile io.Writer) error {
	reader := bufio.NewReaderSize(inputFile, 1024)
	writer := bufio.NewWriterSize(outputFile, 1024)

	var insideScriptTag bool
	var insideLiteralTag bool
	var insidePHPTag bool
	var insideMultilineComment bool
	var cm []string
	var mlm []string

	parse := func(line string, lf bool) string {
		if line == "" {
			return line
		}

		var anyMatched bool
		var matched bool

		line, matched = parseInlineObject(line)
		anyMatched = anyMatched || matched

		line, matched = parseLeftBrace(line)
		anyMatched = anyMatched || matched

		line, matched = parseRightBrace(line)
		anyMatched = anyMatched || matched

		if anyMatched && lf {
			line += "\n"
		}

		return line
	}

	assembleFragments := func(l string, fragmets []string) string {
		for i, v := range fragmets {
			l = strings.Replace(l, "[FCT-"+strconv.Itoa(i)+"]", v, 1)
		}

		return l
	}

	for {
		var isCommentLine bool
		var firstML bool
		var comment string
		var leftComment string
		var rightComment string

		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if !insideScriptTag {
			insideScriptTag = startOfScriptTag(line)
		}

		if !insideScriptTag {
			writer.WriteString(line)
			continue
		}

		line, fragments := parseCommentFragmets(line)

		if !insideMultilineComment {
			mlm, insideMultilineComment = parseMultilineCommentStart(line, false)

			if !insideMultilineComment {
				mlm, insideMultilineComment = parseMultilineCommentStart(line, true)
			}

			if insideMultilineComment {
				firstML = true
				line = mlm[0]
				rightComment = mlm[1] + "\n"
			}
		}

		if insideMultilineComment && !firstML {
			var endMultiline bool

			mlm, endMultiline = parseMultilineCommentEnd(line, false)

			if !endMultiline {
				mlm, endMultiline = parseMultilineCommentEnd(line, true)
			}

			if endMultiline {
				insideMultilineComment = false
				leftComment = mlm[0]
				line = mlm[1] + "\n"
			} else {
				writer.WriteString(line)
				continue
			}
		}

		if !insideMultilineComment {
			cm, isCommentLine = parseCommentLine(line)

			if isCommentLine {
				line = cm[0]
				comment = cm[1] + "\n"
			}
		}

		if !insidePHPTag {
			insidePHPTag = startOfPHPTag(line)
		}

		if insidePHPTag {
			insidePHPTag = !endOfPHPTag(line)
			writer.WriteString(assembleFragments(leftComment+line+comment+rightComment, fragments))
			continue
		}

		if !insideLiteralTag {
			insideLiteralTag = startOfLiteralTag(line)
		}

		if insideLiteralTag {
			insideLiteralTag = !endOfLiteralTag(line)
			writer.WriteString(assembleFragments(leftComment+line+comment+rightComment, fragments))
			continue
		}

		if isRegExp(line) {
			slice, _ := parseRegExp(line)

			line = parse(slice[0], false) + slice[1] + parse(slice[2], false)

			if !isCommentLine && !firstML {
				line += "\n"
			}
		} else {
			line = parse(line, !isCommentLine && !firstML)
		}

		if insideScriptTag {
			insideScriptTag = !endOfScriptTag(line)
		}

		writer.WriteString(assembleFragments(leftComment+line+comment+rightComment, fragments))
	}

	writer.Flush()

	return nil
}

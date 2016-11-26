// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "testing"

// ------------ SCRIPT TAGS

var openScriptTags = []string{
	`<script src="js/pace.min.js">`,
	`<script src="app.min.js" another="weird-tag">`,
	`  yo-!    <script no-tag>       // random comment`,
}

var closeScriptTags = []string{
	`anything </script> //  comment`,
	`</script>`,
	`  yo-!    </script>       // random comment`,
}

var nonScriptTags = []string{
	`<!doctype html>`,
	`<html class="no-js" lang="en">`,
	`<head>`,
	`<meta charset="utf-8" />`,
	`<meta http-equiv="x-ua-compatible" content="ie=edge">`,
	`<meta name="viewport" content="width=device-width, initial-scale=1.0" />`,
	`<title>React-App</title>`,
	`<link rel="stylesheet" href="/css/pace.css">`,
	`<link rel="stylesheet" href="/css/app.min.css">`,
	`</head>`,
	`<body>`,
	`<div id="app"></div>`,
	`<script src="js/pace.min.js"></script>`,
	`<script src="app.min.js"></script>`,
	`<script src="app.min.js"> function() {}</script>`,
	`</body>`,
	`</html>`,
}

func TestStartScriptTags(t *testing.T) {
	for _, line := range openScriptTags {
		if !startOfScriptTag(line) {
			t.Fatal("Should be script tag")
		}
	}
}

func TestStartScriptTagNonTags(t *testing.T) {
	for _, line := range nonScriptTags {
		if startOfScriptTag(line) {
			t.Fatal("Should not be script tag")
		}
	}
}

func TestEndScriptTags(t *testing.T) {
	for _, line := range closeScriptTags {
		if !endOfScriptTag(line) {
			t.Fatal("Should be script tag")
		}
	}
}

func TestEndScriptTagNonTags(t *testing.T) {
	for _, line := range nonScriptTags {
		if endOfScriptTag(line) {
			t.Fatal("Should not be script tag")
		}
	}
}

// ------------ LEFT BRACKET

var leftBracket = []string{
	"function () {",
	"call({",
	"myObject: { // random comment",
	"let array = [{",
	"{",
	"calling({my: {",
}

var expLeftBracket = []string{
	"function () {ldelim}",
	"call({ldelim}",
	"myObject: {ldelim} // random comment",
	"let array = [{ldelim}",
	"{ldelim}",
	"calling({ldelim}my: {ldelim}",
}

var nonLeftBracket = []string{
	`<body>`,
	`{$some_variable}`,
	`Outside the script tag may be pure html or may not`,
	`<script type="text/javascript">`,
	`let myVar = {json_decode($jsonVariable)}`,
	`let myOtherVar = '{$wuuuu}'`,
	`console.log({include file=$myCustomFile})`,
	`const myObject = {hello: "world", myObject:{one: 1, two: [2, 2]}}`,
	`}`,
	`hello: "world"`,
	`world: "hello"`,
	`})`,
	`hello: "world",`,
	`one: 1,`,
	`two: [2, 2]`,
	`}`,
	`}]`,
	`</script>`,
	`</body>`,
}

func TestIsLeftBracket(t *testing.T) {
	for _, line := range leftBracket {
		if !isLeftBracket(line) {
			t.Fatalf("Should be left bracket: %s", line)
		}
	}
}

func TestIsNotLeftBracket(t *testing.T) {
	for _, line := range nonLeftBracket {
		if isLeftBracket(line) {
			t.Fatalf("Should not be left bracket: %s", line)
		}
	}
}

func TestParseLeftBracket(t *testing.T) {
	for i, line := range leftBracket {
		nl, matched := parseLeftBracket(line)

		if !matched {
			t.Fatalf("Should parse")
		}

		if nl != expLeftBracket[i] {
			t.Fatalf("Expected bracket parsed: %s; got: %s", expLeftBracket[i], nl)
		}
	}
}

// ------------ RIGHT BRACKET

var rightBracket = []string{
	`everthing }`,
	`}, { should`,
	`be }) good`,
	`or }] not?`,
	`... maybe }] }?`,
	`}, {ldelim}`,
}

var expRightBracket = []string{
	`everthing {rdelim}`,
	`{rdelim}, { should`,
	`be {rdelim}) good`,
	`or {rdelim}] not?`,
	`... maybe {rdelim}] {rdelim}?`,
	`{rdelim}, {ldelim}`,
}

var nonRightBracket = []string{
	`<body>`,
	`{$some_variable}`,
	`Outside the script tag may be pure html or may not`,
	`<script type="text/javascript">`,
	`let myVar = {json_decode($jsonVariable)}`,
	`let myOtherVar = '{$wuuuu}'`,
	`console.log({include file=$myCustomFile})`,
	`funcion () {`,
	`const myObject = {hello: "world", myObject:{one: 1, two: [2, 2]}}`,
	`call({`,
	`hello: "world"`,
	`, {`,
	`world: "hello"`,
	`)`,
	`let array = [{`,
	`hello: "world",`,
	`myObject:{`,
	`one: 1,`,
	`two: [2, 2] `,
	`]`,
	`inline_call({hello: "world", myObject:{one: 1, two: [2, 2]}})`,
	`</script>`,
	`</body>`,
}

func TestIsRightBracket(t *testing.T) {
	for _, line := range rightBracket {
		if !isRightBracket(line) {
			t.Fatalf("Should be left bracket: %s", line)
		}
	}
}

func TestIsNotRightBracket(t *testing.T) {
	for _, line := range nonRightBracket {
		if isRightBracket(line) {
			t.Fatalf("Should not be left bracket: %s", line)
		}
	}
}

func TestParseRightBracket(t *testing.T) {
	for i, line := range rightBracket {
		nl, matched := parseRightBracket(line)

		if !matched {
			t.Fatalf("Should parse")
		}

		if nl != expRightBracket[i] {
			t.Fatalf("Expected bracket parsed: %s; got: %s", expRightBracket[i], nl)
		}
	}
}

// ------------ INLINE OBJECT

var inlineObject = []string{
	`var car = {type:"Fiat", model:"500", color:"white"};`,
	`var person = {firstName: '{$firstName}', lastName: "Doe", age: {json_decode($age)}, eyeColor: {$var}};`,
	`const myObject = {hello: "world", myObject:{one: 1, two: [2, 2]}}`,
	`inline_call({hello: "world", myObject:{one: 1, two: [2, 2]}})`,
}

var expInlineObject = []string{
	`var car = {ldelim}type:"Fiat", model:"500", color:"white"{rdelim};`,
	`var person = {ldelim}firstName: '{$firstName}', lastName: "Doe", age: {json_decode($age)}, eyeColor: {$var}{rdelim};`,
	`const myObject = {ldelim}hello: "world", myObject:{ldelim}one: 1, two: [2, 2]{rdelim}{rdelim}`,
	`inline_call({ldelim}hello: "world", myObject:{ldelim}one: 1, two: [2, 2]{rdelim}{rdelim})`,
}

var nonInlineObject = []string{
	`<body>`,
	`{$some_variable}`,
	`Outside the script tag may be pure html or may not`,
	`<script type="text/javascript">`,
	`let myVar = {json_decode($jsonVariable)}`,
	`let myOtherVar = '{$wuuuu}'`,
	`console.log({include file=$myCustomFile})`,
	`funcion () {`,
	`}`,
	`call({`,
	`hello: "world"`,
	`}, {`,
	`world: "hello"`,
	`})`,
	`let array = [{`,
	`hello: "world",`,
	`myObject:{`,
	`one: 1,`,
	`two: [2, 2]`,
	`}`,
	`}]`,
	`const strangeObject = {maybe: {it: {wont: {work: "?"`,
	`}, maybe: ""}, not: "did"}, work: "entirely"}`,
	`</script>`,
	`</body>`,
}

func TestParseInlineObjectMatch(t *testing.T) {
	for i, line := range inlineObject {
		nl, matched := parseInlineObject(line)

		if !matched {
			t.Fatalf("Should parse")
		}

		if nl != expInlineObject[i] {
			t.Fatalf("Expected bracket parsed: %s; got: %s", expInlineObject[i], nl)
		}
	}
}

func TestParseInlineObjectNoMatch(t *testing.T) {
	for _, line := range nonInlineObject {
		_, matched := parseInlineObject(line)

		if matched {
			t.Fatalf("Should not match")
		}
	}
}

// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "testing"

var commentLines = []string{
	`function () { // Copyright 2016 David Lavieri. All rights reserved.`,
	`some code // Use of this source code is governed by a MIT License`,
	`// License that can be found in the LICENSE file.`,
	`<script type="text/javascript"> // Date: 0/0/0`,
	`console.log({include// file=$myCustomFile}) // Time: 0:0 PM`,
	`</script> // @author    David Lavieri (falmar) <daviddlavier@gmail.com>`,
	`// @copyright 2016 David Lavieri`,
	`let some = 0 // @license   http://opensource.org/licenses/MIT The MIT License (MIT)`,
}

var expCommentLines = [][]string{
	[]string{`function () { `, `// Copyright 2016 David Lavieri. All rights reserved.`},
	[]string{`some code `, `// Use of this source code is governed by a MIT License`},
	[]string{``, `// License that can be found in the LICENSE file.`},
	[]string{`<script type="text/javascript"> `, `// Date: 0/0/0`},
	[]string{`console.log({include`, `// file=$myCustomFile}) // Time: 0:0 PM`},
	[]string{`</script> `, `// @author    David Lavieri (falmar) <daviddlavier@gmail.com>`},
	[]string{``, `// @copyright 2016 David Lavieri`},
	[]string{`let some = 0 `, `// @license   http://opensource.org/licenses/MIT The MIT License (MIT)`},
}

var nonCommentLines = []string{
	`<body>`,
	`{$some_variable}`,
	`Outside the script tag may be pure html or may not`,
	`<script type="text/javascript">`,
	`let myVar = {json_decode($jsonVariable)}`,
	`let myOtherVar = '{$wuuuu}'`,
	`console.log({include file=$myCustomFile})`,
	`let some = 0`,
	`{rdelim}`,
	`hello: "world"`,
	`world: "hello"`,
	`{rdelim})`,
	`hello: "world",`,
	`one: 1,`,
	`two: [2, 2]`,
	`{rdelim}`,
	`{rdelim}]`,
	`{rdelim}, maybe: ""{rdelim}, did: "not"{rdelim}, work: "entirely"{rdelim}`,
	`</script>`,
	`</body>`,
}

func TestIsCommentLineMatch(t *testing.T) {
	for _, c := range commentLines {
		if !isCommentLine(c) {
			t.Fatalf("Should be/have a line comment %s", c)
		}
	}
}

func TestIsCommentLineNoMatch(t *testing.T) {
	for _, c := range nonCommentLines {
		if isCommentLine(c) {
			t.Fatalf("Should not be/have a line comment %s", c)
		}
	}
}

func TestCommentLineMatch(t *testing.T) {
	for i, c := range commentLines {
		left, match := parseCommentLine(c)

		if !match {
			t.Fatalf("Expected match %s", c)
		}

		if left[0] != expCommentLines[i][0] {
			t.Fatalf("Expected left most parse: %s; got: %s", expCommentLines[i][0], left[0])
		}

		if left[1] != expCommentLines[i][1] {
			t.Fatalf("Expected right most parse: %s; got: %s", expCommentLines[i][1], left[1])
		}
	}
}

func TestCommentLineNoMatch(t *testing.T) {
	for _, c := range nonCommentLines {
		left, match := parseCommentLine(c)

		if match {
			t.Fatal("Expected no match")
		}

		if left[1] != c {
			t.Fatalf("Expected left most be intact: %s; got: %s", c, left[1])
		}
	}
}

// ------------- Multiline Single
var singleMultilineCommentStart = []string{
	`/** is comment`,
	`function (){} /**`,
	`function (){} /**  is comment`,
}

var singleMultilineCommentEnd = []string{
	`*/ is comment`,
	`function (){} */`,
	`function (){} */  is comment`,
}

var singleMultilineComment = []string{
	`function () { /** Copyright 2016 David Lavieri. /** All rights reserved.`,
	`some code /** Use of this source code is governed by a MIT License`,
	`/** License that can be found in the LICENSE file.`,
	`<script type="text/javascript"> /** Date: 0/0/0`,
	`console.log({include/** file=$myCustomFile}) /** Time: 0:0 PM`,
	`</script> /** @author    David Lavieri (falmar) <daviddlavier@gmail.com>`,
	`/** @copyright 2016 David Lavieri`,
	`let some = 0 /** @license   http://opensource.org/licenses/MIT The MIT License (MIT)`,
}

var expSingleMultilineComment = [][]string{
	[]string{`function () { `, `/** Copyright 2016 David Lavieri. /** All rights reserved.`},
	[]string{`some code `, `/** Use of this source code is governed by a MIT License`},
	[]string{``, `/** License that can be found in the LICENSE file.`},
	[]string{`<script type="text/javascript"> `, `/** Date: 0/0/0`},
	[]string{`console.log({include`, `/** file=$myCustomFile}) /** Time: 0:0 PM`},
	[]string{`</script> `, `/** @author    David Lavieri (falmar) <daviddlavier@gmail.com>`},
	[]string{``, `/** @copyright 2016 David Lavieri`},
	[]string{`let some = 0 `, `/** @license   http://opensource.org/licenses/MIT The MIT License (MIT)`},
}

// ------------- Multiline Smarty
var smartyMultilineCommentStart = []string{
	`{* is comment`,
	`function (){} {*`,
	`function (){} {*  is comment`,
}

var smartyMultilineCommentEnd = []string{
	`*} is comment`,
	`function (){} *}`,
	`function (){} *}  is comment`,
}

var smartyMultilineComment = []string{
	`function () { {* Copyright 2016 David Lavieri. All rights reserved.`,
	`some code {* Use of this source code is governed by a MIT License`,
	`{* License that can be found in the LICENSE file.`,
	`<script type="text/javascript"> {* Date: 0/0/0`,
	`console.log({include{* file=$myCustomFile}) {* Time: 0:0 PM`,
	`</script> {* @author    David Lavieri (falmar) <daviddlavier@gmail.com>`,
	`{* @copyright 2016 David Lavieri`,
	`let some = 0 {* @license   http://opensource.org/licenses/MIT The MIT License (MIT)`,
}

var expSmartyMultilineComment = [][]string{
	[]string{`function () { `, `{* Copyright 2016 David Lavieri. All rights reserved.`},
	[]string{`some code `, `{* Use of this source code is governed by a MIT License`},
	[]string{``, `{* License that can be found in the LICENSE file.`},
	[]string{`<script type="text/javascript"> `, `{* Date: 0/0/0`},
	[]string{`console.log({include`, `{* file=$myCustomFile}) {* Time: 0:0 PM`},
	[]string{`</script> `, `{* @author    David Lavieri (falmar) <daviddlavier@gmail.com>`},
	[]string{``, `{* @copyright 2016 David Lavieri`},
	[]string{`let some = 0 `, `{* @license   http://opensource.org/licenses/MIT The MIT License (MIT)`},
}

var nonMultilineCommentStart = []string{
	`/* /*_* /*\s* /*-* /*text* is comment`,
	`function (){} /* /*_* /*\s* /*-* /*text*`,
	`{ * {_* {\s* {-* is comment *{`,
	`function (){} { * {_* {\s* {-*`,
	`(.*)(\{)(.+(}))?(.*) ------------ (.+)?({)$ ldelim | $1{ldelim}`,
	`(.*)({)(.*)(})(.*) ----- (\s+)(}) rdelim | $1{rdelim}`,
	`({)([^ldelim\n]\w+:\s?\"?\'?.+\"?\'?[^ldelim\n])(}) inline object | {ldelim}$2{rdelim}`,
	`rc (.*)(({).*)({)(.*)(})(.*)`,
	`rc1 (.*)(({).*)(\})(.*)`,
	`r2 (.*[\=\(\[\{]\s?({))(\w+\s*.+:\s*.+)(})`,
	`r1 regex (.*)(['")])?[^\/'")]?\/(.*)[^\/'")]?\/(['")])?(.*)`,
	`r2 regex (.*?[^\/]?)\/(.+)\/(.*)`,
	`commen ml start (.*)\/\*\*(.*) end (.*)\*\/(.*)`,
	`commen ml start (.*)\{\*(.*) end (.*)\*\}(.*) smarty`,
	`line comment (.*)(\/\/.*) [1]check [2]append`,
	`partial comment (.*)(\/\*.*\*\/)(.*) (.*)\/\*\*(.*(\*\*\/))(.*)`,
	`// @author    David Lavieri (falmar) <daviddlavier@gmail.com>`,
	`// @copyright 2016 David Lavieri`,
	`// @license   http://opensource.org/licenses/MIT The MIT License (MIT)`,
}

// ------------- Multiline Single
func TestIsMultilineCommentStartSingleMatch(t *testing.T) {
	for _, l := range singleMultilineCommentStart {
		if !isMultilineCommentStart(l, true) {
			t.Fatalf("Should match start multiline single comment %s", l)
		}
	}
}

func TestIsMultilineCommentStartSingleNoMatch(t *testing.T) {
	for _, l := range nonMultilineCommentStart {
		if isMultilineCommentStart(l, true) {
			t.Fatalf("Should not match start multiline single comment %s", l)
		}
	}
}

func TestIsMultilineCommentEndSingleMatch(t *testing.T) {
	for _, l := range singleMultilineCommentEnd {
		if !isMultilineCommentEnd(l, true) {
			t.Fatalf("Should match end multiline single comment %s", l)
		}
	}
}

func TestIsMultilineCommentEndSingleNoMatch(t *testing.T) {
	for _, l := range nonMultilineCommentStart {
		if isMultilineCommentStart(l, true) {
			t.Fatalf("Should not match end multiline single comment %s", l)
		}
	}
}

func TestParseMultilineCommentSingleStartMatch(t *testing.T) {
	for i, c := range singleMultilineComment {
		left, match := parseMultilineCommentStart(c, true)

		if !match {
			t.Fatalf("Expected match %s", c)
		}

		if left[0] != expSingleMultilineComment[i][0] {
			t.Fatalf("Expected left most parse: %s; got: %s", expSingleMultilineComment[i][0], left[0])
		}

		if left[1] != expSingleMultilineComment[i][1] {
			t.Fatalf("Expected right most parse: %s; got: %s", expSingleMultilineComment[i][1], left[1])
		}
	}
}

func TestParseMultilineCommentSingleStartNoMatch(t *testing.T) {
	for _, c := range nonCommentLines {
		left, match := parseMultilineCommentStart(c, true)

		if match {
			t.Fatal("Expected no match")
		}

		if left[1] != c {
			t.Fatalf("Expected left most be intact: %s; got: %s", c, left[1])
		}
	}
}

// ------------- Multiline Smarty
func TestIsMultilineCommentStartSmartyMatch(t *testing.T) {
	for _, l := range smartyMultilineCommentStart {
		if !isMultilineCommentStart(l, false) {
			t.Fatalf("Should match start multiline single comment %s", l)
		}
	}
}

func TestIsMultilineCommentStartSmartyNoMatch(t *testing.T) {
	for _, l := range nonMultilineCommentStart {
		if isMultilineCommentEnd(l, false) {
			t.Fatalf("Should not match start multiline single comment %s", l)
		}
	}
}

func TestIsMultilineCommentEndSmartyMatch(t *testing.T) {
	for _, l := range smartyMultilineCommentEnd {
		if !isMultilineCommentEnd(l, false) {
			t.Fatalf("Should match end multiline single comment %s", l)
		}
	}
}

func TestIsMultilineCommentEndSmartyNoMatch(t *testing.T) {
	for _, l := range nonMultilineCommentStart {
		if isMultilineCommentEnd(l, false) {
			t.Fatalf("Should not match end multiline single comment %s", l)
		}
	}
}

func TestParseMultilineCommentSmartyStartMatch(t *testing.T) {
	for i, c := range smartyMultilineComment {
		left, match := parseMultilineCommentStart(c, false)

		if !match {
			t.Fatalf("Expected match %s", c)
		}

		if left[0] != expSmartyMultilineComment[i][0] {
			t.Fatalf("Expected left most parse: %s; got: %s", expSmartyMultilineComment[i][0], left[0])
		}

		if left[1] != expSmartyMultilineComment[i][1] {
			t.Fatalf("Expected right most parse: %s; got: %s", expSmartyMultilineComment[i][1], left[1])
		}
	}
}

func TestParseMultilineCommentSmartyStartNoMatch(t *testing.T) {
	for _, c := range nonCommentLines {
		left, match := parseMultilineCommentStart(c, false)

		if match {
			t.Fatal("Expected no match")
		}

		if left[1] != c {
			t.Fatalf("Expected left most be intact: %s; got: %s", c, left[1])
		}
	}
}

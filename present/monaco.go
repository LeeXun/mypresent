// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package present

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
	"regexp"
	"strings"
)

const monacoName = "monaco"

func init() {
	Register(monacoName, parseMonaco)
}

type Monaco struct {
	Text     template.HTML
	Play     bool   // runnable code
	Edit     bool   // editable code
	FileName string // file name
	Ext      string // file extension
	Raw      string // content of the file
}

type monacoTemplateData struct {
	Lines          []codeLine
	Prefix, Suffix []byte
	Edit, Numbers  bool
	Raw            string
}

type monacoLine struct {
	L  string // The line of code.
	N  int    // The line number from the source file.
	HL bool   // Whether the line should be highlighted.
}

func (m Monaco) TemplateName() string { return monacoName }

// The input line is a .code or .play entry with a file name and an optional HLfoo marker on the end.
// Anything between the file and HL (if any) is an address expression, which we treat as a string here.
// We pick off the HL first, for easy parsing.
var (
	mcodeRE = regexp.MustCompile(`\.(monaco)\s+((?:(?:-edit|-numbers)\s+)*)([^\s]+)(?:\s+(.*))?$`)
)

// parseCode parses a code present directive. Its syntax:
//   .code [-numbers] [-edit] <filename> [address] [highlight]
// The directive may also be ".play" if the snippet is executable.
func parseMonaco(ctx *Context, sourceFile string, sourceLine int, cmd string) (Elem, error) {
	cmd = strings.TrimSpace(cmd)

	// Parse the remaining command line.
	// Arguments:
	// args[0]: whole match
	// args[1]: .monaco
	// args[2]: flags ("-edit -numbers")
	// args[3]: file name
	// args[4]: optional address
	args := mcodeRE.FindStringSubmatch(cmd)
	if len(args) != 5 {
		return nil, fmt.Errorf("%s:%d: syntax error for .%s invocation", sourceFile, monacoName, sourceLine)
	}
	_, _, file, addr := args[1], args[2], args[3], strings.TrimSpace(args[4])

	// Read in code file and (optionally) match address.
	filename := filepath.Join(filepath.Dir(sourceFile), file)
	textBytes, err := ctx.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("%s:%d: %v", sourceFile, sourceLine, err)
	}
	lo, hi, err := addrToByteRange(addr, 0, textBytes)
	if err != nil {
		return nil, fmt.Errorf("%s:%d: %v", sourceFile, sourceLine, err)
	}
	if lo > hi {
		// The search in addrToByteRange can wrap around so we might
		// end up with the range ending before its starting point
		hi, lo = lo, hi
	}

	// Acme pattern matches can stop mid-line,
	// so run to end of line in both directions if not at line start/end.
	for lo > 0 && textBytes[lo-1] != '\n' {
		lo--
	}
	if hi > 0 {
		for hi < len(textBytes) && textBytes[hi-1] != '\n' {
			hi++
		}
	}

	lines := monacoLines(textBytes, lo, hi)

	// data := &monacoTemplateData{
	// 	Lines:   mformatLines(lines, highlight),
	// 	Edit:    strings.Contains(flags, "-edit"),
	// 	Numbers: strings.Contains(flags, "-numbers"),
	// 	Raw:     string(mRawCode(lines)),
	// }

	// Include before and after in a hidden span for playground code.
	// if play {
	// 	data.Prefix = textBytes[:lo]
	// 	data.Suffix = textBytes[hi:]
	// }

	// var buf bytes.Buffer
	// if err := monacoTemplate.Execute(&buf, data); err != nil {
	// 	return nil, err
	// }
	return Monaco{
		// Text: template.HTML(buf.String()),
		// Play:     play,
		// Edit:     data.Edit,
		// FileName: filepath.Base(filename),
		// Ext:      filepath.Ext(filename),
		Raw: string(mRawCode(lines)),
	}, nil
}

// mformatLines returns a new slice of codeLine with the given lines
// replacing tabs with spaces and adding highlighting where needed.
// func mformatLines(lines []codeLine, highlight string) []codeLine {
// 	formatted := make([]codeLine, len(lines))
// 	for i, line := range lines {
// 		// Replace tabs with spaces, which work better in HTML.
// 		line.L = strings.Replace(line.L, "\t", "    ", -1)

// 		// Highlight lines that end with "// HL[highlight]"
// 		// and strip the magic comment.
// 		if m := mhlCommentRE.FindStringSubmatch(line.L); m != nil {
// 			line.L = m[1]
// 			line.HL = m[2] == highlight
// 		}

// 		formatted[i] = line
// 	}
// 	return formatted
// }

// mRawCode returns the code represented by the given monacoLines without any kind
// of formatting.
func mRawCode(lines []codeLine) []byte {
	b := new(bytes.Buffer)
	for _, line := range lines {
		b.WriteString(line.L)
		b.WriteByte('\n')
	}
	return b.Bytes()
}

// codeLine represents a line of code extracted from a source file.

// monacoLines takes a source file and returns the lines that
// span the byte range specified by start and end.
// It discards lines that end in "OMIT".
func monacoLines(src []byte, start, end int) (lines []codeLine) {
	startLine := 1
	for i, b := range src {
		if i == start {
			break
		}
		if b == '\n' {
			startLine++
		}
	}
	s := bufio.NewScanner(bytes.NewReader(src[start:end]))
	for n := startLine; s.Scan(); n++ {
		l := s.Text()
		if strings.HasSuffix(l, "OMIT") {
			continue
		}
		lines = append(lines, codeLine{L: l, N: n})
	}
	// Trim leading and trailing blank lines.
	for len(lines) > 0 && len(lines[0].L) == 0 {
		lines = lines[1:]
	}
	for len(lines) > 0 && len(lines[len(lines)-1].L) == 0 {
		lines = lines[:len(lines)-1]
	}
	return
}

// Package strutil provide some string,char,byte util functions
package strutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/axgle/mahonia"
	"strings"
	"text/template"
)

// Position for padding string
const (
	PosLeft uint8 = iota
	PosRight
)

var (
	Ltrim = TrimLeft
	Rtrim = TrimRight
)

/*************************************************************
 * String filtering
 *************************************************************/

// Trim string
func Trim(s string, cutSet ...string) string {
	if len(cutSet) > 0 && cutSet[0] != "" {
		return strings.Trim(s, cutSet[0])
	}

	return strings.TrimSpace(s)
}

// TrimLeft char in the string.
func TrimLeft(s string, cutSet ...string) string {
	if len(cutSet) > 0 {
		return strings.TrimLeft(s, cutSet[0])
	}

	return strings.TrimLeft(s, " ")
}

// TrimRight char in the string.
func TrimRight(s string, cutSet ...string) string {
	if len(cutSet) > 0 {
		return strings.TrimRight(s, cutSet[0])
	}

	return strings.TrimRight(s, " ")
}

// FilterEmail filter email, clear invalid chars.
func FilterEmail(s string) string {
	s = strings.TrimSpace(s)
	i := strings.LastIndex(s, "@")
	if i == -1 {
		return s
	}

	// According to rfc5321, "The local-part of a mailbox MUST BE treated as case sensitive"
	return s[0:i] + "@" + strings.ToLower(s[i+1:])
}

/*************************************************************
 * String operation
 *************************************************************/

// SplitValid string to slice. will filter empty string node.
func SplitValid(s, sep string) (ss []string) { return Split(s, sep) }

// Split string to slice. will filter empty string node.
func Split(s, sep string) (ss []string) {
	if s = strings.TrimSpace(s); s == "" {
		return
	}

	for _, val := range strings.Split(s, sep) {
		if val = strings.TrimSpace(val); val != "" {
			ss = append(ss, val)
		}
	}
	return
}

// SplitNValid string to slice. will filter empty string node.
func SplitNValid(s, sep string, n int) (ss []string) { return SplitN(s, sep, n) }

// SplitN string to slice. will filter empty string node.
func SplitN(s, sep string, n int) (ss []string) {
	if s = strings.TrimSpace(s); s == "" {
		return
	}

	rawList := strings.Split(s, sep)
	for i, val := range rawList {
		if val = strings.TrimSpace(val); val != "" {
			if len(ss) == n-1 {
				ss = append(ss, strings.TrimSpace(strings.Join(rawList[i:], sep)))
				break
			}

			ss = append(ss, val)
		}
	}
	return
}

// SplitTrimmed split string to slice.
// will trim space for each node, but not filter empty
func SplitTrimmed(s, sep string) (ss []string) {
	if s = strings.TrimSpace(s); s == "" {
		return
	}

	for _, val := range strings.Split(s, sep) {
		ss = append(ss, strings.TrimSpace(val))
	}
	return
}

// SplitNTrimmed split string to slice.
// will trim space for each node, but not filter empty
func SplitNTrimmed(s, sep string, n int) (ss []string) {
	if s = strings.TrimSpace(s); s == "" {
		return
	}

	for _, val := range strings.SplitN(s, sep, n) {
		ss = append(ss, strings.TrimSpace(val))
	}
	return
}

// Substr for a string.
func Substr(s string, pos, length int) string {
	runes := []rune(s)
	strLen := len(runes)

	// pos is to large
	if pos >= strLen {
		return ""
	}

	l := pos + length
	if l > strLen {
		l = strLen
	}

	return string(runes[pos:l])
}

// Padding a string.
func Padding(s, pad string, length int, pos uint8) string {
	diff := len(s) - length
	if diff >= 0 { // do not need padding.
		return s
	}

	if pad == "" || pad == " " {
		mark := ""
		if pos == PosRight { // to right
			mark = "-"
		}

		// padding left: "%7s", padding right: "%-7s"
		tpl := fmt.Sprintf("%s%d", mark, length)
		return fmt.Sprintf(`%`+tpl+`s`, s)
	}

	if pos == PosRight { // to right
		return s + Repeat(pad, -diff)
	}

	return Repeat(pad, -diff) + s
}

// PadLeft a string.
func PadLeft(s, pad string, length int) string {
	return Padding(s, pad, length, PosLeft)
}

// PadRight a string.
func PadRight(s, pad string, length int) string {
	return Padding(s, pad, length, PosRight)
}

// Repeat repeat a string
func Repeat(s string, times int) string {
	if times < 2 {
		return s
	}

	var ss []string
	for i := 0; i < times; i++ {
		ss = append(ss, s)
	}

	return strings.Join(ss, "")
}

// RepeatRune repeat a rune char.
func RepeatRune(char rune, times int) (chars []rune) {
	for i := 0; i < times; i++ {
		chars = append(chars, char)
	}
	return
}

// RepeatBytes repeat a byte char.
func RepeatBytes(char byte, times int) (chars []byte) {
	for i := 0; i < times; i++ {
		chars = append(chars, char)
	}
	return
}

// Replaces replace multi strings
//
// 	pairs: {old1: new1, old2: new2, ...}
//
// Can also use:
// 	strings.NewReplacer("old1", "new1", "old2", "new2").Replace(str)
func Replaces(str string, pairs map[string]string) string {
	ss := make([]string, len(pairs)*2)
	for old, newVal := range pairs {
		ss = append(ss, old, newVal)
	}

	return strings.NewReplacer(ss...).Replace(str)
}

// PrettyJSON get pretty Json string
// Deprecated
//  please use fmtutil.PrettyJSON() or jsonutil.Pretty() instead it
func PrettyJSON(v interface{}) (string, error) {
	out, err := json.MarshalIndent(v, "", "    ")
	return string(out), err
}

// RenderTemplate render text template
func RenderTemplate(input string, data interface{}, fns template.FuncMap, isFile ...bool) string {
	return RenderText(input, data, fns, isFile...)
}

// RenderText render text template
func RenderText(input string, data interface{}, fns template.FuncMap, isFile ...bool) string {
	t := template.New("simple-text")
	t.Funcs(template.FuncMap{
		// don't escape content
		"raw": func(s string) string {
			return s
		},
		"trim": func(s string) string {
			return strings.TrimSpace(string(s))
		},
		// join strings
		"join": func(ss []string, sep string) string {
			return strings.Join(ss, sep)
		},
		// lower first char
		"lcFirst": func(s string) string {
			return LowerFirst(s)
		},
		// upper first char
		"upFirst": func(s string) string {
			return UpperFirst(s)
		},
	})

	// custom add template functions
	if len(fns) > 0 {
		t.Funcs(fns)
	}

	if len(isFile) > 0 && isFile[0] {
		template.Must(t.ParseFiles(input))
	} else {
		template.Must(t.Parse(input))
	}

	// use buffer receive rendered content
	buf := new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		panic(err)
	}

	return buf.String()
}

// 按照长度切分字符串
func Chunks(s string, chunkSize int) []string {

	if len(s) == 0 {

		return nil

	}

	if chunkSize >= len(s) {

		return []string{s}

	}

	var chunks = make([]string, 0, (len(s)-1)/chunkSize+1)

	currentLen := 0

	currentStart := 0

	for i := range s {

		if currentLen == chunkSize {

			chunks = append(chunks, s[currentStart:i])

			currentLen = 0

			currentStart = i

		}

		currentLen++

	}

	chunks = append(chunks, s[currentStart:])

	return chunks

}

// 编码字符集转换
func Translate(src string, srcCode string, tagCode string) (string, error) {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, err := tagCoder.Translate([]byte(srcResult), true)
	return string(cdata), err
}

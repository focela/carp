// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by an MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"bytes"
	"strings"
)

var (
	// DefaultTrimChars are the characters that are stripped by Trim* functions by default.
	DefaultTrimChars = string([]byte{
		'\t', // Tab.
		'\v', // Vertical tab.
		'\n', // New line (line feed).
		'\r', // Carriage return.
		'\f', // New page.
		' ',  // Ordinary space.
		0x00, // NUL-byte.
		0x85, // Delete.
		0xA0, // Non-breaking space.
	})
)

// IsLetterUpper checks whether the given byte b is uppercase.
func IsLetterUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

// IsLetterLower checks whether the given byte b is lowercase.
func IsLetterLower(b byte) bool {
	return b >= 'a' && b <= 'z'
}

// IsLetter checks whether the given byte b is a letter.
func IsLetter(b byte) bool {
	return IsLetterUpper(b) || IsLetterLower(b)
}

// IsNumeric checks whether the given string s is numeric.
// Note that strings representing floating-point numbers like "123.456" are also considered numeric.
func IsNumeric(s string) bool {
	var (
		dotCount = 0
		length   = len(s)
	)
	if length == 0 {
		return false
	}
	for i := 0; i < length; i++ {
		if s[i] == '-' && i == 0 {
			continue
		}
		if s[i] == '.' {
			dotCount++
			if i > 0 && i < length-1 {
				continue
			} else {
				return false
			}
		}
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return dotCount <= 1
}

// UcFirst returns a copy of the string s with the first letter mapped to uppercase.
func UcFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

// ReplaceByMap returns a copy of `origin`,
// where substrings matching keys in the map are replaced with the corresponding values, case-sensitively.
func ReplaceByMap(origin string, replaces map[string]string) string {
	for k, v := range replaces {
		origin = strings.ReplaceAll(origin, k, v)
	}
	return origin
}

// RemoveSymbols removes all symbols from the string and leaves only numbers and letters.
func RemoveSymbols(s string) string {
	var b = make([]rune, 0, len(s))
	for _, c := range s {
		if c > 127 {
			b = append(b, c)
		} else if (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			b = append(b, c)
		}
	}
	return string(b)
}

// EqualFoldWithoutChars checks if strings `s1` and `s2` are equal, ignoring case and the characters '-', '_', '.', and ' '.
func EqualFoldWithoutChars(s1, s2 string) bool {
	return strings.EqualFold(RemoveSymbols(s1), RemoveSymbols(s2))
}

// SplitAndTrim splits string `str` by a string `delimiter` into an array,
// and trims each element of this array. It ignores elements that are empty after trimming.
func SplitAndTrim(str, delimiter string, characterMask ...string) []string {
	array := make([]string, 0)
	for _, v := range strings.Split(str, delimiter) {
		v = Trim(v, characterMask...)
		if v != "" {
			array = append(array, v)
		}
	}
	return array
}

// Trim strips whitespace (or other characters) from the beginning and end of a string.
// The optional parameter `characterMask` specifies the additional characters to strip.
func Trim(str string, characterMask ...string) string {
	trimChars := DefaultTrimChars
	if len(characterMask) > 0 {
		trimChars += characterMask[0]
	}
	return strings.Trim(str, trimChars)
}

// FormatCmdKey formats string `s` as a command key using a uniform format.
func FormatCmdKey(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, "_", "."))
}

// FormatEnvKey formats string `s` as an environment key using a uniform format.
func FormatEnvKey(s string) string {
	return strings.ToUpper(strings.ReplaceAll(s, ".", "_"))
}

// StripSlashes un-quotes a quoted string created by AddSlashes.
func StripSlashes(str string) string {
	var buf bytes.Buffer
	l, skip := len(str), false
	for i, char := range str {
		if skip {
			skip = false
		} else if char == '\\' {
			if i+1 < l && str[i+1] == '\\' {
				skip = true
			}
			continue
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

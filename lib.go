//	This file is part of Fwew.
//	Fwew is free software: you can redistribute it and/or modify
// 	it under the terms of the GNU General Public License as published by
// 	the Free Software Foundation, either version 3 of the License, or
// 	(at your option) any later version.
//
//	Fwew is distributed in the hope that it will be useful,
//	but WITHOUT ANY WARRANTY; without even implied warranty of
//	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//	GNU General Public License for more details.
//
//	You should have received a copy of the GNU General Public License
//	along with Fwew.  If not, see http://gnu.org/licenses/

// Package main contains all the things. lib.go handles common functions.
package main

import "unicode"

// ContainsStr returns true if q is in s
func ContainsStr(s []string, q string) bool {
	if len(q) == 0 || len(s) == 0 {
		return false
	}
	for _, x := range s {
		if q == x {
			return true
		}
	}
	return false
}

// DeleteElement "deletes" all occurrences of q in s
// actually returns a new string slice containing the original minus all q
func DeleteElement(s []string, q string) []string {
	if len(s) == 0 {
		return s
	}
	var r []string
	for _, str := range s {
		if str != q {
			r = append(r, str)
		}
	}
	return r
}

// DeleteEmpty "deletes" all empty string entries in s
// actually returns a new string slice containing all non-empty strings in s
func DeleteEmpty(s []string) []string {
	return DeleteElement(s, "")
}

// IndexStr return the index of q in s
func IndexStr(s string, q rune) int {
	for i, v := range s {
		if v == q {
			return i
		}
	}
	return -1
}

// IsLetter returns true if s is an alphabet character or apostrophe
func IsLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) || r == '\'' || r == '‘' {
			return true
		}
	}
	return false
}

// Reverse returns the reversed version of s
func Reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, r := range s {
		n--
		runes[n] = r
	}
	return string(runes[n:])
}

// Valid validates range of integers for input
func Valid(input int64, reverse bool) bool {
	const (
		maxIntDec int64 = 32767
		maxIntOct int64 = 77777
	)
	if reverse {
		if 0 <= input && input <= maxIntDec {
			return true
		}
		return false
	}
	if 0 <= input && input <= maxIntOct {
		return true
	}
	return false
}

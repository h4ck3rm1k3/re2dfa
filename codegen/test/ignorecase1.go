// Code generated by re2dfa (https://github.com/opennota/re2dfa).

package test

import "unicode/utf8"

func matchIgnoreCase1(s string) (end int) {
	end = -1
	var r rune
	var rlen int
	i := 0
	_, _, _ = r, rlen, i
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r == 65 || r == 97:
		goto s2
	}
	return
s2:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		return
	}
	i += rlen
	switch {
	case r == 90 || r == 122:
		end = i
	}
	return
}
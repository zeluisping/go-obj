package main

import (
	"strconv"
)

func parseFloat(v string) (val float32, rest string, ok bool) {
	if v == "" {
		ok = false
		return
	}
	i := 0
	if v[0] == '+' || v[0] == '-' {
		i++
	}
	if i == len(v) {
		ok = false
		return
	}
	if v[i] < '0' || v[i] > '9' {
		ok = false
		return
	}
	i++
	dot := false
	for i < len(v) && (v[i] >= '0' && v[i] <= '9' || v[i] == '.') {
		if v[i] == '.' {
			if dot == true {
				ok = false
				return
			}
			dot = true
		}
		i++
	}
	if dot == true && v[i-1] == '.' {
		ok = false
		return
	}
	val64, err := strconv.ParseFloat(v[:i], 32)
	if err != nil {
		ok = false
		return
	}
	val = float32(val64)
	rest = v[i:]
	ok = true
	return
}

func parseInteger(v string) (val int32, rest string, ok bool) {
	if v == "" {
		ok = false
		return
	}
	i := 0
	if v[0] == '+' || v[0] == '-' {
		i++
	}
	if i >= len(v) {
		ok = false
		return
	}
	if v[i] < '0' || v[i] > '9' {
		ok = false
		return
	}
	i++
	for i < len(v) && v[i] >= '0' && v[i] <= '9' {
		i++
	}
	val64, err := strconv.ParseInt(v[:i], 10, 32)
	if err != nil {
		ok = false
		return
	}
	val = int32(val64)
	rest = v[i:]
	ok = true
	return
}

func skipSpaces(s string) (rest string, ok bool) {
	if s == "" {
		ok = false
		return
	}
	if s[0] != ' ' && s[0] != '\t' {
		ok = false
		return
	}
	i := 1
	for i < len(s) && (s[i] == ' ' || s[i] == '\t') {
		i++
	}
	rest = s[i:]
	ok = true
	return
}

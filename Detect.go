package main

import (
	"strconv"
	"strings"
	"unicode"
)

func IsSpaceExceptNewline(r rune) bool {
	return unicode.IsSpace(r) && r != '\n'
}

func Detect(txt string) int {
	runes := []rune(txt)
	flag := ""
	end := ""
	found := false
	cp := 0
	for i := 0; i < len(runes); i++ {
		if runes[i] == '(' {
			found = true
			flag += string(runes[i])
			cp++
			continue
		}
		if found {
			if runes[i] == ')' {
				flag += string(runes[i])
				cp++
				break
			}
			flag += string(runes[i])
		}
		cp++
	}
	for j := cp; j < len(runes); j++ {
		if runes[j] == '(' {
			break
		}
		end += string(runes[j])
	}
	onlyspace := true
	for _, ch := range end {
		if !IsSpaceExceptNewline(ch) {
			onlyspace = false
			break
		}
	}
	if onlyspace {
		end = ""
	}
	runes = []rune(flag)
	if strings.HasPrefix(string(runes), "(up") {
		if strings.HasPrefix(string(runes[3:]), ")") {
			return 1
		}
		if strings.HasPrefix(string(runes[3:]), ",") {
			numstr := ""
			for _, ch := range string(runes[4 : len(runes)-1]) {
				if unicode.IsNumber(ch) {
					numstr += string(ch)
					continue
				}
				if ch != ' ' {
					return 0
				}
			}
			num, err := strconv.Atoi(numstr)
			if err != nil {
				return 0
			}
			if num > 0 {
				return num
			}
			return 0
		}
	} else if strings.HasPrefix(string(runes), "(low") || strings.HasPrefix(string(runes), "(cap") {
		if strings.HasPrefix(string(runes[4:]), ")") {
			return 1
		}
		if strings.HasPrefix(string(runes[4:]), ",") {
			numstr := ""
			for _, ch := range string(runes[5 : len(runes)-1]) {
				if unicode.IsNumber(ch) {
					numstr += string(ch)
					continue
				}
				if ch != ' ' {
					return 0
				}
			}
			num, err := strconv.Atoi(numstr)
			if err != nil {
				return 0
			}
			if num > 0 {
				return num
			}
			return 0
		}
	}
	return 0
}

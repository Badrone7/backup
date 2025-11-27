package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Flagreplacer(str string) string {
	tempstr := StringPlitter([]rune(str))
	final := ""
	fmt.Println(strings.Join(tempstr, "||"))
	for i := 0; i < len(tempstr); i++ {
		if IsValid(tempstr[i]) {
			continue
		}
		final += tempstr[i]
	}
	// runes := []rune(str)
	// final := ""
	// found := false
	// for i := 0; i < len(runes); i++ {
	// 	if runes[i] == '(' {
	// 		found = true
	// 	}
	// 	if found {
	// 		if runes[i] == ')' {
	// 			found = false
	// 		}
	// 		continue
	// 	}
	// 	final += string(runes[i])
	// }
	// onlyspace := true
	// for _, ch := range final {
	// 	if !IsSpaceExceptNewline(ch) {
	// 		onlyspace = false
	// 		break
	// 	}
	// }
	// if onlyspace {
	// 	final = ""
	// }
	return final
}

func Hex(str string) string {
	final := ""
	suff := ""
	runes := []rune(str)
	i := 0
	for _, ch := range str {
		if unicode.IsNumber(ch) || unicode.IsLetter(ch) {
			break
		}
		final += string(ch)
		i++
	}
	runes = runes[i:]
	for j := len(runes) - 1; j >= 0; j-- {
		if unicode.IsNumber(runes[j]) || unicode.IsLetter(runes[j]) {
			break
		}
		suff += string(runes[j])
		runes = runes[:len(runes)-1]
	}
	dec, err := strconv.ParseInt(string(runes), 16, 64)
	if err != nil {
		return str
	}
	final += strconv.FormatInt(dec, 10) + suff
	return final
}

func Bin(str string) string {
	final := ""
	suff := ""
	runes := []rune(str)
	i := 0
	for _, ch := range str {
		if unicode.IsNumber(ch) || unicode.IsLetter(ch) {
			break
		}
		final += string(ch)
		i++
	}
	runes = runes[i:]
	for j := len(runes) - 1; j >= 0; j-- {
		if unicode.IsNumber(runes[j]) || unicode.IsLetter(runes[j]) {
			break
		}
		suff += string(runes[j])
		runes = runes[:len(runes)-1]
	}
	dec, err := strconv.ParseInt(string(runes), 2, 64)
	if err != nil {
		return str
	}
	final += strconv.FormatInt(dec, 10) + suff
	return final
}

func Cap(str string) string {
	final := ""
	capi := true
	for _, ch := range str {
		if unicode.IsLetter(ch) {
			if capi {
				final += string(unicode.ToUpper(ch))
				capi = false
			} else {
				final += string(unicode.ToLower(ch))
			}
			continue
		}
		final += string(ch)
	}
	return final
}

func Low(str string) string {
	final := strings.ToLower(str)
	return final
}

func Up(str string) string {
	final := strings.ToUpper(str)
	return final
}

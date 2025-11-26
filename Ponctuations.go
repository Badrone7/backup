package main

import "unicode"

func IsQuote(r rune) bool {
	return r == '\''
}

func IsPonc(r rune) bool {
	if r == '.' || r == ',' || r == ';' || r == ':' || r == '!' || r == '?' {
		return true
	}
	return false
}

func IsSp(r rune) bool {
	if r == ' ' || r == '\t' || r == '\r' {
		return true
	}
	return false
}

func PonctuationFixer(text []rune) string {
	temprune := []rune{}
	tempstr := ""
	for i := 0; i < len(text); i++ {
		if i-1 < 0 && IsPonc(text[i]) {
			temprune = append(temprune, text[i])
			continue
		}
		if i+1 >= len(text) && IsPonc(text[i]) {
			temprune = append(temprune, text[i])
			continue
		}
		if i+1 < len(text) && IsPonc(text[i+1]) {
			if IsSp(text[i]) {
				continue
			} else {
				temprune = append(temprune, text[i])
				continue
			}
		}
		if IsPonc(text[i]) {
			temprune = append(temprune, text[i])
			if i+1 < len(text) && !IsSp(text[i+1]) && text[i+1] != '\n' && text[i+1] != '\'' {
				temprune = append(temprune, ' ')
			}
			continue
		}
		temprune = append(temprune, text[i])
	}
	for i := 0; i < len(temprune); i++ {
		tempstr += string(temprune[i])
	}
	return tempstr
}

func OnlySP(str string) bool {
	if str == "" {
		return false
	}
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		if !unicode.IsSpace(runes[i]) || runes[i] == '\n' {
			return false
		}
	}
	return true
}

package main

import (
	"regexp"
)

func IsValid(str string) bool {
	if IsHex(str) || IsBin(str) || IsUp(str) || IsCap(str) || IsLow(str) {
		return true
	}
	return false
}

func IsHex(str string) bool {
	re := regexp.MustCompile(`^\(hex\)`)
	return re.MatchString(str)
}

func IsBin(str string) bool {
	re := regexp.MustCompile(`^\(bin\)`)
	return re.MatchString(str)
}

func IsUp(str string) bool {
	re := regexp.MustCompile(`\(up,\s*\d+\s*\)`)
	de := regexp.MustCompile(`\(up\)$`)
	return re.MatchString(str) || de.MatchString(str)
}

func IsLow(str string) bool {
	re := regexp.MustCompile(`^\(low,?`)
	de := regexp.MustCompile(`^\(low\)`)
	return re.MatchString(str) || de.MatchString(str)
}

func IsCap(str string) bool {
	re := regexp.MustCompile(`^\(cap,?`)
	de := regexp.MustCompile(`^\(cap\)`)
	return re.MatchString(str) || de.MatchString(str)
}

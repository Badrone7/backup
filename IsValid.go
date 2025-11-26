package main

import "strings"

func IsHex(str string) bool {
	if len(str) < 5 {
		return false
	}
	if len(str) < 5 {
		return false
	}
	runes := []rune(str)
	flag := ""
	found := false
	for i := 0; i < len(runes); i++ {
		if runes[i] == '(' {
			found = true
			flag += string(runes[i])
			continue
		}
		if found {
			if runes[i] == ')' {
				flag += string(runes[i])
				break
			}
			flag += string(runes[i])
		}
	}
	runes = []rune(flag)
	if string(runes[1:4]) == "hex" && strings.HasPrefix(string(runes[4:]), ")") {
		return true
	}
	return false
}

func IsBin(str string) bool {
	if len(str) < 5 {
		return false
	}
	runes := []rune(str)
	flag := ""
	found := false
	for i := 0; i < len(runes); i++ {
		if runes[i] == '(' {
			found = true
			flag += string(runes[i])
			continue
		}
		if found {
			if runes[i] == ')' {
				flag += string(runes[i])
				break
			}
			flag += string(runes[i])
		}
	}
	runes = []rune(flag)
	if string(runes[1:4]) == "bin" && strings.HasPrefix(string(runes[4:]), ")") {
		return true
	}
	return false
}

func IsUp(str string) bool {
	runes := []rune(str)
	flag := ""
	found := false
	for i := 0; i < len(runes); i++ {
		if runes[i] == '(' {
			found = true
			flag += string(runes[i])
			continue
		}
		if found {
			if runes[i] == ')' {
				flag += string(runes[i])
				break
			}
			flag += string(runes[i])
		}
	}
	runes = []rune(flag)
	if string(runes[1:3]) == "up" && (strings.HasPrefix(string(runes[3:]), ")") || strings.HasPrefix(string(runes[3:]), ", ")) {
		return true
	}
	return false
}

func IsLow(str string) bool {
	if len(str) < 5 {
		return false
	}
	runes := []rune(str)
	flag := ""
	found := false
	for i := 0; i < len(runes); i++ {
		if runes[i] == '(' {
			found = true
			flag += string(runes[i])
			continue
		}
		if found {
			if runes[i] == ')' {
				flag += string(runes[i])
				break
			}
			flag += string(runes[i])
		}
	}
	runes = []rune(flag)
	if string(runes[1:4]) == "low" && (strings.HasPrefix(string(runes[4:]), ")") || strings.HasPrefix(string(runes[4:]), ", ")) {
		return true
	}
	return false
}

func IsCap(str string) bool {
	if len(str) < 5 {
		return false
	}
	runes := []rune(str)
	flag := ""
	found := false
	for i := 0; i < len(runes); i++ {
		if runes[i] == '(' {
			found = true
			flag += string(runes[i])
			continue
		}
		if found {
			if runes[i] == ')' {
				flag += string(runes[i])
				break
			}
			flag += string(runes[i])
		}
	}
	runes = []rune(flag)
	if string(runes[1:4]) == "cap" && (strings.HasPrefix(string(runes[4:]), ")") || strings.HasPrefix(string(runes[4:]), ", ")) {
		return true
	}
	return false
}

func IsValid(str string) bool {
	if len(str) < 4 {
		return false
	}
	if IsHex(str) || IsBin(str) || IsUp(str) || IsCap(str) || IsLow(str) {
		return true
	}
	return false
}

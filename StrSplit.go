package main

import (
	"regexp"
)

func IsBrack(str string) bool {
	runes := []rune(str)
	bracket := 0
	for i := 0; i < len(runes); i++ {
		if runes[i] == '(' {
			bracket++
		}
		if bracket > 1 {
			return true
		}
		if runes[i] == ')' {
			bracket--
		}
	}
	return false
}

func IsSuufix(str string) bool {
	re := regexp.MustCompile(`\s*\d+\s*\)`)
	return re.MatchString(str)
}

func IsPrefix(str string) bool {
	re := regexp.MustCompile(`^\((hex|bin|up|low|cap),?`)
	return re.MatchString(str)
}

func StringPlitter(text []rune) []string {
	words := []string{}
	word := ""
	nested := 0
	nestedSlice := []string{}
	for i := 0; i < len(text); i++ {
		if text[i] == '\n' {
			if nested > 0 {
				if word != "" {
					words = append(words, word)
					word = ""
				}
				nested = 0
			} else if word != "" {
				words = append(words, word)
				word = ""
			}
			words = append(words, "\n")
			continue
		}
		if word == "" && (text[i] == ' ' || text[i] == '\r') {
			continue
		}
		if text[i] == '(' {
			if word != "" && word != "(" {
				words = append(words, word)
				word = ""
			}
			word += string(text[i])
			nested++
			continue
		}
		if nested > 1 {
			if text[i] == ')' {
				nested--
			}
			word += string(text[i])
			continue
		}
		if nested == 0 && len(word) >= 2 {
			newrunes := []rune(word[1 : len(word)-1])
			nestedSlice = StringPlitter(newrunes)
			nested--
		}
		if len(nestedSlice) != 0 {
			words = append(words, nestedSlice...)
		}
		if text[i] != ' ' && text[i] != '\t' && text[i] != '\n' && text[i] != '\r' {
			word += string(text[i])
			continue
		}
		word += string(text[i])
		if word != "" {
			words = append(words, word)
			word = ""
		}
	}
	if word != "" {
		words = append(words, word)
	}
	newwords := []string{}
	// here we gonna merge any flags splitted
	for i := 0; i < len(words); i++ {
		if IsPrefix(words[i]) && (i+1) < len(words) && IsSuufix(words[i+1]) {
			newwords = append(newwords, words[i]+words[i+1])
			i++
			if (i+1) < len(words) && words[i+1] == "\n" {
				newwords = append(newwords, "\n")
				i++
			}
			continue
		}
		newwords = append(newwords, words[i])
	}
	words = []string{}
	for i := 0; i < len(newwords); i++ {
		if !OnlySP(newwords[i]) {
			words = append(words, newwords[i])
		} else {
			if len(words) == 0 || words[len(words)-1] == "" {
				words = append(words, newwords[i])
				continue
			}
			j := i
			total := newwords[j]
			for {
				if j+1 >= len(newwords) || !OnlySP(newwords[j+1]) {
					break
				}
				if OnlySP(newwords[j+1]) {
					total += newwords[j+1]
				}
				j++
			}
			words = append(words, total)
			i = j
		}
	}
	return words
}

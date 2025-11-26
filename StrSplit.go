package main

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

func StringPlitter(text []rune) []string {
	words := []string{}
	word := ""
	nested := 0
	nestedSlice := []string{}
	for i := 0; i < len(text); i++ {
		if word == "" && (text[i] == ' ' || text[i] == '\t' || text[i] == '\n' || text[i] == '\r') {
			continue
		}
		if text[i] == '(' {
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
		if (nested == 0 || text[i] == '\n') && len(word) >= 2 {
			newrunes := []rune(word[1 : len(word)-1])
			if IsBrack(word) {
				nestedSlice = StringPlitter(newrunes)
			}
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
	return words
}

// func StringPlitter(text []rune) []string {
// 	words := []string{}
// 	word := ""
// 	count := 0
// 	for i := 0; i < len(text); i++ {
// 		if word == "" && (text[i] == ' ' || text[i] == '\t' || text[i] == '\n' || text[i] == '\r') {
// 			continue
// 		}
// 		if text[i] == '(' {
// 			word += string(text[i])
// 			count = 1
// 			continue
// 		}
// 		if count == 1 {
// 			if text[i] == '(' {
// 				word += string(text[i])
// 				if word != "" {
// 					words = append(words, word)
// 					word = ""
// 				}
// 				count = 1
// 				continue
// 			}
// 			if text[i] == ')' {
// 				word += string(text[i])
// 				count = 0
// 				continue
// 			}
// 			word += string(text[i])
// 			continue
// 		}
// 		if text[i] != ' ' && text[i] != '\t' && text[i] != '\n' && text[i] != '\r' {
// 			word += string(text[i])
// 			continue
// 		}
// 		word += string(text[i])
// 		if word != "" {
// 			words = append(words, word)
// 			word = ""
// 		}
// 	}
// 	if word != "" {
// 		words = append(words, word)
// 	}
// 	return words
// }

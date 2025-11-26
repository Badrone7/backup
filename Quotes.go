package main

func QuotesFixer(text []rune) string {
	foundsingle := 0
	temprune := []rune{}
	tempstr := ""
	for i := 0; i < len(text); i++ {
		if text[i] == '\n' {
			foundsingle = 0
			temprune = append(temprune, text[i])
			continue
		}
		if text[i] == '\'' {
			if i-1 >= 0 && i+1 < len(text) && ((text[i-1] >= 'a' && text[i-1] <= 'z') || (text[i-1] >= 'A' && text[i-1] <= 'Z')) && ((text[i+1] >= 'a' && text[i+1] <= 'z') || (text[i+1] >= 'A' && text[i+1] <= 'Z')) {
				temprune = append(temprune, text[i])
				continue
			}
			if foundsingle%2 == 0 {
				if i-1 < 0 {
					foundsingle++
					temprune = append(temprune, text[i])
					if i+1 < len(text) && (IsPonc(text[i+1]) || IsSp(text[i+1]) || IsQuote(text[i+1])) {
						if IsSp(text[i+1]) {
							i++
						}
						continue
					}
					continue
				}
				if IsPonc(text[i-1]) || IsQuote(text[i-1]) {
					if IsQuote(text[i-1]) && text[i-1]!='\n'{
						temprune=append(temprune,' ')
					}
					temprune = append(temprune, text[i])
					foundsingle++
				} else {
					if !IsSp(text[i-1]) && !IsPonc(text[i-1]) && !IsQuote(text[i-1]) {
						temprune = append(temprune, ' ')
					}
					temprune = append(temprune, text[i])
					foundsingle++
				}
				if i+1 < len(text) && (IsPonc(text[i+1]) || IsSp(text[i+1]) || IsQuote(text[i+1])) {
					if IsSp(text[i+1]) {
						i++
					}
					continue
				}
				continue
			} else {
				if IsPonc(text[i-1]) || IsQuote(text[i-1]) || text[i-1] == '\n' {
					temprune = append(temprune, text[i])
					foundsingle++
				} else if IsSp(text[i-1]) {
					if (i+2 >= 0 && text[i-2] != '\'') || i-2 < 0 {
						temprune = temprune[:len(temprune)-1]
					}
					temprune = append(temprune, text[i])
					foundsingle++
				} else {
					temprune = append(temprune, text[i])
					foundsingle++
				}
				if i+1 < len(text) && (IsPonc(text[i+1]) || IsSp(text[i+1]) || IsQuote(text[i+1]) || text[i] == '\n') {
					if i+2 < len(text) && IsSp(text[i+1]) && IsQuote(text[i+1]) {
						temprune = append(temprune,' ')
						i++
					}
					continue
				} else if i+1 < len(text) && !IsPonc(text[i+1]) && !IsSp(text[i+1]) && !IsQuote(text[i+1]) {
					temprune = append(temprune, ' ')
				}
				continue
			}
		}
		temprune = append(temprune, text[i])
	}
	for i := 0; i < len(temprune); i++ {
		tempstr += string(temprune[i])
	}
	return tempstr
}

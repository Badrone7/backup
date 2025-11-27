package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/*
	Our Program is meant to Correct a TEXT file from any extra spaces or syntax errors
	like quotes and punctuations, also it does the transformations if necessary
*/

// our main function
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not Enough Args")
		return
	}
	if len(os.Args) == 2 {
		fmt.Println("MIssing the Result File Name")
		return
	}
	if len(os.Args) > 3 {
		fmt.Println("Too Many Args")
		return
	}
	for i := len(os.Args[1]) - 1; i >= 0; i-- {
		if os.Args[1][i] == '.' {
			if i+1 < len(os.Args[1]) && (os.Args[1][i:] != ".txt" || os.Args[1][i:] == ".") {
				fmt.Println("Input File must be a .txt file")
				return
			}
			if os.Args[1][i:] == "." {
				fmt.Println("Input File must be a .txt file")
				return
			}
			break
		}
		if i == 0 {
			fmt.Println("Input File must be a .txt file")
			return
		}
	}
	for i := len(os.Args[2]) - 1; i >= 0; i-- {
		if os.Args[2][i] == '.' {
			if i+1 < len(os.Args[2]) && (os.Args[2][i:] != ".txt" || os.Args[2][i:] == ".") {
				fmt.Println("Resulted File must be a .txt file")
				return
			}
			if os.Args[2][i:] == "." {
				fmt.Println("Resulted File must be a .txt file")
				return
			}
			break
		}
		if i == 0 {
			fmt.Println("Resulted File must be a .txt file")
			return
		}
	}
	file, ferr := os.Open(os.Args[1])
	if ferr != nil {
		fmt.Println("Error opening file:", ferr)
		return
	}
	defer file.Close()
	temptext, terr := io.ReadAll(file)
	if terr != nil {
		fmt.Println("Error reading file:", terr)
		return
	}
	if string(temptext) == "" {
		fmt.Println("File is Empty")
		return
	}
	for i := 0; i < len(temptext); i++ {
		if !IsSp(rune(temptext[i])) && temptext[i] != '\n' && (temptext[i] < 127 && temptext[i] > 31) {
			break
		}
		if i == len(temptext)-1 {
			fmt.Println("File is Empty or contains only spaces/newlines/tabulations/non printable characters")
			return
		}
	}
	text := []rune(string(temptext))
	tempstr := QuotesFixer(text)
	tempstr = PonctuationFixer([]rune(tempstr))
	text = []rune(tempstr)
	txt := StringPlitter(text)
	fmt.Println(strings.Join(txt, "||"))
	final := []string{}
	tempostr := ""
	for i := 0; i < len(txt); i++ {
		count := 0
		if IsValid(txt[i]) {
			fmt.Println(txt[i])
			tempostr = Flagreplacer(txt[i])
			if txt[i] == "" {
				continue
			}
			// here we would do the transformation
			if IsHex(txt[i]) {
				if len(final) > 0 {
					for j := len(final) - 1; j >= 0; j-- {
						if !OnlySP(final[j]) {
							final[j] = Hex(final[j])
							break
						}
					}
				}
				if tempostr != "" {
					final = append(final, tempostr)
				}
				continue
			}
			if IsBin(txt[i]) {
				if len(final) > 0 {
					for j := len(final) - 1; j >= 0; j-- {
						if !OnlySP(final[j]) {
							final[j] = Bin(final[j])
							break
						}
					}
				}
				if tempostr != "" {
					final = append(final, tempostr)
				}
				continue
			}
			if IsUp(txt[i]) {
				count = Detect(txt[i])
				if count == 0 {
					if tempostr != "" {
						final = append(final, tempostr)
					}
					continue
				}
				wordsAvailable := len(final)
				if wordsAvailable <= 0 {
					if tempostr != "" {
						final = append(final, tempostr)
					}
					continue
				}
				if count > wordsAvailable {
					count = wordsAvailable
				}
				wordsProcessed := 0
				for j := 1; j <= len(final) && wordsProcessed < count; j++ {
					idx := len(final) - j
					if !OnlySP(final[idx]) && !NonValid(final[idx]) {
						final[idx] = Up(final[idx])
						wordsProcessed++
					}
				}
				if tempostr != "" {
					final = append(final, tempostr)
				}
				continue
			}
			if IsLow(txt[i]) {
				count = Detect(txt[i])
				if count == 0 {
					if tempostr != "" {
						final = append(final, tempostr)
					}
					continue
				}
				wordsAvailable := len(final)
				if wordsAvailable <= 0 {
					if tempostr != "" {
						final = append(final, tempostr)
					}
					continue
				}
				if count > wordsAvailable {
					count = wordsAvailable
				}
				wordsProcessed := 0
				for j := 1; j <= len(final) && wordsProcessed < count; j++ {
					idx := len(final) - j
					if !OnlySP(final[idx]) && !NonValid(final[idx]) {
						final[idx] = Low(final[idx])
						wordsProcessed++
					}
				}
				if tempostr != "" {
					final = append(final, tempostr)
				}
				continue
			}
			if IsCap(txt[i]) {
				count = Detect(txt[i])
				if count == 0 {
					if tempostr != "" {
						final = append(final, tempostr)
					}
					continue
				}
				wordsAvailable := len(final)
				if wordsAvailable <= 0 {
					if tempostr != "" {
						final = append(final, tempostr)
					}
					continue
				}
				if count > wordsAvailable {
					count = wordsAvailable
				}
				wordsProcessed := 0
				for j := 1; j <= len(final) && wordsProcessed < count; j++ {
					idx := len(final) - j
					if !OnlySP(final[idx]) && !NonValid(final[idx]) {
						final[idx] = Cap(final[idx])
						wordsProcessed++
					}
				}
				if tempostr != "" {
					final = append(final, tempostr)
				}
				continue
			}
		}
		final = append(final, txt[i])
	}
	FinalResult := ""
	for i := 0; i < len(final); i++ {
		FinalResult += final[i]
	}
	text = []rune(FinalResult)
	FinalResult = ""
	// here we are going to fix the a and an problem
	for i := 0; i < len(text); i++ {
		if i+2 < len(text) && (text[i] == 'a' || text[i] == 'A') && text[i+1] == ' ' {
			if i-1 >= 0 && (text[i-1] >= 'a' && text[i-1] <= 'z') && (text[i-1] >= 'A' && text[i-1] <= 'Z') {
				continue
			}
			if text[i+2] == 'a' || text[i+2] == 'e' || text[i+2] == 'i' || text[i+2] == 'o' || text[i+2] == 'u' || text[i+2] == 'A' || text[i+2] == 'E' || text[i+2] == 'I' || text[i+2] == 'O' || text[i+2] == 'U' {
				FinalResult += "an "
				i++
				continue
			}
		}
		FinalResult += string(text[i])
	}
	FinalResult = QuotesFixer([]rune(FinalResult))
	FinalResult = PonctuationFixer([]rune(FinalResult))
	err := os.WriteFile(os.Args[2], []byte(FinalResult), 0o644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Print("File Processed Successfully\n")
}

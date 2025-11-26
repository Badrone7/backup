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
	tempstr := string(temptext)
	text := []rune(tempstr)
	tempstr = ""
	tempcount := 0
	// here we are going to remove any extra newline or pace or tabulation
	for i := 0; i < len(text); i++ {
		if text[i] != ' ' && text[i] != '\t' && text[i] != '\r' && text[i] != '\n' {
			tempstr += string(text[i])
			tempcount = 0
			continue
		}
		if tempcount == 0 {
			tempstr += string(text[i])
			tempcount++
		}
		if i == len(text)-1 && text[i] != '\n' {
			tempstr += string(text[i])
			tempstr += "\n"
		}
	}
	text = []rune(tempstr)
	tempstr = ""
	tempcount = 0
	text = []rune(PonctuationFixer(text))
	// here we are going to replace if necessary
	txt := StringPlitter(text)
	fmt.Println(strings.Join(txt, "||"))
	final := []string{}
	vld := 1
	tempostr := ""
	for i := 0; i < len(txt); i++ {
		count := 0
		if IsValid(txt[i]) {
			fmt.Println(vld)
			tempostr = Flagreplacer(txt[i])
			if tempostr != ""{
				final = append(final, tempostr)
			}else {
				vld++
			}
			fmt.Println(vld)
			if txt[i] == "" {
				continue
			}
			if txt[i][len(txt[i])-1] == '\'' || txt[i][len(txt[i])-2] == '\'' {
				final = append(final, "'")
			}
			// here we would do the transformation
			if IsHex(txt[i]) {
				if vld <= i {
					final[len(final)-1-vld] = Hex(final[len(final)-1-vld])
				}
				continue
			}
			if IsBin(txt[i]) {
				tempostr = Flagreplacer(txt[i])
				if vld <= i {
					final[i-vld] = Bin(final[i-vld])
				}
				continue
			}
			if IsUp(txt[i]) {
				count = Detect(txt[i])
				if count == 0 {
					continue
				}
				fmt.Println(":", vld)
				if count == 1 {
					if vld <= i {
						final[i-vld] = Up(final[i-vld])
					}
					continue
				}
				if len(final) > 0 {
					if count > len(final) {
						count = len(final)
					}
					for j := 0; j < count; j++ {
						if j-vld > i {
							break
						}
						final[i-j-vld] = Up(final[i-j-vld])
					}
				}
				continue
			}
			if IsLow(txt[i]) {
				count = Detect(txt[i])
				if count == 0 {
					continue
				}
				if count == 1 {
					if vld <= i {
						final[i-vld] = Low(final[i-vld])
					}
					continue
				}
				if len(final) > 0 {
					if count > len(final) {
						count = len(final)
					}
					for j := 0; j < count; j++ {
						if j-vld > i {
							break
						}
						final[i-j-vld] = Low(final[i-j-vld])
					}
				}
				continue
			}
			if IsCap(txt[i]) {
				count = Detect(txt[i])
				if count == 0 {
					continue
				}
				if count == 1 {
					if vld <= i && len(final) > 1 {
						final[i-vld] = Cap(final[i-vld])
					}
					continue
				}
				if len(final) > 0 {
					if count > len(final) {
						count = len(final)
					}
					for j := 0; j < count; j++ {
						if j-vld > i {
							break
						}
						final[i-j-vld] = Cap(final[i-j-vld])
					}
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
	// here we are going to process the quotes
	FinalResult = QuotesFixer([]rune(FinalResult))
	// here we are going to process the punctuations
	FinalResult = PonctuationFixer([]rune(FinalResult))
	err := os.WriteFile(os.Args[2], []byte(FinalResult), 0o644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Print("File Processed Successfully\n")
}

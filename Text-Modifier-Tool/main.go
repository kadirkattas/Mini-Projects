package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fileName := os.Args[1]
	toFile := os.Args[2]

	content, _ := os.ReadFile(fileName)
	f, _ := os.Create(toFile)
	sepContent := strings.Fields(string(content))

	res := ReplaceVowel(FixAgain(FixSpaceForQuote(FixPunctuation(FixQuotes(UseFunc(sepContent))))))

	_, err := f.WriteString(string(res))
	if err != nil {
		log.Fatal(err)
	}
}

func HexToDec(s string) string {
	decimal_num, _ := strconv.ParseInt(s, 16, 64)
	res := strconv.Itoa(int(decimal_num))
	return res
}

func BinToDec(s string) string {
	decimal_num, _ := strconv.ParseInt(s, 2, 64)
	res := strconv.Itoa(int(decimal_num))
	return res
}

func ToUp(s string) string {
	res := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			res += string(char - 32)
		} else {
			res += string(char)
		}
	}
	return res
}

func ToLow(s string) string {
	res := ""
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			res += string(char + 32)
		} else {
			res += string(char)
		}
	}
	return res
}

func ToCap(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if i == 0 && s[i] >= 'a' && s[i] <= 'z' {
			res += string(s[i] - 32)
		} else if i == 0 && s[i] >= 'A' && s[i] <= 'Z' {
			res += string(s[i])
		}
		if i != 0 && s[i] >= 'A' && s[i] <= 'Z' {
			res += string(s[i] + 32)
		} else if i != 0 && s[i] >= 'a' && s[i] <= 'z' {
			res += string(s[i])
		}
	}
	return res
}

func UseFunc(sr []string) string {
	res := ""
	for i := len(sr) - 1; i >= 0; i-- {
		if sr[i] == "(up)" {
			sr[i-1] = ToUp(sr[i-1])
			i--
		}
		if strings.Contains(sr[i], ")") && sr[i-1] == "(up," {
			nStr := strings.Trim(sr[i], ")")
			n, _ := strconv.Atoi(nStr)
			for a := 0; a < n; a++ {
				sr[i-n+a-1] = ToUp(sr[i-n+a-1])
			}
			i--
			continue
		}
		if sr[i] == "(low)" {
			sr[i-1] = ToLow(sr[i-1])
			i--
		}
		if strings.Contains(sr[i], ")") && sr[i-1] == "(low," {
			nStr := strings.Trim(sr[i], ")")
			n, _ := strconv.Atoi(nStr)
			for a := 0; a < n; a++ {
				sr[i-n+a-1] = ToLow(sr[i-n+a-1])
			}
			i--
			continue
		}
		if sr[i] == "(cap)" {
			sr[i-1] = ToCap(sr[i-1])
			i--
		}
		if strings.Contains(sr[i], ")") && sr[i-1] == "(cap," {
			nStr := strings.Trim(sr[i], ")")
			n, _ := strconv.Atoi(nStr)
			for a := 0; a < n; a++ {
				sr[i-n+a-1] = ToCap(sr[i-n+a-1])
			}
			i--
			continue
		}
		if sr[i] == "(hex)" {
			sr[i-1] = HexToDec(sr[i-1])
			i--
		}
		if strings.Contains(sr[i], ")") && sr[i-1] == "(hex," {
			nStr := strings.Trim(sr[i], ")")
			n, _ := strconv.Atoi(nStr)
			for a := 0; a < n; a++ {
				sr[i-n+a-1] = HexToDec(sr[i-n+a-1])
			}
			i--
			continue
		}
		if sr[i] == "(bin)" {
			sr[i-1] = BinToDec(sr[i-1])
			i--
		}
		if strings.Contains(sr[i], ")") && sr[i-1] == "(bin," {
			nStr := strings.Trim(sr[i], ")")
			n, _ := strconv.Atoi(nStr)
			for a := 0; a < n; a++ {
				sr[i-n+a-1] = BinToDec(sr[i-n+a-1])
			}
			i--
			continue
		}
		res = sr[i] + " " + res
	}
	return res
}

func FixPunctuation(s string) string {
	sepStr := strings.Fields(s)
	tempRes := ""
	res := ""

	for i := len(sepStr) - 1; i >= 0; i-- {
		if strings.ContainsAny(sepStr[i], ".!,?:;") {
			if i < len(sepStr)-1 && unicode.IsLetter(rune(sepStr[i+1][0])) {
				tempRes = sepStr[i] + " " + tempRes
			} else {
				tempRes = sepStr[i] + tempRes
			}
		} else {
			tempRes = sepStr[i] + " " + tempRes
		}
	}
	// is there any punct after space
	for i := 0; i < len(tempRes); i++ {
		if unicode.IsSpace(rune(tempRes[i])) {
			if i < len(tempRes)-1 && strings.ContainsAny(string(tempRes[i+1]), ".!,?:;") {
				res += string(tempRes[i+1])
				i++
			} else {
				res += string(tempRes[i])
			}
		} else {
			res += string(tempRes[i])
		}
	}
	// is there any letter after punct
	actualRes := ""
	for i := 0; i < len(res); i++ {
		if i != len(res)-1 {
			if string(res[i]) == "." || string(res[i]) == "," || string(res[i]) == "!" || string(res[i]) == "?" || string(res[i]) == ":" || string(res[i]) == ";" {
				if res[i+1] >= 'a' && res[i+1] <= 'z' || res[i+1] >= 'A' && res[i+1] <= 'Z' {
					actualRes += string(res[i]) + " "
				} else {
					actualRes += string(res[i])
				}
			} else {
				actualRes += string(res[i])
			}
		} else {
			actualRes += string(res[i])
		}
	}

	return actualRes
}

func FixQuotes(s string) string {
	sepStr := strings.Fields(s)
	quotesCount := 0
	res := ""

	for i := 0; i < len(sepStr); i++ {
		if sepStr[i] == "'" {
			quotesCount++
			if quotesCount%2 == 0 {
				res += sepStr[i] + " "
			} else {
				res += sepStr[i]
			}
		} else {
			if quotesCount%2 == 0 {
				res += sepStr[i] + " "
			} else {
				if sepStr[i+1] == "'" {
					res += sepStr[i]
				} else {
					res += sepStr[i] + " "
				}
			}
		}
	}
	return res
}

// adding space for the first quote
func FixSpaceForQuote(s string) string {
	res := ""
	quoteCount := 0

	for i := len(s) - 1; i >= 0; i-- {
		if quoteCount%2 == 1 && string(s[i]) == "'" && string(s[i-1]) != " " && string(s[i+1]) != " " {
			res = " " + string(s[i]) + res
		} else {
			res = string(s[i]) + res
		}
		if string(s[i]) == " " && string(s[i-1]) == "'" {
			quoteCount++
		}
	}
	return res
}

func FixAgain(res string) string {
	actualRes := ""
	quotesCount := 1
	for i := 0; i < len(res); i++ {
		if (string(res[i]) == "." || string(res[i]) == "," || string(res[i]) == "!" || string(res[i]) == "?" || string(res[i]) == ":" || string(res[i]) == ";") && i != len(res)-1 {
			if string(res[i+1]) == "'" && quotesCount%2 == 1 {
				actualRes += string(res[i]) + " "
				quotesCount++
			} else {
				actualRes += string(res[i])
			}
		} else {
			actualRes += string(res[i])
		}
	}
	return actualRes
}

func ReplaceVowel(s string) string {
	sepStr := strings.Fields(s)
	res := ""

	for i := 0; i < len(sepStr); i++ {
		if sepStr[i] == "A" || sepStr[i] == "AN" {
			if i != len(sepStr)-1 {
				switch string(sepStr[i+1][0]) {
				case "a", "e", "i", "o", "u", "h":
					res += "An" + " "
				case "A", "E", "I", "O", "U", "H":
					res += "AN" + " "
				default:
					if unicode.IsUpper(rune(sepStr[i][0])) {
						res += "A" + " "
					} else {
						res += "a" + " "
					}
				}
			} else {
				res += sepStr[i]
			}
		} else if sepStr[i] == "a" || sepStr[i] == "an" {
			if i != len(sepStr)-1 {
				switch string(sepStr[i+1][0]) {
				case "a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H":
					res += "an" + " "
				default:
					res += "a" + " "
				}
			} else {
				res += sepStr[i]
			}
		} else if sepStr[i] == "An" {
			if i != len(sepStr)-1 {
				switch string(sepStr[i+1][0]) {
				case "a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H":
					res += "An" + " "
				default:
					res += "A" + " "
				}
			} else {
				res += sepStr[i]
			}
		} else {
			if i != len(sepStr)-1 {
				res += sepStr[i] + " "
			} else {
				res += sepStr[i]
			}
		}
	}
	return res
}

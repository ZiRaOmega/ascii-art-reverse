package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var line1Content string
	var line2Content string
	var line3Content string
	var line4Content string
	var line5Content string
	var line6Content string
	var line7Content string
	var line8Content string
	var linesContent []string
	var count int
	var realSpace int
	var runeAlphabet []string
	for i := 0; i <= 95; i++ {
		runeAlphabet = append(runeAlphabet, string(i+32))
	}
	if len(os.Args) == 2 {
		if strings.Contains(os.Args[1], "--reverse") && strings.Contains(os.Args[1], "=") {
			os.Args[1] = strings.Replace(os.Args[1], "--reverse", "", 1)
			os.Args[1] = strings.Replace(os.Args[1], "=", "", 10)
		}
		var s []string
		var standard []string
		file, err := os.Open(os.Args[1])
		err = err
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			s = append(s, scanner.Text())
		}
		file2, err := os.Open("standard.txt")
		err = err
		scanner2 := bufio.NewScanner(file2)
		for scanner2.Scan() {
			standard = append(standard, scanner2.Text())
		}
		standard = standard
		line1 := s[0]
		line2 := s[1]
		line3 := s[2]
		line4 := s[3]
		line5 := s[4]
		line6 := s[5]
		line7 := s[6]
		line8 := s[7]
		for j := 0; j < len(s[0]); j++ {
			count = 0
			count += spaceCheck(line1, j)
			count += spaceCheck(line2, j)
			count += spaceCheck(line3, j)
			count += spaceCheck(line4, j)
			count += spaceCheck(line5, j)
			count += spaceCheck(line6, j)
			count += spaceCheck(line7, j)
			count += spaceCheck(line8, j)
			realSpace += realSpaceCheck(line1, j)
			realSpace += realSpaceCheck(line2, j)
			realSpace += realSpaceCheck(line3, j)
			realSpace += realSpaceCheck(line4, j)
			realSpace += realSpaceCheck(line5, j)
			realSpace += realSpaceCheck(line6, j)
			realSpace += realSpaceCheck(line7, j)
			realSpace += realSpaceCheck(line8, j)
			line1Content += string(line1[j])
			line2Content += string(line2[j])
			line3Content += string(line3[j])
			line4Content += string(line4[j])
			line5Content += string(line5[j])
			line6Content += string(line6[j])
			line7Content += string(line7[j])
			line8Content += string(line8[j])
			if count == 8 {
				linesContent = append(linesContent, line1Content)
				linesContent = append(linesContent, line2Content)
				linesContent = append(linesContent, line3Content)
				linesContent = append(linesContent, line4Content)
				linesContent = append(linesContent, line5Content)
				linesContent = append(linesContent, line6Content)
				linesContent = append(linesContent, line7Content)
				linesContent = append(linesContent, line8Content)
				/*
					fmt.Println(line1Content)
					fmt.Println(line2Content)
					fmt.Println(line3Content)
					fmt.Println(line4Content)
					fmt.Println(line5Content)
					fmt.Println(line6Content)
					fmt.Println(line7Content)
					fmt.Println(line8Content)
				*/
				line1Content = ""
				line2Content = ""
				line3Content = ""
				line4Content = ""
				line5Content = ""
				line6Content = ""
				line7Content = ""
				line8Content = ""
			}
			if realSpaceCheck(line1, j) == 1 && realSpaceCheck(line2, j) == 1 && realSpaceCheck(line3, j) == 1 && realSpaceCheck(line4, j) == 1 && realSpaceCheck(line5, j) == 1 && realSpaceCheck(line6, j) == 1 && realSpaceCheck(line7, j) == 1 && realSpaceCheck(line8, j) == 1 {
				linesContent = append(linesContent, "      ")
				linesContent = append(linesContent, "      ")
				linesContent = append(linesContent, "      ")
				linesContent = append(linesContent, "      ")
				linesContent = append(linesContent, "      ")
				linesContent = append(linesContent, "      ")
				linesContent = append(linesContent, "      ")
				linesContent = append(linesContent, "      ")
			}
		}
		results := getArtFromFile(standard, linesContent)
		for i := 0; i < len(results); i++ {
			fmt.Print(string(results[i]/9 + 32))
		}
		fmt.Println()
	} else {
		fmt.Println("Usage: go run . [OPTION]")
		fmt.Println("EX: go run . --reverse=file.txt")
		return
	}
}

func getArtFromFile(file []string, lines []string) []int {
	var check int
	var countCheck int
	var resultsLines []int
	for i := 0; i < len(lines); i += 8 {
		for j := 1; j < 855; j += 9 {
			/*
				fmt.Println(j)
				fmt.Print(string(file[j]))
				fmt.Print(string(file[j+1]))
				fmt.Print(string(file[j+2]))
				fmt.Print(string(file[j+3]))
				fmt.Print(string(file[j+4]))
				fmt.Print(string(file[j+5]))
				fmt.Print(string(file[j+6]))
				fmt.Print(string(file[j+7]))
			*/
			letterLine1 := lines[i]
			letterLine2 := lines[i+1]
			letterLine3 := lines[i+2]
			letterLine4 := lines[i+3]
			letterLine5 := lines[i+4]
			letterLine6 := lines[i+5]
			letterLine7 := lines[i+6]
			letterLine8 := lines[i+7]
			/*
				fmt.Println("start")
				fmt.Print(letterLine1)
				fmt.Print(letterLine2)
				fmt.Print(letterLine3)
				fmt.Print(letterLine4)
				fmt.Print(letterLine5)
				fmt.Print(letterLine6)
				fmt.Print(letterLine7)
				fmt.Print(letterLine8)
			*/
			if letterLine1 == string(file[j]) {
				countCheck++
			}
			if letterLine2 == string(file[j+1]) {
				countCheck++
			}
			if letterLine3 == string(file[j+2]) {
				countCheck++
			}
			if letterLine4 == string(file[j+3]) {
				countCheck++
			}
			if letterLine5 == string(file[j+4]) {
				countCheck++
			}
			if letterLine6 == string(file[j+5]) {
				countCheck++
			}
			if letterLine7 == string(file[j+6]) {
				countCheck++
			}
			if letterLine8 == string(file[j+7]) {
				countCheck++
			}
			if countCheck == 8 {
				check++
				countCheck = 0
				resultsLines = append(resultsLines, j)
			} else {
				countCheck = 0
			}
		}
	}
	return resultsLines
}

func spaceCheck(line string, char int) int {
	if string(line[char]) == " " {
		return 1
	}
	return 0
}

func realSpaceCheck(line string, char int) int {
	if string(line[char]) == " " && string(line[char+1]) == " " && string(line[char+2]) == " " && string(line[char+3]) == " " && string(line[char+4]) == " " && string(line[char+5]) == " " && string(line[char+6]) == " " {
		return 1
	}
	return 0
}

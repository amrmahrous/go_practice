package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	original_area := scanFile()
	new_area := convertAcre(original_area)
	for i := 0; i < 8; i++ {
		new_area = convertAcre(new_area)
	}
	outputArea(new_area)
}

func outputArea(new_area [50][50]string) {
	data := ""
	for i := 0; i < 50; i++ {
		for k := 0; k < 50; k++ {
			data = data + string(new_area[i][k])
			if k == 49 {
				data = data + "\r\n"
			}
		}
	}
	writeToOutput(data)
}
func writeToOutput(data string) {
	f, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(data)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "file written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func convertAcre(original [50][50]string) [50][50]string {
	new_area := [50][50]string{}
	for i := 0; i < 50; i++ {
		for k := 0; k < 50; k++ {
			adjacent_acres := getAdjacentAcres(original, i, k)
			if original[i][k] == "." {
				new_area[i][k] = getNewAcreOpen(adjacent_acres)
			}
			if original[i][k] == "|" {
				new_area[i][k] = getNewAcreTree(adjacent_acres)
			}
			if original[i][k] == "#" {
				new_area[i][k] = getNewAcreLumber(adjacent_acres)
			}
		}
	}
	return new_area
}
func getNewAcreOpen(adjacent_acres [8]string) string {
	count_tree := stringCount("|", adjacent_acres)
	if count_tree > 2 {
		return "|"
	}
	return "."
}

func getNewAcreTree(adjacent_acres [8]string) string {
	count_lumber := stringCount("#", adjacent_acres)
	if count_lumber > 2 {
		return "#"
	}
	return "|"
}

func getNewAcreLumber(adjacent_acres [8]string) string {
	count_tree := stringCount("|", adjacent_acres)
	count_lumber := stringCount("#", adjacent_acres)
	if count_tree > 0 && count_lumber > 0 {
		return "#"
	}
	return "|"
}

func stringCount(str string, list [8]string) int {
	count := 0
	for _, v := range list {
		if v == str {
			count++
		}
	}
	return count
}

func getAdjacentAcres(area [50][50]string, i int, k int) [8]string {
	acres := [8]string{}
	if i < 49 && k < 49 {
		acres[0] = area[i+1][k+1]
	}
	if i < 49 {
		acres[1] = area[i+1][k]
	}
	if k < 49 {
		acres[2] = area[i][k+1]
	}
	if i > 0 && k > 0 {
		acres[3] = area[i-1][k-1]
	}
	if k > 0 {
		acres[4] = area[i][k-1]
	}
	if i > 0 {
		acres[5] = area[i-1][k]
	}
	if i > 0 && k < 49 {
		acres[6] = area[i-1][k+1]
	}
	return acres
}

func scanFile() [50][50]string {
	filename := "input.text"
	filebuffer, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputdata := string(filebuffer)
	data := bufio.NewScanner(strings.NewReader(inputdata))
	data.Split(bufio.ScanRunes)

	area := [50][50]string{}
	for i := 0; i < 50; i++ {
		for k := 0; k < 50; k++ {
			data.Scan()
			letter := data.Text()
			if letter == "\n" || letter == "\r" {
				data.Scan()
			}
			area[i][k] = data.Text()
		}
	}
	return area
}

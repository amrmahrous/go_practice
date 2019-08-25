package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const ACRE_OPEN string = "."
const ACRE_TREE string = "|"
const ACRE_LAMBER string = "#"
const INPUT_FILE string = "input.text"
const OUTPUT_FILE string = "output.txt"
const MINUTES_TO_RUN int = 10
const AREA_WIDTH int = 50
const AREA_LENGHT int = 50

func main() {
	file_content := scanFile()
	area := stringToArea(file_content)
	for i := 0; i < MINUTES_TO_RUN; i++ {
		area = generateNewArea(area)
	}
	resources_count := countResources(area)
	fmt.Println(resources_count, "resources found")
	final_area_string := areaToString(area)
	writeToFile(final_area_string, OUTPUT_FILE)
}
func countResources(final_area [AREA_WIDTH][AREA_LENGHT]string) int {
	tree_count := 0
	labmer_count := 0
	for i := 0; i < AREA_WIDTH; i++ {
		for k := 0; k < AREA_LENGHT; k++ {
			current_acre := final_area[i][k]
			if current_acre == ACRE_LAMBER {
				labmer_count++
			}
			if current_acre == ACRE_TREE {
				tree_count++
			}
		}
	}
	return labmer_count * tree_count
}
func areaToString(area [AREA_WIDTH][AREA_LENGHT]string) string {
	data := ""
	for i := 0; i < AREA_WIDTH; i++ {
		for k := 0; k < AREA_LENGHT; k++ {
			current_acre := area[i][k]
			data = data + string(current_acre)
			if k == AREA_LENGHT-1 {
				data = data + "\r\n"
			}
		}
	}
	return data
}

func generateNewArea(original [AREA_WIDTH][AREA_LENGHT]string) (new_area [AREA_WIDTH][AREA_LENGHT]string) {
	for i := 0; i < AREA_WIDTH; i++ {
		for k := 0; k < AREA_LENGHT; k++ {
			current_arce := original[i][k]
			adjacent_acres := getAdjacentAcres(original, i, k)
			new_area[i][k] = getNewArce(current_arce, adjacent_acres)
		}
	}
	return new_area
}
func getNewArce(current_arce string, adjacent_acres [8]string) string {
	if current_arce == ACRE_OPEN {
		return getNewAcreOpen(adjacent_acres)
	}
	if current_arce == ACRE_TREE {
		return getNewAcreTree(adjacent_acres)
	}
	if current_arce == ACRE_LAMBER {
		return getNewAcreLumber(adjacent_acres)
	}
	fmt.Println("input file has invalid character:" + current_arce)
	os.Exit(1)
	return ""
}
func getNewAcreOpen(adjacent_acres [8]string) string {
	count_tree := stringCount(ACRE_TREE, adjacent_acres)
	if count_tree > 2 {
		return ACRE_TREE
	}
	return ACRE_OPEN
}

func getNewAcreTree(adjacent_acres [8]string) string {
	count_lumber := stringCount(ACRE_LAMBER, adjacent_acres)
	if count_lumber > 2 {
		return ACRE_LAMBER
	}
	return ACRE_TREE
}

func getNewAcreLumber(adjacent_acres [8]string) string {
	count_tree := stringCount(ACRE_TREE, adjacent_acres)
	count_lumber := stringCount(ACRE_LAMBER, adjacent_acres)
	if count_tree > 0 && count_lumber > 0 {
		return ACRE_LAMBER
	}
	return ACRE_OPEN
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

func getAdjacentAcres(area [AREA_WIDTH][AREA_LENGHT]string, i int, k int) [8]string {
	acres := [8]string{}
	if i < AREA_WIDTH-1 && k < AREA_LENGHT-1 {
		acres[0] = area[i+1][k+1]
	}
	if i < AREA_WIDTH-1 {
		acres[1] = area[i+1][k]
	}
	if k < AREA_LENGHT-1 {
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
	if i > 0 && k < AREA_LENGHT-1 {
		acres[6] = area[i-1][k+1]
	}
	if k > 0 && i < AREA_WIDTH-1 {
		acres[7] = area[i+1][k-1]
	}
	return acres
}

func stringToArea(inputdata string) (area [AREA_WIDTH][AREA_LENGHT]string) {
	data := bufio.NewScanner(strings.NewReader(inputdata))
	data.Split(bufio.ScanRunes)
	for i := 0; i < AREA_WIDTH; i++ {
		for k := 0; k < AREA_LENGHT; k++ {
			data.Scan()
			letter := data.Text()
			if letter == "\n" || letter == "\r" {
				data.Scan()
				letter = data.Text()
			}
			area[i][k] = letter
		}
	}
	return area
}

func scanFile() string {
	filebuffer, err := ioutil.ReadFile(INPUT_FILE)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputdata := string(filebuffer)
	return inputdata
}

func writeToFile(data string, file_name string) {
	f, err := os.Create(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println("file written : " + file_name)
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

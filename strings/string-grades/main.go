package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	dataString := "15 25 35 45 55 65 75 85 95"
	lines := AnalyseGrades(dataString)
	for _, line := range lines {
		fmt.Println(line)
	}
	dataString2 := "10 20 30 40 50 60 70 80 90 100"
	lines2 := AnalyseGrades(dataString2)
	for _, line := range lines2 {
		fmt.Println(line)
	}
	dataString3 := "-1 101"
	lines3 := AnalyseGrades(dataString3)
	for _, line := range lines3 {
		fmt.Println(line)
	}
}

type Category struct {
	Count int
}

func AnalyseGrades(gragesString string) []string {
	lines := []string{}
	categoriesIndexes := []string{"0-20", "21-40", "41-60", "61-80", "81-100"}
	categories := make(map[string]Category)

	errMsg := "Нет корректных оценок"
	grades := strings.Split(gragesString, " ")
	countAllGrades := len(grades)
	if countAllGrades <= 1 {
		return []string{errMsg}
	}
	for _, grage := range grades {
		grageInt, err := strconv.Atoi(grage)
		if err != nil {
			return []string{"error"}
		}
		if grageInt >= 0 && grageInt <= 100 {
			for _, category := range categoriesIndexes {
				parts := strings.Split(category, "-")
				if len(parts) == 2 {
					first, err := strconv.Atoi(parts[0])
					if err != nil {
						return []string{"error"}
					}
					second, err := strconv.Atoi(parts[1])
					if err != nil {
						return []string{"error"}
					}

					if grageInt >= first && grageInt <= second {
						c := categories[category]
						c.Count++
						categories[category] = c
					}
				}
			}
		}
	}

	noResult := true
	for _, category := range categoriesIndexes {
		if categories[category].Count != 0 {
			noResult = false
			percent := float64(categories[category].Count) / float64(countAllGrades) * 100
			line := fmt.Sprintf("%s: %d оценок, %.1f%%\n", category, categories[category].Count, percent)

			lines = append(lines, line)
		}
	}
	if noResult {
		return []string{errMsg}
	}
	return lines
}

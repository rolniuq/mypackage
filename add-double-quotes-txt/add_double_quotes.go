package adddoublequotestxt

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Output struct {
	Data []string `json:"data"`
}

func Run() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	res := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	d := Output{res}
	db, _ := json.Marshal(d)

	os.WriteFile("output.json", db, 0644)
}

package adddoublequotestxt

import (
	"bufio"
	"encoding/json"
	"os"
)

const (
	DEFAULT_OUT_DIR   = ".out"
	DEFAULT_FILE_NAME = "output.json"
)

type Output struct {
	Data []string `json:"data"`
}

func Run(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	res := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	d := Output{res}
	db, _ := json.Marshal(d)

	err = WriteFile(GetFileName(fileName), db)
	if err != nil {
		return err
	}

	return nil
}

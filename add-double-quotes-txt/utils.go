package adddoublequotestxt

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func WriteFile(fileName string, data []byte) error {
	return os.WriteFile(fileName, data, 0644)
}

func GetFileName(fileName string) string {
	splits := strings.Split(filepath.Base(fileName), ".")
	if len(splits) < 2 {
		return fmt.Sprintf("%s/%s.json", DEFAULT_OUT_DIR, DEFAULT_FILE_NAME)
	}

	res := fmt.Sprintf("%s/%s.json", DEFAULT_OUT_DIR, splits[0])
	return res
}

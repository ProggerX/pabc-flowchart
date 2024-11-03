package extra

import (
	"io/fs"
	"os"
	"strings"
)

func ReadLines(filename string) ([]string, error) {
	bts, err := os.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}
	str := string(bts)
	lines := strings.Split(str, "\n")
	for i := 0; i < len(lines); i++ {
		lines[i] = strings.TrimSpace(lines[i])
		lines[i] = strings.Trim(lines[i], "\r")
	}
	return lines, nil
}

func WriteLines(filename string, lines []string) error {
	str := strings.Join(lines, "\n")
	err := os.WriteFile(filename, []byte(str), fs.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

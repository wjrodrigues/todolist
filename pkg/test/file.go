package test

import (
	"fmt"
	"os"
	"strings"
)

func FilePath(fileName string) string {
	absolutePath, _ := os.Getwd()
	basePath := strings.Split(absolutePath, "app")[0]

	return fmt.Sprintf("%sapp/%s", basePath, fileName)
}

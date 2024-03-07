package file

import (
	"fmt"
	"os"
	"strings"
)

func Path(fileName string) string {
	absolutePath, _ := os.Getwd()
	basePath := strings.Split(absolutePath, "app")[0]

	return fmt.Sprintf("%sapp/%s", basePath, fileName)
}

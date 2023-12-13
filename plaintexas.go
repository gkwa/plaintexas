package plaintexas

import (
	"fmt"
	"log/slog"
	"path/filepath"
)

const DataDir = "data"

var DataDirAbsPath string

func init() {
	var err error
	DataDirAbsPath, err = filepath.Abs(DataDir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func Main() int {
	slog.Debug("plaintexas", "test", true)
	run()
	return 0
}

func run() {
	fmt.Println("Absolute path of 'data' directory:", DataDirAbsPath)
}

package generation

import (
	"fmt"
	"os/exec"
	"runtime"
)

// Generate takes a path and length and returns the output from the text generator
func Generate(path, length string) (string, error) {
	cmd, err := exec.Command("../generation/gen-"+runtime.GOOS, "gen", path, "-length", length).Output()
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return string(cmd), nil
}

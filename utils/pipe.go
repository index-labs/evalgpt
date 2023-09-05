package utils

import (
	"io"
	"os"
)

func ReadStdinPipeData() ([]byte, error) {
	fileInfo, _ := os.Stdin.Stat()
	if (fileInfo.Mode() & os.ModeNamedPipe) != os.ModeNamedPipe {
		return nil, nil
	}
	return io.ReadAll(os.Stdin)
}

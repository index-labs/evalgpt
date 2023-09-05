package utils

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func RunCmdWithTimeout(workDir string, timeout time.Duration, name string, arg ...string) (string, error) {
	ctx := context.Background()
	if timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
		defer cancel()
	}

	cmd := exec.CommandContext(ctx, name, arg...)
	cmd.Dir = workDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("cmd err: %v, output: %s", err, string(output))
		return "", err
	}
	return string(output), nil
}

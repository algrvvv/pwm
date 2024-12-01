package gpg

import (
	"bytes"
	"os"
	"os/exec"
)

func Decrypt(data string) (string, error) {
	cmd := exec.Command("gpg", "--decrypt", "--pinentry-mode", "loopback")
	cmd.Env = append(os.Environ(), "GPG_TTY=/dev/tty")
	cmd.Stdin = bytes.NewBufferString(data)

	var output bytes.Buffer
	cmd.Stdout = &output

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return output.String(), nil
}

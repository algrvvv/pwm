package gpg

import (
	"bytes"
	"os/exec"

	"github.com/spf13/viper"
)

func Encrypt(data string) (string, error) {
	r := viper.GetString("gpg_user")
	cmd := exec.Command("gpg", "--encrypt", "--recipient", r, "--armor")
	cmd.Stdin = bytes.NewBufferString(data)

	var output bytes.Buffer
	cmd.Stdout = &output

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return output.String(), nil
}

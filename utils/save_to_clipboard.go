package utils

import "golang.design/x/clipboard"

// Copy save data to clipboard.
// TODO: добавить таймер после которого очищать буфер
func Copy(data string) {
	clipboard.Write(clipboard.FmtText, []byte(data))
}

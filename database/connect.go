package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/algrvvv/pwm/utils"
)

var C *sql.DB

func Open() (err error) {
	path, err := utils.GetDBPath()
	if err != nil {
		return err
	}

	C, err = sql.Open("sqlite3", path)
	if err != nil {
		return err
	}
	return
}

func Close() error {
	return C.Close()
}

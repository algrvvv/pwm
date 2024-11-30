// TODO: ДОБАВИТЬ ПРОВЕРКУ НА СУЩЕСТВОВАНИЕ ТАКОЙ ЗАПИСИ
// И ДОБАВИТЬ ОШИБКУ, КОТОРАЯ БУДЕТ ГОВОРИТЬ О ТОМ, ЧТО ТАКАЯ ЗАПИСЬ УЖЕ ЕСТЬ И ЕЕ МОЖНО
// ПЕРЕЗАПИСАТЬ ПРИ НАДОБНОСТИ

package storage

import "github.com/algrvvv/pwm/database"

type SQLiteStorage struct{}

func NewSqliteStorage() Storage {
	return &SQLiteStorage{}
}

func (s *SQLiteStorage) GetNotes() ([]Note, error) {
	var notes []Note

	query := "select id, name, value, created_at from notes"
	rows, err := database.C.Query(query)
	if err != nil {
		return notes, err
	}
	defer rows.Close()

	for rows.Next() {
		var n Note
		if err := rows.Scan(&n.ID, &n.Name, &n.Value, &n.CreatedAt); err != nil {
			return notes, err
		}

		notes = append(notes, n)
	}

	return notes, nil
}

func (s *SQLiteStorage) GetNoteByName(name string) (Note, error) {
	var n Note
	query := "select id, name, value, created_at from notes where name = ?"
	err := database.C.QueryRow(query, name).Scan(&n.ID, &n.Name, &n.Value, &n.CreatedAt)
	return n, err
}

func (s *SQLiteStorage) SaveNote(note Note) error {
	query := "insert or replace into notes(name, value) values(?, ?)"
	_, err := database.C.Exec(query, note.Name, note.Value)
	return err
}

func (s *SQLiteStorage) DeleteNoteByName(name string) error {
	query := "delete from notes where name = ?"
	_, err := database.C.Exec(query, name)
	return err
}

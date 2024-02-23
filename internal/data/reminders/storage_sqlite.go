package reminders

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type SQLiteStorage struct {
	db *sql.DB
}

func NewSqlite(ds string) (*SQLiteStorage, error) {
	if ds != "" {
		ds = "reminders.db"
	}

	db, err := sql.Open("sqlite3", ds)
	if err != nil {
		return nil, fmt.Errorf("failed to open db %w", err)
	}

	_, err = db.Exec(CREATE_TABLE)
	if err != nil {
		return nil, fmt.Errorf("failed to create a table %w", err)
	}

	return &SQLiteStorage{
		db: db,
	}, nil
}

func (s *SQLiteStorage) Save(r *Reminder) error {
	flag := 0
	if r.Flag {
		flag = 1
	}

	var dueDate int64
	if r.DueDate != nil {
		d := *r.DueDate
		dueDate = d.UnixMilli()
	}

	_, err := s.db.Exec(
		INSERT_QUERY,
		r.Id,
		r.Title,
		r.Status.ToString(),
		r.Notes,
		r.Category,
		r.Priority.ToString(),
		flag,
		TagsToString(r.Tags),
		dueDate,
	)
	if err != nil {
		return fmt.Errorf("failed to insert: %w", err)
	}
	return nil
}

func (s *SQLiteStorage) ReadAll() ([]Reminder, error) {
	rows, err := s.db.Query(SELECT_ALL_QUERY)
	if err != nil {
		return nil, fmt.Errorf("failed query records: %w", err)
	}

	result := make([]Reminder, 0)

	var status string
	var priority string
	var tags string
	var dueDate int64
	var flag int16

	for rows.Next() {
		r := Reminder{}

		err := rows.Scan(&r.Id, &r.Title, &status, &r.Notes, &r.Category, &priority, &flag, &tags, &dueDate)

		if err != nil {
			return nil, fmt.Errorf("failed to query one of the record: %w", err)
		}

		r.Status = ToStatusType(status)
		r.Priority = ToPriorityType(priority)
		r.Tags = strings.Split(tags, ",")

		if flag > 0 {
			r.Flag = true
		} else {
			r.Flag = false
		}

		if dueDate != 0 {
			t := time.UnixMilli(int64(dueDate))
			r.DueDate = &t
		}
		result = append(result, r)
	}

	return result, nil
}

func (s *SQLiteStorage) Read(id string) (*Reminder, error) {
	row := s.db.QueryRow(SELECT_BY_ID, id)

	var status string
	var priority string
	var tags string
	var dueDate int64
	r := Reminder{}

	err := row.Scan(&r.Id, &r.Title, &status, &r.Notes, &r.Category, &priority, &r.Flag, &tags, &dueDate)
	if err != nil {
		return nil, fmt.Errorf("failed to query record by id: %w", err)
	}

	r.Status = ToStatusType(status)
	r.Priority = ToPriorityType(priority)
	r.Tags = strings.Split(tags, ",")
	if dueDate != 0 {
		t := time.UnixMilli(int64(dueDate))
		r.DueDate = &t
	}

	return &r, nil
}

func (s *SQLiteStorage) FlipStatus(id string) (*Reminder, error) {
	fmt.Println("Update Triggered")
	r, err := s.Read(id)
	if err != nil {
		return nil, fmt.Errorf("reminder does not exist %w", err)
	}

	ns := PENDING
	if r.Status == PENDING {
		ns = COMPLETED
	}
	r.Status = ns

	s.db.Exec(UPDATE_FLIP_STATUS, ns.ToString(), id)
	if err != nil {
		return &Reminder{}, fmt.Errorf("failed to flip status: %w", err)
	}

	return r, nil
}

func (s *SQLiteStorage) FlipFlag(id string) (*Reminder, error) {
	fmt.Println("Update Triggered")
	r, err := s.Read(id)
	if err != nil {
		return nil, fmt.Errorf("reminder does not exist %w", err)
	}

	// r.Flag = !r.Flag

	s.db.Exec(UPDATE_FLIP_FLAG, !r.Flag, id)
	if err != nil {
		return &Reminder{}, fmt.Errorf("failed to flip flag: %w", err)
	}

	return r, nil
}

func (s *SQLiteStorage) Remove(id string) error {
	_, err := s.db.Exec(DELETE_BY_ID, id)
	if err != nil {
		return fmt.Errorf("failed to delete record with id %s", id)
	}
	return nil
}

const CREATE_TABLE = `
	CREATE TABLE IF NOT EXISTS reminders (
		id TEXT NOT NULL PRIMARY KEY,
		title TEXT NOT NULL,
		status TEXT NOT NULL,
		notes TEXT,
		category TEXT,
		priority TEXT,
		flag INTEGER,
		tags TEXT,
		due_date INTEGER
	);
`

const SELECT_ALL_QUERY = `
	SELECT *
	FROM reminders;
`

const SELECT_BY_ID = `
	SELECT *
	FROM reminders
	WHERE id = ?;
`

const INSERT_QUERY = `
	INSERT INTO reminders (id,title,status,notes,category,priority,flag,tags,due_date) VALUES (?,?,?,?,?,?,?,?,?)
`

const DELETE_BY_ID = `
	DELETE FROM reminders
	WHERE id = ?;
`

const UPDATE_FLIP_STATUS = `
	UPDATE reminders
	SET status=?
	WHERE id=?;
`
const UPDATE_FLIP_FLAG = `
	UPDATE reminders
	SET flag=?
	WHERE id=?;
`

package reminders

import (
	"fmt"
	"time"
)

type RemindersOps interface {
	Save(reminder *Reminder) error
	ReadAll() ([]Reminder, error)
	Read(id string) (*Reminder, error)
	FlipStatus(id string) (*Reminder, error)
	FlipFlag(id string) (*Reminder, error)
	Remove(id string) error
}

type ReminderModel struct {
	Storage RemindersOps
}

func New(drv string, dsn string) (*ReminderModel, error) {
	switch drv {
	case "sqlite":
		ss, err := NewSqlite(dsn)
		if err != nil {
			return nil, fmt.Errorf("failed to instantiate %q DB Connection %w", drv, err)
		}
		return &ReminderModel{Storage:ss}, nil
	case "json":
		ss, err := NewJsonStorage("reminders.json")
		if err != nil {
			return nil, ErrFailedToOpenDB
		}
		return &ReminderModel{Storage:ss}, nil
		//todo: case "postgres":
		//todo: case "mysql":
		//todo: case "csv":
	default:
		return nil, ErrStorageNotSupported
	}
}

type Reminder struct {
	Id            string       `json:"id,omitempty"`
	Title         string       `json:"title,omitempty"`
	Status        StatusType   `json:"status,omitempty"`
	Notes         string       `json:"notes,omitempty"`
	Category      string       `json:"category,omitempty"`
	Priority      PriorityType `json:"priority,omitempty"`
	Flag          bool         `json:"flag"`
	Tags          []string     `json:"tags,omitempty"`
	DueDate       *time.Time   `json:"dueOn,omitempty"`
	CreatedDate   *time.Time   `json:"createdOn,omitempty"`
	CompletedDate *time.Time   `json:"completedOn,omitempty"`
}

func TagsToString(ts []string) string {
	result := ""
	for i := 0; i < len(ts); i++ {
		result += ts[i]
		if i < len(ts)-1 {
			result += ","
		}
	}
	return result
}

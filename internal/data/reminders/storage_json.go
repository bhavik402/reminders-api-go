package reminders

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type JsonStorage struct {
	FileName string
}

func NewJsonStorage(filename string) (*JsonStorage, error) {
	return &JsonStorage{
		FileName: filename,
	}, nil
}

func (s *JsonStorage) Save(r *Reminder) error {
	reminders, err := readFromFile(s.FileName)
	if err != nil {
		switch {
		case !errors.Is(err, ErrFailedToOpenFile):
			return err
		}
	}
	reminders = append(reminders, *r)

	b, err := json.Marshal(reminders)
	if err != nil {
			return err
	}
	fmt.Println(string(b))

	err = os.WriteFile(s.FileName, b, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}

func (s *JsonStorage) ReadAll() ([]Reminder, error) {
	reminders, err := readFromFile(s.FileName)
	if err != nil {
		return []Reminder{}, err
	}
	b, err := json.Marshal(reminders)
	if err != nil {
			return []Reminder{}, err
	}
	fmt.Println(string(b))
	return reminders, nil
}

func (db *JsonStorage) Read(id string) (*Reminder, error) {
	return &Reminder{}, nil
}

func (db *JsonStorage) FlipStatus(id string) (*Reminder, error) {
	return &Reminder{}, nil
}

func (db *JsonStorage) FlipFlag(id string) (*Reminder, error) {
	return &Reminder{}, nil
}

func (db *JsonStorage) Remove(id string) error {
	return nil
}

func readFromFile(filename string) ([]Reminder, error) {
	result := make([]Reminder, 0)

	file, err := os.Open(filename)
	if err != nil {
		return nil, ErrFailedToOpenFile
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, ErrFailedToReadFile
	}

	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, ErrFailedToUnmarshalJson
	}

	return result, nil
}

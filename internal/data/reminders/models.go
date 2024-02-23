package reminders

import (
	"encoding/json"
	"fmt"
)

type Models struct {
	Reminders ReminderModel
}

type PriorityType uint8

const (
	UNASSIGNED PriorityType = iota
	LOW
	MEDIUM
	HIGH
)

func (pt *PriorityType)ToString() string {
	switch *pt {
	case LOW:
		return "LOW"
	case MEDIUM:
		return "MEDIUM"
	case HIGH:
		return "HIGH"
	default:
		return ""
	}
}

func ToPriorityType(s string) PriorityType {
	switch s {
	case "LOW":
		return LOW
	case "MEDIUM":
		return MEDIUM
	case "HIGH":
		return HIGH
	default:
		return UNASSIGNED
	}
}

func (s PriorityType) MarshalJSON() ([]byte, error) {
	if s == UNASSIGNED {
		return []byte{}, fmt.Errorf("status type UNASSIGNED")
	}
	return json.Marshal(s.ToString())
}

func (s *PriorityType) UnmarshalJSON(data []byte) (err error) {
	var PriorityType string
	if err := json.Unmarshal(data, &PriorityType); err != nil {
		return err
	}
	*s = ToPriorityType(PriorityType)
	return nil
}


type StatusType uint8

const (
	UNKNOWN StatusType = iota
	PENDING
	COMPLETED
)

func (s StatusType) ToString() string {
	switch s {
	case COMPLETED:
		return "COMPLETED"
	default:
		return "PENDING"
	}
}

func ToStatusType(s string) StatusType {
	switch s {
	case "COMPLETED":
		return COMPLETED
	default:
		return PENDING
	}
}

func (s StatusType) MarshalJSON() ([]byte, error) {
	if s == UNKNOWN {
		return []byte{}, fmt.Errorf("status type UNKNOWN")
	}
	return json.Marshal(s.ToString())
}

func (s *StatusType) UnmarshalJSON(data []byte) (err error) {
	var statusType string
	if err := json.Unmarshal(data, &statusType); err != nil {
		return err
	}
	*s = ToStatusType(statusType)
	return nil
}
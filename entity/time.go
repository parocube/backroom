package entity

import (
	"database/sql"
	"encoding/json"
	"time"
)

// NullTime is a global reusable nullable time type.
// Any struct can use this type directly.
type NullTime struct {
	sql.NullTime
}

// MarshalJSON serializes to a RFC3339 time string or null for the frontend.
func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if nt == nil || !nt.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(nt.Time.Format(time.RFC3339))
}

// UnmarshalJSON deserializes a time value or null from the frontend.
func (nt *NullTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		nt.Valid = false
		return nil
	}

	var t time.Time
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	nt.Time = t
	nt.Valid = true
	return nil
}

// NewTime creates a NullTime with a valid time value.
func NewTime(t time.Time) NullTime {
	return NullTime{
		NullTime: sql.NullTime{
			Time:  t,
			Valid: true,
		},
	}
}

// NewNullTime creates a NullTime with a null value.
func NewNullTime() NullTime {
	return NullTime{
		NullTime: sql.NullTime{
			Valid: false,
		},
	}
}

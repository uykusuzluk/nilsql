package nilsql

import (
	"database/sql"
	"encoding/json"
)

// Time wraps sql.Nulltime for proper JSON Marshaling
type Time struct {
	N sql.NullTime
}

// MarshalJSON implements the Marshaler interface for Time
func (n Time) MarshalJSON() ([]byte, error) {
	if !n.N.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.N.Time.Format("06-01-02 15:04:05"))
}

// Scan calls sql.Nulltime Scan method that implements the Scanner interface.
func (n *Time) Scan(value interface{}) error {
	err := n.N.Scan(value)
	if err != nil {
		return err
	}
	return nil
}

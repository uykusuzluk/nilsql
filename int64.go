package nilsql

import (
	"database/sql"
	"encoding/json"
)

// Int64 wraps sql.NullInt64 for proper JSON Marshaling
type Int64 struct {
	N sql.NullInt64
}

// MarshalJSON implements the Marshaler interface for Int64
func (n Int64) MarshalJSON() ([]byte, error) {
	if !n.N.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.N.Int64)
}

// Scan calls sql.NullInt64 Scan method that implements the Scanner interface.
func (n *Int64) Scan(value interface{}) error {
	err := n.N.Scan(value)
	if err != nil {
		return err
	}
	return nil
}

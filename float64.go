package nilsql

import (
	"database/sql"
	"encoding/json"
)

// Float wraps sql.NullInt64 for proper JSON Marshaling
type Float struct {
	N sql.NullFloat64
}

// MarshalJSON implements the Marshaler interface for Float
func (n Float) MarshalJSON() ([]byte, error) {
	if !n.N.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.N.Float64)
}

// Scan calls sql.NullFloat64 Scan method that implements the Scanner interface.
func (n *Float) Scan(value interface{}) error {
	err := n.N.Scan(value)
	if err != nil {
		return err
	}
	return nil
}

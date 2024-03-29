package nilsql

import (
	"database/sql"
	"encoding/json"
)

// Bool wraps sql.NullBool for proper JSON Marshaling
type Bool struct {
	N sql.NullBool
}

// MarshalJSON implements the Marshaler interface for Bool
func (n Bool) MarshalJSON() ([]byte, error) {
	if !n.N.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.N.Bool)
}

// Scan calls sql.NullBool Scan method that implements the Scanner interface.
func (n *Bool) Scan(value interface{}) error {
	err := n.N.Scan(value)
	if err != nil {
		return err
	}
	return nil
}

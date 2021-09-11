package nilsql

import (
	"database/sql"
	"encoding/json"
)

// String wraps sql.NullString for proper JSON Marshaling
type String struct {
	N sql.NullString
}

// MarshalJSON implements the Marshaler interface for String
func (n String) MarshalJSON() ([]byte, error) {
	if !n.N.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.N.String)
}

// Scan calls sql.Nulltime Scan method that implements the Scanner interface.
func (n *String) Scan(value interface{}) error {
	var sqlNT sql.NullString
	err := sqlNT.Scan(value)
	if err != nil {
		return err
	}
	n.N.String = sqlNT.String
	n.N.Valid = sqlNT.Valid
	return nil
}

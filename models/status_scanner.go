package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Scan JSONB data from DB into a Webook struct
func (m *Status) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := Status{}
	err := json.Unmarshal(bytes, &result)
	*m = Status(result)
	return err
}

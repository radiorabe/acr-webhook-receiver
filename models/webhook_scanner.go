package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Scan JSONB data from DB into a Webook struct
func (m *Webhook) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := Webhook{}
	err := json.Unmarshal(bytes, &result)
	*m = Webhook(result)
	return err
}

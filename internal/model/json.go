package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JSON map[string]interface{}

func (j JSON) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan JSON: %v", value)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		return err
	}
	*j = JSON(result)
	return nil
}

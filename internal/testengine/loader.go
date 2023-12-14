package testengine

import (
	"encoding/json"
	"fmt"
	"os"
	"test-engine/internal/model"
)

// LoadTestFromFile loads a test from a file
func LoadTestFromFile(filePath string) (model.Test, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return model.Test{}, fmt.Errorf("error reading file: %v", err)
	}

	var test model.Test
	if err := json.Unmarshal(fileContent, &test); err != nil {
		return model.Test{}, fmt.Errorf("error unmarshalling test: %v", err)
	}

	return test, nil
}

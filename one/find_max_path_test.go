package one

import (
	"encoding/json"
	"os"
	"testing"
)

func TestFindMaxPath(t *testing.T) {
	data, err := os.ReadFile("input.json")
	if err != nil {
		panic(err)
	}

	var triangle [][]int
	if err := json.Unmarshal(data, &triangle); err != nil {
		panic(err)
	}

	expected := 7273
	result := FindMaxPath(triangle)

	if result != expected {
		t.Errorf("actual = %v, expected %v", result, expected)
	}
}

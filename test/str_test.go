package test

import (
	"php4go/string"
	"strconv"
	"testing"
)

func TestStrLen(t *testing.T) {
	str := string.StrInstance()

	result, err := str.StrLen("Arash")

	if err != nil {
		t.Fatalf("StrLen returned an error: %v", err)
	}

	convertInt, err := strconv.Atoi(result)
	if err != nil {
		t.Fatalf("Failed to convert result to int: %v", err)
	}

	expected := 5

	if convertInt != expected {
		t.Errorf("StrLen(\"Arash\") = %d; want %d", convertInt, expected)
	}

}

func TestStrPos(t *testing.T) {
	str := string.StrInstance()

	value := "arash narimani"
	found := "narimani"

	result := str.StrPos(value, found)

	expected := 6

	if result != expected {
		t.Errorf("StrPos(\"arash narimani\") = %d; want %d", result, expected)
	}

}

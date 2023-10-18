package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"regexp"
	"strings"
	"testing"
)

func TestGenerateStructFromJSON(t *testing.T) {
	json := `{
		"name": "John",
		"age": 30,
		"addresses": [
			{"street": "123 Main St"},
			{"street": "456 Oak Rd"}
		]
	}`

	expected := `
		package main
		
		type User struct {
			name string
			age int
			addresses []struct {
				street string
			}
		}
	`

	file, err := GenerateStructFromJSON(json)
	if err != nil {
		t.Fatal(err)
	}

	actual := file.GoString()

	actual = strings.ReplaceAll(actual, "\n", "")
	expected = strings.ReplaceAll(expected, "\n", "")
	actual = strings.ToLower(actual)
	expected = strings.ToLower(expected)
	actual = strings.ReplaceAll(actual, " ", "")
	expected = strings.ReplaceAll(expected, " ", "")

	pattern := regexp.QuoteMeta(expected)
	match, _ := regexp.MatchString(pattern, actual)

	if match {
		fmt.Println("Strings match!")
	} else {
		fmt.Println("Strings do not match.")
	}

	if !cmp.Equal(actual, expected) {
		t.Errorf("Expected: %s, got: %s", expected, actual)
	}
}

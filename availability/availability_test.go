package availability

import (
	"strings"
	"testing"
)

func TestEcho(t *testing.T) {
	expected := []string{"a", "b", "c", "d"}
	actual := echo("a", "b", "c", "d")
	for index, value := range expected {
		if value != actual[index] {
			t.Error("Test failed", "Expected", value, "and received", actual[index])
		}
	}
}

func TestFindAvailability(t *testing.T) {
	expected := [][]string{{"Hello, Go!"}}
	actual := findAvailability("a", "b", "c", "d")
	if strings.ToLower(actual[0][0]) != strings.ToLower(expected[0][0]) {
		t.Error("Test failed", "Expected", expected, "and received", actual)
	}
}

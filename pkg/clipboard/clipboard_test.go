package clipboard_test

import (
	. "github.com/feranwq/ppt/pkg/clipboard"
	"testing"
)

func TestCopyAndPaste(t *testing.T) {
	expected := "hello world"

	err := WriteAll(expected)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Errorf("want %s, got %s", expected, actual)
	}
}

func TestMultiCopyAndPaste(t *testing.T) {
	expected1 := "French: éèêëàùœç"
	expected2 := "Weird UTF-8: 💩☃"

	err := WriteAll(expected1)
	if err != nil {
		t.Fatal(err)
	}

	actual1, err := ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if actual1 != expected1 {
		t.Errorf("want %s, got %s", expected1, actual1)
	}

	err = WriteAll(expected2)
	if err != nil {
		t.Fatal(err)
	}

	actual2, err := ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if actual2 != expected2 {
		t.Errorf("want %s, got %s", expected2, actual2)
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestStringParser(t *testing.T) {
	want := "hello world!"
	str, _ := decode("12:hello world!")
	if str != want {
		t.Fatalf("Does not work :(, %s", str)
	}
}

func TestIntegerZeroParser(t *testing.T) {
	want := 0
	num, _ := decode("i0e")
	if num != want {
		t.Fatalf("Does not work :(, %d", num)
	}
}

func TestPositiveIntegerParser(t *testing.T) {
	want := 42
	num, _ := decode("i42e")
	if num != want {
		t.Fatalf("Does not work :(, %d", num)
	}
}

func TestNegativeIntegerParser(t *testing.T) {
	want := -42
	num, _ := decode("i-42e")
	if num != want {
		t.Fatalf("Does not work :(, %d", num)
	}
}

func TestEmptyListParser(t *testing.T) {
	result, _ := decode("le")
	if len(result.([]interface{})) != 0 {
		t.Fatalf("Does not work :(, %d", len(result.([]interface{})))
	}
}

func TestListParser(t *testing.T) {
	var res []interface{}
	res = append(res, "bencode")
	res = append(res, -20)
	result, _ := decode("l7:bencodei-20ee")

	if result.([]interface{})[0] != res[0] || result.([]interface{})[1] != res[1] {
		t.Fatalf("Does not work :(, %#v", res)
	}
}

func TestDictionaryParser(t *testing.T) {
	result, _ := decode("d4:wiki7:bencode7:meaningi42ee")
	expected := map[string]interface{}{
		"wiki":    "bencode",
		"meaning": 42,
	}

	if reflect.DeepEqual(result, expected) != true {
		t.Fatalf("Does not work :(, %#v", result)
	}
}

package main

import "testing"

func TestStringParser(t *testing.T) {
  want := "hello world!"
  str, _ := parseString("12:hello world!", 0)
  if str != want {
    t.Fatalf("Does not work :(, %s", str)
  }
}

func TestIntegerParser(t *testing.T) {
}

func TestListParser(t *testing.T) {
}

func TestDictionaryParser(t *testing.T) {
}

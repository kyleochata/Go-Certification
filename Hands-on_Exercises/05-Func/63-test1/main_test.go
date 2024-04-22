package main

import "testing"

func TestAdd(t *testing.T) {
	//part of the same package as Add() ==> able to see fxn eventhough it's on a different file.
	total := Add(5, 5)
	if total != 10 {
		t.Errorf("Sum incorrect. Expected: %v\t Output: %v\n", 10, total)
	}
}

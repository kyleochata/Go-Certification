package main

import "testing"

/*
Indv components or units of a software are tested
unit - smallest testable part of any software
Go "unit test" = subset of "test".
Tests in software can take many forms such as unit, integration, functional, system etc.
Integration test = test the interaction b/w multiple components of the system to ensure they work together correctly.
*/
// func TestAdd(t *testing.T) {
// 	//part of the same package as Add() ==> able to see fxn eventhough it's on a different file.
// 	total := Add(5, 5)
// 	if total != 10 {
// 		t.Errorf("Sum incorrect. Expected: %v\t Output: %v\n", 10, total)
// 	}
// }

func TestDoMath(t *testing.T) {
	result := doMath(20, 20, add)
	if result != 40 {
		t.Errorf("Sum incorrect. Expected: %v\t Output: %v\n", 40, result)
	} else {
		t.Logf("Sum correct. Expected: %v\t Output: %v\n", 40, result)
	}
	result = doMath(15, 4, subtract)
	expected := 11
	if result != expected {
		t.Errorf("Difference incorrect. Expected: %v\t Output: %v\n", expected, result)
	} else {
		t.Logf("Difference correct. Expected: %v\t Output: %v\n", expected, result)
	}

}
func TestAdd(t *testing.T) {
	result := add(2, 10)
	expected := 12
	if result != expected {
		t.Errorf("Sum incorrect. Expected: %v\t Output: %v\n", expected, result)
	} else {
		t.Logf("Sum correct. Expected: %v\t Output: %v\n", expected, result)
	}
}
func TestSub(t *testing.T) {
	result := subtract(11, 1)
	expected := 10
	if result != expected {
		t.Errorf("Difference incorrect. Expected: %v\t Output: %v\n", expected, result)
	} else {
		t.Logf("Difference correct. Expected: %v\t Output: %v\n", expected, result)
	}
}

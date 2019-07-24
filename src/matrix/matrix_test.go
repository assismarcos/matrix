package main

import (
	"testing"
)

type testpair struct {
	values[][] string
	result string
}

var prodTest = []testpair{
	{[][]string{}, "0"},
	{[][]string{{"1","1"},{"1","1"}}, "1"},
	{[][]string {{"abc"}}, "Invalid input. Matrix does not contain valid integers"},
}

var sumTest = []testpair{
	{[][]string{}, "0"},
	{[][]string{{"1","1"},{"1","1"}}, "4"},
	{[][]string {{"abc"}}, "Invalid input. Matrix does not contain valid integers"},
}

var flatTest = []testpair{
	{[][]string{}, ""},
	{[][]string{{"1","1"},{"1","1"}}, "1,1,1,1"},
	{[][]string {{"1", "2", "3"}, {"4", "5", "6"}}, "1,2,3,4,5,6"},

}

var invertTest = []testpair{
	{[][]string{{}}, ""},
	{[][]string{{"1","1"},{"1","1"}}, "1,1\n"+"1,1\n"},
	{[][]string {{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8","9"}}, "1,4,7\n"+"2,5,8\n"+"3,6,9\n"},

}

func TestSum(t *testing.T){
	t.Logf("Testing sum of elements of matrix")
	for _, pair := range sumTest {
		i := sum(pair.values)
		if i !=  pair.result {
			t.Errorf("Sum was incorrect, got '%s', want '%s'", i, pair.result)
		}
	}
}

func TestMultiply(t *testing.T){
	t.Logf("Testing multiplication of elements of matrix")
	for _, pair := range prodTest {
		i := multiply(pair.values)
		if i !=  pair.result {
			t.Errorf("Product was incorrect, got %s, want %s", i, pair.result)
		}
	}
}

func TestFlatten(t *testing.T){
	t.Logf("Testing flatten elements of matrix")
	for _, pair := range flatTest {
		i := flatten(pair.values)
		if i != pair.result {
			t.Errorf("Flatten outut was incorrect, got %s, want %s", i, pair.result)
		}
	}
}

func TestInvert(t *testing.T){
	t.Logf("Testing inverting elements of matrix")
	for _, pair := range invertTest {
		i := invert(pair.values)
		if i != pair.result {
			t.Errorf("Invert outut was incorrect, got '%s', want '%s'", i, pair.result)
		}
	}
}


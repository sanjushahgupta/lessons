// white-box testing
package lesson_20201013

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		title  string // the name of the scenario
		x, y   int    // input
		expect int    // expected output
	}{
		{
			title:  "add 0 and 0",
			x:      0,
			y:      0,
			expect: 0,
		},
		{
			title:  "add small numbers",
			x:      7,
			y:      5,
			expect: 12,
		},
		{
			title:  "add big numbers",
			x:      1234,
			y:      1234,
			expect: 2468,
		},
		{
			title:  "add negative numbers",
			x:      -1,
			y:      -3,
			expect: -4,
		},
	}

	for _, test := range tests {

		t.Run(test.title, func(t *testing.T) {
			result := Add(test.x, test.y)
			if result != test.expect {
				t.Errorf("expected %d, got %d", test.expect, result)
			}
		})
	}
}

// func TestSomething
// TODO add tests

// testing for subtraction
func TestSubtract(t *testing.T) {

	test := []struct {
		about  string
		a, b   int
		expect int
	}{
		{
			about:  "subtract",
			a:      8,
			b:      5,
			expect: 3,
		},
		{about: "Subtract negative number",
			a:      -3,
			b:      -3,
			expect: 0,
		},
	}
	for _, test1 := range test {
		t.Run(test1.about, func(t *testing.T) {
			result := Subtract(test1.a, test1.b)
			if result != test1.expect {
				t.Errorf("expected %d, got %d", test1.expect, result)
			}
		})
	}
}

// TESTING FOR Multiplication

func TestMulti(t *testing.T) {
	test := []struct {
		about  string
		a, b   int
		expect int
	}{
		{
			about:  "multiply negative numbers",
			a:      -4,
			b:      -5,
			expect: 20,
		},
		{
			about:  "multiply postive number with neagative number",
			a:      8,
			b:      -8,
			expect: -64,
		},
	}
	for _, test1 := range test {

		t.Run(test1.about, func(t *testing.T) {
			result := Multi(test1.a, test1.b)
			if result != test1.expect {
				t.Errorf("expected %d, got %d", test1.expect, result)
			}
		})
	}
}

// EXTRA test coverage
// https://blog.golang.org/cover
// template
// t.Run(t.title, func(t *testing.T){
// 	the actual function you want to test
// })

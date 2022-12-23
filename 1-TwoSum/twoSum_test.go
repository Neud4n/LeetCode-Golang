package twosum

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	name     string
	input    []int
	target   int
	expected []int
}{
	{"Prueba1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
	{"Prueba2", []int{3, 2, 4}, 6, []int{1, 2}},
	{"Prueba3", []int{3, 3}, 6, []int{0, 1}},
	{"Prueba4", []int{2, 5, 5, 11}, 10, []int{1, 2}},
}

func TestTwoSum(t *testing.T) {

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			if !reflect.DeepEqual(twoSum(tc.input, tc.target), tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, twoSum(tc.input, tc.target))
			}
		})
	}

}

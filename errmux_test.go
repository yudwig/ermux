package ermux

import (
	"errors"
	"testing"
)

var master = []error{
	errors.New("error 1"),
	errors.New("error 2"),
	errors.New("error 3"),
}

type Truth struct {
	first  error
	last   error
	some   bool
	filter []error
}

type Dataset struct {
	errors []error
	truth  Truth
}

var tests = map[string]Dataset{
	"dataset 01": {
		errors: []error{nil, nil, nil},
		truth: Truth{
			first:  nil,
			last:   nil,
			some:   false,
			filter: []error{},
		},
	},
	"dataset 02": {
		errors: []error{master[0], nil, nil},
		truth: Truth{
			first:  master[0],
			last:   master[0],
			some:   true,
			filter: []error{master[0]},
		},
	},
	"dataset 03": {
		errors: []error{nil, master[1], nil},
		truth: Truth{
			first:  master[1],
			last:   master[1],
			some:   true,
			filter: []error{master[1]},
		},
	},
	"dataset 04": {
		errors: []error{master[0], master[1], nil},
		truth: Truth{
			first:  master[0],
			last:   master[1],
			some:   true,
			filter: []error{master[0], master[1]},
		},
	},
	"dataset 05": {
		errors: []error{nil, nil, master[2]},
		truth: Truth{
			first:  master[2],
			last:   master[2],
			some:   true,
			filter: []error{master[2]},
		},
	},
	"dataset 06": {
		errors: []error{master[0], nil, master[2]},
		truth: Truth{
			first:  master[0],
			last:   master[2],
			some:   true,
			filter: []error{master[0], master[2]},
		},
	},
	"dataset 07": {
		errors: []error{nil, master[1], master[2]},
		truth: Truth{
			first:  master[1],
			last:   master[2],
			some:   true,
			filter: []error{master[1], master[2]},
		},
	},
	"dataset 08": {
		errors: []error{master[0], master[1], master[2]},
		truth: Truth{
			first:  master[0],
			last:   master[2],
			some:   true,
			filter: []error{master[0], master[1], master[2]},
		},
	},
	"dataset 09": {
		errors: []error{},
		truth: Truth{
			first:  nil,
			last:   nil,
			some:   false,
			filter: []error{},
		},
	},
	"dataset 10": {
		errors: []error{nil, master[0], nil, master[1], nil, master[2], nil},
		truth: Truth{
			first:  master[0],
			last:   master[2],
			some:   true,
			filter: []error{master[0], master[1], master[2]},
		},
	},
}

func TestFirst(test *testing.T) {
	for name, data := range tests {
		test.Run(name, func(t *testing.T) {
			res := First(data.errors)
			if res != data.truth.first {
				t.Logf("res: %v", res)
				t.Logf("truth: %v", data.truth.first)
				t.Error(name)
			}
		})
	}
}

func TestLast(test *testing.T) {
	for name, data := range tests {
		test.Run(name, func(t *testing.T) {
			res := Last(data.errors)
			if res != data.truth.last {
				t.Logf("res: %v", res)
				t.Logf("truth: %v", data.truth.last)
				t.Error(name)
			}
		})
	}
}

func TestSome(test *testing.T) {
	for name, data := range tests {
		test.Run(name, func(t *testing.T) {
			res := Some(data.errors)
			if res != data.truth.some {
				t.Logf("res: %v", res)
				t.Logf("truth: %v", data.truth.some)
				t.Error(name)
			}
		})
	}
}

func TestFilter(test *testing.T) {
	for name, data := range tests {
		test.Run(name, func(t *testing.T) {
			res := Filter(data.errors)
			if !isEqual(res, data.truth.filter) {
				t.Logf("res: %v", res)
				t.Logf("truth: %v", data.truth.filter)
				t.Error(name)
			}
		})
	}
}

func isEqual(errs1 []error, errs2 []error) bool {
	if len(errs1) != len(errs2) {
		return false
	}
	for i := 0; i < len(errs1); i++ {
		if errs1[i] != errs2[i] {
			return false
		}
	}
	return true
}

package main

import "testing"

func testEqual(t *testing.T) {
	tests := []struct {
		pos1    positionCartesian
		pos2    positionCartesian
		expects bool
	}{
		{pos1: positionCartesian{x: 0, y: 0, z: 0}, pos2: positionCartesian{x: 0, y: 0, z: 0}, expects: true},
		{pos1: positionCartesian{x: 2, y: 3, z: 4}, pos2: positionCartesian{x: 2, y: 3, z: 4}, expects: true},
		{pos1: positionCartesian{x: 1, y: 0, z: 0}, pos2: positionCartesian{x: 0, y: 0, z: 0}, expects: false},
		{pos1: positionCartesian{x: 0, y: 1, z: 0}, pos2: positionCartesian{x: 0, y: 0, z: 0}, expects: false},
		{pos1: positionCartesian{x: 0, y: 0, z: 1}, pos2: positionCartesian{x: 0, y: 0, z: 0}, expects: false},
		{pos1: positionCartesian{x: 0, y: 0, z: 0}, pos2: positionCartesian{x: 1, y: 0, z: 0}, expects: false},
		{pos1: positionCartesian{x: 0, y: 0, z: 0}, pos2: positionCartesian{x: 0, y: 1, z: 0}, expects: false},
		{pos1: positionCartesian{x: 0, y: 0, z: 0}, pos2: positionCartesian{x: 0, y: 0, z: 1}, expects: false},
	}

	for _, test := range tests {

		if res := test.pos1.equal(test.pos2); res != test.expects {
			if test.expects == true {
				t.Errorf("%v and %v should match, but they don't", test.pos1, test.pos2)
			} else {
				t.Errorf("%v and %v shouldn't match, but they do", test.pos1, test.pos2)
			}
		}
	}
}

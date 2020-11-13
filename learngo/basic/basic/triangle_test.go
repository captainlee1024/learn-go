package main

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{a: 3,
			b: 4,
			c: 5},
		{a: 5,
			b: 12,
			c: 13},
		{a: 12,
			b: 35,
			c: 3},
		{a: 30000,
			b: 40000,
			c: 5000},
	}

	for _, tt := range tests {
		if angle := calcTriangle(tt.a, tt.b); angle != tt.c {
			t.Errorf("calcTriangle(%d %d);"+
				"got %d; expected %d", tt.a, tt.b, angle, tt.c)
		}
	}
}

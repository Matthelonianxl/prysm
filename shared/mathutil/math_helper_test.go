package mathutil

import (
	"testing"
)

func TestIntegerSquareRoot(t *testing.T) {
	tt := []struct {
		number uint64
		root   uint64
	}{
		{
			number: 20,
			root:   4,
		},
		{
			number: 200,
			root:   14,
		},
		{
			number: 1987,
			root:   44,
		},
		{
			number: 34989843,
			root:   5915,
		},
		{
			number: 97282,
			root:   311,
		},
	}

	for _, testVals := range tt {
		root := IntegerSquareRoot(testVals.number)
		if testVals.root != root {
			t.Fatalf("expected root and computed root are not equal %d, %d", testVals.root, root)
		}
	}
}

func TestCeilDiv8(t *testing.T) {
	tests := []struct {
		number int
		div8   int
	}{
		{
			number: 20,
			div8:   3,
		},
		{
			number: 200,
			div8:   25,
		},
		{
			number: 1987,
			div8:   249,
		},
		{
			number: 1,
			div8:   1,
		},
		{
			number: 97282,
			div8:   12161,
		},
	}

	for _, tt := range tests {
		div8 := CeilDiv8(tt.number)
		if tt.div8 != div8 {
			t.Fatalf("Div8 was not an expected value. Wanted: %d, got: %d", tt.div8, div8)
		}
	}
}

func TestIsPowerOf2(t *testing.T) {
	tests := []struct {
		a uint64
		b bool
	}{
		{
			a: 2,
			b: true,
		},
		{
			a: 64,
			b: true,
		},
		{
			a: 100,
			b: false,
		},
		{
			a: 1024,
			b: true,
		},
		{
			a: 0,
			b: false,
		},
	}
	for _, tt := range tests {
		if tt.b != IsPowerOf2(tt.a) {
			t.Fatalf("IsPowerOf2(%d) = %v, wanted: %v", tt.a, IsPowerOf2(tt.a), tt.b)
		}
	}
}

func TestPowerOf2(t *testing.T) {
	tests := []struct {
		a uint64
		b uint64
	}{
		{
			a: 3,
			b: 8,
		},
		{
			a: 20,
			b: 1048576,
		},
		{
			a: 11,
			b: 2048,
		},
		{
			a: 8,
			b: 256,
		},
	}
	for _, tt := range tests {
		if tt.b != PowerOf2(tt.a) {
			t.Fatalf("PowerOf2(%d) = %d, wanted: %d", tt.a, PowerOf2(tt.a), tt.b)
		}
	}
}

func TestClosestPowerOf2(t *testing.T) {
	tests := []struct {
		a uint64
		b uint64
	}{
		{
			a: 10,
			b: 8,
		},
		{
			a: 300,
			b: 256,
		},
		{
			a: 1200,
			b: 1024,
		},
		{
			a: 4500,
			b: 4096,
		},
	}
	for _, tt := range tests {
		if tt.b != ClosestPowerOf2(tt.a) {
			t.Fatalf("ClosestPowerOf2(%d) = %d, wanted: %d", tt.a, ClosestPowerOf2(tt.a), tt.b)
		}
	}
}

func TestMaxValue(t *testing.T) {
	tests := []struct {
		a      uint64
		b      uint64
		result uint64
	}{
		{
			a:      10,
			b:      8,
			result: 10,
		},
		{
			a:      300,
			b:      256,
			result: 300,
		},
		{
			a:      1200,
			b:      1024,
			result: 1200,
		},
		{
			a:      4500,
			b:      4096,
			result: 4500,
		},
		{
			a:      9999,
			b:      9999,
			result: 9999,
		},
	}
	for _, tt := range tests {
		if tt.result != Max(tt.a, tt.b) {
			t.Fatalf("Max(%d) = %d, wanted: %d", tt.a, Max(tt.a, tt.b), tt.result)
		}
	}
}

func TestMinValue(t *testing.T) {
	tests := []struct {
		a      uint64
		b      uint64
		result uint64
	}{
		{
			a:      10,
			b:      8,
			result: 8,
		},
		{
			a:      300,
			b:      256,
			result: 256,
		},
		{
			a:      1200,
			b:      1024,
			result: 1024,
		},
		{
			a:      4500,
			b:      4096,
			result: 4096,
		},
		{
			a:      9999,
			b:      9999,
			result: 9999,
		},
	}
	for _, tt := range tests {
		if tt.result != Min(tt.a, tt.b) {
			t.Fatalf("Min(%d) = %d, wanted: %d", tt.a, Min(tt.a, tt.b), tt.result)
		}
	}
}

func TestNextPowerOf2(t *testing.T) {
	tests := []struct {
		input int
		want int
	}{
		{input:0, want:0},
		{input:1, want:1},
		{input:2, want:2},
		{input:3, want:4},
		{input:5, want:8},
		{input:9, want:16},
		{input:20, want:32},
	}
	for _, tt := range tests {
			if got := NextPowerOf2(tt.input); got != tt.want {
				t.Errorf("NextPowerOf2() = %v, want %v", got, tt.want)
			}
	}
}

func TestPrevPowerOf2(t *testing.T) {
	tests := []struct {
		input int
		want int
	}{
		{input:0, want:0},
		{input:1, want:1},
		{input:2, want:2},
		{input:3, want:2},
		{input:5, want:4},
		{input:9, want:8},
		{input:20, want:16},
	}
	for _, tt := range tests {
		if got := PrevPowerOf2(tt.input); got != tt.want {
			t.Errorf("PrevPowerOf2() = %v, want %v", got, tt.want)
		}
	}
}
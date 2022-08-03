package math

import "testing"

func TestAdd(t *testing.T) {
	type AddData struct {
		x, y   int64
		result int64
	}

	data := []AddData{
		{1, 2, 3},
		{3, 4, 7},
		{1, -1, 0},
	}

	for _, dt := range data {
		result := Add(dt.x, dt.y)
		if result != dt.result {
			t.Errorf("Add(%d, %d) FAILED. expcted %d, got %d", dt.x, dt.y, dt.result, result)
		} else {
			t.Logf("Add(%d, %d) PASSED. expcted %d, got %d", dt.x, dt.y, dt.result, result)
		}
	}
}

func TestDivide(t *testing.T) {
	result := Divide(5, 0)

	if result != 0 {
		t.Errorf("Divide(5, 0) FAILED. expected %d, got %f", 0, result)
	} else {
		t.Logf("Divide(5, 0) PASSED. expected %d, got %f", 0, result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(3, 1)
	if result != 2 {
		t.Errorf("Subtract(3, 1) FAILED. expected %d, got %d", 2, result)
	} else {
		t.Logf("Subtract(3, 1) PASSED. expected %d, got %d", 2, result)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(5, 1)
	if result != 5 {
		t.Errorf("Multiply(5, 1) FAILED. expected %d, got %d", 5, result)
	} else {
		t.Logf("Multiply(5, 1) PASSED. expected %d, got %d", 5, result)
	}
}

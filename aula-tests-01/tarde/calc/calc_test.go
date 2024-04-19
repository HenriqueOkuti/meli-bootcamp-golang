package calc

import "testing"

func TestSum(t *testing.T) {
	result := Sum(5, 5)
	if result != 10 {
		t.Errorf("Sum(5, 5) = %d; want 10", result)
	}

	result = Sum(5, 0)
	if result != 5 {
		t.Errorf("Sum(5, 0) = %d; want 5", result)
	}

	result = Sum(0, 5)
	if result != 5 {
		t.Errorf("Sum(0, 5) = %d; want 5", result)
	}

	result = Sum(0, 0)
	if result != 0 {
		t.Errorf("Sum(0, 0) = %d; want 0", result)
	}

	result = Sum(-5, 5)
	if result != 0 {
		t.Errorf("Sum(-5, 5) = %d; want 0", result)
	}
}

func TestSub(t *testing.T) {
	result := Sub(5, 5)
	if result != 0 {
		t.Errorf("Sub(5, 5) = %d; want 0", result)
	}

	result = Sub(5, 0)
	if result != 5 {
		t.Errorf("Sub(5, 0) = %d; want 5", result)
	}

	result = Sub(0, 5)
	if result != -5 {
		t.Errorf("Sub(0, 5) = %d; want -5", result)
	}

	result = Sub(0, 0)
	if result != 0 {
		t.Errorf("Sub(0, 0) = %d; want 0", result)
	}

	result = Sub(-5, 5)
	if result != -10 {
		t.Errorf("Sub(-5, 5) = %d; want -10", result)
	}
}

func TestDivide(t *testing.T) {
	divide, err := Divide(5, 5)
	if divide != 1 || err != nil {
		t.Errorf("Divide(5, 5) = %d, %v; want 1, nil", divide, err)
	}

	divide, err = Divide(5, 0)
	if divide != 0 || err == nil {
		t.Errorf("Divide(5, 0) = %d, %v; want 0, error", divide, err)
	}

	divide, err = Divide(0, 5)
	if divide != 0 || err != nil {
		t.Errorf("Divide(0, 5) = %d, %v; want 0, nil", divide, err)
	}

}

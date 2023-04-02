package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.00
	expectedTax := 5.00

	result := CalculateTax(amount)

	if result != expectedTax {
		t.Errorf("Expected %f but go %f", expectedTax, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expectedTax float64
	}

	table := []calcTax{
		{100, 5},
		{500, 5},
		{1000, 10},
		{1500, 10},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)

		if result != item.expectedTax {
			t.Errorf("Expected %f but go %f", item.expectedTax, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500)
	}
}

func BenchmarkCalculateTaxWithSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTaxWithSleep(500)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -3.5, 0, 500.0, 1000.0, 1500.0}

	for _, amount := range seed {
		f.Add(amount)
	}

	f.Fuzz(
		func(t *testing.T, amount float64) {
			result := CalculateTax(amount)

			if amount <= 0 && result != 0 {
				t.Errorf("Expected 0 but got %f", result)
			}

			if amount >= 20000 && result != 20 {
				t.Errorf("Expected 20 but got %f", result)
			}
		},
	)
}

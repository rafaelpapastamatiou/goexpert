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

package pricing

import (
	"testing"
)

// TestBasicCalculation tests the happy path with all rules applied
func TestBasicCalculation(t *testing.T) {
	order := Order{
		BasePrice:    1000,
		CountryCode:  "TH",
		CustomerType: "regular",
		IsFirstOrder: true,
	}

	rules := PricingRules{
		Taxes:              map[string]float64{"TH": 0.07},
		FirstOrderDiscount: 100,
		CustomerDiscounts:  map[string]float64{"regular": 0},
	}

	result := CalculateFinalPrice(order, rules)
	expected := 970.00 // 1000 × 1.07 = 1070 → 1070 - 100 = 970

	if result != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, result)
	}
}

// TestDiscountOrder verifies customer discount is applied before first order discount
func TestDiscountOrder(t *testing.T) {
	order := Order{
		BasePrice:    1000,
		CountryCode:  "TH",
		CustomerType: "vip",
		IsFirstOrder: true,
	}

	rules := PricingRules{
		Taxes:              map[string]float64{"TH": 0.07},
		FirstOrderDiscount: 100,
		CustomerDiscounts:  map[string]float64{"vip": 0.10},
	}

	result := CalculateFinalPrice(order, rules)
	expected := 863.00 // 1000 × 1.07 = 1070 → 1070 × 0.9 = 963 → 963 - 100 = 863

	if result != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, result)
	}
}

// TestNegativePricePrevention ensures final price cannot be negative
func TestNegativePricePrevention(t *testing.T) {
	order := Order{
		BasePrice:    50,
		CountryCode:  "TH",
		CustomerType: "vip",
		IsFirstOrder: true,
	}

	rules := PricingRules{
		Taxes:              map[string]float64{"TH": 0.07},
		FirstOrderDiscount: 100,
		CustomerDiscounts:  map[string]float64{"vip": 0.10},
	}

	result := CalculateFinalPrice(order, rules)
	expected := 0.00 // 50 × 1.07 = 53.5 → 53.5 × 0.9 = 48.15 → 48.15 - 100 = -51.85 → 0

	if result != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, result)
	}
}

// TestDifferentCountries verifies tax rates are applied correctly per country
func TestDifferentCountries(t *testing.T) {
	rules := PricingRules{
		Taxes:              map[string]float64{"TH": 0.07, "FR": 0.20},
		FirstOrderDiscount: 0,
		CustomerDiscounts:  map[string]float64{"regular": 0},
	}

	// Test TH
	orderTH := Order{
		BasePrice:    1000,
		CountryCode:  "TH",
		CustomerType: "regular",
		IsFirstOrder: false,
	}
	resultTH := CalculateFinalPrice(orderTH, rules)
	expectedTH := 1070.00 // 1000 × 1.07

	if resultTH != expectedTH {
		t.Errorf("TH: Expected %.2f, got %.2f", expectedTH, resultTH)
	}

	// Test FR
	orderFR := Order{
		BasePrice:    1000,
		CountryCode:  "FR",
		CustomerType: "regular",
		IsFirstOrder: false,
	}
	resultFR := CalculateFinalPrice(orderFR, rules)
	expectedFR := 1200.00 // 1000 × 1.20

	if resultFR != expectedFR {
		t.Errorf("FR: Expected %.2f, got %.2f", expectedFR, resultFR)
	}
}

// TestRounding verifies proper rounding to 2 decimal places
func TestRounding(t *testing.T) {
	order := Order{
		BasePrice:    100,
		CountryCode:  "TH",
		CustomerType: "vip",
		IsFirstOrder: false,
	}

	rules := PricingRules{
		Taxes:              map[string]float64{"TH": 0.07},
		FirstOrderDiscount: 0,
		CustomerDiscounts:  map[string]float64{"vip": 0.10},
	}

	result := CalculateFinalPrice(order, rules)
	expected := 96.30 // 100 × 1.07 = 107 → 107 × 0.9 = 96.3

	if result != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, result)
	}
}

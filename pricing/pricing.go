package pricing

import "math"

type Order struct {
	BasePrice    float64 // Base price (e.g., 1000)
	CountryCode  string  // Country code (e.g., "TH")
	CustomerType string  // Customer type (e.g., "vip")
	IsFirstOrder bool    // First order flag (e.g., true)
}

type PricingRules struct {
	Taxes              map[string]float64 // Tax rates by country (e.g., {"TH": 0.07})
	FirstOrderDiscount float64            // Fixed first order discount (e.g., 100)
	CustomerDiscounts  map[string]float64 // Discount rates by customer type (e.g., {"vip": 0.10})
}

// CalculateFinalPrice calculates the final price with tax and discounts
// Rules applied in order:
// 1. Apply tax based on country
// 2. Apply customer discount after tax
// 3. Apply first order discount after customer discount
// 4. Ensure final price cannot be negative
// 5. Round to 2 decimals
func CalculateFinalPrice(order Order, rules PricingRules) float64 {
	// Step 1: Apply tax based on country (e.g., 1000 × 1.07 = 1070)
	taxRate := rules.Taxes[order.CountryCode]
	price := order.BasePrice * (1 + taxRate)

	// Step 2: Apply customer discount after tax (e.g., 1070 × 0.9 = 963 for vip)
	customerDiscountRate := rules.CustomerDiscounts[order.CustomerType]
	price = price * (1 - customerDiscountRate)

	// Step 3: Apply first order discount (e.g., 1070 - 100 = 970)
	if order.IsFirstOrder {
		price = price - rules.FirstOrderDiscount
	}

	// Step 4: Ensure final price cannot be negative
	price = math.Max(0, price)

	// Step 5: Round to 2 decimals
	return roundToTwoDecimals(price)
}

// roundToTwoDecimals rounds a float64 to 2 decimal places
// Formula: Round(value × 100) / 100
func roundToTwoDecimals(value float64) float64 {
	return math.Round(value*100) / 100
}

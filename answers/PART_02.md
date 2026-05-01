# Part 2 – Code Review

## AI-Generated Code

```javascript
function calculateFinalPrice(order) {
  let price = order.basePrice;

  if (order.countryCode === "TH") {
    price = price + price * 0.07;
  } else {
    price = price + price * 0.20;
  }

  if (order.customerType === "vip") {
    price = price - 100;
  }

  if (order.isFirstOrder) {
    price = price * 0.9;
  }

  return price.toFixed(2);
}
```

## 1. Bugs & Specification Mismatches

- **Missing `pricingRules` parameter**: Tax rates hardcoded (violates rule #7)
- **Wrong discount order**: VIP discount before first order (spec requires opposite)
- **Wrong discount types**: First order uses % (should be fixed 100), VIP uses fixed (should be 10%)
- **Returns string**: `toFixed(2)` returns string, spec requires number
- **No negative protection**: Can return negative prices
- **Hardcoded fallback**: Non-TH countries default to 20% tax

## 2. Production Risks

- **Silent failures**: Missing countries default to wrong tax rate
- **Type mismatch**: String return breaks downstream calculations
- **Inflexible**: New countries/customer types require code changes
- **Wrong rounding**: `toFixed(2)` rounds differently than spec

## 3. Refactoring Approach

Already implemented in Part 1 (`pricing/pricing.go`):
- Accept `pricingRules` parameter
- Correct order: tax → customer% → firstOrder fixed
- Non-negative guarantee with `math.Max(0, price)`
- Return number type (`float64`)
- Proper rounding: `Math.round(value*100)/100`

## 4. First Test to Write

```go
func TestCorrectDiscountOrder(t *testing.T) {
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
    
    // Expected: 1000 × 1.07 = 1070 → 1070 × 0.9 = 963 → 963 - 100 = 863
    assert.Equal(t, 863.00, result)
}
```

**Why this test?**
- Catches the most critical bug: wrong discount order
- Verifies all pricing rules interact correctly
- Would fail with AI code
- High confidence in core logic if it passes

# Part 1 – Implementation

## Design Decisions

### Type System
Used Go structs (`Order` and `PricingRules`) for type safety and clear function signature.

### Implementation Approach
Straightforward sequential calculation following the specification exactly:
1. Apply tax → 2. Customer discount → 3. First order discount → 4. Ensure non-negative → 5. Round

### Key Choices
- **No hardcoded values**: All rates come from `PricingRules`
- **Pure function**: No side effects, easy to test
- **Helper function**: `roundToTwoDecimals` for clarity
- **Idiomatic Go**: Used `math.Max` and `math.Round`

## Code

See: `pricing/pricing.go`

## Example Calculation

```
Input: basePrice=1000, country=TH, type=regular, firstOrder=true

1. Tax (7%):           1000 × 1.07 = 1070
2. Customer discount:  1070 × (1-0) = 1070  (regular = 0%)
3. First order:        1070 - 100 = 970
4. Non-negative:       max(0, 970) = 970
5. Round:              970.00

Result: 970.00
```

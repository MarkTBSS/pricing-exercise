# Part 3 – Testing Strategy

## Selected Tests

### 1. TestBasicCalculation
**Scenario:** Happy path with regular customer, first order, single country  
**Why selected:** Validates the complete flow works end-to-end  
**Catches:** Integration issues, basic calculation errors

### 2. TestDiscountOrder
**Scenario:** VIP customer with first order discount  
**Why selected:** Most critical business logic - wrong order = wrong price  
**Catches:** Discount sequence bugs (customer% before firstOrder fixed)

### 3. TestNegativePricePrevention
**Scenario:** Small base price with large discounts  
**Why selected:** Edge case that could cause negative prices  
**Catches:** Missing `math.Max(0, price)` protection

### 4. TestDifferentCountries
**Scenario:** Same order in TH (7%) vs FR (20%)  
**Why selected:** Verifies tax rates come from rules, not hardcoded  
**Catches:** Hardcoded tax values, wrong tax lookup

### 5. TestRounding
**Scenario:** Price that produces decimal places  
**Why selected:** Ensures proper 2-decimal rounding  
**Catches:** Incorrect rounding logic, floating-point precision issues

## Cases Intentionally Not Tested

### Missing Country Code
**Why:** Spec says "input validation intentionally excluded" - focus on core logic, not defensive coding

### Invalid Customer Type
**Why:** Same reason - not testing edge cases outside core pricing logic

### Nil/Empty PricingRules
**Why:** Assumes valid input per spec instructions

### Multiple Decimal Edge Cases
**Why:** One rounding test is sufficient - Go's `math.Round` is well-tested

## Confidence for Production Release

### Required
1. ✅ All 5 tests pass
2. ✅ Code review approved
3. ✅ Manual calculation verification matches test expectations

### Additional Confidence
- Run tests against real pricing data samples
- Deploy to staging with production-like data
- Monitor first few transactions closely
- Have rollback plan ready

### Why This Is Enough
- Tests cover all critical paths (tax, discounts, edge cases)
- Spec explicitly excludes defensive coding
- Simple, pure function with no side effects
- Easy to verify manually

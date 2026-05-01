# Part 7 – Coding Style & Trade-offs

## 1. How Did You Structure Your Code?

**Structure:**
- **Data types** (`Order`, `PricingRules`) - separate from logic
- **Pure function** (`CalculateFinalPrice`) - takes input, returns output, no side effects
- **Helper function** (`roundToTwoDecimals`) - single responsibility

**Key characteristics:**
- Sequential steps (tax → discount → discount → validate → round)
- Clear separation: data vs logic
- No state, no mutations
- Easy to test

---

## 2. OOP, Functional, or Mix?

**Functional style with Go structs**

**Functional aspects:**
- Pure function (no side effects)
- Immutable inputs (doesn't modify Order or PricingRules)
- Composable (could extract each step as separate function)
- Predictable (same input = same output)

**Not OOP because:**
- No methods on structs
- No inheritance
- No encapsulation of behavior
- Structs are just data containers

**Why not full OOP:**
```go
// Didn't do this (OOP style):
type PricingCalculator struct {
    rules PricingRules
}
func (p *PricingCalculator) Calculate(order Order) float64 { ... }
```

---

## 3. Why Did You Choose This Approach?

**Reasons:**

**1. Simplicity**
- Spec says "keep the solution simple and production-minded"
- Pure function is simplest approach

**2. Testability**
- No setup needed
- No mocks needed
- Just call function with different inputs

**3. Predictability**
- No hidden state
- No side effects
- Easy to reason about

**4. Go conventions**
- Go favors composition over inheritance
- Simple functions over complex objects
- "Accept interfaces, return structs"

**5. Spec requirements**
- "Focus on core logic, not defensive coding"
- Pure function fits this perfectly

---

## 4. If Pricing Rules Become Much More Complex?

**Current approach works until:**
- 10+ different discount types
- Complex conditional logic (if A and (B or C) then D)
- Rules depend on external data (inventory, time of day)
- Need to audit which rules were applied

**Evolution path:**

**Phase 1: Extract steps (still functional)**
```go
price = applyTax(order, rules)
price = applyCustomerDiscount(price, order, rules)
price = applyFirstOrderDiscount(price, order, rules)
price = applyConditionalDiscounts(price, order, rules)
```
- Still pure functions
- Easier to test each step
- Clear order of operations

**Phase 2: Rule chain pattern**
- Each rule is independent function
- Chain rules together
- Easy to add/remove rules
- Still functional

**Phase 3: Rule engine (if really complex)**
- Config-driven rules
- Rule evaluation engine
- Audit trail of applied rules
- But only if complexity justifies it

**Would NOT move to OOP** because:
- Doesn't solve complexity problem
- Adds boilerplate
- Harder to test
- Go doesn't favor it

---

## Summary

**Current approach:**
- Functional style with Go structs
- Pure function, no side effects
- Simple, testable, predictable

**Why:**
- Matches spec requirements
- Go idioms
- Simplest solution that works

**Evolution:**
- Extract steps → Rule chain → Rule engine (only if needed)
- Stay functional, avoid OOP complexity

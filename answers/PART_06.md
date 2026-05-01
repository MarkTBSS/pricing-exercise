# Part 6 – Scaling the Design

## Scenario
Support 10 more countries next month (12 total). Keep it simple.

---

## 1. Would You Keep the Current Design?

**Yes.** The core logic is solid and doesn't need to change:
- `PricingRules` already uses maps (flexible for any number of countries)
- Pure function, easy to test
- No hardcoded values

**What works:**
```go
Taxes: map[string]float64  // Can hold 2 or 200 countries
```

---

## 2. What Would You Change?

**Move rules from code to external config:**

**Current:** Rules passed as parameter (good), but created in code (not scalable)

**Change:** Load rules from config file at startup

**Why:**
- Adding countries doesn't require code changes
- Non-developers can update tax rates
- Version controlled (Git tracks changes)

---

## 3. Configuration, Database, or Code?

**Start with: JSON configuration file**

**Example: `pricing_rules.json`**
```json
{
  "taxes": {
    "TH": 0.07,
    "FR": 0.20,
    "US": 0.08,
    "JP": 0.10,
    "SG": 0.07
  },
  "firstOrderDiscount": 100,
  "customerDiscounts": {
    "vip": 0.10,
    "regular": 0
  }
}
```

**Why config file (not database):**
- ✅ Simple to implement
- ✅ Fast (load once, cache in memory)
- ✅ Version controlled
- ✅ Easy to review changes
- ✅ No database dependency

**When to use database:**
- Rules change multiple times per day
- Need per-customer custom pricing
- Need audit trail with timestamps

**For 10 countries:** Config file is enough.

---

## 4. Complex Rule: TH + VIP + Order > 5000

**Scenario:** Extra 5% discount for VIP in Thailand with orders > 5000

**Approach: Add conditional rules to config**

```json
{
  "conditionalDiscounts": [
    {
      "name": "TH VIP Large Order",
      "conditions": {
        "country": "TH",
        "customerType": "vip",
        "minAmount": 5000
      },
      "discount": 0.05
    }
  ]
}
```

**How it works:**
1. After applying standard discounts
2. Check if order matches all conditions
3. If yes, apply additional discount

**Code change needed:**
- Add loop to check conditional rules
- Apply matching discounts

**Limitation:** Can't handle very complex logic (e.g., "A OR B", nested conditions)
- For that, need rule engine (but spec says keep it simple)

---

## Summary

**Simplest approach:**
1. Keep current `CalculateFinalPrice` logic (no changes)
2. Move `PricingRules` to JSON config file
3. Load config at startup, cache in memory
4. Add conditional discount support for complex rules

**Result:** Adding 10 countries = updating JSON file (no code changes)

# Part 8 – AI Usage

## Did You Use AI?

**Yes**

---

## 1. What Did AI Help With?

**Speed of writing small functions:**
- AI helped generate small function structures quickly
- I used it to create and check function logic
- I manually copy-pasted the code myself (never let AI edit code directly)

**Reduced human error:**
- AI doesn't get tired or lazy
- Helps catch mistakes from fatigue
- Consistent output quality

---

## 2. What Did You Verify Yourself?

**Code correctness:**
- Ran `go build ./...` and `go test ./...` to verify everything works
- Manually calculated expected results (e.g., 1000 × 1.07 - 100 = 970)
- Read through all generated code line by line

**Specification compliance:**
- Re-read exercise requirements multiple times
- Verified all 7 rules in Part 1 are implemented correctly
- Checked discount order matches specification exactly

**Logic verification:**
- Traced through calculation steps manually
- Tested edge cases (negative prices, rounding)

---

## 3. One Example Where AI Could Be Wrong

**AI can hallucinate** - I understand this well because we're using generative AI.

**"Generative" means creative, not deterministic.**

**Most common mistake: Missing context**
- AI doesn't understand the full context
- Doesn't see all related components
- Doesn't know all related functions
- May suggest code that looks correct but doesn't fit the system

**Example in this exercise:**
- AI might suggest wrong discount order (first order before customer discount)
- AI might not understand the relationship between tax, discounts, and rounding
- AI might miss that specification explicitly requires specific sequence

**How I prevent this:**
- Always verify against specification
- Manually trace through logic
- Test with real examples
- Never blindly trust AI output

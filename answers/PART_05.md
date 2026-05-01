# Part 5 – Production Incident

## Scenario
Pricing logic was wrong in production for 2 hours. Some customers were overcharged.

---

## 1. What Do You Do First?

1. **Stop the bleeding** - Rollback to previous version immediately
2. **Verify rollback** - Confirm pricing is correct again
3. **Preserve evidence** - Capture logs, database snapshots, time window

**Priority:** Prevent more damage, investigate later

---

## 2. Who Do You Inform?

**Immediate (within 15 min):**
- Engineering manager
- On-call team lead
- Customer support team

**Within 1 hour:**
- Product manager
- Finance team (refund budget)
- Legal/Compliance (if required)

**After investigation:**
- Affected customers (apology + refund plan)

---

## 3. How Do You Identify Affected Orders?

**Query:**
```sql
SELECT * FROM orders 
WHERE created_at BETWEEN '2026-05-01 10:00' AND '2026-05-01 12:00'
  AND status = 'completed'
```

**Verification:**
1. Recalculate price using correct logic
2. Compare with charged amount
3. Calculate overcharge per order
4. Have another engineer verify the query

---

## 4. How Do You Fix the Issue Safely?

**Refund process:**
1. Generate refund list with amounts
2. Get approval from finance/manager
3. Test with 1-2 orders first
4. Process in batches, monitor success rate

**Customer communication:**
- Email affected customers
- Brief explanation + refund timeline
- Apologize

---

## 5. How Do You Prevent It from Happening Again?

**Immediate:**
- Add monitoring alert for price anomalies
- Add test for the bug that caused this

**Short-term:**
- Canary deployment (1% traffic first)
- Require 2+ engineer code review for pricing changes
- Add price comparison check (new vs old logic)

**Long-term:**
- Incident review meeting (no blame)
- Feature flag for pricing logic (instant rollback)
- Pricing validation in CI/CD pipeline

---

## 6. How Would You Have Deployed This Change?

**Safe deployment strategy:**

**Canary → Gradual rollout → Full**
- Start with 1% traffic for 1 hour
- Monitor: error rate, price distribution, complaints
- Gradually increase: 10% → 50% → 100% over 6-12 hours
- Auto-rollback if anomaly detected

**Additional safeguards:**
- Feature flag (disable instantly without deployment)
- Price sanity check (alert if price > 2x or < 0.5x expected)
- Shadow mode (compare new vs old logic without affecting customers)

**Result:** Bug caught in canary phase (1% affected instead of 100%)

# Part 4 – External Payment API

## 1. Boundary/Interface

**Approach:** Create abstraction layer with standard methods (`Charge`, `Refund`, `GetStatus`)

**Why:**
- Swap providers (Stripe → Omise) without changing business logic
- Mock for testing without real API calls
- Isolate external dependencies

## 2. Prevent Double Charging

**Idempotency Key:** Unique identifier per payment attempt (e.g., `orderID-timestamp`)

**Flow:**
1. Check database: if key exists → return existing result
2. Store key before calling provider
3. Provider detects duplicate key → returns original result (no double charge)

## 3. Data to Store Before Calling Provider

**Payment Intent Record:**
- Order ID
- Idempotency key
- Amount
- Status (pending/success/failed)
- Timestamp

**Why store BEFORE calling:**
- Audit trail
- Prevent duplicates
- Recover from timeout (can query provider later with key)
- Show customer "payment pending" status

## 4. Unit Tests with Fake Provider

**Test scenarios:**
1. Successful payment → verify status updated
2. Provider failure → verify error handling
3. Timeout → verify stays "pending"
4. Duplicate key → verify provider called only once
5. Retry logic → verify exponential backoff

**Fake provider benefits:**
- No real API calls
- Simulate failures/timeouts
- Fast, deterministic tests

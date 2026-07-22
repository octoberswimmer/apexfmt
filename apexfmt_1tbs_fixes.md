# apexfmt — 1TBS / Formatting Fix Specification

**Repo:** `github.com/octoberswimmer/apexfmt`
**Verified against:** `master` @ `b7843d2` (July 2026)
**Context:** apexfmt is gofmt-inspired and aims for 1TBS (cuddled opening braces, mandatory braces on single-statement bodies, one statement per line, closing brace on its own line). The issues below are the places it currently deviates. All were confirmed by building the binary and running it against minimal repro inputs; line numbers were verified with `grep -n` against the commit above.

Each fix below is **atomic**: independently implementable, independently testable, and mergeable in any order. FIX-1 and FIX-2 restore idempotency and should be prioritized. FIX-5 and FIX-6 are design decisions, not bugs — implement only if the maintainer confirms the desired behavior.

All formatter tests live in `formatter/format_test.go` (table-driven `{input, expected}` pairs). Add repro cases there. Run with:

```sh
cd formatter && go test ./...
```

---

## FIX-1: Brace-less `else` body — closing brace cuddled to statement

**Severity:** Bug (breaks 1TBS, breaks idempotency)
**File:** `formatter/visitors.go`
**Function:** `VisitIfStatement`
**Line:** 233

### Current code

```go
out.WriteString(fmt.Sprintf(" else {\n%s}", indent(v.visitRule(ctx.Statement(1)).(string))))
```

The format string is missing a `\n` before the closing `}`.

### Repro

Input:

```apex
public class T {
	public void run(Boolean x) {
		if (x) doThing(); else doOther();
	}
}
```

Actual output (body):

```apex
if (x) {
	doThing();
} else {
	doOther();}
```

Expected output:

```apex
if (x) {
	doThing();
} else {
	doOther();
}
```

### Fix

Change the format string to `" else {\n%s\n}"`, matching the correct sibling branch at lines 224–225 (`"if %s {\n%s\n}"`).

### Acceptance criteria

- Closing `}` of a brace-inserted `else` block is on its own line at the enclosing indent level.
- Output is idempotent: formatting the output again produces identical text.
- Existing tests pass (the current test suite only exercises already-braced `else` inputs, which take the `Block()` path at line 229 and are unaffected).

### Test to add

Table entry in `format_test.go`: input `if (x) doThing(); else doOther();` (inside a class/method wrapper matching the existing test harness), expected output with both branches braced and each `}` on its own line.

---

## FIX-2: Brace-less `while` body — no indentation and cuddled closing brace

**Severity:** Bug (breaks 1TBS, breaks idempotency)
**File:** `formatter/visitors.go`
**Function:** `VisitWhileStatement`
**Line:** 246

### Current code

```go
return fmt.Sprintf("while %s {\n%s}", v.visitRule(ctx.ParExpression()), v.visitRule(ctx.Statement()))
```

Two defects in one format string:
1. Missing `indent(...)` around the statement — the body renders at the same level as the `while` keyword (actually shallower once the enclosing block indents everything else).
2. Missing `\n` before the closing `}`.

### Repro

Input:

```apex
while (x) spin();
```

Actual output:

```apex
		while (x) {
		spin();}
```

Expected output:

```apex
		while (x) {
			spin();
		}
```

### Fix

```go
return fmt.Sprintf("while %s {\n%s\n}", v.visitRule(ctx.ParExpression()), indent(v.visitRule(ctx.Statement()).(string)))
```

This matches the pattern used by `VisitIfStatement`'s non-block branch (lines 224–225) and `VisitDoWhileStatement` (line ~254).

### Acceptance criteria

- Brace-less `while` body is indented one level relative to `while`.
- Closing `}` is on its own line at the `while` indent level.
- Idempotent round-trip.
- No change for already-braced `while` bodies (line 245 `Block()` path).

### Test to add

Input `while (x) spin();`; also a variant with a wrapped SOQL condition (`while (![SELECT ...].isEmpty()) drain();`) since that combination was observed broken in testing.

---

## FIX-3: Brace-less `for` body — stray trailing newline creates blank line

**Severity:** Cosmetic bug
**File:** `formatter/visitors.go`
**Function:** `VisitForStatement`
**Line:** 263

### Current code

```go
return fmt.Sprintf("for (%s) {\n%s\n}\n", v.visitRule(ctx.ForControl()), indent(v.visitRule(ctx.Statement()).(string)))
```

The trailing `\n` after `}` is wrong. Statements are joined with `"\n"` by the enclosing block visitor (`VisitBlock`, line 56), so any statement that emits its own trailing newline produces a doubled newline — a blank line after the loop.

### Repro

Input:

```apex
for (Integer i = 0; i < 10; i++) step(i);
if (x) other();
```

Actual output has a blank line between `}` and `if (x) {`. Expected output has none. Same artifact appears after brace-less SOQL for loops (`for (Account a : [SELECT ...]) process(a);`).

### Fix

Remove the trailing `\n`: `"for (%s) {\n%s\n}"`.

### Acceptance criteria

- No blank line between a brace-inserted `for` loop and the following statement.
- Braced `for` loops (line 261 path) unaffected.
- Idempotent round-trip.

### Test to add

Input with a brace-less `for` followed immediately by another statement; assert exactly one `\n` between `}` and the next statement.

---

## FIX-4: Empty-accessor property flattening uses fragile length check and inconsistent spacing

**Severity:** Code-quality bug (works today, semantically fragile) + style inconsistency
**File:** `formatter/visitors.go`
**Function:** `VisitPropertyDeclaration`
**Lines:** 160–161

### Current code

```go
// Flatten empty getter/setter
if len(strings.Join(propertyBlocks, "")) == 8 {
	return fmt.Sprintf("%s %s {%s}", v.visitRule(ctx.TypeRef()), ctx.Id().GetText(), strings.Join(propertyBlocks, " "))
}
```

Problems:

1. **Fragile detection.** "The joined blocks are exactly 8 characters" is a proxy for "the blocks are exactly `get;` + `set;`". It works by accident: it also matches a lone hypothetical 8-char block, and it *misses* legitimate flatten candidates like `{ get; }` (4 chars) or `{ set; }`, and any accessor with a visibility modifier (`private set;` — 12+ chars), so `public Integer other { get; private set; }` expands to four lines while `{ get; set; }` flattens. Whether modifier'd accessors *should* flatten is a style call (see design note), but the detection should at minimum be explicit, not length-based.
2. **Spacing inconsistency.** Output is `{get; set;}` — no inner padding — while single-line enums render `{ RED, GREEN, BLUE }` with padding, and inline collection literals render `{ a, b }` with padding.

### Repro

Input:

```apex
public Integer counter { get; set; }
public Integer other { get; private set; }
```

Actual output:

```apex
public Integer counter {get; set;}
public Integer other {
	get;
	private set;
}
```

### Fix

Replace the length check with an explicit predicate: flatten iff **every** property block is a bare `get;` or `set;` (no body, no modifiers). Suggested shape:

```go
allEmpty := len(propertyBlocks) > 0
for _, b := range propertyBlocks {
	if b != "get;" && b != "set;" {
		allEmpty = false
		break
	}
}
if allEmpty {
	return fmt.Sprintf("%s %s { %s }", v.visitRule(ctx.TypeRef()), ctx.Id().GetText(), strings.Join(propertyBlocks, " "))
}
```

Note the added inner padding (`{ get; set; }`) for consistency with enum/collection rendering. **The padding change alters existing expected output** — grep `format_test.go` for `{get; set;}` and update expectations in the same commit. If the maintainer prefers zero behavior change, split this into 4a (predicate refactor, byte-identical output) and 4b (padding, expectation updates).

### Acceptance criteria

- `{ get; set; }` flattens (with padding); `{ get; }` and `{ set; }` alone flatten too (new behavior — confirm with maintainer, or exclude and keep pairs-only).
- Any accessor with a body or modifier expands to multiline, as today.
- Detection no longer depends on string length.

### Test to add

Cases: `{ get; set; }`, `{ get; }`, `{ set; }`, `{ get; private set; }`, `{ get { return x; } set; }`.

---

## FIX-5 (design decision): SOQL wrap trigger ignores line length entirely

**Severity:** Design gap, not a bug — confirm intent before implementing
**Files:** `formatter/visitors.go` (`VisitSoqlLiteral`, line 751; `VisitQuery`, line 761), `formatter/chain.go` (`VisitQuery` scoring, line 83)

### Current behavior

A SOQL literal wraps iff its `ChainVisitor` complexity score exceeds 3. Scoring: 1 per SELECT field (5 for a subquery, 3 for `TYPEOF`), 1 per FROM entity, WHERE = condition count + `AND`/`OR` count, 1 per ORDER BY field, +1 each for LIMIT / OFFSET / USING SCOPE / FOR UPDATE. Raw line length is never consulted.

### Consequence (verified)

- `[SELECT Id, Name, Type, Phone FROM Account]` (44 chars, score 5) explodes to 10 lines.
- A ~210-char single-field query with one WHERE condition (score 3) stays on one line, unboundedly long.

This is inconsistent with the rest of the formatter, where wrapping is length-triggered (`VisitFormalParameters` line ~1561: wrap if `len > 40 && params > 2`, or `len > 60`; similar thresholds in expression-list and binary-expression visitors at lines 569 and 659).

### Proposed fix (pending maintainer decision on thresholds)

In `VisitSoqlLiteral` and `VisitQuery`, make the wrap decision `n > 3 || len(ctx.GetText()) > L` for some threshold `L` consistent with the rest of the codebase (the existing precedents are 40/60/150; SOQL likely wants the higher end, e.g. 80–120). Keep the complexity trigger — it produces good output for structurally rich queries — and add length as a second trigger so no query line grows without bound.

Do **not** change the score-based *layout* (one clause per line); only the *trigger*.

### Acceptance criteria

- Long-but-simple queries wrap once past the agreed threshold.
- Short complex queries continue to wrap as today (no regressions in existing SOQL tests).
- Wrapped output remains idempotent (score and length are stable across round-trips).

### Test to add

A score-3 query longer than the threshold (must wrap after the fix); a score-3 short query (must stay inline); a score-5 short query (must wrap, as today).

---

## FIX-6 (design decision): Nested SOQL subquery closing paren cuddles to last token

**Severity:** Style inconsistency — confirm intent before implementing
**File:** `formatter/visitors.go` (subquery rendering within the SELECT-entry / subquery visitor path)

### Current behavior (verified)

Outer wrapped query gives the closing `]` its own line, but a wrapped subquery cuddles `)` onto the final token:

```apex
for (Account a : [
	SELECT
		Id,
		(SELECT
			Id
		FROM
			Contacts)
	FROM
		Account
]) {
```

### Proposed fix

For symmetry with the outer `[` / `]`, place the subquery's closing `)` on its own line aligned with the opening `(`:

```apex
		(SELECT
			Id
		FROM
			Contacts
		)
```

(Alternatively, the maintainer may prefer the current Lisp-style close as intentional; if so, document it and close this item.)

### Acceptance criteria

- Consistent closing-delimiter treatment between outer query brackets and subquery parens.
- Idempotent round-trip.
- Update any existing subquery expectations in `format_test.go` / `soql_test.go` in the same commit.

---

## Cross-cutting verification (run after any fix lands)

1. **Idempotency sweep:** for every test input, assert `format(format(x)) == format(x)`. FIX-1 and FIX-2 are the two known violations today; a general idempotency assertion in the test harness would prevent regressions. This is cheap to add as a wrapper around the existing table-driven tests and is the single highest-value test change in this document.
2. **Brace audit:** grep `visitors.go` for `fmt.Sprintf` format strings containing `{\n%s}` (missing pre-`}` newline) — after FIX-1/FIX-2 there should be zero matches:

   ```sh
   grep -n '{\\n%s}' formatter/visitors.go   # expect no output
   ```
3. Full suite: `go test ./...` from the repo root.

## Suggested commit sequence

Each fix is one commit with its test(s):

1. `fix: put closing brace of brace-inserted else on its own line` (FIX-1)
2. `fix: indent brace-inserted while body and close brace on own line` (FIX-2)
3. `fix: remove stray blank line after brace-inserted for loop` (FIX-3)
4. `refactor: detect empty property accessors explicitly` (FIX-4, optionally split 4a/4b)
5. `feat: add line-length trigger to SOQL wrapping` (FIX-5, after maintainer sign-off)
6. `style: dedent subquery closing paren` (FIX-6, after maintainer sign-off)
7. `test: add idempotency assertion to formatter test harness` (can land any time; will fail until 1–2 are in)
# Consuming Test Results

This guide explains how to consume output from CCL test runners that emit the [`test-results-format.json`](../schemas/test-results-format.json) schema, and how to derive the summaries, breakdowns, and dashboards that implementations historically baked into their runners.

## Design Principle: Raw Outcomes, Consumer-Side Aggregation

The results schema deliberately contains **only raw per-test outcomes** along with their full tag metadata (`features`, `behaviors`, `variants`). It does **not** pre-compute pass/fail counts, feature-support matrices, or compliance scores.

This is intentional: different consumers want different aggregations, and the information needed to compute any of them is already present per test. A runner that also emits pre-computed summaries would lock consumers into one particular view.

> **For runner authors:** Do not add a "feature support summary" printout to your runner. Emit the results document; let consumers (dashboards, CI checks, docs sites) derive whatever breakdown they need.

## Schema at a Glance

Every `TestOutcome` carries the tags needed for grouping:

```json
{
  "name": "multiline_key_basic",
  "validation": "parse",
  "features": ["multiline_keys", "whitespace"],
  "behaviors": [],
  "variants": [],
  "outcome": "fail",
  "error": "expected 2 entries, got 1"
}
```

The `outcome` field is one of `"pass"`, `"fail"`, `"skip"`, or `"todo"`. `implementation.implementedFunctions` in the top-level document tells consumers which functions have real implementations (vs. `todo` stubs) — this is the one derived fact that cannot be recovered from outcomes alone.

## Deriving a Feature-Support Summary

A feature-support summary answers: "Of the tests that exercise feature X, how many pass?"

```javascript
function featureSupport(results) {
  const summary = new Map(); // feature -> { pass, fail, skip, todo, total }

  for (const test of results.tests) {
    for (const feature of test.features) {
      let row = summary.get(feature);
      if (!row) {
        row = { pass: 0, fail: 0, skip: 0, todo: 0, total: 0 };
        summary.set(feature, row);
      }
      row[test.outcome]++;
      row.total++;
    }
  }

  return summary;
}
```

Rendered as a table:

```
Feature                    Pass  Fail  Skip  Todo  Total  Support
─────────────────────────────────────────────────────────────────
comments                     42     0     0     0     42  100%
multiline_keys                3     7     0     0     10   30%
tab_in_value_preserved       15     0     0     0     15  100%
unicode                      18     0     2     0     20   90%
```

A feature is "fully supported" when `pass == total - skip - todo` and `fail == 0`. Partial support (some pass, some fail) usually points at an incomplete implementation of that feature.

## Deriving a Behavior/Variant Breakdown

Behaviors and variants encode implementation *choices*, not capabilities. Group them the same way to sanity-check that an implementation's declared variant matches how it actually behaves:

```javascript
function groupByTag(results, tagField) {
  const out = new Map();
  for (const test of results.tests) {
    for (const tag of test[tagField]) {
      let row = out.get(tag) ?? { pass: 0, fail: 0, skip: 0, todo: 0 };
      row[test.outcome]++;
      out.set(tag, row);
    }
  }
  return out;
}

const behaviorBreakdown = groupByTag(results, "behaviors");
const variantBreakdown = groupByTag(results, "variants");
```

A failure inside a variant the implementation claims to target is an alignment bug; a failure inside a variant it does *not* target is expected and can be filtered out at the runner level via `conflicts`.

## Classifying Functions: Implemented / Todo / Unsupported

`outcome == "todo"` alone does not tell you whether a function is unimplemented or just partially incorrect. Combine outcomes with `implementation.implementedFunctions`:

```javascript
function classifyFunctions(results) {
  const implemented = new Set(results.implementation.implementedFunctions);
  const seen = new Map(); // function -> { pass, fail, skip, todo }

  for (const test of results.tests) {
    let row = seen.get(test.validation);
    if (!row) {
      row = { pass: 0, fail: 0, skip: 0, todo: 0 };
      seen.set(test.validation, row);
    }
    row[test.outcome]++;
  }

  const classification = {};
  for (const [fn, row] of seen) {
    if (!implemented.has(fn)) {
      classification[fn] = "unsupported";
    } else if (row.fail === 0 && row.pass > 0) {
      classification[fn] = "implemented";
    } else {
      classification[fn] = "partial";
    }
  }
  return classification;
}
```

## Cross-Implementation Matrix

When comparing multiple implementations, load each results document and pivot by feature:

```javascript
function featureMatrix(resultsByImpl) {
  const features = new Set();
  const rows = {};

  for (const [name, results] of Object.entries(resultsByImpl)) {
    const support = featureSupport(results);
    rows[name] = support;
    for (const f of support.keys()) features.add(f);
  }

  return { features: [...features].sort(), rows };
}
```

The [CCL docs site](https://ccl.tylerbutler.com) uses this pattern to generate the implementation comparison tables.

## See Also

- [`schemas/test-results-format.json`](../schemas/test-results-format.json) — the full schema definition
- [`docs/test-filtering.md`](test-filtering.md) — filtering inputs *before* a run (as opposed to aggregating outputs *after*)
- [`docs/implementing-ccl.md`](implementing-ccl.md) — guidance for implementers emitting results documents

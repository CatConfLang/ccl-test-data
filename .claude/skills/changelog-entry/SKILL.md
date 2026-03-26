---
name: changelog-entry
description: Create properly formatted commit messages that render correctly in the project's auto-generated changelog. Use this skill whenever committing changes, writing commit messages, preparing squash merge descriptions, or when the user mentions "changelog", "commit message", "release notes", or asks how changes will appear in the changelog.
---

# Changelog Entry Skill

This project generates its changelog automatically from git commit messages using [Python Semantic Release](https://python-semantic-release.readthedocs.io/). The changelog template (`templates/CHANGELOG.md.j2`) renders commit descriptions as markdown nested under version and type headings. Getting the commit format right matters because the commit message *is* the changelog entry.

## How the changelog is structured

```
# Changelog                          ← h1
## [0.6.0] - 2026-02-24             ← h2 (version)
### Features                         ← h3 (commit type)
- **scope:** First line of commit    ← list item
  Rest of commit body                ← indented under list item
```

Commit bodies are rendered indented under a list item, nested three heading levels deep. This has direct implications for formatting — see the heading level rule below.

## Commit message format

Use [conventional commits](https://www.conventionalcommits.org/):

```
type(scope): short description

Optional body with more detail.
```

**Types that appear in the changelog:**
- `feat` — new functionality (bumps minor version)
- `fix` — bug fixes (bumps patch version)
- `perf` — performance improvements (bumps patch version)
- `docs` — documentation changes
- `refactor` — code restructuring
- `test` — test changes

**Types excluded from the changelog** (still valid commits, just filtered out):
- `chore`, `ci`, `style`, `build`

**Scope** is optional but encouraged — it appears bold in the changelog entry (e.g., `**tests:**`). Common scopes in this project: `tests`, `schema`, `generator`, `cli`, `mock`.

## The heading level rule

Commit bodies render nested inside `### Type` (h3), so the changelog template (`templates/CHANGELOG.md.j2`) automatically bumps heading levels by 2 to maintain correct hierarchy:

- `##` in commit body → `####` in changelog
- `###` → `#####`
- `####` → `######`

Headings are fine to use in commit bodies — they'll render correctly up to `####` (which becomes `######`, the maximum markdown heading level). Avoid `#####` or `######` in commit bodies since they can't be bumped further.

```
fix(tests): resolve three test inconsistencies (#85)

## Summary

Fixes three related issues with test generation.

### Changes

* fix(generator): add canonical_format to compositeFunctionMap
* fix(schema): add filter to CRLF behavior affectedFunctions
* fix(tests): remove trailing whitespace from print expectations

Closes #83, closes #84, closes #85
```

## Squash merge format

When a PR contains multiple logical changes that should each appear as separate changelog entries, format the squash merge commit body with `*`-prefixed sub-commits:

```
feat(scope): main PR title (#123)

Brief summary of the overall changes.

* fix(tests): first logical change

Description of what this change does and why.

* fix(schema): second logical change

Description of what this change does and why.

* feat(cli): third logical change

Description of what this change does and why.
```

Each `* type(scope): description` becomes a separate changelog entry with its body text preserved. The `chore` type is excluded from individual sub-entries too.

## Writing good changelog entries

The first line (after `type(scope):`) becomes the changelog entry title. It should be concise and describe the *what* from a user's perspective. The body should explain the *why* — what was broken, what motivated the change, what the user-visible impact is.

**Good:**
```
fix(tests): correct behavior tag on key_with_tabs test

The reference_compliant variant strips tabs around the `=` delimiter,
which is tabs_as_whitespace behavior, not tabs_as_content. Fixes #83.
```

**Weak:**
```
fix(tests): update test file

Changed some tags.
```

## Useful commands

```bash
# Preview what the next version would look like
semantic-release --noop version

# Regenerate the changelog from git history
semantic-release changelog

# See the Jinja2 template that renders the changelog
cat templates/CHANGELOG.md.j2
```

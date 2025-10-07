# Documentation Guide for Copilot Agents

This guide provides specialized instructions for agents updating documentation in the CCL Test Suite repository.

## Prerequisites

Before working on documentation, ensure you've read:
- [onboarding.md](onboarding.md) - General repository understanding

## Documentation Structure

```
ccl-test-data/
├── README.md                      # Main repository README
├── CLAUDE.md                      # Claude AI assistant guidance
└── docs/
    ├── README.md                  # Documentation index
    ├── ARCHITECTURE.md            # System architecture
    ├── API.md                     # API reference
    ├── DEVELOPER_GUIDE.md         # Development workflows
    ├── CLI_REFERENCE.md           # CLI command reference
    ├── TROUBLESHOOTING.md         # Common issues
    ├── MOCK_IMPLEMENTATION.md     # Reference implementation
    ├── implementing-ccl.md        # Implementation guide
    ├── test-selection-guide.md    # Test filtering
    ├── test-architecture.md       # Test design
    ├── test-filtering.md          # Advanced filtering
    └── schema-reference.md        # Schema documentation
```

## Documentation Categories

### User-Facing Documentation

**Purpose**: Help users understand and use the repository

**Files**:
- `README.md` - Overview, quick start, basic usage
- `docs/implementing-ccl.md` - Guide for CCL implementers
- `docs/CLI_REFERENCE.md` - Command-line interface reference
- `docs/test-selection-guide.md` - How to select and filter tests

**Audience**: CCL implementers, test suite users

**Update when**:
- Adding new features
- Changing user-facing behavior
- Adding new commands
- Updating statistics

### Developer Documentation

**Purpose**: Help contributors work on the repository

**Files**:
- `docs/DEVELOPER_GUIDE.md` - Development workflow and standards
- `docs/ARCHITECTURE.md` - System design and structure
- `docs/API.md` - Internal API documentation
- `docs/MOCK_IMPLEMENTATION.md` - Reference implementation details

**Audience**: Repository contributors, developers extending the suite

**Update when**:
- Changing internal APIs
- Modifying architecture
- Adding new packages
- Changing development workflow

### Reference Documentation

**Purpose**: Provide detailed technical specifications

**Files**:
- `docs/schema-reference.md` - JSON schema documentation
- `docs/test-architecture.md` - Test suite design
- `docs/test-filtering.md` - Advanced filtering patterns
- `CLAUDE.md` - AI assistant guidance

**Audience**: Technical users, automation tools, AI assistants

**Update when**:
- Changing schemas
- Adding new test types
- Modifying filtering logic
- Updating automation guidance

## Documentation Standards

### Markdown Style

**Headers**:
```markdown
# Top-level heading (one per document)

## Major section

### Subsection

#### Minor subsection
```

**Lists**:
```markdown
- Unordered list item
- Another item
  - Nested item

1. Ordered list item
2. Another numbered item
```

**Code blocks**:
````markdown
```bash
# Shell commands
just test
```

```go
// Go code
func Example() {}
```

```json
{
  "example": "JSON data"
}
```
````

**Links**:
```markdown
[Link text](relative/path/file.md)
[External link](https://example.com)
[Section link](#section-heading)
```

**Emphasis**:
```markdown
**Bold text**
*Italic text*
`code or command`
```

### Writing Style

**Be clear and concise**:
```markdown
❌ The test suite, which is comprehensive, contains many tests
✅ The test suite contains 452 assertions across 167 tests
```

**Use active voice**:
```markdown
❌ Tests are run by executing the command
✅ Run tests with: just test
```

**Provide examples**:
```markdown
❌ You can filter tests
✅ Filter tests by feature:
    just test --features comments,parsing
```

**Be specific**:
```markdown
❌ Run the necessary commands
✅ Run these commands:
    1. just lint
    2. just reset
    3. just validate
```

### Structure Guidelines

**Start with overview**:
```markdown
# Document Title

Brief description of what this document covers.

## Overview

High-level explanation of the topic.
```

**Use progressive disclosure**:
```markdown
## Quick Start
(Basic usage)

## Common Tasks
(Typical workflows)

## Advanced Topics
(Complex scenarios)

## Reference
(Detailed specifications)
```

**Include navigation**:
```markdown
## See Also

- [Related Doc 1](path/to/doc1.md)
- [Related Doc 2](path/to/doc2.md)
```

## Updating Specific Documentation

### Updating README.md

**When to update**:
- Major feature additions
- Test statistics change
- Quick start process changes
- Repository structure changes

**Key sections**:
1. **Overview** - High-level description
2. **Test Statistics** - Current test counts
3. **Quick Start** - Essential first commands
4. **Test Files** - Organization and categories
5. **Contributing** - How to contribute

**Example update** (test statistics):
```markdown
## Test Statistics

The test suite provides comprehensive coverage with **452 assertions** across **167 tests**:

```bash
# View detailed statistics
just stats
```
```

### Updating DEVELOPER_GUIDE.md

**When to update**:
- Changing development workflow
- Adding new commands
- Modifying quality standards
- Adding new test types

**Key sections**:
1. **Project Architecture** - High-level structure
2. **Development Workflow** - Standard processes
3. **Adding New Tests** - Test creation guide
4. **Pre-Commit Checklist** - Required steps

**Example update** (adding new command):
```markdown
### Essential Commands

| Command | Purpose |
|---------|---------|
| `just lint` | Format and lint (REQUIRED before commits) |
| `just reset` | Generate basic tests and verify passing |
| `just newcommand` | Description of new command |
```

### Updating CLI_REFERENCE.md

**When to update**:
- Adding CLI commands
- Changing command options
- Modifying command behavior
- Adding examples

**Key sections**:
1. **Installation** - How to install tools
2. **Main Command** - Primary CLI documentation
3. **Utility Commands** - Helper tools
4. **Example Workflows** - Common usage patterns

**Example update** (new command):
```markdown
### Command: newcommand

Brief description of what the command does.

#### Options

| Option | Description | Default |
|--------|-------------|---------|
| `--option` | What this option does | `default` |

#### Examples

```bash
# Basic usage
ccl-test-runner newcommand

# With options
ccl-test-runner newcommand --option value
```

#### Output

Description of command output.
```

### Updating API.md

**When to update**:
- Changing function signatures
- Adding new packages
- Modifying public APIs
- Adding new types

**Key sections**:
1. **Package Overview** - Package organization
2. **Core Types** - Data structures
3. **Functions** - Public API documentation
4. **Examples** - Usage examples

**Example update** (new function):
```markdown
### NewFunction

```go
func NewFunction(param string) (result string, error)
```

Description of what the function does.

**Parameters**:
- `param` - Description of parameter

**Returns**:
- `result` - Description of return value
- `error` - Error conditions

**Example**:
```go
result, err := NewFunction("input")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```
```

### Updating test-selection-guide.md

**When to update**:
- Adding new features
- Adding new behaviors
- Changing tag structure
- Adding filtering examples

**Key sections**:
1. **Overview** - Test selection introduction
2. **Typed Fields Architecture** - Metadata explanation
3. **Language-Specific Integration** - Per-language examples
4. **Progressive Implementation Guide** - Phased approach

**Example update** (new feature):
```markdown
### Features Array

| Feature | Description | Example Use |
|---------|-------------|-------------|
| `comments` | Requires comment support | Tests with `#` comments |
| `new_feature` | Description of new feature | Usage example |
```

## Common Documentation Tasks

### Adding Examples

**When adding examples**:
1. Test the example to ensure it works
2. Include expected output
3. Explain what the example demonstrates
4. Keep examples minimal but complete

**Format**:
```markdown
### Example: Descriptive Title

Brief description of what this example shows.

```bash
# Command to run
just command --option value
```

**Output**:
```
Expected output here
```

**Explanation**: What this demonstrates and why it's useful.
```

### Updating Command References

**When documenting commands**:
1. Show command syntax
2. List all options with descriptions
3. Provide practical examples
4. Include common use cases

**Template**:
```markdown
### Command: command-name

Brief description.

#### Syntax

```bash
command-name [OPTIONS] [ARGS]
```

#### Options

| Option | Description | Default |
|--------|-------------|---------|
| `--option` | What it does | `value` |

#### Examples

```bash
# Example 1
command-name --option value

# Example 2
command-name arg1 arg2
```
```

### Adding Troubleshooting Sections

**Format**:
```markdown
### Issue: Brief Description

**Symptoms**:
- What the user sees
- Error messages

**Cause**:
Why this happens

**Solution**:
1. Step-by-step fix
2. Another step
3. Verification step

**Prevention**:
How to avoid this issue
```

### Cross-Referencing Documents

**Always link related content**:
```markdown
For more details on test generation, see [DEVELOPER_GUIDE.md](DEVELOPER_GUIDE.md).

Related topics:
- [Test Architecture](test-architecture.md) - Test suite design
- [CLI Reference](CLI_REFERENCE.md) - Command documentation
```

## Documentation Workflow

### Standard Update Process

```bash
# 1. Make documentation changes
# Edit relevant .md files

# 2. Check for broken links (if tool available)
# markdown-link-check docs/*.md

# 3. Preview changes (if possible)
# Open in Markdown viewer

# 4. Verify examples work
# Run any commands shown in examples

# 5. Check consistency
# Ensure terminology is consistent across docs

# 6. Commit changes
git add docs/ README.md CLAUDE.md
git commit -m "Update documentation for [feature]"
```

### Pre-Commit Documentation Checklist

✅ **Content**:
- [ ] Information is accurate
- [ ] Examples are tested and work
- [ ] Terminology is consistent
- [ ] Links are valid

✅ **Style**:
- [ ] Headers follow hierarchy
- [ ] Code blocks have language tags
- [ ] Lists are formatted consistently
- [ ] Tables are aligned

✅ **Completeness**:
- [ ] All sections updated
- [ ] Cross-references added
- [ ] Examples included
- [ ] Related docs updated

### Synchronizing Documentation

**When changing functionality, update**:

1. **API changes** → Update `API.md` and `DEVELOPER_GUIDE.md`
2. **CLI changes** → Update `CLI_REFERENCE.md`
3. **Test changes** → Update `test-selection-guide.md` and `DEVELOPER_GUIDE.md`
4. **Statistics changes** → Update `README.md`
5. **Workflow changes** → Update `DEVELOPER_GUIDE.md` and `CLAUDE.md`

**Check list of affected docs**:
```bash
# Search for references to changed item
grep -r "changed_item" docs/ README.md CLAUDE.md
```

## Documentation Testing

### Verifying Examples

**Test all code examples**:
```bash
# Extract and test commands from documentation
# Example: test commands from README.md

just deps
just reset
just stats
# ... etc
```

### Checking Links

**Manual link verification**:
1. Click each link in rendered markdown
2. Verify it points to correct location
3. Check external links are accessible

**Automated link checking** (if available):
```bash
# Using markdown-link-check
markdown-link-check docs/*.md README.md
```

### Consistency Checks

**Terminology**:
- Use consistent names for concepts
- Define acronyms on first use
- Use same command names throughout

**Formatting**:
- Consistent code block languages
- Consistent list formatting
- Consistent header capitalization

**Cross-references**:
- All mentioned docs are linked
- Links use relative paths
- Section links use correct anchors

## Special Documentation Files

### CLAUDE.md

**Purpose**: Guide Claude AI assistant when working on repository

**Update when**:
- Adding new workflows
- Changing test structure
- Adding new commands
- Modifying quality standards

**Key sections**:
- Quick Start
- Development Workflow
- Command Reference
- Test Architecture
- Build System

**Keep concise**: Focus on practical commands and workflows

### README.md

**Purpose**: First impression and quick start

**Update when**:
- Major changes to repository
- Statistics change significantly
- Installation process changes
- Key features are added

**Keep focused**: Don't duplicate detailed docs, link to them instead

### docs/README.md

**Purpose**: Documentation index and navigation

**Update when**:
- Adding new documentation files
- Reorganizing docs
- Changing documentation structure

**Keep organized**: Group related docs, provide clear navigation

## Best Practices for Documentation

### DO:
✅ Test all examples before documenting
✅ Use consistent terminology throughout
✅ Link to related documentation
✅ Include practical examples
✅ Update all affected docs together
✅ Keep examples minimal but complete
✅ Use clear, simple language
✅ Structure content progressively

### DON'T:
❌ Duplicate content across documents
❌ Include untested examples
❌ Use jargon without explanation
❌ Create orphaned documentation
❌ Break existing links
❌ Skip cross-referencing
❌ Use vague descriptions
❌ Forget to update related docs

## Advanced Topics

### Documentation Generation

Some documentation is auto-generated:

**Generated sections**:
- Test statistics in README.md
- Some API documentation

**When updating generated docs**:
1. Check if there's a generation script
2. Update the source, not the output
3. Regenerate using appropriate tool

### Documentation Templates

**Creating consistent sections**:
```markdown
## Section Name

Brief introduction to the section.

### Subsection 1

Content with examples.

### Subsection 2

Content with examples.

## See Also

- [Related Doc](path.md)
```

### Maintaining Documentation Quality

**Regular reviews**:
1. Check for outdated information
2. Test examples periodically
3. Verify links are still valid
4. Update statistics and numbers
5. Improve clarity based on feedback

## Common Issues and Solutions

### Broken Links

**Issue**: Links don't work after file moves

**Solution**:
```bash
# Find all references to moved file
grep -r "old-filename" docs/ *.md

# Update all references
# Use find-and-replace carefully
```

### Outdated Examples

**Issue**: Examples no longer work

**Solution**:
1. Test all examples
2. Update commands/code
3. Verify output matches
4. Document changes

### Inconsistent Terminology

**Issue**: Same concept called different names

**Solution**:
1. Choose canonical term
2. Search and replace consistently
3. Add definitions if needed
4. Update glossary (if exists)

### Duplicate Information

**Issue**: Same content in multiple docs

**Solution**:
1. Choose primary location
2. Replace duplicates with links
3. Keep DRY (Don't Repeat Yourself)

## Pre-Commit Checklist for Documentation

```bash
# 1. Review changes
git diff docs/ README.md CLAUDE.md

# 2. Test examples
# Run all commands shown in updated docs

# 3. Check links
# Verify all links work

# 4. Verify consistency
# Check terminology matches across docs

# 5. Preview if possible
# View in Markdown renderer

# 6. Stage changes
git add docs/ README.md CLAUDE.md

# 7. Commit with clear message
git commit -m "Update documentation: [brief description]

- Updated [file] to reflect [changes]
- Added examples for [feature]
- Fixed links to [document]"
```

## References

- **[DEVELOPER_GUIDE.md](/docs/DEVELOPER_GUIDE.md)** - Contributing guidelines
- **[docs/README.md](/docs/README.md)** - Documentation index
- **Markdown Guide**: https://www.markdownguide.org/
- **GitHub Flavored Markdown**: https://github.github.com/gfm/

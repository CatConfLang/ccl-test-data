---
title: Test Architecture
type: note
permalink: architecture/test-architecture
---

# CCL Test Architecture & Tagging System

## Multi-Level CCL Implementation
- **Level 1**: Raw parsing (text → flat entries) - `Parse()`
- **Level 2**: Entry processing (indentation, comments) - `Filter()`, `Compose()`, `ExpandDotted()`  
- **Level 3**: Object construction (flat → nested objects) - `MakeObjects()`
- **Level 4**: Typed access (type-safe value extraction) - `GetString()`, `GetInt()`, etc.
- **Level 5**: Validation/formatting - `PrettyPrint()`

## Feature-Based Tagging System

### Structured Tags
**Function Tags** (`function:*`) - Required CCL functions:
- `function:parse`, `function:parse-value`, `function:filter`, `function:compose`
- `function:make-objects`, `function:get-string`, `function:get-int`, etc.

**Feature Tags** (`feature:*`) - Optional language features:
- `feature:comments`, `feature:dotted-keys`, `feature:empty-keys`
- `feature:multiline`, `feature:unicode`, `feature:whitespace`

**Behavior Tags** (`behavior:*`) - Implementation choices:
- `behavior:crlf-preserve` vs `behavior:crlf-normalize`
- `behavior:tabs-preserve` vs `behavior:tabs-to-spaces`

**Variant Tags** (`variant:*`) - Specification variants:
- `variant:proposed-behavior` vs `variant:reference-compliant`

## Test Organization
- **Essential**: `api-essential-parsing.json` - Basic Level 1 functionality
- **Comprehensive**: `api-comprehensive-parsing.json` - Edge cases, whitespace
- **Processing**: `api-processing.json` - Level 2 composition and filtering
- **Objects**: `api-object-construction.json` - Level 3 nested object creation
- **Typed**: `api-typed-access.json` - Level 4 type-safe access
- **Comments**: `api-comments.json` - Comment syntax and filtering
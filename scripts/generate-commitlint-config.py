#!/usr/bin/env python3
"""Generate commitlint.config.cjs from commit-types.json (single source of truth)."""

import json
import sys
from pathlib import Path

COMMITLINT_TEMPLATE = '''/** @type {{import('@commitlint/types').UserConfig}} */
// Auto-generated from commit-types.json - edit that file instead
module.exports = {{
  extends: ['@commitlint/config-conventional'],
  plugins: ['selective-scope'],
  rules: {{
    // Disable default scope rules - we use selective-scope instead
    'scope-empty': [0],
    'scope-enum': [0],

    // Allowed scopes per type
    // - Types listed with array: scope REQUIRED, must be from the list
    // - Types not listed: scope not enforced
    'selective-scope': [
      2,
      'always',
      {{
{selective_scope_rules}
      }},
    ],
  }},
}};
'''


def generate_commitlint_config(config: dict) -> str:
    """Generate commitlint.config.cjs content from config."""
    types = config["types"]
    scopes = config["scopes"]

    # Get list of scopes for selective-scope
    scope_list = list(scopes.keys())

    # Generate selective-scope rules
    scope_required_types = []
    for type_name, type_config in types.items():
        if type_config.get("scope_required", False):
            scope_required_types.append(type_name)

    # Format the rules
    rules_lines = []
    for type_name in scope_required_types:
        rules_lines.append(f"        {type_name}: {json.dumps(scope_list)},")

    return COMMITLINT_TEMPLATE.format(selective_scope_rules="\n".join(rules_lines))


def main():
    root = Path(__file__).parent.parent
    config_path = root / "commit-types.json"

    with open(config_path) as f:
        config = json.load(f)

    dry_run = "--dry-run" in sys.argv

    # Generate commitlint.config.cjs
    commitlint_content = generate_commitlint_config(config)
    commitlint_path = root / "commitlint.config.cjs"

    if dry_run:
        print(f"=== {commitlint_path} ===")
        print(commitlint_content)
    else:
        commitlint_path.write_text(commitlint_content)
        print(f"Wrote {commitlint_path}")


if __name__ == "__main__":
    main()

#!/usr/bin/env python3
"""Generate cliff.toml and commitlint.config.cjs from commit-types.json (single source of truth)."""

import json
import sys
from pathlib import Path

CLIFF_TEMPLATE = '''# git-cliff config for ccl-test-data
# Auto-generated from commit-types.json - edit that file instead

[changelog]
header = """
# Changelog\\n
All notable changes to the CCL test data will be documented in this file.\\n
"""
body = """
{{% if version %}}\\
    ## [{{{{ version | trim_start_matches(pat="v") }}}}] - {{{{ timestamp | date(format="%Y-%m-%d") }}}}
{{% else %}}\\
    ## [Unreleased]
{{% endif %}}\\
{{% set scope_order = {scope_order} %}}\\
{{% for scope_name in scope_order %}}\\
    {{% set scope_commits = commits | filter(attribute="scope", value=scope_name) %}}\\
    {{% if scope_commits | length > 0 %}}\\
    ### {{{{ scope_name | replace(from="-", to=" ") | title }}}}
    {{% for group, grouped_commits in scope_commits | group_by(attribute="group") %}}\\
        #### {{{{ group | striptags | trim | upper_first }}}}
        {{% for commit in grouped_commits %}}\\
            - {{{{ commit.message | upper_first }}}}
        {{% endfor %}}
    {{% endfor %}}
    {{% endif %}}\\
{{% endfor %}}
"""
footer = ""
trim = true

[changelog.scopes]
{scope_mappings}

[git]
conventional_commits = true
filter_unconventional = true
split_commits = false
commit_parsers = [
    # Skip noise
{skip_parsers}

    # Group by type
{group_parsers}
]
filter_commits = false
tag_pattern = "v[0-9].*"

[bump]
features_always_bump_minor = true
breaking_always_bump_major = true
'''

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


def generate_cliff_config(config: dict) -> str:
    """Generate cliff.toml content from config."""
    types = config["types"]
    scopes = config["scopes"]
    scope_order = config.get("changelog_scope_order", list(scopes.keys()))

    # Generate scope mappings
    scope_mappings = "\n".join(
        f'{scope} = "{info["display_name"]}"' for scope, info in scopes.items()
    )

    # Generate skip parsers for types without changelog groups
    skip_lines = []
    for type_name, type_config in types.items():
        if type_config.get("changelog_group") is None:
            skip_lines.append(f'    {{ message = "^{type_name}", skip = true }},')

    # Generate skip parsers for scopes not in changelog
    for scope, info in scopes.items():
        if not info.get("in_changelog", True):
            skip_lines.append(f'    {{ scope = "{scope}", skip = true }},')

    # Generate group parsers for types with changelog groups
    group_lines = []
    for type_name, type_config in types.items():
        group = type_config.get("changelog_group")
        if group:
            group_lines.append(f'    {{ message = "^{type_name}", group = "{group}" }},')

    return CLIFF_TEMPLATE.format(
        scope_order=json.dumps(scope_order),
        scope_mappings=scope_mappings,
        skip_parsers="\n".join(skip_lines),
        group_parsers="\n".join(group_lines),
    )


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

    # Generate cliff.toml
    cliff_content = generate_cliff_config(config)
    cliff_path = root / "cliff.toml"

    if dry_run:
        print(f"=== {cliff_path} ===")
        print(cliff_content)
    else:
        cliff_path.write_text(cliff_content)
        print(f"Wrote {cliff_path}")

    # Generate commitlint.config.cjs
    commitlint_content = generate_commitlint_config(config)
    commitlint_path = root / "commitlint.config.cjs"

    if dry_run:
        print(f"\n=== {commitlint_path} ===")
        print(commitlint_content)
    else:
        commitlint_path.write_text(commitlint_content)
        print(f"Wrote {commitlint_path}")


if __name__ == "__main__":
    main()

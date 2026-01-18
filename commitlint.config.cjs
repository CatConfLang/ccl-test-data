/** @type {import('@commitlint/types').UserConfig} */
// Auto-generated from commit-types.json - edit that file instead
module.exports = {
  extends: ['@commitlint/config-conventional'],
  plugins: ['selective-scope'],
  rules: {
    // Disable default scope rules - we use selective-scope instead
    'scope-empty': [0],
    'scope-enum': [0],

    // Allowed scopes per type
    // - Types listed with array: scope REQUIRED, must be from the list
    // - Types not listed: scope not enforced
    'selective-scope': [
      2,
      'always',
      {
        feat: ["schema", "tests", "test-reader", "cli", "generation", "config", "build", "release"],
        fix: ["schema", "tests", "test-reader", "cli", "generation", "config", "build", "release"],
        perf: ["schema", "tests", "test-reader", "cli", "generation", "config", "build", "release"],
        refactor: ["schema", "tests", "test-reader", "cli", "generation", "config", "build", "release"],
        test: ["schema", "tests", "test-reader", "cli", "generation", "config", "build", "release"],
        revert: ["schema", "tests", "test-reader", "cli", "generation", "config", "build", "release"],
        style: ["schema", "tests", "test-reader", "cli", "generation", "config", "build", "release"],
      },
    ],
  },
};

/** @type {import('@commitlint/types').UserConfig} */
module.exports = {
  extends: ['@commitlint/config-conventional'],
  plugins: ['selective-scope'],
  rules: {
    // Disable default scope rules - we use selective-scope instead
    'scope-empty': [0],
    'scope-enum': [0],

    // Allowed scopes per type
    // - Types listed with array: scope REQUIRED, must be from the list
    // - Types listed with [null, ...]: scope OPTIONAL, if present must be from list
    // - Types not listed: scope not enforced
    'selective-scope': [
      2,
      'always',
      {
        // These types REQUIRE a scope from the allowed list
        feat: ['schema', 'tests', 'generation', 'config', 'test-reader', 'cli'],
        fix: ['schema', 'tests', 'generation', 'config', 'test-reader', 'cli'],
        refactor: ['schema', 'tests', 'generation', 'config', 'test-reader', 'cli'],
        perf: ['schema', 'tests', 'generation', 'config', 'test-reader', 'cli'],
        test: ['schema', 'tests', 'generation', 'config', 'test-reader', 'cli'],
        revert: ['schema', 'tests', 'generation', 'config', 'test-reader', 'cli'],
        style: ['schema', 'tests', 'generation', 'config', 'test-reader', 'cli'],

        // These types: scope NOT enforced (not listed)
        // - chore
        // - ci
        // - build
        // - docs
      },
    ],
  },
};

#!/usr/bin/env node

import { readFileSync, writeFileSync } from 'fs';

const filePath = './source_tests/experimental/api_experimental.json';

console.log('Removing expand_dotted function tests from experimental file...');

const content = JSON.parse(readFileSync(filePath, 'utf8'));

// Remove expand_dotted tests from each test case
const updatedTests = content.tests.map(test => ({
    ...test,
    tests: test.tests.filter(t => t.function !== 'expand_dotted')
}));

const updatedContent = {
    ...content,
    tests: updatedTests
};

writeFileSync(filePath, JSON.stringify(updatedContent, null, 2) + '\n');

console.log('Removed expand_dotted function tests from experimental file');
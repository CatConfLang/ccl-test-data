#!/usr/bin/env node

import { readFileSync, writeFileSync, readdirSync, existsSync, rmSync } from 'fs';
import { join } from 'path';

const CORE_DIR = './source_tests/core';
const PROPOSED_DIR = './source_tests/proposed';
const REFERENCE_DIR = './source_tests/reference';

console.log('Consolidating variant tests into single files...');

// Collect all proposed behavior tests
const proposedTests = [];
if (existsSync(PROPOSED_DIR)) {
    const proposedFiles = readdirSync(PROPOSED_DIR).filter(file => file.endsWith('.json'));
    for (const filename of proposedFiles) {
        const content = JSON.parse(readFileSync(join(PROPOSED_DIR, filename), 'utf8'));
        proposedTests.push(...content.tests);
    }
}

// Collect all reference compliant tests
const referenceTests = [];
if (existsSync(REFERENCE_DIR)) {
    const referenceFiles = readdirSync(REFERENCE_DIR).filter(file => file.endsWith('.json'));
    for (const filename of referenceFiles) {
        const content = JSON.parse(readFileSync(join(REFERENCE_DIR, filename), 'utf8'));
        referenceTests.push(...content.tests);
    }
}

// Create consolidated proposed behavior file
if (proposedTests.length > 0) {
    const proposedContent = {
        $schema: "file:///Volumes/Code/claude-workspace-ccl/ccl-test-data/schemas/source-format.json",
        tests: proposedTests
    };
    writeFileSync(
        join(CORE_DIR, 'api_proposed_behavior.json'),
        JSON.stringify(proposedContent, null, 2) + '\n'
    );
    console.log(`Created api_proposed_behavior.json with ${proposedTests.length} tests`);
}

// Create consolidated reference compliant file
if (referenceTests.length > 0) {
    const referenceContent = {
        $schema: "file:///Volumes/Code/claude-workspace-ccl/ccl-test-data/schemas/source-format.json",
        tests: referenceTests
    };
    writeFileSync(
        join(CORE_DIR, 'api_reference_compliant.json'),
        JSON.stringify(referenceContent, null, 2) + '\n'
    );
    console.log(`Created api_reference_compliant.json with ${referenceTests.length} tests`);
}

// Remove the separate variant directories
if (existsSync(PROPOSED_DIR)) {
    rmSync(PROPOSED_DIR, { recursive: true });
    console.log('Removed proposed/ directory');
}

if (existsSync(REFERENCE_DIR)) {
    rmSync(REFERENCE_DIR, { recursive: true });
    console.log('Removed reference/ directory');
}

console.log('Consolidation complete!');
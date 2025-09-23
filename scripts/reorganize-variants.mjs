#!/usr/bin/env node

import { readFileSync, writeFileSync, mkdirSync, readdirSync, existsSync } from 'fs';
import { join, dirname, basename } from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const PROJECT_ROOT = join(__dirname, '..');
const CORE_DIR = join(PROJECT_ROOT, 'source_tests', 'core');
const PROPOSED_DIR = join(PROJECT_ROOT, 'source_tests', 'proposed');
const REFERENCE_DIR = join(PROJECT_ROOT, 'source_tests', 'reference');

// Create directories if they don't exist
mkdirSync(PROPOSED_DIR, { recursive: true });
mkdirSync(REFERENCE_DIR, { recursive: true });

console.log('Reorganizing variant tests...');

// Get all JSON files in core directory
const coreFiles = readdirSync(CORE_DIR).filter(file => file.endsWith('.json'));

for (const filename of coreFiles) {
    const filePath = join(CORE_DIR, filename);
    console.log(`Processing ${filename}...`);

    const content = JSON.parse(readFileSync(filePath, 'utf8'));

    // Extract tests with proposed_behavior variants
    const proposedTests = content.tests.filter(test =>
        test.variants && test.variants.includes('proposed_behavior')
    );

    if (proposedTests.length > 0) {
        console.log(`  Extracting ${proposedTests.length} proposed behavior variants...`);
        const proposedContent = {
            $schema: content.$schema,
            tests: proposedTests
        };
        writeFileSync(
            join(PROPOSED_DIR, filename),
            JSON.stringify(proposedContent, null, 2) + '\n'
        );
    }

    // Extract tests with reference_compliant variants
    const referenceTests = content.tests.filter(test =>
        test.variants && test.variants.includes('reference_compliant')
    );

    if (referenceTests.length > 0) {
        console.log(`  Extracting ${referenceTests.length} reference compliant variants...`);
        const referenceContent = {
            $schema: content.$schema,
            tests: referenceTests
        };
        writeFileSync(
            join(REFERENCE_DIR, filename),
            JSON.stringify(referenceContent, null, 2) + '\n'
        );
    }

    // Remove variant tests from core file
    const coreTests = content.tests.filter(test =>
        !test.variants || test.variants.length === 0
    );

    console.log(`  Removing variants from core file, keeping ${coreTests.length} core tests...`);
    const updatedContent = {
        $schema: content.$schema,
        tests: coreTests
    };

    writeFileSync(filePath, JSON.stringify(updatedContent, null, 2) + '\n');
}

console.log('Reorganization complete!');
console.log(`Created directories:`);
console.log(`  ${PROPOSED_DIR}`);
console.log(`  ${REFERENCE_DIR}`);
console.log('');
console.log('Updated core files to remove variant tests');
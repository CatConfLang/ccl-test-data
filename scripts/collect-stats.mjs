#!/usr/bin/env node

import { readdir, readFile } from 'fs/promises';
import { join, basename } from 'path';
import { fileURLToPath } from 'url';
import { dirname } from 'path';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);
const ROOT_DIR = join(__dirname, '..');

// ANSI color codes
const colors = {
  reset: '\x1b[0m',
  bold: '\x1b[1m',
  dim: '\x1b[2m',
  red: '\x1b[31m',
  green: '\x1b[32m',
  yellow: '\x1b[33m',
  blue: '\x1b[34m',
  magenta: '\x1b[35m',
  cyan: '\x1b[36m',
  gray: '\x1b[90m'
};

// Feature-based categories - based on meta.feature field in test files
const FEATURE_CATEGORIES = {
  'parsing': 'core-parsing',
  'processing': 'advanced-processing', 
  'comments': 'advanced-processing',
  'object-construction': 'object-construction',
  'dotted-keys': 'object-construction',
  'typed-parsing': 'type-system',
  'pretty-printing': 'output-validation',
  'error-handling': 'output-validation'
};

async function findTestFiles() {
  const testDir = join(ROOT_DIR, 'tests');
  const files = await readdir(testDir, { recursive: true });
  
  return files
    .filter(file => file.endsWith('.json'))
    .map(file => join(testDir, file));
}

async function countTests(filePath) {
  try {
    const content = await readFile(filePath, 'utf8');
    const data = JSON.parse(content);
    
    if (!data.tests || !Array.isArray(data.tests)) {
      return 0;
    }
    
    return data.tests.length;
  } catch (error) {
    if (process.stdout.isTTY) {
      console.error(`${colors.red}Error reading ${filePath}: ${error.message}${colors.reset}`);
    }
    return 0;
  }
}

async function getFileMetadata(filePath) {
  try {
    const content = await readFile(filePath, 'utf8');
    const data = JSON.parse(content);
    
    // Get the feature from the first test's meta.feature field
    const feature = data.tests && data.tests[0] && data.tests[0].meta && data.tests[0].meta.feature;
    
    return {
      description: data.description || '',
      feature: feature || null
    };
  } catch {
    return { description: '', feature: null };
  }
}

function categorizeByFeature(feature) {
  if (!feature) return 'other';
  return FEATURE_CATEGORIES[feature] || 'other';
}

async function collectStats(silent = false) {
  const isTTY = process.stdout.isTTY;
  const showOutput = !silent && isTTY;
  
  if (showOutput) {
    console.log(`${colors.cyan}${colors.bold}🔍 Collecting CCL test statistics...${colors.reset}\n`);
  }
  
  const testFiles = await findTestFiles();
  const stats = {
    structure: 'feature-based',
    categories: {
      'core-parsing': { total: 0, files: {} },
      'advanced-processing': { total: 0, files: {} },
      'object-construction': { total: 0, files: {} },
      'type-system': { total: 0, files: {} },
      'output-validation': { total: 0, files: {} },
      'other': { total: 0, files: {} }
    },
    totalTests: 0,
    totalFiles: 0
  };
  
  if (showOutput) {
    console.log(`${colors.gray}Categorizing test files by metadata...${colors.reset}`);
  }
  
  for (const file of testFiles) {
    const count = await countTests(file);
    if (count > 0) {
      const metadata = await getFileMetadata(file);
      const category = categorizeByFeature(metadata.feature);
      const fileName = basename(file, '.json');
      
      stats.categories[category].files[fileName] = count;
      stats.categories[category].total += count;
      stats.totalTests += count;
      stats.totalFiles++;
      
      if (showOutput) {
        const categoryDisplay = category.replace('-', ' ').split(' ').map(w => w.charAt(0).toUpperCase() + w.slice(1)).join(' ');
        console.log(`  ${colors.gray}${categoryDisplay}: ${fileName} (${count} tests) [${metadata.feature || 'unknown'}]${colors.reset}`);
      }
    }
  }
  
  // Display summary or JSON based on TTY and silent mode
  if (!silent && isTTY) {
    console.log(`\n${colors.magenta}${colors.bold}📊 Test Statistics Summary${colors.reset}\n`);
    console.log(`${colors.blue}${colors.bold}Feature-Based Structure:${colors.reset}`);
    console.log(`  ${colors.gray}Core Parsing: ${stats.categories['core-parsing'].total} tests${colors.reset}`);
    console.log(`  ${colors.gray}Advanced Processing: ${stats.categories['advanced-processing'].total} tests${colors.reset}`);
    console.log(`  ${colors.gray}Object Construction: ${stats.categories['object-construction'].total} tests${colors.reset}`);
    console.log(`  ${colors.gray}Type System: ${stats.categories['type-system'].total} tests${colors.reset}`);
    console.log(`  ${colors.gray}Output & Validation: ${stats.categories['output-validation'].total} tests${colors.reset}`);
    
    if (stats.categories.other.total > 0) {
      console.log(`  ${colors.gray}Other: ${stats.categories.other.total} tests${colors.reset}`);
    }
    
    console.log(`\n${colors.blue}${colors.bold}Total: ${stats.totalTests} tests across ${stats.totalFiles} files${colors.reset}\n`);
    console.log(`${colors.blue}${colors.bold}JSON Output:${colors.reset}`);
    console.log(JSON.stringify(stats, null, 2));
  } else if (!silent) {
    // When piped (not silent), output only JSON
    console.log(JSON.stringify(stats, null, 2));
  }
  
  return stats;
}

// Run if executed directly
if (process.argv[1] === __filename) {
  collectStats().catch(error => {
    console.error(`${colors.red}${colors.bold}Error: ${error.message}${colors.reset}`);
    process.exit(1);
  });
}

export { collectStats, countTests, categorizeByFeature, getFileMetadata };
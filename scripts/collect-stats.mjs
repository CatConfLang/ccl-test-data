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

// Categories for test files - based on actual test file names and descriptions
const CATEGORIES = {
  core: ['core', 'fundamental', 'basic', 'essential', 'comprehensive', 'parsing validation'],
  features: ['feature', 'functionality', 'capability', 'dotted', 'typed', 'access', 'comment'],
  integration: ['integration', 'e2e', 'end-to-end', 'workflow', 'error', 'processing', 'composition'],
  utilities: ['util', 'helper', 'utility', 'tool', 'common', 'pretty', 'print', 'format']
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

async function getDescription(filePath) {
  try {
    const content = await readFile(filePath, 'utf8');
    const data = JSON.parse(content);
    return data.description || '';
  } catch {
    return '';
  }
}

function categorizeFile(filePath, description = '') {
  const fileName = basename(filePath).toLowerCase();
  const searchText = `${fileName} ${description}`.toLowerCase();
  
  for (const [category, keywords] of Object.entries(CATEGORIES)) {
    if (keywords.some(keyword => searchText.includes(keyword))) {
      return category;
    }
  }
  
  return 'other';
}

async function collectStats(silent = false) {
  const isTTY = process.stdout.isTTY;
  const showOutput = !silent && isTTY;
  
  if (showOutput) {
    console.log(`${colors.cyan}${colors.bold}🔍 Collecting CCL test statistics...${colors.reset}\n`);
  }
  
  const testFiles = await findTestFiles();
  const stats = {
    structure: 'flat',
    categories: {
      core: { total: 0, files: {} },
      features: { total: 0, files: {} },
      integration: { total: 0, files: {} },
      utilities: { total: 0, files: {} },
      other: { total: 0, files: {} }
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
      const description = await getDescription(file);
      const category = categorizeFile(file, description);
      const fileName = basename(file, '.json');
      
      stats.categories[category].files[fileName] = count;
      stats.categories[category].total += count;
      stats.totalTests += count;
      stats.totalFiles++;
      
      if (showOutput) {
        const categoryDisplay = category.charAt(0).toUpperCase() + category.slice(1);
        console.log(`  ${colors.gray}${categoryDisplay}: ${fileName} (${count} tests)${colors.reset}`);
      }
    }
  }
  
  // Display summary or JSON based on TTY and silent mode
  if (!silent && isTTY) {
    console.log(`\n${colors.magenta}${colors.bold}📊 Test Statistics Summary${colors.reset}\n`);
    console.log(`${colors.blue}${colors.bold}Feature-Based Structure (Flat):${colors.reset}`);
    console.log(`  ${colors.gray}Core: ${stats.categories.core.total} tests${colors.reset}`);
    console.log(`  ${colors.gray}Features: ${stats.categories.features.total} tests${colors.reset}`);
    console.log(`  ${colors.gray}Integration: ${stats.categories.integration.total} tests${colors.reset}`);
    console.log(`  ${colors.gray}Utilities: ${stats.categories.utilities.total} tests${colors.reset}`);
    
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

export { collectStats, countTests, categorizeFile };
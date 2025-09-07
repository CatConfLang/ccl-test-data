#!/usr/bin/env node

import { readFile, writeFile } from 'fs/promises';
import { join, dirname } from 'path';
import { fileURLToPath } from 'url';
import { collectStats } from './collect-stats.mjs';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);
const ROOT_DIR = join(__dirname, '..');
const README_PATH = join(ROOT_DIR, 'README.md');

// ANSI color codes
const colors = {
  reset: '\x1b[0m',
  bold: '\x1b[1m',
  red: '\x1b[31m',
  green: '\x1b[32m',
  yellow: '\x1b[33m',
  blue: '\x1b[34m',
  magenta: '\x1b[35m',
  cyan: '\x1b[36m',
  gray: '\x1b[90m'
};

async function updateReadme() {
  const isTTY = process.stdout.isTTY;
  
  if (isTTY) {
    console.log(`${colors.magenta}${colors.bold}📝 CCL README Updater${colors.reset}`);
    console.log(`${colors.gray}Updating README.md with current test statistics${colors.reset}\n`);
  } else {
    console.log('📝 Updating README.md with current test statistics...');
  }
  
  try {
    // Collect current stats (silent mode to avoid output)
    const stats = await collectStats(true);
    
    // Extract counts
    const coreTotal = stats.categories.core.total;
    const featuresTotal = stats.categories.features.total;
    const integrationTotal = stats.categories.integration.total;
    const utilitiesTotal = stats.categories.utilities.total;
    const total = stats.totalTests;
    
    if (isTTY) {
      console.log(`${colors.blue}${colors.bold}Current Test Counts:${colors.reset}`);
      console.log(`  ${colors.gray}Core: ${coreTotal}${colors.reset}`);
      console.log(`  ${colors.gray}Features: ${featuresTotal}${colors.reset}`);
      console.log(`  ${colors.gray}Integration: ${integrationTotal}${colors.reset}`);
      console.log(`  ${colors.gray}Utilities: ${utilitiesTotal}${colors.reset}`);
      console.log(`  ${colors.gray}Total: ${total}${colors.reset}\n`);
    } else {
      console.log(`📊 Current counts: Core:${coreTotal} Features:${featuresTotal} Integration:${integrationTotal} Utilities:${utilitiesTotal} Total:${total}`);
    }
    
    // Read current README
    let readme = await readFile(README_PATH, 'utf8');
    const originalReadme = readme;
    
    // Update total count in main description
    readme = readme.replace(
      /includes \*\*\d+ test cases\*\* total:/,
      `includes **${total} test cases** total:`
    );
    
    // Update individual category sections
    readme = readme.replace(
      /### Core \(\d+ tests\)/,
      `### Core (${coreTotal} tests)`
    );
    readme = readme.replace(
      /### Features \(\d+ tests\)/,
      `### Features (${featuresTotal} tests)`
    );
    readme = readme.replace(
      /### Integration \(\d+ tests\)/,
      `### Integration (${integrationTotal} tests)`
    );
    readme = readme.replace(
      /### Utilities \(\d+ tests\)/,
      `### Utilities (${utilitiesTotal} tests)`
    );
    
    // Check if anything changed
    if (readme === originalReadme) {
      if (isTTY) {
        console.log(`${colors.green}✅ README.md is already up to date${colors.reset}`);
      } else {
        console.log('✅ README.md is already up to date');
      }
      return { updated: false };
    }
    
    // Write updated README
    await writeFile(README_PATH, readme);
    
    if (isTTY) {
      console.log(`${colors.green}✅ README.md updated successfully${colors.reset}`);
      
      // Show a summary of changes
      console.log(`\n${colors.blue}${colors.bold}🔄 Changes Applied:${colors.reset}`);
      
      const changes = [];
      if (originalReadme.match(/### Core \((\d+) tests\)/)?.[1] !== String(coreTotal)) {
        changes.push(`Core: ${originalReadme.match(/### Core \((\d+) tests\)/)?.[1]} → ${coreTotal}`);
      }
      if (originalReadme.match(/### Features \((\d+) tests\)/)?.[1] !== String(featuresTotal)) {
        changes.push(`Features: ${originalReadme.match(/### Features \((\d+) tests\)/)?.[1]} → ${featuresTotal}`);
      }
      if (originalReadme.match(/### Integration \((\d+) tests\)/)?.[1] !== String(integrationTotal)) {
        changes.push(`Integration: ${originalReadme.match(/### Integration \((\d+) tests\)/)?.[1]} → ${integrationTotal}`);
      }
      if (originalReadme.match(/### Utilities \((\d+) tests\)/)?.[1] !== String(utilitiesTotal)) {
        changes.push(`Utilities: ${originalReadme.match(/### Utilities \((\d+) tests\)/)?.[1]} → ${utilitiesTotal}`);
      }
      
      changes.forEach(change => {
        console.log(`  ${colors.gray}${change}${colors.reset}`);
      });
    } else {
      console.log('✅ README.md updated successfully');
      console.log('🔄 Changes made:');
      
      // Simple diff output for non-TTY
      const oldCore = originalReadme.match(/### Core \((\d+) tests\)/)?.[1];
      const oldFeatures = originalReadme.match(/### Features \((\d+) tests\)/)?.[1];
      const oldIntegration = originalReadme.match(/### Integration \((\d+) tests\)/)?.[1];
      const oldUtilities = originalReadme.match(/### Utilities \((\d+) tests\)/)?.[1];
      
      if (oldCore !== String(coreTotal)) console.log(`- Core: ${oldCore} → ${coreTotal}`);
      if (oldFeatures !== String(featuresTotal)) console.log(`- Features: ${oldFeatures} → ${featuresTotal}`);
      if (oldIntegration !== String(integrationTotal)) console.log(`- Integration: ${oldIntegration} → ${integrationTotal}`);
      if (oldUtilities !== String(utilitiesTotal)) console.log(`- Utilities: ${oldUtilities} → ${utilitiesTotal}`);
    }
    
    return { updated: true };
    
  } catch (error) {
    console.error(`${colors.red}${colors.bold}Error: ${error.message}${colors.reset}`);
    process.exit(1);
  }
}

// Run if executed directly
if (process.argv[1] === __filename) {
  updateReadme().then(result => {
    // Exit with code 0 if successful (whether updated or not)
    process.exit(0);
  }).catch(error => {
    console.error(`${colors.red}${colors.bold}Error: ${error.message}${colors.reset}`);
    process.exit(1);
  });
}

export { updateReadme };
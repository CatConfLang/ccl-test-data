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
    
    // Extract counts from feature-based categories
    const coreParsingTotal = stats.categories['core-parsing'].total;
    const advancedProcessingTotal = stats.categories['advanced-processing'].total;
    const objectConstructionTotal = stats.categories['object-construction'].total;
    const typeSystemTotal = stats.categories['type-system'].total;
    const outputValidationTotal = stats.categories['output-validation'].total;
    const total = stats.totalTests;
    
    if (isTTY) {
      console.log(`${colors.blue}${colors.bold}Current Test Counts:${colors.reset}`);
      console.log(`  ${colors.gray}Core Parsing: ${coreParsingTotal}${colors.reset}`);
      console.log(`  ${colors.gray}Advanced Processing: ${advancedProcessingTotal}${colors.reset}`);
      console.log(`  ${colors.gray}Object Construction: ${objectConstructionTotal}${colors.reset}`);
      console.log(`  ${colors.gray}Type System: ${typeSystemTotal}${colors.reset}`);
      console.log(`  ${colors.gray}Output & Validation: ${outputValidationTotal}${colors.reset}`);
      console.log(`  ${colors.gray}Total: ${total}${colors.reset}\n`);
    } else {
      console.log(`📊 Current counts: CoreParsing:${coreParsingTotal} AdvancedProcessing:${advancedProcessingTotal} ObjectConstruction:${objectConstructionTotal} TypeSystem:${typeSystemTotal} OutputValidation:${outputValidationTotal} Total:${total}`);
    }
    
    // Read current README
    let readme = await readFile(README_PATH, 'utf8');
    const originalReadme = readme;
    
    // Update total count in main description
    readme = readme.replace(
      /includes \*\*\d+ test cases\*\* total:/,
      `includes **${total} test cases** total:`
    );
    
    // Update individual category sections with new feature-based structure
    readme = readme.replace(
      /\*\*Core Parsing\*\* \(\d+ tests\)/,
      `**Core Parsing** (${coreParsingTotal} tests)`
    );
    readme = readme.replace(
      /\*\*Advanced Processing\*\* \(\d+ tests\)/,
      `**Advanced Processing** (${advancedProcessingTotal} tests)`
    );
    readme = readme.replace(
      /\*\*Object Construction\*\* \(\d+ tests\)/,
      `**Object Construction** (${objectConstructionTotal} tests)`
    );
    readme = readme.replace(
      /\*\*Type System\*\* \(\d+ tests\)/,
      `**Type System** (${typeSystemTotal} tests)`
    );
    readme = readme.replace(
      /\*\*Output & Validation\*\* \(\d+ tests\)/,
      `**Output & Validation** (${outputValidationTotal} tests)`
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
      const oldCoreParsing = originalReadme.match(/\*\*Core Parsing\*\* \((\d+) tests\)/)?.[1];
      const oldAdvancedProcessing = originalReadme.match(/\*\*Advanced Processing\*\* \((\d+) tests\)/)?.[1];
      const oldObjectConstruction = originalReadme.match(/\*\*Object Construction\*\* \((\d+) tests\)/)?.[1];
      const oldTypeSystem = originalReadme.match(/\*\*Type System\*\* \((\d+) tests\)/)?.[1];
      const oldOutputValidation = originalReadme.match(/\*\*Output & Validation\*\* \((\d+) tests\)/)?.[1];
      
      if (oldCoreParsing !== String(coreParsingTotal)) {
        changes.push(`Core Parsing: ${oldCoreParsing} → ${coreParsingTotal}`);
      }
      if (oldAdvancedProcessing !== String(advancedProcessingTotal)) {
        changes.push(`Advanced Processing: ${oldAdvancedProcessing} → ${advancedProcessingTotal}`);
      }
      if (oldObjectConstruction !== String(objectConstructionTotal)) {
        changes.push(`Object Construction: ${oldObjectConstruction} → ${objectConstructionTotal}`);
      }
      if (oldTypeSystem !== String(typeSystemTotal)) {
        changes.push(`Type System: ${oldTypeSystem} → ${typeSystemTotal}`);
      }
      if (oldOutputValidation !== String(outputValidationTotal)) {
        changes.push(`Output & Validation: ${oldOutputValidation} → ${outputValidationTotal}`);
      }
      
      changes.forEach(change => {
        console.log(`  ${colors.gray}${change}${colors.reset}`);
      });
    } else {
      console.log('✅ README.md updated successfully');
      console.log('🔄 Changes made:');
      
      // Simple diff output for non-TTY
      const oldCoreParsing = originalReadme.match(/\*\*Core Parsing\*\* \((\d+) tests\)/)?.[1];
      const oldAdvancedProcessing = originalReadme.match(/\*\*Advanced Processing\*\* \((\d+) tests\)/)?.[1];
      const oldObjectConstruction = originalReadme.match(/\*\*Object Construction\*\* \((\d+) tests\)/)?.[1];
      const oldTypeSystem = originalReadme.match(/\*\*Type System\*\* \((\d+) tests\)/)?.[1];
      const oldOutputValidation = originalReadme.match(/\*\*Output & Validation\*\* \((\d+) tests\)/)?.[1];
      
      if (oldCoreParsing !== String(coreParsingTotal)) console.log(`- Core Parsing: ${oldCoreParsing} → ${coreParsingTotal}`);
      if (oldAdvancedProcessing !== String(advancedProcessingTotal)) console.log(`- Advanced Processing: ${oldAdvancedProcessing} → ${advancedProcessingTotal}`);
      if (oldObjectConstruction !== String(objectConstructionTotal)) console.log(`- Object Construction: ${oldObjectConstruction} → ${objectConstructionTotal}`);
      if (oldTypeSystem !== String(typeSystemTotal)) console.log(`- Type System: ${oldTypeSystem} → ${typeSystemTotal}`);
      if (oldOutputValidation !== String(outputValidationTotal)) console.log(`- Output & Validation: ${oldOutputValidation} → ${outputValidationTotal}`);
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
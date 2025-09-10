#!/usr/bin/env node

import { readFile, writeFile } from 'fs/promises';
import { join, dirname } from 'path';
import { fileURLToPath } from 'url';
import { unified } from 'unified';
import remarkParse from 'remark-parse';
import remarkStringify from 'remark-stringify';
import remarkGfm from 'remark-gfm';
import { visit } from 'unist-util-visit';
import { execSync } from 'child_process';

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

/**
 * Custom remark plugin to update test counts in the README
 */
function remarkUpdateCounts(stats) {
  return (tree) => {
    let updatedCount = 0;

    // Update assertion count in list items
    visit(tree, 'listItem', (node) => {
      visit(node, 'text', (textNode) => {
        const text = textNode.value;
        
        // Match: ✅ **148+ test assertions**
        const assertionMatch = text.match(/✅ \*\*(\d+)\+ test assertions\*\*/);
        if (assertionMatch) {
          const oldCount = assertionMatch[1];
          if (oldCount !== String(stats.totalAssertions)) {
            textNode.value = text.replace(
              /✅ \*\*\d+\+ test assertions\*\*/,
              `✅ **${stats.totalAssertions}+ test assertions**`
            );
            updatedCount++;
          }
        }
      });
    });

    // Update category counts in headings
    visit(tree, 'text', (textNode) => {
      const text = textNode.value;
      const updates = [
        { pattern: /\*\*Core Parsing\*\* \((\d+) tests\)/, count: stats.categories['core-parsing'].total },
        { pattern: /\*\*Advanced Processing\*\* \((\d+) tests\)/, count: stats.categories['advanced-processing'].total },
        { pattern: /\*\*Object Construction\*\* \((\d+) tests\)/, count: stats.categories['object-construction'].total },
        { pattern: /\*\*Type System\*\* \((\d+) tests\)/, count: stats.categories['type-system'].total },
        { pattern: /\*\*Output & Validation\*\* \((\d+) tests\)/, count: stats.categories['output-validation'].total }
      ];

      for (const update of updates) {
        const match = text.match(update.pattern);
        if (match) {
          const oldCount = match[1];
          if (oldCount !== String(update.count)) {
            textNode.value = text.replace(update.pattern, (fullMatch) => {
              return fullMatch.replace(/\(\d+ tests\)/, `(${update.count} tests)`);
            });
            updatedCount++;
          }
        }
      }
    });

    // Return metadata about updates made
    tree._updateInfo = { updatedCount };
    return tree;
  };
}

async function updateReadme() {
  const isTTY = process.stdout.isTTY;
  
  if (isTTY) {
    console.log(`${colors.magenta}${colors.bold}📝 CCL README Updater (Remark)${colors.reset}`);
    console.log(`${colors.gray}Updating README.md with current test statistics using AST-based approach${colors.reset}\n`);
  } else {
    console.log('📝 Updating README.md with current test statistics...');
  }
  
  try {
    // Collect current stats using Go stats command
    const statsOutput = execSync('go run ./cmd/ccl-test-runner stats --format json', { 
      encoding: 'utf8',
      cwd: ROOT_DIR,
      stdio: ['ignore', 'pipe', 'ignore'] // Ignore stderr to suppress status messages
    });
    const stats = JSON.parse(statsOutput);
    
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
      console.log(`  ${colors.gray}Total: ${total} tests (${stats.totalAssertions} assertions)${colors.reset}\n`);
    } else {
      console.log(`📊 Current counts: CoreParsing:${coreParsingTotal} AdvancedProcessing:${advancedProcessingTotal} ObjectConstruction:${objectConstructionTotal} TypeSystem:${typeSystemTotal} OutputValidation:${outputValidationTotal} Total:${total} Assertions:${stats.totalAssertions}`);
    }
    
    // Read current README
    const readme = await readFile(README_PATH, 'utf8');
    
    // Create unified processor with our custom plugin
    const processor = unified()
      .use(remarkParse)
      .use(remarkGfm)
      .use(remarkUpdateCounts, stats)
      .use(remarkStringify, {
        bullet: '-',
        emphasis: '*',
        strong: '*',
        listItemIndent: 'one',
        fences: true,
        incrementListMarker: false
      });

    // Process the README
    const result = await processor.process(readme);
    const updatedReadme = String(result);
    const updateInfo = result.data.tree?._updateInfo || { updatedCount: 0 };
    
    // Check if anything changed
    if (updatedReadme === readme) {
      if (isTTY) {
        console.log(`${colors.green}✅ README.md is already up to date${colors.reset}`);
      } else {
        console.log('✅ README.md is already up to date');
      }
      return { updated: false };
    }
    
    // Write updated README
    await writeFile(README_PATH, updatedReadme);
    
    if (isTTY) {
      console.log(`${colors.green}✅ README.md updated successfully${colors.reset}`);
      console.log(`${colors.blue}${colors.bold}🔄 AST-based Updates Applied: ${updateInfo.updatedCount}${colors.reset}`);
    } else {
      console.log('✅ README.md updated successfully');
      console.log(`🔄 AST-based updates: ${updateInfo.updatedCount}`);
    }
    
    return { updated: true, updatedCount: updateInfo.updatedCount };
    
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
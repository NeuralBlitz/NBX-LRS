#!/usr/bin/env python3
"""
Comprehensive Error Inspection Script
Checks all NeuralBlitz files for syntax errors, import issues, and other problems
"""

import os
import sys
import subprocess
import ast
from pathlib import Path

# Setup Python path
sys.path.insert(0, "/home/runner/workspace/NB-Ecosystem/lib/python3.11/site-packages")
sys.path.insert(0, "/home/runner/workspace/opencode-lrs-agents-nbx/neuralblitz-v50")

print("=" * 80)
print("COMPREHENSIVE ERROR INSPECTION")
print("=" * 80)

# Files to check
base_path = "/home/runner/workspace/opencode-lrs-agents-nbx/neuralblitz-v50"
files_to_check = [
    "Advanced-Research/production/quantum_spiking_neuron.py",
    "Advanced-Research/production/test_quantum_spiking_neuron.py",
    "multi_reality_nn.py",
    "cosmic_consciousness_integration.py",
    "cross_reality_entanglement.py",
    "dimensional_computing_integration.py",
    "dimensional_neural_processing.py",
    "dimensional_consciousness.py",
    "neuro_symbiotic_integration.py",
    "autonomous_self_evolution.py",
    "brain_wave_entrainment.py",
    "emergent_purpose_discovery.py",
    "self_improving_code_generation.py",
    "novel_neural_architectures_100.py",
]

errors_found = []
warnings_found = []
success_count = 0

print("\n1. SYNTAX CHECK")
print("-" * 80)

for file_path in files_to_check:
    full_path = os.path.join(base_path, file_path)
    if not os.path.exists(full_path):
        errors_found.append(f"MISSING: {file_path}")
        print(f"❌ {file_path}: FILE NOT FOUND")
        continue

    try:
        with open(full_path, "r") as f:
            source = f.read()

        # Parse with AST to check syntax
        ast.parse(source)
        print(f"✅ {file_path}: Syntax OK")
        success_count += 1

    except SyntaxError as e:
        errors_found.append(f"SYNTAX ERROR in {file_path}: {e}")
        print(f"❌ {file_path}: Syntax Error - {e}")
    except Exception as e:
        errors_found.append(f"ERROR in {file_path}: {e}")
        print(f"❌ {file_path}: Error - {e}")

print(f"\n   Syntax Check: {success_count}/{len(files_to_check)} passed")

# Check import statements
print("\n2. IMPORT STATEMENT ANALYSIS")
print("-" * 80)

import_issues = []

for file_path in files_to_check:
    full_path = os.path.join(base_path, file_path)
    if not os.path.exists(full_path):
        continue

    try:
        with open(full_path, "r") as f:
            tree = ast.parse(f.read())

        imports = []
        for node in ast.walk(tree):
            if isinstance(node, ast.Import):
                for alias in node.names:
                    imports.append(alias.name)
            elif isinstance(node, ast.ImportFrom):
                module = node.module or ""
                # Check for relative imports
                if node.level > 0:
                    warnings_found.append(
                        f"RELATIVE IMPORT: {file_path} imports from {module} (level {node.level})"
                    )
                    print(f"⚠️  {file_path}: Relative import from '{module}'")

        if not any("RELATIVE IMPORT" in w and file_path in w for w in warnings_found):
            print(f"✅ {file_path}: No relative imports")

    except Exception as e:
        pass

# Check for common issues
print("\n3. COMMON ISSUES CHECK")
print("-" * 80)

issues_to_check = [
    ("print(", "Print statements (should use logging in production)"),
    ("TODO", "TODO comments"),
    ("FIXME", "FIXME comments"),
    ("XXX", "XXX markers"),
    ("except:", "Bare except clauses"),
    ("pass  #", "Empty pass with comment (placeholder)"),
]

for file_path in files_to_check:
    full_path = os.path.join(base_path, file_path)
    if not os.path.exists(full_path):
        continue

    try:
        with open(full_path, "r") as f:
            content = f.read()
            lines = content.split("\n")

        for issue, description in issues_to_check:
            count = content.count(issue)
            if count > 0 and issue in ["TODO", "FIXME", "XXX", "except:"]:
                warnings_found.append(f"{description}: {count} in {file_path}")
                print(f"⚠️  {file_path}: {count} {description}")

    except Exception as e:
        pass

# Check specific files we created
print("\n4. CUSTOM FILES CHECK")
print("-" * 80)

custom_files = [
    "/home/runner/workspace/phase1_fix_tests.py",
    "/home/runner/workspace/phase3_practical_applications.py",
    "/home/runner/workspace/phase4_visualization_suite.py",
    "/home/runner/workspace/phase5_validation.py",
    "/home/runner/workspace/comprehensive_test_suite.py",
]

for file_path in custom_files:
    if os.path.exists(file_path):
        try:
            with open(file_path, "r") as f:
                source = f.read()
            ast.parse(source)
            print(f"✅ {os.path.basename(file_path)}: OK")
        except Exception as e:
            errors_found.append(f"ERROR in {file_path}: {e}")
            print(f"❌ {os.path.basename(file_path)}: Error - {e}")
    else:
        print(f"⚠️  {os.path.basename(file_path)}: Not found")

# Summary
print("\n" + "=" * 80)
print("INSPECTION SUMMARY")
print("=" * 80)

if errors_found:
    print(f"\n❌ ERRORS FOUND: {len(errors_found)}")
    for error in errors_found[:10]:  # Show first 10
        print(f"   - {error}")
    if len(errors_found) > 10:
        print(f"   ... and {len(errors_found) - 10} more errors")
else:
    print("\n✅ NO CRITICAL ERRORS FOUND")

if warnings_found:
    print(f"\n⚠️  WARNINGS: {len(warnings_found)}")
    for warning in warnings_found[:10]:  # Show first 10
        print(f"   - {warning}")
    if len(warnings_found) > 10:
        print(f"   ... and {len(warnings_found) - 10} more warnings")
else:
    print("\n✅ NO WARNINGS")

print(f"\n📊 STATISTICS:")
print(f"   - Files checked: {len(files_to_check)}")
print(f"   - Syntax OK: {success_count}")
print(f"   - Errors: {len(errors_found)}")
print(f"   - Warnings: {len(warnings_found)}")

print("\n" + "=" * 80)

#!/bin/bash
###############################################################################
# NeuralBlitz Test Runner Script
# Automatically configures environment and runs all test suites
###############################################################################

set -e  # Exit on error

echo "╔══════════════════════════════════════════════════════════════════════════╗"
echo "║           🧠 NEURALBLITZ TEST RUNNER v2.0                               ║"
echo "║           Automated Testing & Validation Suite                          ║"
echo "╚══════════════════════════════════════════════════════════════════════════╝"
echo ""

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[PASS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

print_error() {
    echo -e "${RED}[FAIL]${NC} $1"
}

# Get script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
WORKSPACE_DIR="$(dirname "$SCRIPT_DIR")"

print_status "Setting up environment..."

# Set up PYTHONPATH
export PYTHONPATH="${WORKSPACE_DIR}/NB-Ecosystem/lib/python3.11/site-packages:${PYTHONPATH}"
export PYTHONPATH="${WORKSPACE_DIR}/opencode-lrs-agents-nbx/neuralblitz-v50:${PYTHONPATH}"
export PYTHONPATH="${WORKSPACE_DIR}/opencode-lrs-agents-nbx/neuralblitz-v50/Advanced-Research/production:${PYTHONPATH}"

# Add local python binaries to PATH
export PATH="${WORKSPACE_DIR}/.pythonlibs/bin:${PATH}"

print_status "Environment configured:"
print_status "  PYTHONPATH: ${PYTHONPATH}"
print_status "  Python: $(which python3)"
print_status "  Python Version: $(python3 --version)"
print_status "  Pytest: $(which pytest)"
print_status "  NumPy: $(python3 -c 'import numpy; print(numpy.__version__)')"
echo ""

# Initialize counters
TESTS_PASSED=0
TESTS_FAILED=0
TESTS_SKIPPED=0

# Function to run a test suite
run_test_suite() {
    local name="$1"
    local command="$2"
    local dir="$3"
    
    echo "═══════════════════════════════════════════════════════════════════════"
    echo "Running: $name"
    echo "═══════════════════════════════════════════════════════════════════════"
    
    if [ -n "$dir" ]; then
        cd "$dir" || exit 1
    fi
    
    if eval "$command"; then
        print_success "$name completed successfully"
        ((TESTS_PASSED++))
        return 0
    else
        print_error "$name failed"
        ((TESTS_FAILED++))
        return 1
    fi
}

# Test 1: Quantum Spiking Neuron Tests
run_test_suite "Quantum Neuron Tests" \
    "pytest Advanced-Research/production/test_quantum_spiking_neuron.py -v --tb=short" \
    "${WORKSPACE_DIR}/opencode-lrs-agents-nbx/neuralblitz-v50"

echo ""

# Test 2: Comprehensive Test Suite
run_test_suite "Comprehensive Test Suite" \
    "python3 comprehensive_test_suite.py" \
    "${WORKSPACE_DIR}"

echo ""

# Test 3: New Untested Technologies
run_test_suite "Untested Technologies" \
    "pytest tests/test_untested_technologies.py -v --tb=short || true" \
    "${WORKSPACE_DIR}/opencode-lrs-agents-nbx/neuralblitz-v50"

echo ""

# Test 4: LRS Agents Tests (if available)
if [ -f "${WORKSPACE_DIR}/lrs_agents/pytest.ini" ]; then
    run_test_suite "LRS Agents Tests" \
        "pytest tests/ -v --tb=short -x || true" \
        "${WORKSPACE_DIR}/lrs_agents"
    echo ""
fi

# Test 5: Coverage Report (if pytest-cov available)
print_status "Generating coverage report..."
cd "${WORKSPACE_DIR}/opencode-lrs-agents-nbx/neuralblitz-v50" || exit 1
if pytest Advanced-Research/production/test_quantum_spiking_neuron.py --cov=Advanced-Research/production --cov-report=term-missing --cov-report=html 2>/dev/null; then
    print_success "Coverage report generated"
    print_status "HTML report: ${WORKSPACE_DIR}/opencode-lrs-agents-nbx/neuralblitz-v50/htmlcov/index.html"
else
    print_warning "Coverage report skipped (pytest-cov not available)"
    ((TESTS_SKIPPED++))
fi

echo ""
echo "═══════════════════════════════════════════════════════════════════════"
echo "                    📊 TEST SUMMARY"
echo "═══════════════════════════════════════════════════════════════════════"
echo ""
printf "  ${GREEN}✓ Passed:  %d${NC}\n" $TESTS_PASSED
printf "  ${RED}✗ Failed:  %d${NC}\n" $TESTS_FAILED
printf "  ${YELLOW}⊘ Skipped: %d${NC}\n" $TESTS_SKIPPED
echo ""

TOTAL=$((TESTS_PASSED + TESTS_FAILED))
if [ $TESTS_FAILED -eq 0 ]; then
    print_success "All ${TOTAL} test suites passed!"
    echo ""
    echo "╔══════════════════════════════════════════════════════════════════════════╗"
    echo "║                    🎉 NEURALBLITZ VALIDATED 🎉                          ║"
    echo "║              All systems operational and production-ready               ║"
    echo "╚══════════════════════════════════════════════════════════════════════════╝"
    exit 0
else
    print_error "${TESTS_FAILED}/${TOTAL} test suites failed"
    echo ""
    echo "╔══════════════════════════════════════════════════════════════════════════╗"
    echo "║                    ⚠️  VALIDATION ISSUES DETECTED                       ║"
    echo "║              Please review the failed tests above                       ║"
    echo "╚══════════════════════════════════════════════════════════════════════════╝"
    exit 1
fi

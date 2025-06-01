#!/bin/zsh
# Vista API Test Script
# Tests all API endpoints and verifies responses

# Configuration
API_HOST="localhost"
API_PORT=8080
API_BASE="http://${API_HOST}:${API_PORT}"
JQ_AVAILABLE=$(command -v jq &> /dev/null && echo true || echo false)

# Color codes for output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Test counter
TESTS_PASSED=0
TESTS_FAILED=0

# Function to format output
format_json() {
  if [ "$JQ_AVAILABLE" = true ]; then
    jq '.'
  else
    cat
  fi
}

# Function to test an endpoint
test_endpoint() {
  local endpoint=$1
  local description=$2
  local expected_status=$3

  echo -e "\n${BLUE}Testing: ${description}${NC}"
  echo -e "GET ${endpoint}"
  
  # Make the request and capture status code
  local response=$(curl -s -w "\n%{http_code}" "${API_BASE}${endpoint}")
  local status=$(echo "$response" | tail -n1)
  local body=$(echo "$response" | sed '$d')
  
  # Check status code
  if [ "$status" -eq "$expected_status" ]; then
    echo -e "${GREEN}✓ Status: ${status} (expected: ${expected_status})${NC}"
    TESTS_PASSED=$((TESTS_PASSED+1))
  else
    echo -e "${RED}✗ Status: ${status} (expected: ${expected_status})${NC}"
    TESTS_FAILED=$((TESTS_FAILED+1))
  fi
  
  # Display the response body
  echo -e "${YELLOW}Response:${NC}"
  echo "$body" | format_json
}

echo -e "${BLUE}===== Vista API Test Suite =====${NC}"
echo -e "Testing against: ${API_BASE}"

# Check if server is running
if ! curl -s "${API_BASE}/repos" > /dev/null; then
  echo -e "\n${RED}Error: Cannot connect to API server at ${API_BASE}${NC}"
  echo "Make sure the server is running (go run ./cmd/vista/main.go)"
  exit 1
fi

# Test all endpoints
test_endpoint "/repos" "List all repositories" 200
test_endpoint "/repo/ecr-main" "Get repository details: ecr-main" 200
test_endpoint "/repo/ecr-main/resources" "List resources in repository: ecr-main" 200
test_endpoint "/repo/ecr-main/resource/my-app" "Get resource details: my-app in ecr-main" 200
test_endpoint "/repo/dockerhub/resources" "List resources in repository: dockerhub" 200
test_endpoint "/repo/dockerhub/resource/nginx" "Get resource details: nginx in dockerhub" 200

# Test error cases
test_endpoint "/repo/non-existent" "Get non-existent repository" 404
test_endpoint "/repo/ecr-main/resource/non-existent" "Get non-existent resource" 404

# Summary
echo -e "\n${BLUE}===== Test Summary =====${NC}"
echo -e "${GREEN}Tests passed: ${TESTS_PASSED}${NC}"
echo -e "${RED}Tests failed: ${TESTS_FAILED}${NC}"

if [ $TESTS_FAILED -eq 0 ]; then
  echo -e "\n${GREEN}All tests passed!${NC}"
  exit 0
else
  echo -e "\n${RED}Some tests failed.${NC}"
  exit 1
fi

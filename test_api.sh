#!/bin/bash

# IPENPOTO API - Testing Script
# Complete API testing automation

set -e  # Exit on error

# Color codes
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
API_URL="${API_URL:-http://localhost:8080}"
API_BASE="$API_URL/api"
TEST_USER="${TEST_USER:-testuser}"
TEST_PASSWORD="${TEST_PASSWORD:-Password123!@#}"

# Counters
TESTS_RUN=0
TESTS_PASSED=0
TESTS_FAILED=0

# Helper functions
print_header() {
    echo -e "\n${BLUE}═══════════════════════════════════════════════════${NC}"
    echo -e "${BLUE}$1${NC}"
    echo -e "${BLUE}═══════════════════════════════════════════════════${NC}\n"
}

print_test() {
    echo -e "${YELLOW}▶ Testing: $1${NC}"
}

print_success() {
    echo -e "${GREEN}✓ $1${NC}"
    ((TESTS_PASSED++))
}

print_error() {
    echo -e "${RED}✗ $1${NC}"
    ((TESTS_FAILED++))
}

# Main testing functions
test_api_health() {
    print_header "API HEALTH CHECK"

    print_test "Health check endpoint"
    RESPONSE=$(curl -s -w "\n%{http_code}" "$API_URL/")
    HTTP_CODE=$(echo "$RESPONSE" | tail -1)
    BODY=$(echo "$RESPONSE" | sed '$d')

    if [ "$HTTP_CODE" = "200" ]; then
        print_success "API is running"
        echo "Response: $BODY"
    else
        print_error "API health check failed (HTTP $HTTP_CODE)"
        return 1
    fi
    ((TESTS_RUN++))
}

test_login() {
    print_header "AUTHENTICATION TESTS"

    print_test "Valid login"
    RESPONSE=$(curl -s -X POST "$API_BASE/auth/login" \
        -H "Content-Type: application/json" \
        -d "{\"username\":\"$TEST_USER\",\"password\":\"$TEST_PASSWORD\"}" \
        -w "\n%{http_code}")

    HTTP_CODE=$(echo "$RESPONSE" | tail -1)
    BODY=$(echo "$RESPONSE" | sed '$d')

    if [ "$HTTP_CODE" = "200" ]; then
        print_success "Login successful"

        # Extract tokens
        ACCESS_TOKEN=$(echo "$BODY" | grep -o '"access_token":"[^"]*' | head -1 | cut -d'"' -f4)
        REFRESH_TOKEN=$(echo "$BODY" | grep -o '"refresh_token":"[^"]*' | head -1 | cut -d'"' -f4)

        if [ -n "$ACCESS_TOKEN" ]; then
            print_success "Access token received"
            echo "Token: ${ACCESS_TOKEN:0:50}..."
        else
            print_error "No access token in response"
        fi

        # Save for later use
        echo "$ACCESS_TOKEN" > /tmp/access_token.txt
        echo "$REFRESH_TOKEN" > /tmp/refresh_token.txt
    else
        print_error "Login failed (HTTP $HTTP_CODE)"
        echo "Response: $BODY"
    fi
    ((TESTS_RUN++))
}

test_invalid_login() {
    print_header "INVALID LOGIN TESTS"

    print_test "Invalid password"
    RESPONSE=$(curl -s -X POST "$API_BASE/auth/login" \
        -H "Content-Type: application/json" \
        -d "{\"username\":\"$TEST_USER\",\"password\":\"WrongPassword\"}" \
        -w "\n%{http_code}")

    HTTP_CODE=$(echo "$RESPONSE" | tail -1)

    if [ "$HTTP_CODE" = "401" ]; then
        print_success "Correctly rejected invalid credentials"
    else
        print_error "Should reject invalid credentials"
    fi
    ((TESTS_RUN++))

    print_test "Missing password"
    RESPONSE=$(curl -s -X POST "$API_BASE/auth/login" \
        -H "Content-Type: application/json" \
        -d "{\"username\":\"$TEST_USER\"}" \
        -w "\n%{http_code}")

    HTTP_CODE=$(echo "$RESPONSE" | tail -1)

    if [ "$HTTP_CODE" = "400" ]; then
        print_success "Correctly rejected missing password"
    else
        print_error "Should reject missing password"
    fi
    ((TESTS_RUN++))
}

test_protected_routes() {
    print_header "PROTECTED ROUTES TESTS"

    if [ ! -f /tmp/access_token.txt ]; then
        print_error "No access token found. Run login test first."
        return 1
    fi

    ACCESS_TOKEN=$(cat /tmp/access_token.txt)

    print_test "Access without token"
    RESPONSE=$(curl -s -X GET "$API_BASE/user/profile" \
        -w "\n%{http_code}")

    HTTP_CODE=$(echo "$RESPONSE" | tail -1)

    if [ "$HTTP_CODE" = "401" ]; then
        print_success "Correctly requires authentication"
    else
        print_error "Should require authentication"
    fi
    ((TESTS_RUN++))

    print_test "Access with valid token"
    RESPONSE=$(curl -s -X GET "$API_BASE/user/profile" \
        -H "Authorization: Bearer $ACCESS_TOKEN" \
        -w "\n%{http_code}")

    HTTP_CODE=$(echo "$RESPONSE" | tail -1)
    BODY=$(echo "$RESPONSE" | sed '$d')

    if [ "$HTTP_CODE" = "200" ]; then
        print_success "Successfully accessed protected route"
        echo "User data: $(echo $BODY | jq '.data.username' 2>/dev/null || echo 'N/A')"
    else
        print_error "Failed to access protected route (HTTP $HTTP_CODE)"
    fi
    ((TESTS_RUN++))
}

test_rate_limiting() {
    print_header "RATE LIMITING TESTS"

    print_test "Check rate limit headers"
    RESPONSE=$(curl -s -i -X POST "$API_BASE/auth/login" \
        -H "Content-Type: application/json" \
        -d "{\"username\":\"$TEST_USER\",\"password\":\"$TEST_PASSWORD\"}")

    RATE_LIMIT=$(echo "$RESPONSE" | grep -i "X-RateLimit-Limit" | cut -d' ' -f2 | tr -d '\r')
    REMAINING=$(echo "$RESPONSE" | grep -i "X-RateLimit-Remaining" | cut -d' ' -f2 | tr -d '\r')

    if [ -n "$RATE_LIMIT" ]; then
        print_success "Rate limit headers present"
        echo "Limit: $RATE_LIMIT, Remaining: $REMAINING"
    else
        print_error "Rate limit headers missing"
    fi
    ((TESTS_RUN++))
}

test_cors() {
    print_header "CORS TESTS"

    print_test "CORS headers present"
    RESPONSE=$(curl -s -i -X OPTIONS "$API_BASE/auth/login" \
        -H "Origin: http://localhost:3000")

    ALLOWED_ORIGIN=$(echo "$RESPONSE" | grep -i "Access-Control-Allow-Origin" | cut -d' ' -f2 | tr -d '\r')

    if [ -n "$ALLOWED_ORIGIN" ]; then
        print_success "CORS headers present"
        echo "Allowed Origin: $ALLOWED_ORIGIN"
    else
        print_error "CORS headers missing"
    fi
    ((TESTS_RUN++))
}

print_summary() {
    print_header "TEST SUMMARY"

    echo "Tests run:    $TESTS_RUN"
    echo -e "Tests passed: ${GREEN}$TESTS_PASSED${NC}"
    echo -e "Tests failed: ${RED}$TESTS_FAILED${NC}"

    if [ $TESTS_FAILED -eq 0 ]; then
        echo -e "\n${GREEN}✓ All tests passed!${NC}\n"
        return 0
    else
        echo -e "\n${RED}✗ Some tests failed${NC}\n"
        return 1
    fi
}

# Main execution
main() {
    clear
    echo -e "${BLUE}╔════════════════════════════════════════════════════╗${NC}"
    echo -e "${BLUE}║         IPENPOTO API - Automated Testing           ║${NC}"
    echo -e "${BLUE}╚════════════════════════════════════════════════════╝${NC}"

    echo -e "\n${YELLOW}Configuration:${NC}"
    echo "API URL: $API_URL"
    echo "Test User: $TEST_USER"
    echo "Test Password: ****"

    # Run tests
    test_api_health
    test_login
    test_invalid_login
    test_protected_routes
    test_rate_limiting
    test_cors

    # Print summary
    print_summary
}

# Run main function
main "$@"

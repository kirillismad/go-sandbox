#!/bin/bash

# Repository Setup Script
# This script helps set up branch protection and other repository settings
# Requires GitHub CLI (gh) to be installed and authenticated

set -e

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
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if GitHub CLI is installed
check_gh_cli() {
    if ! command -v gh &> /dev/null; then
        print_error "GitHub CLI (gh) is not installed. Please install it first:"
        echo "  https://cli.github.com/"
        exit 1
    fi
    
    # Check if authenticated
    if ! gh auth status &> /dev/null; then
        print_error "GitHub CLI is not authenticated. Please run 'gh auth login' first."
        exit 1
    fi
    
    print_success "GitHub CLI is installed and authenticated"
}

# Get repository information
get_repo_info() {
    if ! REPO_INFO=$(gh repo view --json owner,name 2>/dev/null); then
        print_error "Could not get repository information. Make sure you're in a Git repository with a GitHub remote."
        exit 1
    fi
    
    OWNER=$(echo "$REPO_INFO" | jq -r '.owner.login')
    REPO=$(echo "$REPO_INFO" | jq -r '.name')
    
    print_status "Repository: $OWNER/$REPO"
}

# Setup branch protection for master branch
setup_branch_protection() {
    print_status "Setting up branch protection for master branch..."
    
    # Check if master branch exists
    if ! git show-ref --verify --quiet refs/heads/master; then
        print_warning "Master branch does not exist locally. Checking remote..."
        if ! git ls-remote --exit-code --heads origin master &> /dev/null; then
            print_error "Master branch does not exist. Please create it first or change the script to use 'main' branch."
            return 1
        fi
    fi
    
    # Create branch protection rule
    gh api repos/$OWNER/$REPO/branches/master/protection \
        --method PUT \
        --field required_status_checks='{"strict":true,"contexts":["test","lint","build","security","validate-dependencies"]}' \
        --field enforce_admins=true \
        --field required_pull_request_reviews='{"required_approving_review_count":1,"dismiss_stale_reviews":true,"require_code_owner_reviews":true}' \
        --field restrictions=null \
        --field required_conversation_resolution=true \
        --silent
        
    print_success "Branch protection rules applied to master branch"
}

# Enable repository security features
enable_security_features() {
    print_status "Enabling security features..."
    
    # Enable vulnerability alerts
    gh api repos/$OWNER/$REPO \
        --method PATCH \
        --field has_vulnerability_alerts=true \
        --silent
    
    # Enable automated security fixes (Dependabot)
    gh api repos/$OWNER/$REPO/automated-security-fixes \
        --method PUT \
        --silent
    
    print_success "Security features enabled"
}

# Create or update repository description
update_repository_settings() {
    print_status "Updating repository settings..."
    
    gh api repos/$OWNER/$REPO \
        --method PATCH \
        --field description="A comprehensive Go development sandbox with examples, algorithms, and various integrations" \
        --field has_issues=true \
        --field has_projects=true \
        --field has_wiki=true \
        --field delete_branch_on_merge=true \
        --field allow_squash_merge=true \
        --field allow_merge_commit=true \
        --field allow_rebase_merge=false \
        --silent
    
    print_success "Repository settings updated"
}

# Display current protection status
show_protection_status() {
    print_status "Current branch protection status for master:"
    
    if PROTECTION_STATUS=$(gh api repos/$OWNER/$REPO/branches/master/protection 2>/dev/null); then
        echo "$PROTECTION_STATUS" | jq '.required_status_checks.contexts[]' 2>/dev/null || echo "No required status checks"
        
        if echo "$PROTECTION_STATUS" | jq -e '.required_pull_request_reviews' > /dev/null 2>&1; then
            REQUIRED_REVIEWS=$(echo "$PROTECTION_STATUS" | jq -r '.required_pull_request_reviews.required_approving_review_count')
            print_success "Pull request reviews required: $REQUIRED_REVIEWS"
        fi
        
        if echo "$PROTECTION_STATUS" | jq -e '.enforce_admins.enabled' > /dev/null 2>&1; then
            print_success "Admin enforcement: enabled"
        fi
    else
        print_warning "No branch protection found for master branch"
    fi
}

# Main function
main() {
    echo "🛡️  Repository Branch Protection Setup"
    echo "======================================"
    echo
    
    check_gh_cli
    get_repo_info
    
    echo
    print_status "This script will set up the following:"
    echo "  • Branch protection rules for master branch"
    echo "  • Required status checks (CI/CD pipeline)"
    echo "  • Pull request review requirements"
    echo "  • Security features (vulnerability alerts, Dependabot)"
    echo "  • Repository settings optimization"
    echo
    
    read -p "Continue? (y/N): " -n 1 -r
    echo
    
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        print_status "Setup cancelled"
        exit 0
    fi
    
    echo
    print_status "Setting up repository protection..."
    
    # Show current status first
    show_protection_status
    echo
    
    # Apply settings
    setup_branch_protection
    enable_security_features
    update_repository_settings
    
    echo
    print_success "Repository setup completed successfully!"
    echo
    
    print_status "Next steps:"
    echo "  1. Review the protection rules in GitHub repository settings"
    echo "  2. Ensure all team members understand the new workflow"
    echo "  3. Update any automation that directly pushes to master"
    echo "  4. Consider setting up additional status checks as needed"
    echo
    
    print_status "For more information, see: docs/BRANCH_PROTECTION.md"
}

# Run main function
main "$@"
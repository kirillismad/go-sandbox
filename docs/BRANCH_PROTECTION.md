# Branch Protection Guide

This document explains how to manage who can push to the master branch in this repository.

## Overview

Branch protection rules help maintain code quality and enforce collaboration workflows by restricting direct pushes to important branches like `master` (or `main`). Instead of allowing direct pushes, contributors must create pull requests that can be reviewed before merging.

## Setting Up Branch Protection Rules

### Via GitHub Web Interface

1. **Navigate to Repository Settings**
   - Go to your repository on GitHub
   - Click on the "Settings" tab
   - Select "Branches" from the left sidebar

2. **Add Branch Protection Rule**
   - Click "Add rule" button
   - Enter `master` (or `main`) as the branch name pattern
   - Configure the following recommended settings:

#### Required Settings
- ✅ **Require a pull request before merging**
  - ✅ Require approvals (recommended: 1 or more)
  - ✅ Dismiss stale pull request approvals when new commits are pushed
  - ✅ Require review from code owners (if CODEOWNERS file exists)

- ✅ **Require status checks to pass before merging**
  - ✅ Require branches to be up to date before merging
  - Select specific status checks (CI/CD workflows)

- ✅ **Require conversation resolution before merging**

#### Additional Security Settings
- ✅ **Restrict pushes that create files**
- ✅ **Require signed commits** (recommended for enhanced security)
- ✅ **Include administrators** (applies rules to repository administrators)
- ✅ **Allow force pushes** (❌ disable for maximum protection)
- ✅ **Allow deletions** (❌ disable to prevent accidental branch deletion)

### Via GitHub CLI

```bash
# Install GitHub CLI if not already installed
# https://cli.github.com/

# Create branch protection rule
gh api repos/:owner/:repo/branches/master/protection \
  --method PUT \
  --field required_status_checks='{"strict":true,"contexts":[]}' \
  --field enforce_admins=true \
  --field required_pull_request_reviews='{"required_approving_review_count":1,"dismiss_stale_reviews":true}' \
  --field restrictions=null
```

### Via API

```bash
curl -X PUT \
  -H "Accept: application/vnd.github.v3+json" \
  -H "Authorization: token YOUR_TOKEN" \
  https://api.github.com/repos/OWNER/REPO/branches/master/protection \
  -d '{
    "required_status_checks": {
      "strict": true,
      "contexts": []
    },
    "enforce_admins": true,
    "required_pull_request_reviews": {
      "required_approving_review_count": 1,
      "dismiss_stale_reviews": true,
      "require_code_owner_reviews": true
    },
    "restrictions": null
  }'
```

## Managing Access Permissions

### Repository Collaborators

1. **Adding Collaborators**
   - Go to Settings → Manage access
   - Click "Invite a collaborator"
   - Set appropriate permission level:
     - **Read**: Can view and clone
     - **Triage**: Can manage issues and pull requests
     - **Write**: Can push, but subject to branch protection
     - **Maintain**: Can manage repository settings
     - **Admin**: Full access (can bypass branch protection if not configured)

2. **Team Access** (for Organizations)
   - Go to Settings → Manage access
   - Add teams with appropriate permission levels
   - Teams inherit permissions from organization membership

### Code Owners

Create a `CODEOWNERS` file to automatically request reviews from specific people or teams:

```bash
# Example CODEOWNERS file
# Global owners
* @username @team-name

# Specific paths
/src/ @backend-team
/docs/ @documentation-team
*.go @go-experts
```

## Workflow Examples

### Standard Contribution Workflow

1. **Fork the repository** (for external contributors)
2. **Create a feature branch**
   ```bash
   git checkout -b feature/my-new-feature
   ```
3. **Make changes and commit**
   ```bash
   git add .
   git commit -m "Add new feature"
   ```
4. **Push to your branch**
   ```bash
   git push origin feature/my-new-feature
   ```
5. **Create a Pull Request**
   - Go to GitHub and create a PR from your branch to master
   - Request review from appropriate team members
   - Ensure all status checks pass
6. **Review and Merge**
   - Reviewers approve the changes
   - Merge the PR (only possible after meeting all requirements)

### Emergency Hotfix Workflow

For critical production fixes, you may need a streamlined process:

1. Create hotfix branch from master
2. Make minimal necessary changes
3. Create PR with "hotfix" label
4. Fast-track review process
5. Merge after required approvals

## Best Practices

### For Repository Administrators

1. **Always enable branch protection** on main branches
2. **Include administrators** in branch protection rules
3. **Require status checks** for automated testing
4. **Use CODEOWNERS** for automatic review assignment
5. **Regularly audit** repository access permissions
6. **Enable security features** like required signed commits
7. **Monitor repository activity** through audit logs

### For Contributors

1. **Never force push** to protected branches
2. **Keep pull requests small** and focused
3. **Write descriptive commit messages**
4. **Respond promptly** to review feedback
5. **Test changes locally** before creating PR
6. **Update documentation** when needed

### For Reviewers

1. **Review code thoroughly** for security and quality
2. **Test changes locally** when possible
3. **Provide constructive feedback**
4. **Approve only when confident** in the changes
5. **Check for breaking changes**
6. **Verify documentation updates**

## Troubleshooting

### Common Issues

**"Push to master blocked"**
- Solution: Create a pull request instead of pushing directly

**"Required status checks must pass"**
- Solution: Fix failing tests/workflows before merging

**"Pull request needs approval"**
- Solution: Request review from team members or code owners

**"Branch is not up to date"**
- Solution: Rebase or merge latest master into your branch

### Bypassing Protection (Emergency Only)

Repository administrators can temporarily disable branch protection for emergency situations:

1. Go to Settings → Branches
2. Edit the protection rule
3. Temporarily disable specific requirements
4. **Remember to re-enable** after the emergency

## Monitoring and Auditing

### Repository Activity

Monitor repository activity through:
- **Insights tab** for contribution graphs
- **Security tab** for alerts and analysis
- **Actions tab** for workflow runs
- **Settings → Audit log** for administrative changes

### Automated Monitoring

Consider setting up automated monitoring for:
- Failed status checks
- Direct pushes to protected branches
- Changes to branch protection rules
- Unusual repository activity

## Integration with CI/CD

### GitHub Actions Integration

Example workflow that integrates with branch protection:

```yaml
name: CI
on:
  pull_request:
    branches: [ master ]
  push:
    branches: [ master ]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3
    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version: 1.21
    - name: Run tests
      run: go test ./...
    - name: Run linter
      run: golangci-lint run
```

This workflow will run on every pull request to master and must pass before merging is allowed.

## Conclusion

Proper branch protection is essential for maintaining code quality and security. By following this guide, you can ensure that only reviewed and tested code makes it into your master branch, while still maintaining an efficient development workflow.

For more information, see:
- [GitHub Branch Protection Documentation](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/defining-the-mergeability-of-pull-requests/about-protected-branches)
- [GitHub Permissions Documentation](https://docs.github.com/en/organizations/managing-access-to-your-organizations-repositories/repository-permission-levels-for-an-organization)
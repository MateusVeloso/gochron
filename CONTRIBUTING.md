# Contributing to gochron

First of all, thank you for considering contributing to **gochron**! We appreciate your interest in making this library better. The following guidelines will help ensure a smooth process for everyone involved.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [How to Contribute](#how-to-contribute)
- [Setting Up the Development Environment](#setting-up-the-development-environment)
- [Reporting Issues](#reporting-issues)
- [Submitting Changes](#submitting-changes)
- [Style Guide](#style-guide)
- [Commit Message Guidelines](#commit-message-guidelines)
- [Pull Request Process](#pull-request-process)

## Code of Conduct

By participating in this project, you agree to abide by our [Code of Conduct](CODE_OF_CONDUCT.md). Please read it to understand the standards we expect from all contributors.

## How to Contribute

There are several ways to contribute to **gochron**:

1. **Report Bugs**: If you encounter any issues, please [open an issue](https://github.com/mateusveloso/gochron/issues) with details about the problem.
2. **Suggest Features**: Have an idea for a new feature? We'd love to hear it! [Open a feature request](https://github.com/mateusveloso/gochron/issues) to discuss your proposal.
3. **Submit Code**: We welcome code contributions, whether they are bug fixes, improvements, or new features.

## Setting Up the Development Environment

To get started with contributing:

1. Fork the repository to your own GitHub account.
2. Clone the forked repository:
  ```bash
    git clone https://github.com/mateusveloso/gochron.git
  ```
3. Navigate to the project directory:
  ```bash
  cd gochron
  ```
4. Create a new branch for your feature or bug fix:
  ```bash
git checkout -b feature/your-feature-name
  ```

## Reporting Issues
When reporting issues, please include:

- A clear and descriptive title.
- Steps to reproduce the issue.
- The expected and actual behavior.
- Any relevant logs, screenshots, or additional information.

## Submitting Changes

1. Make Your Changes
  - Follow the [Style Guide](#style-guide) to keep the code consistent.
  - Write tests for new features or bug fixes. We aim for high code quality, and testing is a crucial part of this process.
2. Run Tests
  - Before submitting a pull request, ensure that all tests pass:
    ```bash
      go test ./...
    ```
3. Commit Changes
   - Follow the [Commit Message Guidelines](#commit-message-guidelines) when making commits:
  ```bash
  git commit -m "feat: add feature X to enhance functionality Y"
  ```
4. Push to Your Branch
  ```bash
  git push origin feature/your-feature-name
  ```
5. Open a Pull Request
  - Once your changes are pushed, open a pull request (PR) against the appropriate version branch (v0, v1, etc.). In the PR description, provide:
    - A summary of the changes.
    - Any related issue numbers.
    - Details on testing performed.
   
## Style Guide
To keep the codebase consistent and readable:
- Follow Go's standard conventions (e.g., gofmt and golint).
- Organize imports using the standard library first, followed by external packages, and then internal packages.
- Write clear and descriptive comments for functions, methods, and packages.
- Run gofmt to format your code automatically:
```bash
gofmt -w .
```

## Commit Message Guidelines
We follow [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) for commit messages. 
This helps us generate changelogs and version numbers automatically. 
Examples:
- ```bashfeat: add support for new time zone parsing```
- ```bashfix: correct duration calculation logic```
- ```bashdocs: update README with usage examples```
- ```bashtest: add tests for date parsing functionality```

## Pull Request Process
- Ensure your PR targets the appropriate version branch (v0, v1, etc.).
- Your PR should have a clear title and description summarizing the changes.
- Ensure that all tests pass and that your code meets the style guidelines.
- Wait for a review from a maintainer. We may ask for changes or further tests.
- Once approved, your PR will be merged!

### Thank you for your contribution! We are excited to work with you on building a better gochron library!

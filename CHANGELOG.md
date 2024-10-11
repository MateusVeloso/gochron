# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v0.0.1] - 2024-10-10
### Added
- Initial implementation of the `gochron` library for Go.
- `Date` structure with functions for creation (`NewDate`), adding days, subtracting days, comparing dates, and formatting.
- `Time` structure with functions for creation (`NewTime`), adding seconds, subtracting seconds, comparing times, and formatting.
- `DateTime` structure with functions for creation (`NewDateTime`), adding durations, subtracting durations, comparing datetimes, and formatting.
- Unit tests for `Date`, `Time`, and `DateTime` covering validation, addition/subtraction, comparison, and formatting.
- Benchmark tests for `Date`, `Time`, and `DateTime` structures to ensure high performance under various operations.
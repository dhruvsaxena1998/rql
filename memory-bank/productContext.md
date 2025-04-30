# Product Context

## Purpose and Value
- Purpose: Offer a concise, human-readable language for defining conditional logic and seamlessly convert it into standard JSONLogic, eliminating manual translation and reducing errors.
- Value to Users: Speeds up rule definition, improves maintainability of complex conditions, and enables non-experts to author business rules without deep JSONLogic knowledge.

## Target Audience
- Developers: Building validation layers, rule engines, or configuration-driven applications where dynamic conditional logic is required.
- Low‑Code/No‑Code Platforms: Teams that need an intermediate representation for user-defined rules.
- Compliance & QA Engineers: Who write and validate business logic for data access, security policies, or audit checks.

> Needs & Pain Points
- Avoiding verbose and error-prone JSONLogic syntax.
- Quick iteration on rule definitions without deep coding.
- Clear, actionable error messages when rules are invalid.

## Problem Statement
- Current Challenge: Writing JSONLogic by hand is technical, verbose, and error-prone—especially for those unfamiliar with its structure.
- Illustrative Example: A user manually crafting nested and/or blocks in JSON often misplaces brackets or keys, leading to silent failures. Nested conditions become hard to read and maintain.
- Solution: A DSL with intuitive syntax (@age > 18 AND @age <= 22) that abstracts the complexity of JSONLogic, allowing users to write rules in a more intuitive format. The tool will parse this DSL and output valid JSONLogic, ensuring correctness and clarity.

## User Experience Goals
- Intuitive Syntax: DSL should read like plain English, minimizing the learning curve.
- Clear Feedback: Syntax errors pinpoint exact location and offer corrective hints.
- Fast Feedback Loop: Near-instant parsing and translation in CLI and library.
- Seamless Integration: Simple API and CLI flags for embedding in CI pipelines or build scripts.

## Success Metrics
- Adoption: GitHub stars, downloads of CLI binaries, and go module imports.
- Reliability: ≥95% of valid DSL snippets translate without errors; <5% reported edge-case failures.
- Performance: Average parse-and-translate time <10ms per expression for typical rule sets.
- User Satisfaction: Positive feedback from early adopters, tracked via surveys or issue sentiments.

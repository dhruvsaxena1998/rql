# Project Brief

## Core Requirements and Goals
- Parse DSL to JSONLogic: Interpret a human-friendly syntax (e.g., @age > 18 AND @age <= 22) and translate into a valid JSONLogic structure.
- Support Fundamental Operators: 
    - Logic and Boolean Operations ( if, ==, ===, !=, !==, !, !!, OR, AND, NOT ) 
    - Numeric Operations ( >, <, >=, <=, BETWEEN, MAX, MIN, AVG, +, -, *, /, % )
    - Array Operations ( IN, CONTAINS, MAP, FILTER, ALL, NONE, SOME )
    - String Operations ( IN )
    - Accessor ( @ ) -> @ will be use as variable
- DSL will support `.rql` file extension.

## Project Scope
> MVP

- Parsing basic comparison and logical expressions.
- Abstract Syntax Tree (AST) construction in Go.
- Visitor or recursive translator converting AST nodes into JSONLogic maps.
- Command‑line interface with flags (--pretty, --infile, --outfile).
- Unit and integration tests validating both parser and JSONLogic output.
- CI pipeline (GitHub Actions) for tests and linting.

- Logic and Boolean Operations
- Numeric Operations
- Accessor ( @ ) -> @ will be use as variable

> Phase 2

- String Operations
- Array Operations
- Custom or user‑defined functions/plugins.
- Performance optimizations beyond typical Go performance

> Phase 3

- Advanced DSL features: custom functions or custom operations.
- Support for more complex JSONLogic constructs.
- Evaluation of DSL and Value via HTTP API.
- Implementation of Rule Engine (CRUD) for managing rules.

## Deliverables
- Source Code Repository: Go modules, clear folder structure, documentation.
- README & Documentation: Installation, usage examples, API reference.
- CLI Tool: Executable that reads DSL and emits JSONLogic.
- Test Suites: Unit tests for parser/translator, integration tests with sample files.
- CI Configuration: Automated testing and linting on each push on main.

> Success Criteria:

- All tests pass with ≥90% coverage.
- DSL examples in README translate correctly to expected JSONLogic.
- Error messages are clear and pinpoint syntax issues.
- CI pipeline green on default branch.

## Stakeholders
@dhruvsaxena119: Project Owner & Lead Developer – responsible for design, implementation, and releases.

## Risks and Mitigation Strategies
- Ambiguous Grammar: Risk of parsing conflicts.
    - Mitigation: Draft clear EBNF/PEG spec, write exhaustive parser tests early.

- Library Limitations: Chosen parser library may not support needed constructs.
    - Mitigation: Prototype grammar integration in PoC, evaluate alternatives.

- Edge Cases in JSONLogic: Complex nested expressions may miscompile.
    - Mitigation: Integration tests against a JSONLogic runner, peer code reviews.

- Scope Creep: Adding non‑MVP features too early.
    - Mitigation: Strict backlog grooming, prioritize core items first.

## Assumptions
- Users will write only numeric and variable comparisons in MVP.
- Input variables (@x) map directly to JSONLogic var paths.
- Environment has Go ≥1.18 installed.
- Consumers of output accept standard JSONLogic format.

# Progress

## What Works
- List the features or components that are working as expected.
- Provide a brief description of each working feature or component.

## What's Left to Build
- **MVP Development**
  - [X] Create a new Git repo
  - [X] Setup memory bank
  - [ ] Add a README file
  - [ ] Set up linters (e.g. golangci‐lint) and formatters.
  - [ ] Define project layout
  - [ ] Define the grammar using Participle.
  - [ ] Build the AST model.
  - [ ] Implement the translator to convert AST to JSONLogic.
  - [ ] Develop the CLI interface for parsing DSL expressions.
  - [ ] Set up the CI pipeline for automated testing and linting.

- **Phase 2 Development**
  - [ ] Add support for string operations.
  - [ ] Add support for array operations.
  - [ ] Implement custom or user-defined functions/plugins.
  - [ ] Perform performance optimizations.

- **Phase 3 Development**
  - [ ] Implement advanced DSL features.
  - [ ] Support more complex JSONLogic constructs.
  - [ ] Evaluate DSL and Value via HTTP API.
  - [ ] Implement a Rule Engine (CRUD) for managing rules.
  
## Current Status
- The project is currently in the MVP development phase.
- The initial parser using Participle has been created.
- The AST model with basic node types is defined.
- The translator to convert AST to JSONLogic is being implemented.
- The CLI interface with basic commands is developed.
- The CI pipeline using GitHub Actions is configured.
## Known Issues
- **Parser Grammar Ambiguity**: The initial grammar definition may have ambiguities that need to be resolved.
- **Translator Inefficiency**: The current translator implementation may not be optimized for performance.
- **CLI Command Handling**: Some CLI commands may not handle edge cases correctly.
- **CI Pipeline Configuration**: The CI pipeline may need adjustments to handle new dependencies or configurations.
## Evolution of Project Decisions
- **Initial Project Setup**: Chose Go for the project due to its performance, static typing, and strong ecosystem. Decided to use Participle for parsing due to its declarative grammar and ease of embedding in Go code.
- **AST Model Introduction**: Introduced an explicit AST layer to simplify complex transformations and enable future optimizations.
- **Translator Implementation**: Used visitor-like recursion for translation to keep translator code modular and extensible.
- **CLI Framework Selection**: Chose Cobra and Viper for CLI framework and configuration management.
- **HTTP API Framework Selection**: Selected go-chi for the lightweight HTTP API framework.
- **JSONLogic Serialization**: Used JSONLogic-go for serializing Go maps into compliant JSONLogic output.
- **CI Pipeline Configuration**: Configured GitHub Actions for CI pipeline to ensure automated testing and linting.
- **Containerization**: Used Docker for containerization to ensure consistent environments.
- **Performance Targeting**: Targeted <10ms per expression parse-and-translate on typical inputs to ensure performance.
- **Deterministic Behavior**: Avoided reflection or runtime code generation to ensure reproducible builds.
- **Concurrency Strategy**: Kept the parser and translator single-threaded, applying concurrency only at the HTTP handler level.
- **Dependency Management**: Pinned dependency versions in go.mod to ensure consistent builds and avoid breaking changes.
- **Code Quality Assurance**: Ran golangci-lint run and go test ./... before commits to enforce code quality.
- **Code Formatting**: Used gofmt and goimports via pre-commit to enforce consistent formatting.
- **Environment-Specific Configuration**: Kept environment-specific config (e.g., port) in .env or CLI flags, not in code.
- **Error Reporting**: Provided clear user feedback and integration in automated pipelines.
- **Stateless Processing**: Ensured each parse/translate request is independent to support deployment in constrained environments.
- **Memory Footprint Management**: Kept in-memory data small to support deployment in constrained environments.

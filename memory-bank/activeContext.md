# Active Context

## Current Work Focus
- Current focus is on implementing the DSL parser and translator in Go.
- Tasks include:
  - Defining the grammar using Participle.
  - Building the AST model.
  - Implementing the translator to convert AST to JSONLogic.
  - Developing the CLI interface for parsing DSL expressions.
  - Setting up the CI pipeline for automated testing and linting.
## Recent Changes
- Initialized the project structure with core files.
- Set up the Go project.
## Next Steps
- Create the initial parser using Participle.
- Define the AST model with basic node types.
- Start implementing the translator to convert AST to JSONLogic.
- Develop the CLI interface with basic commands.
- Configured GitHub Actions for CI pipeline.
- Complete the implementation of the parser and translator.
- Add support for more complex DSL constructs (e.g., string and array operations).
- Implement error handling and validation for the parser.
- Develop the HTTP API for parsing DSL expressions.
- Write unit and integration tests for the parser and translator.
- Refine the CLI interface with additional flags and options.
- Document the project with README and usage examples.
- Deploy the HTTP API using Docker.
## Active Decisions and Considerations
- Chose Participle for parsing due to its declarative grammar and ease of embedding in Go code.
- Introduced an explicit AST layer to simplify complex transformations and enable future optimizations.
- Used visitor-like recursion for translation to keep translator code modular and extensible.
- Leveraged Participle’s built-in position tracking to provide line/col in syntax errors.
- Decided to use Go for the project due to its performance, static typing, and strong ecosystem.
- Chose Cobra and Viper for CLI framework and configuration management.
- Selected go-chi for the lightweight HTTP API framework.
- Used JSONLogic-go for serializing Go maps into compliant JSONLogic output.
- Configured GitHub Actions for CI pipeline to ensure automated testing and linting.
- Used Docker for containerization to ensure consistent environments.
## Important Patterns and Preferences
- **Visitor-like Recursion**: Separates operations (translation) from data structures (AST nodes).
- **Factory Methods**: Centralizes instantiation and default values for AST nodes.
- **Command Pattern**: Decouples CLI commands (parse, translate, version) from execution logic.
- **Singleton**: Single configuration object (flags) shared across CLI execution.
- **Pin Dependency Versions**: Ensures consistent builds and avoids breaking changes.
- **Run Linters and Tests**: Enforces code quality and reliability.
- **Use gofmt and goimports**: Ensures consistent code formatting.
- **Environment-Specific Config**: Keeps environment-specific config (e.g., port) in .env or CLI flags, not in code.
- **Error Reporting**: Provides clear user feedback and integration in automated pipelines.
- **Stateless Processing**: Ensures each parse/translate request is independent.
- **Memory Footprint**: Keeps in-memory data small for constrained environments.
- **Performance**: Targets <10ms per expression parse-and-translate on typical inputs.
- **Deterministic Behavior**: Avoids reflection or runtime code generation for reproducible builds.
- **Single-threaded Core**: Parser and translator run sequentially; concurrency may be applied at HTTP handler level only.
## Learnings and Project Insights
- **Grammar Design**: Clear and well-defined grammar is crucial for parsing accuracy.
- **Error Handling**: Robust error handling improves user experience and debugging.
- **Testing**: Comprehensive testing ensures reliability and prevents regressions.
- **Modular Design**: Separation of concerns (parser, AST, translator) simplifies maintenance and future enhancements.
- **Documentation**: Clear documentation is essential for onboarding new contributors and users.
- **CI/CD**: Automated testing and linting improve code quality and reduce manual effort.
- **Containerization**: Docker simplifies deployment and ensures consistent environments.
- **Tooling**: Leveraging existing tools (Participle, Cobra, Viper, go-chi) accelerates development.
- **Performance Optimization**: Early focus on performance ensures scalability.
- **User Feedback**: Continuous user feedback helps refine the product and meet user needs.

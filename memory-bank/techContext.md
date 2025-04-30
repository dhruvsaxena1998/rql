# Tech Context

## Technologies Used
- Go (≥1.24): Primary language for parser, AST modeling, translator, CLI, and HTTP API.
- [Participle](https://github.com/alecthomas/participle): Declarative parsing library to define and parse the DSL grammar.
- [Cobra](https://github.com/spf13/cobra) & [Viper](https://github.com/spf13/viper): CLI framework (Cobra) and configuration management (Viper) for command parsing and flags.
- [go-chi](https://go-chi.io/): Lightweight go-chi framework to expose the HTTP API.
- [JSONLogic-go](https://github.com/diegoholiveira/jsonlogic): Utility or custom code to serialize Go maps into compliant JSONLogic output.
- GitHub Actions: CI pipeline for testing, linting, and build validation.
- Docker: Containerization of the HTTP API for development and deployment.

## Development Setup
1. Prerequisites
    - Install Go (≥1.24) and Docker.
    - Install Git for version control.
2. Clone the repository:
    ```bash
    git clone git@github.com/dhruvsaxena119/rql.git
    cd rql
    go mod tidy
    ```
3. Environment Variables
    - Set any required environment variables for the project using `.env.example` as a template.

4. Local Development
    - Use Docker for local development to ensure consistent environments.
    - Run the HTTP API locally:
    ```bash
    docker-compose up
    ```

    - Access the API at `http://localhost:8080`.
    ```
    curl -X POST http://localhost:8080/translate \
        -H "Content-Type: application/json" \
        -d '{"dsl":"@age > 18 AND @age <= 22"}'
    ```

    - Access the CLI tool:
    ```bash
    go build -o bin/rql ./cmd/rql/main.go
    ./bin/rql translate --inline "@age > 18 AND @age <= 22" --pretty
    ./bin/rql translate input.rql --out output.json 
    ```

## Technical Constraints
- Stateless Processing: Each parse/translate request is independent; no session state.
- Memory Footprint: Keep in-memory data small to support deployment in constrained environments.
- Performance: Target <10ms per expression parse-and-translate on typical inputs.
- Deterministic Behavior: No use of reflection or runtime code generation to ensure reproducible builds.
- Single-threaded Core: Parser and translator run sequentially; concurrency may be applied at HTTP handler level only.

## Dependencies
- Go Modules (go.mod)
- Devtools
    - golangci-lint (https://github.com/golangci/golangci-lint)

## Tool Usage Patterns
- Cli Usage
    - `rql translate <file>`: Parses the DSL expression and outputs JSONLogic.
    ```
        # input from file, output to file
        rql translate <input.rql> --out <output.json> [--pretty]

        # output to file
        rql translate --inline "<expression>" --out <output.json> [--pretty]

        # output to stdout
        rql translate --inline "<expression>" [--pretty]
    ```
- HTTP Usage
    - `POST /translate`: Accepts a JSON body with the DSL expression and returns the JSONLogic output.
    ```jsonc
    <!-- Request: -->
    {
        "dsl": "@age > 18 AND @age <= 22"
    }
    ```
    ```jsonc
    <!-- Response: -->
    {
        "success": true,
        "json_logic": {
             "and": [
                { ">": [ { "var": "age" }, 18 ] },
                { "<=": [ { "var": "age" }, 22 ] }
            ]
        }
    }
    ```
- Best Practices:
    - Pin dependency versions in go.mod.
    - Run golangci-lint run and go test ./... before commits.
    - Use gofmt and goimports via pre-commit to enforce formatting.
    - Keep environment-specific config (e.g., port) in .env or CLI flags, not in code.


# RQL: Rule Query Language
## Product Requirements Document (PRD)

## 1. Introduction

### 1.1 Document Purpose
This Product Requirements Document (PRD) defines the requirements for Rule Query Language (RQL), a domain-specific language for defining conditional logic that translates to JSONLogic. This document will serve as the primary reference for the development team and stakeholders.

### 1.2 Product Overview
RQL is a human-readable language that allows users to define conditional logic in a natural syntax, which is then translated into JSONLogic-compatible JSON. This product includes both a language definition, a parser library in Go, a command-line interface, and an HTTP API for integration with other systems.

### 1.3 Scope
This document covers the requirements for the minimum viable product (MVP) of RQL, with references to future phases for context. The MVP focuses on core language features, parsing, translation to JSONLogic, and essential developer tools.

## 2. Vision and Goals

### 2.1 Product Vision
To create the standard tool for defining business rules and conditional logic in an intuitive, human-readable syntax that seamlessly translates to industry-standard JSONLogic.

### 2.2 Business Goals
1. Reduce the time and complexity required to author business rules by 70%
2. Eliminate errors from manual JSONLogic creation
3. Enable non-technical stakeholders to participate in rule creation and validation
4. Establish RQL as the go-to intermediate representation for rule engines

### 2.3 User Goals
1. Write conditional logic in a natural, readable syntax
2. Avoid verbose JSON structures and complex nesting
3. Receive clear error messages when rules are invalid
4. Integrate rule validation into existing workflows and pipelines

### 2.4 Success Metrics
1. Adoption: GitHub stars, downloads of CLI binaries, and Go module imports
2. Reliability: ≥95% of valid DSL snippets translate without errors; <5% reported edge-case failures
3. Performance: Average parse-and-translate time <10ms per expression for typical rule sets
4. User Satisfaction: Positive feedback from early adopters, tracked via surveys or issue sentiments

## 3. User Personas

### 3.1 Primary Persona: Application Developer
**Name:** Alex
**Role:** Software Developer
**Technical Level:** High (proficient in Go, JavaScript, etc.)
**Goals:** 
- Integrate dynamic conditional logic into applications
- Avoid writing complex nested JSON structures by hand
- Maintain a clean, readable codebase

**Pain Points:**
- Writing JSONLogic directly is verbose and error-prone
- Complex boolean logic becomes unreadable in pure JSON
- Current tools require significant manual work or code generation

### 3.2 Secondary Persona: Low-Code Platform Developer
**Name:** Jordan
**Role:** Product Developer at a Low-Code/No-Code Platform
**Technical Level:** Medium to High
**Goals:**
- Create an intermediate representation for user-defined rules
- Validate rule logic before execution
- Translate business requirements into technical implementations

**Pain Points:**
- Current rule expression formats are either too technical or too limited
- Need a balance between power and understandability

### 3.3 Tertiary Persona: Compliance Engineer
**Name:** Morgan
**Role:** Compliance & QA Engineer
**Technical Level:** Medium
**Goals:**
- Author and validate business logic for data access, security policies, or audit checks
- Understand and maintain complex rule systems
- Communicate rule logic to non-technical stakeholders

**Pain Points:**
- Current rule systems are too technical to explain to stakeholders
- Need to frequently translate between natural language and code

## 4. Requirements

### 4.1 Language Specification

#### 4.1.1 Syntax Overview
RQL must support a human-readable syntax for defining conditional logic that resembles natural language expressions. The basic structure should be familiar to anyone with programming experience while still being approachable to non-developers.

#### 4.1.2 Operators
The language must support the following operators for MVP:

**Logical Operators:**
- `AND`: Logical AND
- `OR`: Logical OR
- `NOT`: Logical negation

**Comparison Operators:**
- `==`: Equality
- `===`: Strict equality
- `!=`: Inequality
- `!==`: Strict inequality
- `>`: Greater than
- `<`: Less than
- `>=`: Greater than or equal to
- `<=`: Less than or equal to

**Numeric Operations:**
- `+`: Addition
- `-`: Subtraction
- `*`: Multiplication
- `/`: Division
- `%`: Modulo
- `BETWEEN`: Check if a value is between two other values
- `MAX`: Return the maximum of multiple values
- `MIN`: Return the minimum of multiple values
- `AVG`: Return the average of multiple values

**Accessor:**
- `@`: Denotes a variable (e.g., `@age` refers to the "age" variable)

#### 4.1.3 Data Types
RQL must support the following data types in the MVP:
- Numbers (integers and floats)
- Booleans (`true`, `false`)
- Variables (prefixed with `@`)
- Strings (enclosed in double or single quotes)

#### 4.1.4 Expression Structure
Expressions in RQL should follow these structural rules:
- Binary operations (e.g., `@age > 18 AND @name == "John"`)
- Unary operations (e.g., `@role NOT IN ["admin", "guest"] AND @isActive === true`)
- Parentheses for grouping (e.g., `(@age > 18 AND @age < 65) OR @isVIP === true`)
- Function calls (e.g., `MAX(@score1, @score2, @score3)`)

#### 4.1.5 Operator Precedence
RQL must define and respect operator precedence, following standard programming language conventions:
1. Parentheses
2. Function calls
3. Unary operators (NOT)
4. Multiplicative operators (*, /, %)
5. Additive operators (+, -)
6. Comparison operators (>, <, >=, <=, ==, ===, !=, !==, BETWEEN)
7. Logical AND
8. Logical OR

### 4.2 Parser Requirements

#### 4.2.1 Parsing Capabilities
The parser must:
- Parse RQL expressions into an Abstract Syntax Tree (AST)
- Handle all operators and data types specified in the language
- Provide detailed error messages with line and column information
- Be deterministic and produce consistent results for the same input

#### 4.2.2 Error Handling
The parser must provide clear error messages that:
- Indicate the exact location (line and column) of syntax errors
- Describe the nature of the error in plain language
- Suggest possible corrections when applicable

#### 4.2.3 Performance
The parser must:
- Parse typical expressions (<100 tokens) in under 10ms
- Handle complex expressions (up to 1000 tokens) without excessive memory usage
- Be thread-safe for concurrent usage in the API context

### 4.3 Translator Requirements

#### 4.3.1 JSONLogic Output
The translator must:
- Convert the AST into valid JSONLogic structures
- Support all operators and data types in the language specification
- Preserve the original semantics of the RQL expression
- Generate optimized JSONLogic when possible (avoiding unnecessary nesting)

#### 4.3.2 Variable Handling
The translator must:
- Convert `@variable` references to JSONLogic `{"var": "variable"}` format
- Support nested variables (e.g., `@user.name`) with proper JSONLogic paths

#### 4.3.3 Function Handling
The translator must:
- Map RQL functions to their JSONLogic equivalents
- Preserve function arguments in the correct order
- Support all functions specified in the language specification

### 4.4 Command-Line Interface Requirements

#### 4.4.1 Core Functionality
The CLI must:
- Parse RQL expressions from strings or files
- Output JSONLogic to stdout or files
- Support pretty-printing for human readability
- Return appropriate exit codes (0 for success, non-zero for errors)

#### 4.4.2 Usage Patterns
The CLI must support these usage patterns:
```
# Input from file, output to file
rql translate <input.rql> --out <output.json> [--pretty]

# Input from string, output to file
rql translate --inline "<expression>" --out <output.json> [--pretty]

# Input from string, output to stdout
rql translate --inline "<expression>" [--pretty]
```

#### 4.4.3 Error Reporting
The CLI must:
- Display error messages to stderr
- Include line and column information for syntax errors
- Provide context around the error when possible
- Return non-zero exit codes for errors

### 4.5 HTTP API Requirements

#### 4.5.1 Core Functionality
The API must:
- Accept POST requests with RQL expressions
- Return JSONLogic in the response
- Support JSON input and output formats
- Provide appropriate status codes and error messages

#### 4.5.2 Endpoints
The API must implement:
- `POST /translate`: Translate RQL to JSONLogic
  - Request: `{"dsl": "@age > 18 AND @age <= 22"}`
  - Response: `{"success": true, "json_logic": {"and": [{">":[{"var":"age"},18]},{"<=":[{"var":"age"},22]}]}}`

#### 4.5.3 Error Handling
The API must:
- Return 400 status codes for invalid expressions
- Include detailed error information in the response
- Handle and report internal errors appropriately

### 4.6 Non-Functional Requirements

#### 4.6.1 Performance
- Parse and translate typical expressions in under 10ms
- Handle complex expressions (up to 1000 tokens) in under 100ms
- Support concurrent requests in the API (minimum 100 req/sec)

#### 4.6.2 Reliability
- Pass all unit and integration tests
- Achieve >90% code coverage
- Handle edge cases gracefully
- Provide consistent results for the same input

#### 4.6.3 Maintainability
- Follow Go best practices and conventions
- Implement clear separation of concerns
- Document code and APIs
- Provide examples and usage guides

#### 4.6.4 Security
- Validate all input to prevent injection attacks
- Implement appropriate resource limits to prevent DoS attacks
- Secure API endpoints with appropriate authentication (Phase 2)

#### 4.6.5 Scalability
- Support containerized deployment
- Allow horizontal scaling of the API (Phase 2)
- Implement caching for frequent expressions (Phase 2)

## 5. Functional Specifications

### 5.1 Language Grammar
The RQL grammar will be implemented using the Participle library. A simplified EBNF-like grammar is:

```
Expression    = LogicalExpr
LogicalExpr   = ComparisonExpr (("AND" | "OR") ComparisonExpr)*
ComparisonExpr = SimpleExpr (CompOp SimpleExpr)?
CompOp        = "==" | "===" | "!=" | "!==" | ">" | "<" | ">=" | "<="
SimpleExpr    = Term (("+"|"-") Term)*
Term          = Factor (("*"|"/"|"%") Factor)*
Factor        = Number | String | Variable | "(" Expression ")" | FunctionCall | "NOT" Factor
Variable      = "@" Identifier
FunctionCall  = Identifier "(" Arguments ")"
Arguments     = Expression ("," Expression)*
Identifier    = [a-zA-Z_][a-zA-Z0-9_]*
Number        = [0-9]+ ("." [0-9]+)?
String        = "\"" [^"]* "\"" | "'" [^']* "'"
```

### 5.2 AST Structure
The AST will consist of the following node types:
- Expression (base interface)
- BinaryOp (left, operator, right)
- UnaryOp (operator, expression)
- Literal (value)
- Variable (name)
- FunctionCall (name, arguments)

### 5.3 JSONLogic Translation
The translator will map RQL constructs to JSONLogic as follows:

#### 5.3.1 Logical Operators
- `A AND B` → `{"and": [A_json, B_json]}`
- `A OR B` → `{"or": [A_json, B_json]}`
- `NOT A` → `{"!": [A_json]}`

#### 5.3.2 Comparison Operators
- `A == B` → `{"==": [A_json, B_json]}`
- `A === B` → `{"===": [A_json, B_json]}`
- `A != B` → `{"!=": [A_json, B_json]}`
- `A !== B` → `{"!==": [A_json, B_json]}`
- `A > B` → `{">": [A_json, B_json]}`
- `A < B` → `{"<": [A_json, B_json]}`
- `A >= B` → `{">=": [A_json, B_json]}`
- `A <= B` → `{"<=": [A_json, B_json]}`

#### 5.3.3 Numeric Operations
- `A + B` → `{"+": [A_json, B_json]}`
- `A - B` → `{"-": [A_json, B_json]}`
- `A * B` → `{"*": [A_json, B_json]}`
- `A / B` → `{"/": [A_json, B_json]}`
- `A % B` → `{"%": [A_json, B_json]}`
- `BETWEEN(A, min, max)` → `{"and": [{">=":[A_json,min_json]},{"<=":[A_json,max_json]}]}`
- `MAX(A, B, ...)` → `{"max": [A_json, B_json, ...]}`
- `MIN(A, B, ...)` → `{"min": [A_json, B_json, ...]}`
- `AVG(A, B, ...)` → `{"reduce": [{"+":[{"var":"current"},{"var":"accumulator"}]}, [A_json, B_json, ...], 0]}`

#### 5.3.4 Variables and Literals
- `@variable` → `{"var": "variable"}`
- `@user.name` → `{"var": "user.name"}`
- `42` → `42`
- `"text"` → `"text"`
- `true` → `true`

### 5.4 CLI Interface
The CLI will implement the following commands:

#### 5.4.1 Translate Command
```
rql translate [flags] [file]
```

Flags:
- `--inline, -i`: Provide an inline expression instead of a file
- `--out, -o`: Output file path (defaults to stdout)
- `--pretty, -p`: Pretty-print the JSON output

#### 5.4.2 Version Command
```
rql version
```
Outputs the current version of the tool.

#### 5.4.3 Help Command
```
rql help [command]
```
Provides help information for the specified command.

### 5.5 HTTP API Interface
The API will implement the following endpoints:

#### 5.5.1 Translate Endpoint
```
POST /translate
```

Request:
```json
{
  "dsl": "@age > 18 AND @age <= 22"
}
```

Success Response:
```json
{
  "success": true,
  "json_logic": {
    "and": [
      { ">": [{"var": "age"}, 18] },
      { "<=": [{"var": "age"}, 22] }
    ]
  }
}
```

Error Response:
```json
{
  "success": false,
  "error": {
    "message": "Syntax error: unexpected token at line 1, column 10",
    "line": 1,
    "column": 10
  }
}
```

## 6. Technical Architecture

### 6.1 System Components
1. **DSL Parser**: Parses RQL expressions into an AST
2. **AST Model**: Defines the structure of the AST
3. **Translator**: Converts the AST to JSONLogic
4. **CLI Interface**: Provides command-line access
5. **HTTP API**: Provides web API access

### 6.2 Component Relationships
```
CLI Interface → DSL Parser → AST Model → Translator → JSONLogic Output
HTTP API → DSL Parser → AST Model → Translator → JSONLogic Output
```

### 6.3 Technology Stack
- **Language**: Go (≥1.24)
- **Parsing Library**: Participle
- **CLI Framework**: Cobra & Viper
- **API Framework**: go-chi
- **JSONLogic**: Custom implementation or diegoholiveira/jsonlogic
- **CI/CD**: GitHub Actions

### 6.4 Design Patterns
1. **Visitor Pattern**: For AST traversal and translation
2. **Factory Methods**: For creating AST nodes
3. **Command Pattern**: For CLI commands
4. **Singleton**: For configuration

## 7. Implementation Plan

### 7.1 MVP (Phase 1)
- Basic comparison and logical expressions
- AST construction
- JSONLogic translation
- CLI with basic functionality
- Unit and integration tests
- CI pipeline

### 7.2 Phase 2
- String operations
- Array operations
- Custom function support
- Performance optimizations

### 7.3 Phase 3
- Advanced DSL features
- Complex JSONLogic constructs
- HTTP API enhancements
- Rule engine implementation

## 8. Testing and Quality Assurance

### 8.1 Testing Approach
1. **Unit Tests**: For individual components
2. **Integration Tests**: For end-to-end functionality
3. **Performance Tests**: For ensuring performance targets
4. **Usability Tests**: For validating ease of use

### 8.2 Test Cases
1. **Parser Tests**: Validate parsing of various expressions
2. **Translator Tests**: Ensure correct JSONLogic output
3. **CLI Tests**: Verify command behavior
4. **API Tests**: Check endpoint functionality
5. **Error Tests**: Validate error handling

### 8.3 Acceptance Criteria
1. All tests pass with ≥90% coverage
2. DSL examples translate correctly to expected JSONLogic
3. Error messages are clear and pinpoint syntax issues
4. CI pipeline green on default branch
5. Performance targets met

## 9. Release Plan

### 9.1 MVP Release
1. **Alpha**: Internal testing only
2. **Beta**: Limited external testing
3. **Release Candidate**: Feature complete, final testing
4. **General Availability**: Public release

### 9.2 Release Deliverables
1. **Source Code**: GitHub repository
2. **Documentation**: README, API docs, examples
3. **CLI Binary**: For major platforms
4. **Docker Image**: For API deployment

### 9.3 Release Process
1. Code freeze and final testing
2. Build release artifacts
3. Update documentation
4. Tag release in GitHub
5. Publish binaries and images
6. Announce release

## 10. Appendices

### 10.1 Glossary
- **RQL**: Rule Query Language, the DSL defined in this document
- **DSL**: Domain-Specific Language
- **JSONLogic**: A JSON-based format for defining conditional logic
- **AST**: Abstract Syntax Tree, a tree representation of code
- **CLI**: Command-Line Interface
- **API**: Application Programming Interface

### 10.2 References
- JSONLogic specification: https://jsonlogic.com/
- Participle library: https://github.com/alecthomas/participle
- Cobra CLI framework: https://github.com/spf13/cobra
- Go-chi framework: https://go-chi.io/

### 10.3 Examples
#### Example 1: Age Between 18 and 22
```
@age > 18 AND @age <= 22
```
Translates to:
```json
{
  "and": [
    { ">": [{"var": "age"}, 18] },
    { "<=": [{"var": "age"}, 22] }
  ]
}
```

#### Example 2: Complex Condition with Parentheses
```
(@age >= 18 OR @parent_consent == true) AND @country IN ["US", "CA", "MX"]
```
Translates to:
```json
{
  "and": [
    {
      "or": [
        { ">=": [{"var": "age"}, 18] },
        { "==": [{"var": "parent_consent"}, true] }
      ]
    },
    { "in": [{"var": "country"}, ["US", "CA", "MX"]] }
  ]
}
```

#### Example 3: Numeric Operations
```
(@score1 + @score2) / 2 > 75
```
Translates to:
```json
{
  ">": [
    { "/": [
      { "+": [{"var": "score1"}, {"var": "score2"}] },
      2
    ]},
    75
  ]
}
```
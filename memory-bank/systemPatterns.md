# System Patterns

## Architecture Overview
- DSL Parser (dsl/parser.go): Reads DSL input and builds an AST.
- AST Model (ast/*.go): Defines node types (BinaryOp, UnaryOp, Var, Number).
- Translator (translate/jsonlogic.go): Walks AST and emits a Go map[string]interface{} representing JSONLogic.
- CLI Interface (cmd/jsonlogic-dsl/main.go): Accepts input (stdin/file), flags, and outputs JSON to stdout or file.
- API Interface (api/jsonlogic): Provides an HTTP API for parsing string to JSONLogic
- CI Pipeline (GitHub Actions): Runs tests, linting, and formatting checks on each push.

Components interact sequentially: CLI → Parser → AST → Translator → JSON Logic Output.
Components interact sequentially: API → Parser → AST → Translator → JSON Logic Output.

## Key Technical Decisions
- Parsing Library: Chose [Participle](https://github.com/alecthomas/participle) for its declarative grammar and ease of embedding in Go code; hand-rolled parsing considered but deferred for maintainability.
- AST vs. Direct Translation: Introduced an explicit AST layer to simplify complex transformations and enable future optimizations/testing hooks.
- Visitor Pattern: Used visitor-like recursion over AST for translation to keep translator code modular and extensible.
- Error Reporting: Leveraged Participle’s built-in position tracking to provide line/col in syntax errors.

## Design Patterns
- Visitor-like Recursion: Separates operations (translation) from data structures (AST nodes).
- Factory Methods: For creating AST nodes, centralizing instantiation and default values.
- Command Pattern: CLI commands (parse, translate, version) decoupled from execution logic.
- Singleton: Single configuration object (flags) shared across CLI execution.

## Component Relationships
```
graph LR
    CLI[CLI Interface] --> Parser[Parser]
    HTTP_API[HTTP API] --> Parser
    Parser --> AST_Model[AST Model]
    AST_Model --> Translator
    Translator --> Output[JSONLogic Output]

    CI[CI Pipeline] -.-> QA[Quality Assurance]
```

### Critical Implementation Paths
- **Parsing Path**: `Input DSL` → `Parse()` → `AST` → `ToJSONLogic()` → `Serialized JSON`.  
  *Importance*: Core functionality; must handle all valid DSL constructs reliably.

- **Error Path**: Invalid DSL → Parser error with line/column → Exit with non-zero status in CLI.  
  *Importance*: Provides clear user feedback and integration in automated pipelines.
  
- **Testing Path**: On each push → Run `go test ./...` → Validate parser, translator, and CLI behavior.  
  *Importance*: Prevent regressions and maintain project reliability.

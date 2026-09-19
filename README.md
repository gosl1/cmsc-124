# Amadeus

## Creators

- Ethan Sean T. Gapulan (Ethanerz)
- Marc Raven A. Sian (gosl1)

## Overview

Amadeus is a dynamically typed, general-purpose scripting language inspired by the visual novel and anime franchise Steins;Gate, built specifically to have execution-history awareness. Amadeus natively treats the program's running execution path as an immutable, queryable "worldline." Rather than relying strictly on current variable states, developers can native-query the runtime history using custom syntax inspired by Steins;Gate to dynamically reroute logic based on what statements were executed and in what exact order.

## Host language and build

- Host language: go 1.26.7
- Version metadata: go.mod
- Build: `./build.sh`

## Running it


| Command | What it does |
|---|---|
| `./run <file>` | [Executes a program. Available from Lab 4.] |
| `./run --tokenize <file>` | [Prints the token stream.] |
| `./run --parse <file>` | [Prints the parsed tree.] |
| `./run --eval <file>` | [Evaluates each expression and prints its value.] |
| `./run` | [Starts the REPL.] |


Exit codes: 0 on sucessful run, 65 on a static error, 70 on a runtime error.

## File extension

`[.ama]` 

## Lexical structure


### Keywords

| Keyword | Purpose |
|---|---|
| [TBD] | variable declaration |
| [TBD] | print/output statement |
| [TBD] | if |
| [TBD] | else |
| [TBD] | while |
| [TBD] | boolean literal — true |
| [TBD] | boolean literal — false |
| [TBD] | nil / absence-of-value literal |
| [TBD] | function declaration |
| [TBD] | return statement |
| [TBD] | expose the execution-history record as a value (returns the log) |

### Operators


| Operator | Category | Operands | Associativity | Precedence |
|---|---|---|---|---|
| `+` | arithmetic | binary | left | [TBD — set in Lab 2] |
| `-` | arithmetic | binary | left | [TBD] |
| `*` | arithmetic | binary | left | [TBD] |
| `/` | arithmetic | binary | left | [TBD] |
| `%` | arithmetic (modulo) | binary | left | [TBD] |
| `=` | assignment | binary | right | [TBD] |
| `==` | comparison | binary | left | [TBD] |
| `!=` | comparison | binary | left | [TBD] |
| `<` | comparison | binary | left | [TBD] |
| `<=` | comparison | binary | left | [TBD] |
| `>` | comparison | binary | left | [TBD] |
| `>=` | comparison | binary | left | [TBD] |
| `!` | logical (not) | unary | right | [TBD] |
| `[` `]` | indexing (array/log access) | binary | left | [TBD] |
| `.` | field access | binary | left | [TBD] |

### Literals


| Kind | Syntax | Produces |
|---|---|---|
| number | `4`, `4.0` (integers and decimals; no leading-dot numbers, e.g. `.5` is invalid; a trailing dot with no following digit is its own token, so `3.toString` scans as NUMBER, DOT, IDENTIFIER) | numeric value |
| string | `"hello"`, double-quoted, single-line only (no multi-line strings); supports `\n`, `\"`, `\\` escape sequences | string value (decoded — lexeme keeps the raw escaped text, literal holds the decoded characters) |
| boolean | [TBD keyword] / [TBD keyword] | boolean value |
| nil | [TBD keyword] | absence-of-value |


### Identifiers

- Start characters: [which]
- Continue characters: [which]
- Case-sensitive: [yes or no]
- [Reserved patterns, length limits, or other restrictions.]

### Comments

- Line comments: `//`, discarded and not counted, runs to end of line
- Block comments: not supported (documented decision — revisit only if the
  language later needs to annotate nested structure worth commenting on)
- Nesting: n/a
- Harness note: `comment_prefix` in `tests/lab*/manifest.json` is set to `//`

## Whitespace and termination

- Whitespace significant: No
- Statement terminator: newline
- Block delimiters: Braces
- Grouping delimiters: Parentheses

## Token output format

```
[one line of real --tokenize output]
```

[What each field means. Frozen as of Lab 1; changes are recorded in the
changelog.]



## Grammar

```
[Your complete context-free grammar, current as of the latest activity.
Unambiguous, with precedence and associativity encoded in rule structure.]
```

## Parse output format

```
[one line of real --parse output, e.g. (+ 1.0 (* 2.0 3.0))]
```

- Groupings print as: [form]
- Numbers print as: [form]

## Semantics

### Values and types

[What runtime values exist, and how they are represented in the host
language.]

### Value printing

- Numbers: [e.g. 5 rather than 5.0]
- Nil: [spelling]
- Strings: [with or without quotes]

### Truthiness

[The complete rule. Which values are false in a condition; everything else is
true.]

### Operator semantics

- Arithmetic: [accepted operand types]
- `+` on strings: [concatenation, error, or coercion]
- Mixed types: [what happens]
- Comparison: [accepted operand types]
- Equality across types: [false, or an error]
- Division by zero: [value produced, or runtime error]

### Scope and bindings

- Redeclaration in the same scope: [allowed or an error]
- Uninitialized variable holds: [value]
- Shadowing: [behavior]
- Undefined name: [static error with exit 65, or runtime error with exit 70]

### Control flow and functions

- Logical operators return: [booleans, or the operand]
- Dangling else binds to: [which if]
- Closure capture of a loop variable: [per iteration, or shared]
- Function with no return statement produces: [value]
- Arity mismatch: [message and exit code]

## Native functions


| Name | Arguments | Returns | Notes |
|---|---|---|---|
| [name] | [count and types] | [type] | [caveats] |


## Errors and diagnostics

Message format:

```
[one real static error]
[one real runtime error]
```


| Failure | Exit code |
|---|---|
| [lexical error] | 65 |
| [syntax error] | 65 |
| [runtime error] | 70 |


## Testing conventions


| Folder | Activity | Mode | Flag |
|---|---|---|---|
| tests/lab1 | Scanner | sidecar | `--tokenize` |
| tests/lab2 | Parser | sidecar | `--parse` |
| tests/lab3 | Evaluator | inline | `--eval` |
| tests/lab4 | Context | inline | none |
| tests/lab5 | Functions | inline | none |


```
[specific tests]...
```

Run locally with:

```bash
curl -sSL https://raw.githubusercontent.com/WhiteLicorice/cmsc-124-harness/v1.1/run_tests.py -o run_tests.py
./build.sh
python3 run_tests.py tests/lab1
```

## Sample code

```
[a short program]
```

Output:

```
[its output]
```

## Design rationale

[Why the language is the way it is. Cover the choices that surprised you, the
features you cut, and the decisions you reversed. Specific reasons, not
approval of your own work.]

## Known limitations

- [What doesn't work, what is unimplemented, where behavior is worse than you
  would like.]

## Changelog


| Activity | What changed in the language |
|---|---|
| Lab 1 | [entry] |

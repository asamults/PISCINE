# PISCINE (Go)

## Purpose
PISCINE is an educational repository for an intensive Go programming course.  
It develops fundamental programming skills, algorithmic reasoning, and disciplined use of Git and command-line tools.

## Technologies
- Language: Go
- Runtime: CLI
- Version control: Git
- Dependency management: Go Modules

## Repository Structure

PISCINE/
├── go.mod
├── go.sum
├── hello.sh
├── printalphabet/
├── sudoku/
├── piscine-go-s2wft/
└── README.md


- `printalphabet/` — basic exercises (loops, output, control flow)  
- `sudoku/` — algorithmic exercise (problem solving, backtracking)  
- `piscine-go-s2wft/` — module or collection of exercises  
- `hello.sh` — shell exercise  
- `go.mod`, `go.sum` — Go module and dependencies  

## Rules
1. Each exercise must be independent.  
2. Only Go standard library is allowed unless stated.  
3. Do not modify provided function signatures.  
4. Code must be deterministic, reproducible, and readable.  
5. Output must match assignment specifications exactly.  

## Build and Run
```bash
go version
cd printalphabet
go run .

Or for a single file:

go run main.go

Limitations

    No third-party libraries unless explicitly allowed.

    Avoid global side effects.

    Not for production use.

    Optimization only after correct implementation.

Academic Integrity

    Solutions must be produced independently.

    Copying code is prohibited.

    Algorithm discussion is allowed without sharing implementations.

Violations may invalidate results.
Status

Educational project for learning and evaluation purposes only.


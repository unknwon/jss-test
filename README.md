# jss-test

A random content generator CLI tool written in Go. It generates random numbers, inspirational quotes, and RGB colors with a formatted terminal display.

## Usage

```bash
go run . [flags]
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-mode` | `all` | Output mode: `number`, `quote`, `color`, or `all` |

### Examples

Generate all random content (default):

```bash
go run .
```

Generate a random number between 1 and 100:

```bash
go run . -mode number
```

Get a random inspirational quote:

```bash
go run . -mode quote
```

Generate a random RGB color:

```bash
go run . -mode color
```

### Sample Output

```
 ╔══════════════════════════════════╗
 ║   🎲  Random Content Generator  ║
 ╚══════════════════════════════════╝

────────────────────────────────────────
 Number     │ 42
────────────────────────────────────────
 Quote      │ "Talk is cheap. Show me the code. – Linus Torvalds"
────────────────────────────────────────
 Color      │ RGB(128, 0, 255) / #8000ff
────────────────────────────────────────
```

## Build

```bash
go build -o jss-test .
./jss-test -mode quote
```

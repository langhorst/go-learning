# 2. Predeclared Types and Declarations

## The Predeclared Types

- _predeclared_ types are built-in to the language
- Go assigns a default _zero value_ to any variable that is declared but not assigned a value
- A _literal_ is an explicitly specified number, character, or string. There are four kinds:
  1. _integer literal_ - a sequence of numbers
    - base 10 by default
    - different prefixes are used to indicate other bases
      - `0b` for binary (base 2)
      - `0o` for octal (base 8)
      - `0x` for hexadecimal (base 16)
    - using an underscore (`_`) makes it easier to read longer literals
      - `1_234` represents one thousand two hundred thirty four in base 10
  2. _floating point literal_ - has a decimal point to indicate the fractional portion of the value
    - `0x12.34p5` which is equal to `582.5` in base 10
    - you can also user underscores to format your floating-point literals
  3. _rune literal_ - represents a character and is surrounded by single quotes
    - single quotes and double quotes are not interchangeable
    - Unicode characters (`'a'`)
    - 8-bit octal number (`'\141'`)
    - 8-bit hexadecimal numbers (`'\x61'`)
    - 16-bit hexadecimal numbers (`'\u0061'`)
    - 32-bit Unicode numbers (`'\U00000061'`)
    - several backslash-escaped runes:
      - newline: `'\n'`
      - tab: `'\t'` 
      - backslash: `'\\'`
  4. _string literal_ - there are two ways to indicate string literals:
    - use double quotes to create an _interpreted string literal_ (e.g., `"Greetings and Salutations"`)
    - use backquotes to create a _raw string literal_ if you need to include backslashes, double quotes, or newlines in your string
- Boolean variables are `true` or `false` and default to `false`
- Numeric types


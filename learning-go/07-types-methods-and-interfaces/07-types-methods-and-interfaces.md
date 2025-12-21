# 7. Types, Methods, and Interfaces

- Go is designed to encourage best practices:
  - Avoid inheritance
  - Encourage composition

## Types in Go

```go
type Person struct {
  FirstName string
  LastName string
  Age int
}
```

- This should be read as declaring a user-defined type with the name `Person` to have the _underlying type_ of the struct literal that follows
- You can use any primitive type or compound type literal to define a concrete type. Examples:

```go
type Score int
type Converter func(string)Score
type TeamScores map[string]Score
```
```
```

- Types can be declared at any block level
  - From the package block down
  - Bound to the scope of the block they are defined within

## Methods

```go
type Person struct {
  FirstName string
  LastName string
  Age int
}

func (p Person) String() string {
  return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}
```

- Method declarations look like function declarations with one addition: the _receiver_ specification



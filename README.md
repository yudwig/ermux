# ermux

ermux is a minimal GO library for multi-error handling.

## Installation

```sh
go get https://github.com/yudwig/ermux
```

## Usage

* import 
```go
import(
  "github.com/yudwig/ermux"
)
```

* example 

```go
errs := make([]error, 3)
a, errs[0] = exec()
b, errs[1] = exec()
c, errs[2] = exec()

if ermux.Some(errs) {
	return ermux.First(errs)
}
```

## Features

ermux simplify below code.

```go
a, err = exec()
if err != nil {
  return err
}
b, err = exec()
if err != nil {
  return err
}
c, err = exec()
if err != nil {
  return err
}
```

## Docs

ermux has only 4 functions.

| I/F     |  description   |
|:------------------------ |:------------------------------------------- |
| **Some** ([]error) bool      | Returns true if input has some error(!= nil).  |
| **First** ([]error) error    |  Returns the first error(!= nil) of input.  |
| **Last** ([]error) error     | Returns the last error(!= nil) of input.  |
| **Filter** ([]error) []error | Removes empty(= nil) elements from input error slice.|



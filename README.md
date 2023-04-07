# Go CUstom Types for MONgo

Go Custom Types for Mongo (gcutmon) provides a set of custom types for working with MongoDB databases in Go. It includes additional data types that are not natively supported by MongoDB, making it easier to work with complex data types such as decimals in your MongoDB databases.

## Installation

To use gcutmon in your Go project, you can install it using `go get`:

```sh
go get github.com/sunboyy/gcutmon
```

## Custom Types

Here's a list of the custom data types included in gcutmon:

- `Decimal`: Derived from `github.com/shopstring/decimal.Decimal`

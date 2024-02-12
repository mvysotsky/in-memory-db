# In-Memory Database

This is a simple example of in-memory key-value database with support for transactions.

## Features

- Key-value storage with `Get`, `Set`, and `Delete` operations.
- Transaction support with `StartTransaction`, `Commit`, and `RollBack`.
- Nested transactions.

## Usage

To use the database, create a new instance and use the provided methods:

```go
db := NewMemoryDB()
db.Set("key1", "value1")
value, _ := db.Get("key1")
db.Delete("key1")
```

To start a transaction:

```go
db.StartTransaction()
db.Set("key1", "value2")
// Either commit the transaction
db.Commit()
// Or roll back the changes
db.RollBack()
```

## Running

Run usage example:

```sh
go build .
./inmemorydb
```


## Testing

Run tests using the following command:

```sh
go test -v
```

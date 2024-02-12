package main

import "fmt"

func main() {
	db := NewMemoryDB()
	db.Set("key1", "value1")
	value, _ := db.Get("key1")
	fmt.Println("key1:", value)

	db.StartTransaction()
	db.Set("key1", "value2")
	value, _ = db.Get("key1")
	fmt.Println("key1 after transaction change:", value)

	db.RollBack()
	value, _ = db.Get("key1")
	fmt.Println("key1 after rollback:", value)

	db.StartTransaction()
	db.Set("key1", "value3")
	db.Commit()
	value, _ = db.Get("key1")
	fmt.Println("key1 after commit:", value)
}

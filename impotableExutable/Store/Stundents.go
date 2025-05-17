package store

import "fmt"

type store struct {
}

func (s store) SaveRecord(record interface{}) {
	fmt.Println("Record ",record)
}
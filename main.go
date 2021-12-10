package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name struct {
		name string
	}
	err := json.Unmarshal(nil, &name)
	fmt.Println(err)
}

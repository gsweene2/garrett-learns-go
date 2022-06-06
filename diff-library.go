package main

import (
	"fmt"

	"github.com/r3labs/diff"
)

func DiffLibrary() diff.Changelog {
	from := make(map[string]map[string]string)
	item_a_from := map[string]string{"hello": "world"}
	from["a"] = item_a_from
	to := make(map[string]map[string]string)
	item_a_to := map[string]string{"hello": "world1"}
	to["a"] = item_a_to

	changelog, err := diff.Diff(from, to)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return changelog
}

// func main() {
// 	result := DiffLibrary()
// 	fmt.Printf("result: %v\n", result)
// 	b, err := json.Marshal(result)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("result as json: %v\n", result)
// 	fmt.Println(string(b))
// }

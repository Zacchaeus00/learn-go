package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := map[string]string{}
	name := ""
	addr := ""

	_, err := fmt.Scanf("%s %s", &name, &addr)
	if err != nil {
		return
	}
	data["name"] = name
	data["addr"] = addr
	jobj, err := json.Marshal(data)
	fmt.Println(jobj)
}

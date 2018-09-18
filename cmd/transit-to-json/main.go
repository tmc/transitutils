package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/russolsen/transit"
	"github.com/tmc/transitutils"
)

func main() {
	decoder := transit.NewDecoder(os.Stdin)
	value, err := decoder.Decode()
	if err != nil {
		log.Fatal(err)
	}
	v, err := transitutils.ToGo(value)
	if err != nil {
		log.Fatal(err)
	}
	j, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(j))
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := dog{"alex", "salsicha", 3}
	fmt.Println(c)

	dogEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(dogEmJSON)

	abc := bytes.NewBuffer(dogEmJSON)

	fmt.Println(abc)

}

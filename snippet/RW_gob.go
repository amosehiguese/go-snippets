package snippet

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

type GOBPost struct {
	Id        int
	Content   string
	Author    string
}

func Gstore(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(filename, buffer.Bytes(), 0600 )
	if err != nil {
		panic(err)
	}
}

func load(data interface{}, filename string) {
	raw, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func GOBmain() {
	post := GOBPost{Id: 1, Content: "Hello World!", Author: "Amos"}
	Gstore(post, "post1")
	var postRead GOBPost
	load(&postRead, "post1")
	fmt.Println(postRead)
}
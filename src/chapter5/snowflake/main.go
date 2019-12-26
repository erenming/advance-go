package main

import (
	"io/ioutil"
	"log"
)

func main() {
	headerMap := make(map[string][]byte)
	for i := 0; i < 5; i++ {
		name := "/path/to/file"
		data, err := ioutil.ReadFile(name)
		if err != nil {
			log.Fatal(err)
		}
		headerMap[name] = data[:1]  // data切片所指向的底层数组被锁定了
	}
	// do some thing
}

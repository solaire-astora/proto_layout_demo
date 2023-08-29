package main

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/encoding/protowire"
)

func main() {
	bytes, err := os.ReadFile("test.bin")
	if err != nil {
		panic(err)
	}
	idx := 0
	for len(bytes) > 0 {
		number, tag, ret := protowire.ConsumeField(bytes)
		fmt.Printf("field %d, wire type %d, offset %d\n", number, tag, idx)
		if ret < len(bytes) {
			bytes = bytes[ret:]
		} else {
			bytes = []byte{}
		}
		idx += ret
	}
}

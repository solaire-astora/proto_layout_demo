//go:build generate
// +build generate

package main

import (
	"log"
	"os"

	"demo/testpb"

	"google.golang.org/protobuf/proto"
)

func generateBin() {
	msg := testpb.MyMessage{
		Name: "123",
		Age:  10,
		Nested: []*testpb.NestedMessage{
			{
				Country: "China",
				Code:    -10086,
			},
		},
	}

	bytes, err := proto.Marshal(&msg)
	if err != nil {
		log.Fatalln(err)
	}
	err = os.WriteFile("test.bin", bytes, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	generateBin()
}

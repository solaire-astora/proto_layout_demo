package main

import (
	"bytes"
	"demo/testpb"
	"fmt"
	"os"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
)

func marshal() {
	msg := testpb.MyMessage{
		Name: "123",
		Age:  10,
	}
	bytes1, _ := proto.Marshal(&msg)

	bytes2 := []byte{}
	bytes2 = protowire.AppendTag(bytes2, 1, protowire.BytesType)
	bytes2 = protowire.AppendString(bytes2, "123")
	bytes2 = protowire.AppendTag(bytes2, 2, protowire.VarintType)
	bytes2 = protowire.AppendVarint(bytes2, 10)

	fmt.Println(bytes.Equal(bytes1, bytes2))
}

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

	marshal()
}

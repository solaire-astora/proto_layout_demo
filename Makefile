.PHONY: all build

PROTO_FILE = test.proto
GO_OUT_DIR = ./
OUTPUT_FILE = proto_demo

all: build

build:
	protoc --go_out=$(GO_OUT_DIR) $(PROTO_FILE)
	go run -tags=generate gen_bin.go
	go build -o $(OUTPUT_FILE)

clean:
	rm -rf $(GO_OUT_DIR)/testpb
	rm -f test.bin $(OUTPUT_FILE)

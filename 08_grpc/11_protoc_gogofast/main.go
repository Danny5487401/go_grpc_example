package main

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	searchproto "go_package_example/08_grpc/11_protoc_gogofast/proto"
	"log"
)

func main() {
	req := &searchproto.SearchRequestParam{
		QueryText: "danny",
		Limit:     10,
		Type:      searchproto.SearchRequestParam_PC,
	}
	data, err := proto.Marshal(req)
	if err != nil {
		log.Fatal("Marshal err : err")
	}
	// send data
	fmt.Println(string(data))

	var respData []byte
	var result = searchproto.SearchResultPage{}
	if err = proto.Unmarshal(respData, &result); err == nil {
		fmt.Println(result)
	} else {
		log.Fatal("Unmarshal err : err")

	}
}

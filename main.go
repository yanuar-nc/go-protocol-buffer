package main

import (
	"fmt"
	"log"

	proto "github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("Learning Protocol Buffers :)")

	tags := []*Article_Tags{
		&Article_Tags{
			TagId: 1,
			Name:  "Golang",
		},
		&Article_Tags{
			TagId: 2,
			Name:  "Programming",
		},
		&Article_Tags{
			TagId: 3,
			Name:  "Protocol Buffers",
		},
	}

	article := &Article{
		Id:      1,
		Title:   "Cara membuat Protocol Buffer menggunakan Golang",
		Content: "Lorem ipsum color si amet",
		Status:  Article_DRAFT,
		Tags:    tags,
	}

	data, err := proto.Marshal(article)
	if err != nil {
		log.Fatal("Marshalling Error: ", err)
	}

	fmt.Println(string(data), "\n============= Unmarshalling ===============")

	unmarshArticle := &Article{}
	err = proto.Unmarshal(data, unmarshArticle)

	fmt.Println("Title: ", unmarshArticle.GetTitle())
	fmt.Println("Content: ", unmarshArticle.GetContent())
	fmt.Println("Status: ", unmarshArticle.GetStatus())

	fmt.Print("Tags: ")
	for _, tag := range unmarshArticle.GetTags() {
		fmt.Print(tag.GetName(), ", ")
	}
}

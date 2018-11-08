package main

import (
	"fmt"
	"log"
	"os"

	proto "github.com/golang/protobuf/proto"
	"github.com/yanuar-nc/go-protocol-buffer/pb"
)

func main() {
	fmt.Println("Learning Protocol Buffers :)")

	// Set data `tags` protobuf
	tags := []*pb.Article_Tags{
		&pb.Article_Tags{
			TagId: 1,
			Name:  "Golang",
		},
		&pb.Article_Tags{
			TagId: 2,
			Name:  "Programming",
		},
		&pb.Article_Tags{
			TagId: 3,
			Name:  "Protocol Buffers",
		},
	}

	// Set data `maps` protobuf
	socmedStats := make(map[string]*pb.Article_SocialMediaStatisticField)
	socmedStats["facebook"] = &pb.Article_SocialMediaStatisticField{
		Like:     13,
		Share:    40,
		Comments: 2,
	}
	socmedStats["twiiter"] = &pb.Article_SocialMediaStatisticField{
		Like:     413,
		Share:    85,
		Comments: 100,
	}

	// Set data `oneof` protobuf
	process := &pb.Article_Update{
		Update: true,
	}

	// Set data to article message
	article := &pb.Article{
		Id:                   1,
		Title:                "Cara membuat Protocol Buffer menggunakan Golang",
		Content:              "Lorem ipsum color si amet",
		Status:               pb.Article_DRAFT,
		Tags:                 tags,
		ProcessOneof:         process,
		SocialMediaStatistic: socmedStats,
	}

	data, err := proto.Marshal(article)
	if err != nil {
		log.Fatal("Marshalling Error: ", err)
		os.Exit(1)
	}

	fmt.Println(string(data))
	fmt.Println(len(string(data)))
	fmt.Println("\n*=================== Unmarshalling =====================*\n")

	unmarshArticle := &pb.Article{}
	err = proto.Unmarshal(data, unmarshArticle)
	if err != nil {
		log.Fatal("Unmarshalling Error: ", err)
		os.Exit(1)
	}
	fmt.Println("Title: ", unmarshArticle.GetTitle())
	fmt.Println("Content: ", unmarshArticle.GetContent())
	fmt.Println("Status: ", unmarshArticle.GetStatus())

	// get `repeated` protobuf
	fmt.Print("Tags: ")
	for _, tag := range unmarshArticle.GetTags() {
		fmt.Printf("%s, ", tag.GetName())
	}

	// get `map` protobuf
	fmt.Println("Social Media Statistic: ")
	for socmed, v := range unmarshArticle.GetSocialMediaStatistic() {
		fmt.Printf("%s: Like(%b), Share(%b) And Comments(%b) \n", socmed, v.GetLike(), v.GetShare(), v.GetComments())
	}

	// Get `Oneof` protobuf
	switch unmarshArticle.ProcessOneof.(type) {
	case *pb.Article_Insert:
		fmt.Println("Article has been save successfully")
	case *pb.Article_Update:
		fmt.Println("Article has been update successfully")
	case *pb.Article_Delete:
		fmt.Println("Article has been delete successfully")
	case *pb.Article_Archive:
		fmt.Println("Article has been archive successfully")
	}

}

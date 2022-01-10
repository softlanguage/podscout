package main

import (
	"fmt"

	"softlang.net/podsdog/dog"
	"softlang.net/podsdog/models"
)

func main() {
	item := new(models.PodItem)
	item.Name = "hello"
	fmt.Println(item)
	dog.EnsureService(item)
	fmt.Printf("%+v\n", item)
}

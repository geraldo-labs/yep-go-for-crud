package main

import "github.com/geraldo-labs/yep-go-for-crud/pkg/admin"

func main() {
	crud := admin.New()
	crud.Init()
}
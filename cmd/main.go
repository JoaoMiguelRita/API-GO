package main

import (
	"auction_golang/internal/entity/auction_entity"
	"fmt"
)

func main() {
	a := auction_entity.CreateAuction("Opalla", auction_entity.New)

	fmt.Printf("id: %s nome %s", a.Id, a.ProductName)
}

package main

import (
	"./gift"
	"./users"
	"fmt"
)

func main() {
	userShyntas := users.NewUser("Shyntas", "+7(777)777-77-77")

	giftList := gift.GiftList{Recipient: *userShyntas}

	giftList.AddGift(&gift.Book{
		Title:  "Jack London",
		Author: "Martin Eden",
	})

	giftList.AddGift(&gift.GiftCard{
		Store:  "Marwin",
		Amount: 10000,
	})

	err := giftList.PresentRandomGift()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Finished")
}

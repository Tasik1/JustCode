package gift

import (
	"../users"
	"fmt"
	"math/rand"
)

type Gift interface {
	Present()
}

type GiftCollection interface {
	AddGift(g Gift)
	PresentRandomGift() error
}

type GiftList struct {
	Recipient users.User
	Gifts     []Gift
}

func (gl *GiftList) AddGift(g Gift) {
	gl.Gifts = append(gl.Gifts, g)
}

func (gl *GiftList) PresentRandomGift() error {
	if len(gl.Gifts) == 0 {
		fmt.Println("Увы, нет желаемых подарков для", gl.Recipient.Name)
		return fmt.Errorf("empty favorites list")
	}
	randomIndex := rand.Intn(len(gl.Gifts))
	randomGift := gl.Gifts[randomIndex]
	fmt.Println("Подарок для", gl.Recipient.Name+":")
	randomGift.Present()
	return nil
}

package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

const (
	//黑桃
	Spade = 0
	//红桃
	Hearts = 1
	//梅花
	Club = 2
	//方块
	Diamond = 3
)

type Poker struct {
	Num int
	Flower int
}

func (p Poker)PokerSelf()string  {
	var buffer string

	switch p.Flower {
	case Spade:
		buffer += "♤"
	case Hearts:
		buffer += "♡"
	case Club:
		buffer += "♧"
	case Diamond:
		buffer += "♢"
	}
	switch p.Num {
	case 13:
		buffer += "2"
	case 12:
		buffer += "A"
	case 11:
		buffer += "K"
	case 10:
		buffer += "Q"
	case 9:
		buffer += "J"
	default:
		buffer += strconv.Itoa(p.Num+2)
	}

	return buffer
}

func CreatePokers()(pokers Pokers)  {
	for i := 1; i < 14; i++ {
		for j := 0; j < 4; j++ {
			pokers = append(pokers,Poker{
				Num: i,
				Flower: j,
			})
		}
	}
	return
}

type Pokers []Poker

/*func (p Pokers)Print()  {
	for _, i2 := range p {
		fmt.Print(i2.PrintPoker()," ")
	}
	fmt.Println()
}*/

func print(a [52]string){
	for i := 0;i < 52;i++{
		fmt.Println(a[i])
	}
}
func xipai(Pokers []Poker)  {
	for i := 1;i < 10;i++{
		for j := 0;j < 52;j++{
			var temp Poker
			var n int
			n = rand.Intn(51)
			temp = Pokers[j]
			Pokers[j] = Pokers[n]
			Pokers[n] = temp
		}
	}
}
func w (a [52]string,  Pokers []Poker){
	for i := 0;i < 52;i++{
		a[i] = Pokers[i].PokerSelf()
		fmt.Print(a[i])
	}
}
func paixu(Pokers []Poker)  {
	for o := 1;o < 53;o++{
		for i := 0;i < 51;i++{
			if Pokers[i].Num > Pokers[i+1].Num{
				var temp Poker
				temp = Pokers[i]
				Pokers[i] = Pokers[i+1]
				Pokers[i+1] = temp
			}
		}
	}
	for i := 0;i < 52;i+=4{
		for j := i;j <= i+3;j++{
			if Pokers[i].Flower > Pokers[i+1].Flower{
				var temp Poker
				temp = Pokers[i]
				Pokers[i] = Pokers[i+1]
				Pokers[i+1] = temp
			}
		}
	}
}
func main() {
	var a1 [52]string
	var a2 [52]string
	Pokers := CreatePokers()
	xipai(Pokers)
	w(a1,Pokers)
	fmt.Println("")
	paixu(Pokers)
	w(a2,Pokers)
}



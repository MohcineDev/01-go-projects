package main

import (
	"fmt"

	"piscine"
)

func main() {
	// quest 07

	// fmt.Println(piscine.AppendRange(5, 10))

	// fmt.Println(piscine.AppendRange(10, 5))
	// MakeRange(-9145814, -9145796) == []int(nil) instead of []int{-9145814, -9145813, -9145812, -9145811, -9145810, -9145809, -9145808, -9145807, -9145806, -9145805, -9145804, -9145803, -9145802, -9145801, -9145800, -9145799, -9145798, -9145797}
	// fmt.Println(piscine.MakeRange(5, 10))
	// fmt.Println(piscine.MakeRange(10, 5))
	// fmt.Println(piscine.MakeRange(-9145814, -9145796))
	//	MakeRange(-7346006, -7346025) == []int{} instead of []int(nil)
	//	fmt.Println(piscine.MakeRange(-7346006, -7346025))

	// fmt.Println(piscine.Unique("foo", "boo"))
	// fmt.Println(piscine.Unique("abc", "def"))

	/// concatparams
	// test := []string{"Hello", "how", "are", "you?"}
	// fmt.Println(piscine.ConcatParams(test))

	//	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello\nhow\nare\nyou?"))
	// fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello  how"))
	// fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello how are you?"))
	// fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("SOgS8,jERM!P  !mF14cI(ocq* &(Q<WE3iZ9@DS z$m2)%gW|1/|Y [F:':K]?F6KYq 1GNLYv+Q|3ho_ 9 'PtiKd.1B(Y 9C=4a,9W9o\"k3"))
	// fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello\nhow\nare\nyou?"))

	//    SplitWhiteSpaces("SOgS8,jERM!P  !mF14 cI(ocq* &(Q<WE3iZ9@DS z$m2)%gW|1/|Y [F:':K]?F6KYq 1GNLYv+Q|3ho_ 9 'PtiKd.1B(Y 9C=4a,9W9o\"k3")
	//    == []string{"SOgS8,jERM!P","", "!mF14", "cI(ocq*", "&(Q<WE3iZ9@DS", "z$m2)%gW|1/|Y", "[F:':K]?F6KYq", "1GNLYv+Q|3ho_", "9", "'PtiKd.1B(Y", "9C=4a,9W9o\"k3"}
	//     f []string{"SOgS8,jERM!P", "!mF14", "cI(ocq*", "&(Q<WE3iZ9@DS", "z$m2)%gW|1/|Y", "[F:':K]?F6KYq", "1GNLYv+Q|3ho_", "9", "'PtiKd.1B(Y", "9C=4a,9W9o\"k3"}

	// SplitWhiteSpaces("Hello how are you?")
	// == []string{"H", "e", "l", "l", "o", "", "h", "o", "w", "", "a", "r", "e", "", "y", "o", "u", "?"}
	// instead of []string{"Hello", "how", "are", "you?"}

	// a := piscine.SplitWhiteSpaces("Hello how are you?")
	// piscine.PrintWordsTables(a)

	// s := "HelloHAhowHAareHAyou?"
	//s := "Hellobbhowbbarebbyou?"
	// s := "azonQ1i6DRn0p!==!YXChCPb1Fstu0!==!Y6FgwvSLkSjaC!==!JYbxVGt2xUDWj!==!a1EG9YV8rh1me!==!ZnSoQCa5SLBqq!==!fTRpOSNWfnQCO!==!C1U0FjjsR97Kq"
	//	== []string{"azonQ1i6DRn0p", "=!YXChCPb1Fstu0", "=!Y6FgwvSLkSjaC", "=!JYbxVGt2xUDWj", "=!a1EG9YV8rh1me", "=!ZnSoQCa5SLBqq", "=!fTRpOSNWfnQCO", "=!C1U0FjjsR97Kq"}
	//	of []string{"azonQ1i6DRn0p", "YXChCPb1Fstu0", "Y6FgwvSLkSjaC", "JYbxVGt2xUDWj", "a1EG9YV8rh1me", "ZnSoQCa5SLBqq", "fTRpOSNWfnQCO", "C1U0FjjsR97Kq"}
	///	fmt.Printf("%#v\n", piscine.Split(s, "bb"))

	// result := piscine.ConvertBase("101011", "01150", "0123456789")
	// fmt.Println(result)

	// piscine.Sudoku()
	// piscine.PrintNbr(-123)
	// piscine.PrintNbr(0)
	// piscine.PrintNbr(123)
	// z01.PrintRune('\n')

	// a := []int{1, 2, 3, 4, 5, 6}
	// piscine.ForEach(piscine.PrintNbr, a)

	// a := []int{1, 2, 3, 4, 5, 6}
	// result := piscine.Map(piscine.IsPrime, a)
	// fmt.Println(result)

	// a1 := []string{"Hello", "how", "are", "you"}
	// a2 := []string{"This", "is", "4", "you"}

	// result1 := piscine.Any(piscine.IsNumeric, a1)
	// result2 := piscine.Any(piscine.IsNumeric, a2)

	// fmt.Println(result1)
	// fmt.Println(result2)

	// tab1 := []string{"Hello", "how", "are", "you"}
	// tab2 := []string{"This", "1", "is", "4", "you"}
	// answer1 := piscine.CountIf(piscine.IsNumeric, tab1)
	// answer2 := piscine.CountIf(piscine.IsNumeric, tab2)
	// fmt.Println(answer1)
	// fmt.Println(answer2)

	//) == false instead of true
	// a1 := []int{0, 1, 2, 3, 4, 5}
	// result1 := piscine.IsSorted(f, a1)
	// fmt.Println(result1)

	// piscine.DescendComb()

	// a := []int{1, 2, 4, 3, 4, 1, 2, 3, 4}
	// unmatch := piscine.Unmatch(a)
	// fmt.Println(unmatch)

	// fmt.Println(piscine.FoodDeliveryTime("pizza"))
	// fmt.Println(piscine.FoodDeliveryTime("burger"))
	// fmt.Println(piscine.FoodDeliveryTime("chips"))
	// fmt.Println(piscine.FoodDeliveryTime("nuggets"))
	// fmt.Println(piscine.FoodDeliveryTime("burger") + piscine.FoodDeliveryTime("chips") + piscine.FoodDeliveryTime("nuggets"))

	// steps := piscine.CollatzCountdown(12)
	// fmt.Println(steps)
	//  == 1287619551094458777 instead of -2964871301041485964
	//	middle := piscine.Abort(2, 3, 8, 5, 7)
	// middle := piscine.Abort(1287619551094458777, -2964871301041485964, -2275620264057798100, -7192713852289106210, -7502979245297788682)
	// fmt.Println(middle)
	// ShoppingSummaryCounter("Burger Burger Water Coffe    e Water Chips Carrot") ==
	// map[string]int{"Burger":2, "Carrot":1, "Chips":1, "Coffe":1, "Water":2, "e":1}

	// instead of map[string]int{"":3, "Burger":2, "Carrot":1, "Chips":1, "Coffe":1, "Water":2, "e":1}
	// ShoppingSummaryCounter("Burger Burger Water Coffe    e Water Chips Carrot") ==
	// map[string]int{"Burger":2, "Carrot":1, "Chips":1, "Coffe":1, "Water":2, "e":1}
	// instead of map[string]int{"":3, "Burger":2, "Carrot":1, "Chips":1, "Coffe":1, "Water":2, "e":1}

	// ShoppingSummaryCounter("Burger Burger Water Coffe    e Water Chips Carrot") ==
	// map[string]int{"Burger":2, "Carrot":1, "Chips":1, "Coffe":1, "Water":2, "e":1}
	// instead of map[string]int{"":3, "Burger":2, "Carrot":1, "Chips":1, "Coffe":1, "Water":2, "e":1}

	// summary := "Burger Burger Water Coffe    e Water Chips Carrot"
	// // summary := "Burger Water Carrot Coffee Water Water Chips Carrot Carrot Burger Carrot Water"
	// for index, element := range piscine.ShoppingSummaryCounter(summary) {
	// 	fmt.Println(index, "=>", element)
	// }
	// const N = 6

	// a := make([]string, N)

	// a[0] = "a"
	// a[2] = "b"
	// a[4] = "c"

	// for _, v := range a {
	// 	fmt.Println(v)
	// }

	// fmt.Println("Size after compacting:", piscine.Compact(&a))

	// for _, v := range a {
	// 	fmt.Println(v)
	// }

	// deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	// piscine.DealAPackOfCards(deck)

	// fmt.Print(piscine.JumpOver("1010101010"))
	// fmt.Print(piscine.JumpOver(""))
	// fmt.Print(piscine.JumpOver("t w e l v e"))
	// fmt.Print(piscine.JumpOver("12"))

	// fmt.Println(piscine.StringToIntSlice("A quick brown fox jumps over the lazy dog"))
	// fmt.Println(piscine.StringToIntSlice("Converted this string into an int"))
	// fmt.Println(piscine.StringToIntSlice("hello THERE"))

	// x := 5
	// y := &x
	// z := &y
	// a := &z

	// w := 2
	// b := &w

	// u := 7
	// e := &u
	// f := &e
	// g := &f
	// h := &g
	// i := &h
	// j := &i
	// c := &j

	// k := 6
	// l := &k
	// m := &l
	// n := &m
	// d := &n

	// fmt.Println(***a)
	// fmt.Println(*b)
	// fmt.Println(*******c)
	// fmt.Println(****d)

	// piscine.Enigma(a, b, c, d)

	// fmt.Println("After using Enigma")
	// fmt.Println(***a)
	// fmt.Println(*b)
	// fmt.Println(*******c)
	// fmt.Println(****d)

	// fmt.Println(piscine.DescendAppendRange(10, 5))
	// fmt.Println(piscine.DescendAppendRange(5, 10))

	// slice := []string{"Pineapple", "Honey", "Mushroom", "Tea", "Pepper", "Milk"}
	// fmt.Println(piscine.ShoppingListSort(slice))

	// fmt.Println(piscine.ReverseMenuIndex([]string{"desserts", "mains", "drinks", "starters"}))

	// p4 := []string{"4th Place"}
	// p3 := []string{"3rd Place"}
	// p2 := []string{"2nd Place"}
	// p1 := []string{"1st Place"}

	// position := [][]string{p4, p3, p2, p1}
	// fmt.Println(piscine.PodiumPosition(position))

	// fmt.Println(piscine.ActiveBits(7))

	// a := []int{23, 123, 1, 11, 55, 93}
	// max := piscine.Max(a)
	// fmt.Println(max)

	fmt.Println(piscine.RockAndRoll(4))
	fmt.Println(piscine.RockAndRoll(9))
	fmt.Println(piscine.RockAndRoll(6))
}

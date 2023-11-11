package main

import (
	"github.com/ichungelo/mock-interview/mock"
)

func main() {

	// //!MOCK 1
	// var (
	// 	arr    = []int{1, 2, 3, 1, 2, 3, 8, 9}
	// 	target = 9
	// )
	// resOne := mock.GetNumber(arr, target)
	// fmt.Println(resOne)

	// //!MOCK 2
	// var (
	// 	max = 100
	// )

	// resTwo := mock.FizzBuzz(max)
	// fmt.Println(resTwo)

	// //!MOCK 3
	// var (
	// 	input = "[]{}({})]"
	// )

	// resThree := mock.CloseBracket(input)
	// fmt.Println(resThree)

	// //!MOCK 4
	// var (
	// 	num = 1234
	// )

	// resFour := mock.IntReverser(num)
	// fmt.Println(resFour)

	// //! PALINDROME
	// resPalindrome := mock.Palindrome("ABSBA")
	// fmt.Println(resPalindrome)

	// //! LONGEST TIME
	// var(
	// 	arrKeyTime = [][2]int{
	// 		{0,2},
	// 		{1,9},
	// 		{0,12},
	// 		{2,15},
	// 	}
	// )
	// resLongestTime := mock.LongestTime(arrKeyTime)
	// fmt.Println(resLongestTime)

	// //! COIN
	// var (
	// 	maxCoin = 3
	// 	target  = 4
	// )
	// resCoin := mock.Coin(maxCoin, target)
	// fmt.Println(resCoin)

	// //! MINIMUM GROUP
	// var (
	// 	predators = []int{-1, 8, 6, 0, 7, 3, 8, 9, -1, 6}
	// )
	// resMinimumGroup := mock.MinimumGroup(predators)
	// fmt.Println(resMinimumGroup)

	//! SQUARE
	mock.Square(5)

	//! TRIANGLE
	mock.Triangle(5)

	//! TRIANGLE HOLE
	mock.TriangleHole(5)

	//! ISOSCELES TRIANGLE
	mock.IsoscelesTriangle(5)
}

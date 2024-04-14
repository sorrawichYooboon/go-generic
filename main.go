package main

import (
	"fmt"

	studentmodels "github.com/sorrawichYooboon/gogeneric/models/student"
	"github.com/sorrawichYooboon/gogeneric/utils"
)

func main() {
	fmt.Println("----------- log utils generic -----------")
	utils.PrintWithEmoji("Hello, World!")
	utils.PrintWithEmoji(123)
	utils.PrintWithEmojiOnlyStringNum("Hello, World!")
	utils.PrintWithEmojiOnlyStringNum(123)
	utils.PrintWithEmojiOnlyNumFloat(123)

	fmt.Println("----------- cal utils generic -----------")
	fmt.Println("Min >>>", utils.Min(1, 100))
	fmt.Println("Max >>>", utils.Max(1, 100))

	fmt.Println("----------- stack utils generic -----------")
	stringStack := utils.Stack[string]{}
	stringStack.Push("Hello")
	stringStack.Push("World")
	fmt.Println("Stack Pop >>>", stringStack.Pop())
	fmt.Println("Stack Pop >>>", stringStack.Pop())
	fmt.Println("Stack IsEmpty >>>", stringStack.IsEmpty())

	intStack := utils.Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	fmt.Println("Stack Pop >>>", intStack.Pop())
	fmt.Println("Stack Pop >>>", intStack.Pop())
	fmt.Println("Stack IsEmpty >>>", intStack.IsEmpty())

	fmt.Println("----------- student models generic -----------")
	firstStudent := studentmodels.Student[string]{Name: "John", Score: "A", Age: 20}
	fmt.Println("Get score >>>", firstStudent.GetScore())
	firstStudent.SetScore("B")
	fmt.Println("Get score >>>", firstStudent.GetScore())
	fmt.Println("Get score type >>>", firstStudent.GetScoreType())

	secondStudent := studentmodels.Student[int]{Name: "Jane", Score: 100, Age: 21}
	fmt.Println("Get score >>>", secondStudent.GetScore())
	secondStudent.SetScore(90)
	fmt.Println("Get score >>>", secondStudent.GetScore())
	fmt.Println("Get score type >>>", secondStudent.GetScoreType())
}

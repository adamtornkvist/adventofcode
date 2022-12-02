package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var input = "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000"

//var input =

func main() {

	file, _ := os.Open("input.txt")

	b, _ := ioutil.ReadAll(file)

	input = string(b)

	fmt.Println(file)

	//input = "10000\n\n20\n\n"
	//fmt.Println(testinput)

	//fmt.Println("Testing: " + input[14:20])

	//fmt.Println(input[14] == '\n')
	//fmt.Println(testinput)

	start := 0
	end := 0

	highestCalories := 0
	secondHighestCalories := 0
	thirdHighestCalories := 0
	fmt.Println("Input length: " + strconv.Itoa(len(input)))

	for i := 0; i < len(input); i++ {
		//fmt.Println(i)
		if i+1 == len(input) || (input[i] == '\n' && input[i+1] == '\n') {
			end = i

			fmt.Println("Start: " + strconv.Itoa(start))
			fmt.Println("End: " + strconv.Itoa(end))
			//fmt.Println("The input: " + input[start:end])

			var calories = countElf(input[start:end])

			fmt.Println("Current elfs calories: " + strconv.Itoa(calories))
			fmt.Println("Highest calories:  " + strconv.Itoa(highestCalories))
			start = i + 2
			i += 2
			if calories >= highestCalories {
				thirdHighestCalories = secondHighestCalories
				secondHighestCalories = highestCalories
				highestCalories = calories
			} else if calories >= secondHighestCalories {
				thirdHighestCalories = secondHighestCalories
				secondHighestCalories = calories
			} else if calories >= thirdHighestCalories {
				thirdHighestCalories = calories
			}
		}
	}

	var totalCalories = thirdHighestCalories + secondHighestCalories + highestCalories

	fmt.Println("Third highest calories: " + strconv.Itoa(thirdHighestCalories))
	fmt.Println("Second highest calories: " + strconv.Itoa(secondHighestCalories))
	fmt.Println("Highest calories: " + strconv.Itoa(highestCalories))
	fmt.Println("Total calories: " + strconv.Itoa(totalCalories))
	//fmt.Println("Start: " + strconv.Itoa(start))
	//fmt.Println("End: " + strconv.Itoa(end))
	//fmt.Println(highestCalories)
}

func countElf(s string) int {

	start := 0
	end := 0
	totalCalories := 0

	fmt.Println("The count input: " + s)
	fmt.Println("Input length: " + strconv.Itoa(len(s)))
	for i := 0; i < len(s); i++ {
		//fmt.Println("Index: " + strconv.Itoa(i))
		//fmt.Println("Input[i]: " + string(s[i]))
		if s[i] == '\n' {
			end = i
			fmt.Println("Start: " + strconv.Itoa(start))
			fmt.Println("End: " + strconv.Itoa(end))
			fmt.Println("Just one number: " + s[start:end])

			var calories, _ = strconv.Atoi(s[start:end])
			totalCalories = totalCalories + calories
			start = i + 1

		}
	}

	//fmt.Println("Start: " + strconv.Itoa(start))
	//fmt.Println("End: " + strconv.Itoa(end))
	var calories, _ = strconv.Atoi(s[start:])
	totalCalories += calories
	fmt.Println(totalCalories)
	return totalCalories
}

/*TODO
Fixa så den bryter ut separata alver
Konvertera den alvens olika kalorisiffror från string till int
Summera de int:sen
Om den är högre än den hittils högst sparade så är den nu den nya högsta
Gå genom hela listan
Returnera den högsta
*/

package main

import "fmt"

func main () {
	data := [] float64{43,55,77,99,66,32}

	grades := [] float64{77,76,90,85}

	n := average(data...)

	fmt.Println(n) 
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~")
	printGrades(grades...)

}

func average(sf ...float64) float64 {
	total := 0.0

	for _,v := range sf {
		total += v
	}

	return (total/float64(len(sf)))
}

func printGrades(tg ...float64) {
	for _, v := range tg {
		fmt.Println(v)
	}
}
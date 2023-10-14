package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Welcome to the quiz game!")
	csvfile := flag.String("csvfile", "quizzes.csv", "it has the quizzes")
	flag.Parse()
	file, err := os.Open(*csvfile)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	Reader := csv.NewReader(file)
	data, err := Reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}
	var input string
	var crt int64

	for _, value := range data {
		ques := value[0]
		ans := value[1]

		fmt.Printf("\nQuestion : What %s is ?  ", ques)
		fmt.Scanln(&input)

		if input == ans {
			fmt.Println("Your answer is correct :)")
			crt++
		} else {
			fmt.Println("Your answer is wrong :(")
		}
	}

	fmt.Printf("Your score is %d/%d ", crt, len(data))
}

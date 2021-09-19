package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var name string
	var userRating string

	//Front end
	//Take name as input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your full name")
	name, _ = reader.ReadString('\n')

	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please rate between 1 and 5:")
	userRating, _ = reader.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	//Back end
	fmt.Printf("Hello %v \n Thanks for the rating of %v \n Your rating was recorded at %v", name, mynumRating, time.Now().Format(time.Stamp))
	if mynumRating == 5 {
		fmt.Println("\nDeserves Award")

	} else if mynumRating == 4 || mynumRating == 3 {
		fmt.Println("\nStill honourable")
	} else if mynumRating < 3 {
		fmt.Println("\nNeeds some work")
	}

}

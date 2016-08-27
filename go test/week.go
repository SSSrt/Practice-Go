package main

import (
	"fmt"
)

func main() {
	week := map[string]string{"Monday": "Monday", "Tuesday": "Tuesday", "Wensday": "Wensday", "Thursday": "Thursday", "Friday": "Friday", "Saturday": "Saturday", "Sunday": "Sunday"}
	for _, value := range week {
		fmt.Println(value)
	}
	if _, ok := week["Tuesday"]; ok {
		fmt.Printf("Today is %s\n", week["Tuesday"])
	}
	if _, ok := week["Hollyday"]; ok {
		fmt.Println("Have Hollyday")
	} else {
		fmt.Println("Don't have this day")
	}
}

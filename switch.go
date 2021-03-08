package main

import "fmt"
import "time"

func main()  {
	i:=2
	fmt.Println(i)

	switch  i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}

	whatIam:=func (i interface{})  {
		switch t:=i.(type) {
		case bool:
			fmt.Println("boolean")
		case int:
			fmt.Println("integer")
		default:
			fmt.Println("unknown type",t)
		}
	}

	whatIam(true)
	whatIam(1)
	whatIam("hey")
}
package main

import "fmt"

func main()  {
	fmt.Println("hello world")

	elliot:= &Person{
		Name: "Elliot",
		Age: 24,
	}

	data,err := proto.Marshal(elliot)

	if err != nil{
		log.Fatal("Marshalling Error")
	}
	
	fmt.Println(data)

	newElliot := &Person{}

	err = proto.Unmarshal(data, newElliot)

	if err != nil{
		log.Fatal("Unmarshaling error")
	}

	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.GetName())
}

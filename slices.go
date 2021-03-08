package main

import "fmt"

func main()  {
	s:= make([]string, 3)
	fmt.Println("emp",s)

	s[0]="a"
	s[1]="b"
	s[2]="c"

	fmt.Println("set",s)
	fmt.Println("get",s[1])
	fmt.Println("len", len(s))

	s=append(s,"d")
	s=append(s,"e")
	s=append(s,"f")

	fmt.Println("abd",s)

	c:=make([]string, len(s))
	copy(c,s)
	fmt.Println("copy",c)

	l:=s[2:5]
	fmt.Println(l)

	l=s[:5]
	fmt.Println(l)


}
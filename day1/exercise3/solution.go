package main

import "fmt"

type person struct {
	role string
	days int
	hours int
}
type Employee interface {
	cal() int
}

func (p *person)cal() int{
	if p.role == "fulltime"{
		return p.days*500
	} else if p.role == "contractor"{
		return p.days*100
	}else {
		return p.hours*10
	}
}
func main(){
	var e Employee
	fte := &person{role: "fulltime",days: 28}
	con := &person{role: "contractor",days: 28}
	free := &person{role: "freelancer",hours: 20}
	e=fte
	fmt.Println(e.cal())
	e=con
	fmt.Println(e.cal())
	e=free
	fmt.Println(e.cal())
}
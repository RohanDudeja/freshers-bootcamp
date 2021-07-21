package main

import "fmt"

type Employee interface {
	GetSalary() int
}

type FullTimeEmployee struct {
	Name string
	BasicMultiplier int
}

func (f FullTimeEmployee) GetSalary() int{
	return f.BasicMultiplier*500
}

type Contractor struct {
	Name string
	BasicMultiplier int
}

func (c Contractor) GetSalary() int {
	return c.BasicMultiplier*100
}

type PartTimeEmployee struct {
	Name string
	BasicMultiplier int
}

func (p PartTimeEmployee) GetSalary() int {
	return p.BasicMultiplier*10
}

func main(){
	var e Employee
	fte := &FullTimeEmployee{"rohan_fte", 28}
	con := &Contractor{"rohan_con", 28}
	pte := &PartTimeEmployee{"rohan_pte", 20}
	e=fte
	fmt.Println("Salary of FTE: ", e.GetSalary())
	e=con
	fmt.Println("Salary of CON: ", e.GetSalary())
	e=pte
	fmt.Println("Salary of PTE: ", e.GetSalary())


}
package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	Row int
	Col int
	Val [][]int
}

func (m Matrix)getRowCount() int {
	return m.Row
}
func (m Matrix)getColCount() int {
	return m.Col
}
func (m *Matrix)setValAt(i,j,k int) {
	m.Val[i][j]=k
}
func (m *Matrix)add(m2 Matrix) *Matrix {
	var rows = m.getRowCount()
	var cols = m.getColCount()
	sum := Matrix{rows, cols, make([][]int, rows)}
	for r := 0; r < rows; r++ {
		sum.Val[r] = make([]int, cols)
		for c := 0; c < cols; c++ {
			sum.Val[r][c] = m.Val[r][c] + m2.Val[r][c]
		}
	}
	return &sum
}

func createMatrix(R,C int) *Matrix{
	m:= Matrix{Row: R,Col: C,Val: make([][]int,R)}
	for r := 0 ; r < R ; r++ {
		m.Val[r] = make([]int, C)
		for c := 0; c < C; c++ {
			m.Val[r][c] = r + c
		}
	}
	return &m
}

func (m Matrix)print() {
	json,err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json))
}

func main(){
	mat:= createMatrix(2,2)
	fmt.Println("Matrix: ",mat.Val)
	fmt.Println("number of Rows: ",mat.getRowCount())
	fmt.Println("number of Columns: ",mat.getColCount())
	// set value
	mat.setValAt(0,0,1)
	fmt.Println("Updated Matrix: ",mat.Val)
	fmt.Println("Sum of Two Matrix: ",mat.add(*mat).Val)
	mat.print()
}

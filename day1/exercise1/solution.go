package main

<<<<<<< HEAD
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
	return m.Row
}
func (m *Matrix)setValAt(i,j,k int) {
	m.Val[i][j]=k
}
func (m *Matrix)add(m2 Matrix) *Matrix{
	var rows = m.getRowCount()
	var cols = m.getColCount()
	sum:= Matrix{rows,cols,make([][]int,rows)}
	for r:=0;r< rows; r++ {
		sum.Val[r] = make([]int,cols)
		for c:=0; c<cols;c++{
			sum.Val[r][c]= m.Val[r][c]+m2.Val[r][c]
=======
import "fmt"

type matrix struct {
	row, col int
	val [][]int
}

func (m matrix)getRowCount() int {
	return m.row
}
func (m matrix)getColCount() int {
	return m.row
}
func (m *matrix)setValAt(i,j,k int) {
	m.val[i][j]=k
}
func (m *matrix)add(m2 *matrix) *matrix{
	var rows = m.getRowCount()
	var cols = m.getColCount()
	sum:= matrix{rows,cols,make([][]int,rows)}
	for r:=0;r< rows; r++ {
		sum.val[r] = make([]int,cols)
		for c:=0; c<cols;c++{
			sum.val[r][c]= m.val[r][c]+m2.val[r][c]
>>>>>>> 7649901d913d4441ecb2c8edd3259d75239b8ae6
		}
	}
	return &sum
}

<<<<<<< HEAD
func createMatrix(R,C int) *Matrix{
	m:= Matrix{Row: R,Col: C,Val: make([][]int,R)}
	for r := 0 ; r < R ; r++ {
		m.Val[r] = make([]int, C)
		for c := 0; c < C; c++ {
			m.Val[r][c] = r + c
=======
func createMatrix(R,C int) *matrix{
	m:= matrix{row: R,col: C,val: make([][]int,R)}
	for r := 0 ; r < 2 ; r++ {
		m.val[r] = make([]int, C)
		for c := 0; c < C; c++ {
			m.val[r][c] = r + c
>>>>>>> 7649901d913d4441ecb2c8edd3259d75239b8ae6
		}
	}
	return &m
}
<<<<<<< HEAD
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
=======
func (m matrix)print() {
	for r:=0;r<m.getRowCount();r++{
		fmt.Println(m.val[r])
	}
}
func main(){
	m:= createMatrix(2,2)
	fmt.Println(m.val)
	fmt.Println("number of rows: ",m.getRowCount())
	fmt.Println("number of columns: ",m.getColCount())
	m.setValAt(0,0,1)
	fmt.Println(m.val)
	fmt.Println(m.add(m).val)
	m.print()
>>>>>>> 7649901d913d4441ecb2c8edd3259d75239b8ae6
}
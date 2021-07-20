package main

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
		}
	}
	return &sum
}

func createMatrix(R,C int) *matrix{
	m:= matrix{row: R,col: C,val: make([][]int,R)}
	for r := 0 ; r < 2 ; r++ {
		m.val[r] = make([]int, C)
		for c := 0; c < C; c++ {
			m.val[r][c] = r + c
		}
	}
	return &m
}
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
}
package main

import "fmt"

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	//1、生成源数组
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	//2、输出元数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
	//3、转成稀疏数组
	var sparseArr []ValNode
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				//创建一个valNode结点
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d", i, valNode.row, valNode.col, valNode.val)
		fmt.Println()
	}

}

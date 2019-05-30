package main

import "fmt"

func main() {
	count := 0
	for cock := 0; cock <= 20; cock++ {
		for hen := 0; hen <= 33; hen++ {
			//for chicken:=0;chicken<=100;chicken++{
			//	if (cock+hen+chicken==100) && (cock*5+hen*3+chicken*1/3==100){
			//		fmt.Printf("公鸡：%d 母鸡：%d 小鸡：%d",cock,hen,chicken)
			//		fmt.Println()
			//	}
			//	count++
			//}
			//for chicken:=0;chicken<=100;chicken+=3{
			//	if (cock+hen+chicken==100) && (cock*5+hen*3+chicken*1/3==100){
			//		fmt.Printf("公鸡：%d 母鸡：%d 小鸡：%d",cock,hen,chicken)
			//		fmt.Println()
			//	}
			//	count++
			//}
			chicken := 100 - cock - hen
			if cock*5+hen*3+chicken/3 == 100 && chicken%3 == 0 {
				fmt.Printf("公鸡：%d 母鸡：%d 小鸡：%d", cock, hen, chicken)
				fmt.Println()
			}
			count++
		}
	}
	fmt.Println("程序共运行了：", count, "次！")
}

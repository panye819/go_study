package main

import (
	"fmt"
)

func main() {
	/*
		map

		类似其他语言中的哈希表或者字典，以key-value形式存储数据
		key必须是支持==或者!=比较运算的类型，不可以是函数、map或slice
		map查找比线性搜索快很多，但比使用索引访问数据的类型蛮100倍
		map使用make()床及哦啊牛，支持:=这种简写的方式

		make([KeyType]valueType,cap)，cap表示容量，可以省略
		超出容量会自动扩容，但近来提供一个合理的初始值
		使用len()获得元素个数

		键值对不存在时自动添加，使用delete()删除某键值对
		使用for range对map和slice进行迭代操作
	*/

}

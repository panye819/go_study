package unittest

/**
1、测试用例文件名必须以_test.go结尾
2、测试用例函数必须以Test开头，一般来说就是Test+被测函数名，被测函数名的第一个字母要大写
3、测试用例函数的形参类型必须是*testing.T
4、一个测试用例文件中可以有多个测试用例函数
5、go tesst -v 运行测试命令
6、当出现错误时，可以使用t.Fatalf来格式化输出错误信息并退出程序
7、t.Logf方法可以输出相应的日志
8、测试用例函数并没有放在main函数中，也执行了，说明测试用例框架有一个隐藏的mian函数
9、PASS表示测试用例运行成功，FAIL表示测试用例运行失败

**/
import "testing"

func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("addUpper(10) error,期望值:%v, 实际值:%v", 55, res)
	}
	t.Logf("Test addUpper succ!")
}

package main

/*
使用TestMain作为初始化test，并且使用m.Run()来调用其他tests来完成一些需要初始化
操作的testing，比如数据库连接，文件打开，REST服务登录等
*/

package plugin

import (
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func Test() {

	src := `
    package plugins_test

    import "fmt"


    func printlog() {
    fmt.Println("test")
    }
    `
	intp := interp.New(interp.Options{}) // 初始化一个 yaegi 解释器
	intp.Use(stdlib.Symbols)             // 允许脚本调用（几乎）所有的 Go 官方 package 代码
	_, err := intp.Eval(src)             // 执行go代码
	if err != nil {
		panic(err)
	}

	v, _ := intp.Eval("plugins_test.printlog") // 由于上面以执行加载了一段go代码，所以这里可以通过package+funcname的方式来获取函数
	fu := v.Interface().(func())               // 函数类型转换
	fu()                                       // 函数调用

}

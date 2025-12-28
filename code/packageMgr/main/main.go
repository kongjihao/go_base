package main

// 当使用$GOPATH模式时，当自己写的代码都放在$GOPATH/src目录下时，我们导入自己的代码package，需要从$GOPATH/srch后面继续写起
// 当使用go mod模式时，我们需先在本地项目目录下执行go mod init，生成go.mod文件，
// 这样我们导入自己的代码package时就不需要从$GOPATH/src后面写起，而是直接从 项目/xxx/写包名 即可
import (
	"fmt"
	heihei "go_base/code/packageMgr/calc" // 给包起别名
)

func main() {
	fmt.Println("Hello Main Function!")

	res := heihei.Add(1, 2)

	fmt.Println("res =", res)
}

package main

import (
	"fmt"
	"io/ioutil"

	"github.com/lsq51201314/go-astar"
)

func main() {
	//读取数据
	data, err := ioutil.ReadFile("./test.map")
	if err != nil {
		panic(err)
	}
	//新建寻路
	a := astar.NewAstar(1002, 802)
	a.SetData(data)
	//获取路径
	a.Find(110, 100, 846, 674)
	fmt.Println(a.CheckPoint(0, 0))
	fmt.Println(a.CheckPoint(110, 100))
	fmt.Println(len(a.GetPath()))
}

// gvsOPCexact
// OPC运行控制精确调节
// 2019-10-12 Go语言开发 + OPCDAAuto.dll + webui
package main

import (
	"fmt"
	"time"

	"github.com/konimarti/opc"
)

func main() {
	fmt.Println("OPC运行控制精确调节")
	opc.Debug()
	//连接前需要注册 OPCDAAuto.dll
	client, err := opc.NewConnection(
		"Kepware.KEPServerEX.V6", // ProgId
		[]string{"localhost"},    //  OPC servers nodes
		[]string{"通道 1.设备 1.标记 1", "模拟器示例.函数.FlowMod", "通道 1.设备 1.标记 2"}) // slice of OPC tags
	if err != nil {
		fmt.Println(err)
	}
	//模拟器示例.函数.FlowMod
	// client, _ := opc.NewConnection(
	// 	"Graybox.Simulator",   // ProgId
	// 	[]string{"localhost"}, //  OPC servers nodes
	// 	[]string{"numeric.sin.int64", "numeric.saw.float"} // slice of OPC tags
	// )
	defer client.Close()
	fmt.Println("Kepware.KEPServerEX.V6 已连接")

	time.Sleep(1 * time.Second)
	// read single tag: value, quality, timestamp
	fmt.Println(client.ReadItem("通道 1.设备 1.标记 1").Value)
	fmt.Println(client.ReadItem("通道 1.设备 1.标记 2").Value)
	fmt.Println(client.ReadItem("模拟器示例.函数.FlowMod").Value)
	// read all added tags

	fmt.Println(client.Read())
	for i := 1; i < 10; i++ {
		//fmt.Println(client.Read())
		fmt.Print(i, "\t")
		//Read()顺序是变化的，需要按名解析
		// for k, v := range client.Read() {
		// 	fmt.Print(k, "\t")
		// 	fmt.Print(v.Value, "\t")
		// }
		fmt.Print(client.ReadItem("通道 1.设备 1.标记 1").Value, "\t")
		fmt.Print(client.ReadItem("通道 1.设备 1.标记 2").Value, "\t")
		fmt.Print(client.ReadItem("模拟器示例.函数.FlowMod").Value, "\t")
		fmt.Println("")
		time.Sleep(1 * time.Second)
	}
	lvalue := "8614"
	wterr := client.Write("通道 1.设备 1.标记 2", lvalue)
	if wterr != nil {
		fmt.Println(wterr)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Tag2写入值：  ", lvalue)
	fmt.Print("modify:", "\t")
	fmt.Print(client.ReadItem("通道 1.设备 1.标记 1").Value, "\t")
	fmt.Print(client.ReadItem("通道 1.设备 1.标记 2").Value, "\t")
	fmt.Print(client.ReadItem("模拟器示例.函数.FlowMod").Value, "\t")
	fmt.Println("")
}

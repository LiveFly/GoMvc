// SimpleMVC project main.go
package main

import (
	_ "Areas/Admin/Controllers"
	_ "Controllers"
	. "System/Routing"
	. "System/Web"
	"fmt"
)

//当前配置文件的端口为6080,输入http://localhost:6080/可查看运行结果
//注册路由
func init() {
	//Admin域的标准路由
	RouteTable.AddRote(&RouteItem{
		Name:     "admin_area",
		Url:      "admin/{controller}/{action}",
		Defaults: map[string]interface{}{"controller": "home", "action": "index", "area": "admin"},
	})

	//标准路由
	RouteTable.AddRote(&RouteItem{
		Name:        "default",
		Url:         "{controller}/{action}/{id}",
		Defaults:    map[string]interface{}{"controller": "home", "action": "index", "id": 123},
		Constraints: map[string]string{"id": `^(\d+)$`}})

}
func main() {
	err := App.Run()
	fmt.Println(err)
}

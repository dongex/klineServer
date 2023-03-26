package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "serveLine/routers"
)

func init() {

}
func main() {
	beego.Run()
	//models.OpenFileTest()
}

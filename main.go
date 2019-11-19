package main

import (
	_ "weather/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}


package routers

import (
	"weather/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/city_list", &controllers.MainController{},"*:GetCityList")
	beego.Router("/weather", &controllers.MainController{},"*:GetWeather")
}

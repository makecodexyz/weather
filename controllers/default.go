package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"time"
	"weather/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (c *MainController) GetCityList() {
	cityList := getCityList()
	result := models.ApiResult{StatusCode: 0, StatusMsg: "ok", Data: cityList}
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *MainController) GetWeather() {
	city := c.GetString("city")
	weather := fakeWeatherAPI(city)
	result := models.ApiResult{StatusCode: 0, StatusMsg: "ok", Data: weather}
	c.Data["json"] = result
	c.ServeJSON()
}

func fakeWeatherAPI(city string) *models.CityWeather {
	weather := &models.CityWeather{
		City:        city,
		UpdateTime:  time.Now().Format("2006-01-02 15:04:05"),
		Weather:     "Weather " + city,
		Temperature: "Temperature" + city,
		Wind:        "Wind " + city,
	}
	return weather
}

func getCityList() []models.City {
	content, err := ioutil.ReadFile(beego.AppConfig.String("cities"))
	if err != nil {
		logs.Error(err)
		return []models.City{}
	}
	var cityList []models.City
	err = json.Unmarshal(content, &cityList)
	if err != nil {
		logs.Error(err)
		return []models.City{}
	}
	return cityList
}

package test

import (
	"encoding/json"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
	"weather/models"
	_ "weather/routers"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestCityList(t *testing.T) {
	r, _ := http.NewRequest("GET", "/city_list", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	Convey("TestCityList\n", t, func() {
		var excepted models.ApiResult
		err := json.Unmarshal([]byte(w.Body.String()), &excepted)
		So(err, ShouldBeNil)
		So(excepted.StatusMsg, ShouldEqual, "ok")
		So(len(excepted.Data.([]interface{})), ShouldEqual, 3)
	})
}

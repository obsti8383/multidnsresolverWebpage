package routers

import (
	"github.com/obsti8383/multidnsresolverWebpage/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}

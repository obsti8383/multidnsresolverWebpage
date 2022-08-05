package main

import (
	_ "github.com/obsti8383/multidnsresolverWebpage/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

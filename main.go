package main

import (
    _ "github.com/bigbluebutton/bbb-api-recordings/docs"
    _ "github.com/bigbluebutton/bbb-api-recordings/routers"
    "github.com/bigbluebutton/bbb-api-recordings/controllers"

    "github.com/astaxie/beego"
)

func main() {
    beego.ErrorController(&controllers.ErrorController{})
    beego.Run()
}

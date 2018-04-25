package main

import (
    _ "./docs"
    _ "./routers"
    "./controllers"

    "github.com/astaxie/beego"
)

func main() {
    if beego.RunMode == "dev" {
        beego.DirectoryIndex = true
        beego.StaticDir["/swagger"] = "swagger"
    }
    beego.ErrorController(&controllers.ErrorController{})
    beego.Run()
}

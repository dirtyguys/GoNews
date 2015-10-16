package main

import (
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
)

func main() {
    m := martini.Classic()
    // render html templates from templates directory
    m.Use(render.Renderer())
    m.Use(martini.Static("assets"))

    m.Get("/", func(r render.Render) {
        r.HTML(200, "home", "")
    })
    m.Run()
}

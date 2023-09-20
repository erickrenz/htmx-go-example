package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}  status=${status}  uri=${uri}\n",
	}))
	e.Use(middleware.Recover())

	e.File("/favicon.ico", "static/favicon.ico")
	e.Static("/", "static")

	e.GET("/click", func(c echo.Context) error {
		return c.String(http.StatusOK, "I have been clicked.")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Message from Echo.")
	})

	e.GET("/css", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<div id=\"div1\" class=\"red\">New Content</div>")
	})

	blog_post := "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Sint sit id vitae minus provident officia et iste. Dolores, velit obcaecati unde debitis veniam veritatis a consectetur maxime possimus aliquam eaque odit et aperiam asperiores? Perspiciatis minus voluptatibus debitis alias ad excepturi officiis eveniet et, pariatur architecto ipsum odio fugiat est dolores illum optio nostrum dolorum aliquam tempora. Molestiae laboriosam perferendis reprehenderit sint dolore aut ratione, tempore eos voluptatum veritatis eveniet rem rerum eligendi suscipit vero laborum quo ducimus, voluptas enim facilis incidunt exercitationem. Explicabo cum commodi voluptates placeat omnis excepturi accusamus optio, ipsam odio, hic in blanditiis enim perspiciatis ipsa?"
	e.GET("/blog", func(c echo.Context) error {
		return c.String(http.StatusOK, blog_post)
	})

	global_integer := 0
	e.POST("/bar", func(c echo.Context) error {
		return c.String(http.StatusOK, strconv.Itoa(global_integer))
	})
	e.POST("/foo", func(c echo.Context) error {
		global_integer += 1
		return c.String(http.StatusOK, strconv.Itoa(global_integer))
	})

	e.Logger.Fatal(e.Start(":1323"))
}

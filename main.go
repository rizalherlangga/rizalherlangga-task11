package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/public", "public")

	e.GET("/", home)
	e.GET("/myProject", myProject)
	e.GET("/contact", contact)
	e.GET("/project-detail/:id", projectDetail)
	e.GET("/testimonial", testimonial)
	e.GET("/project", addProject)
	e.POST("/addedProject", addedProject)

	e.Logger.Fatal(e.Start("localhost:4000"))
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}
func myProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/my-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}
func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}
func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Id":    id,
		"Title": "Dumbways Web App",
		"Content": "Lorem ipsum dolor sit amet consectetur adipisicing elit. Ducimus,placeat recusandae, doloremque vel et in impedit omnis velit idporro animi, optio alias dolorem. Beatae a facilis earum!Praesentium at voluptates amet assumenda soluta earum eos nesciunt		eaque. Dolorum recusandae quo et delectus eius optio temporibusdebitis voluptates sed nam?",
	}

	var tmpl, err = template.ParseFiles("views/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}
func testimonial(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}
func addProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}
func addedProject(c echo.Context) error {
	name := c.FormValue("input-nama")
	start := c.FormValue("input-start")
	end := c.FormValue("input-end")
	desc := c.FormValue("input-deskripsi")
	js := c.FormValue("input-js")
	bootstrap := c.FormValue("input-bootstrap")
	golang := c.FormValue("input-golang")
	react := c.FormValue("input-react")

	println("name : " + name)
	println("start : " + start)
	println("end : " + end)
	println("desc : " + desc)
	println("js : " + js)
	println("bootstrap : " + bootstrap)
	println("golang : " + golang)
	println("react : " + react)

	return c.Redirect(http.StatusMovedPermanently, "/myProject")
}

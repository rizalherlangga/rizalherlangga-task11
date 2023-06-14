package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

type project struct {
	Title       string
	Start       string
	End         string
	Duration    string
	Description string
	Tech        []string
}

var dataProject = []project{
	{
		Title:       "FAKTA BANGET NICH !",
		Start:       "24/06/2023",
		End:         "27/06/2023",
		Duration:    "3 hari",
		Description: "MAU TAU FAKTA GA BRO ? RAHASIA DAPUR NIH !",
		Tech:        []string{"js"},
	},
	{
		Title:       "GOKIL SIH INI CUY !",
		Start:       "24/06/2023",
		End:         "27/06/2023",
		Duration:    "3 hari",
		Description: "LU KEREN GUA KEREN, KITA BERDUA ? KEREN !",
		Tech:        []string{"js"},
	},
	{
		Title:       "AZEEEEKKKK !",
		Start:       "24/06/2023",
		End:         "27/06/2023",
		Duration:    "3 hari",
		Description: "WAHHH GOKIL PISAN NIH KEKNYA !",
		Tech:        []string{"js"},
	},
}

func main() {
	e := echo.New()

	e.Static("/public", "public")

	e.GET("/", home)
	e.GET("/myProject", myProject)
	e.GET("/contact", contact)
	e.GET("/project-detail/:id", projectDetail)
	e.GET("/testimonial", testimonial)
	e.GET("/project", addProject)

	e.POST("/edit-project/:id", editProject)
	e.POST("/delete-project/:id", deleteProject)
	e.POST("/addedProject", addedProject)
	e.POST("/project", addProject)

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

	projects := map[string]interface{}{
		"projects": dataProject,
	}

	return tmpl.Execute(c.Response(), projects)
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

	var ProjectDetail = project{}

	for index, item := range dataProject {
		if id == index {
			ProjectDetail = project{
				Title:       item.Title,
				Start:       item.Start,
				End:         item.End,
				Duration:    item.Duration,
				Description: item.Description,
				Tech:        item.Tech,
			}
		}
	}

	item := map[string]interface{}{
		"Project": ProjectDetail,
	}

	var tmpl, err = template.ParseFiles("views/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), item)
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
	Tech := c.Request().Form["Tech"]

	startDate, err := time.Parse("2006-01-02", start)
	if err != nil {
		return err
	}

	endDate, err := time.Parse("2006-01-02", end)
	if err != nil {
		return err
	}

	duration := int(endDate.Sub(startDate).Hours() / 24)
	var durationOutput string

	fmt.Println("name : ", name)
	fmt.Println("start : ", start)
	fmt.Println("end : ", end)
	fmt.Println("desc : ", desc)
	fmt.Println("Technologies : ", strings.Join(Tech, ", "))

	if duration <= 7 {
		durationOutput = fmt.Sprintf("%d day(s)", duration)
	} else if duration > 7 && duration <= 30 {
		weeks := duration / 7
		days := duration % 7
		durationOutput = fmt.Sprintf("%d week(s) %d day(s)", weeks, days)
	} else if duration > 30 && duration <= 365 {
		months := duration / 30
		durationOutput = fmt.Sprintf("%d month(s)", months)
	} else {
		years := duration / 365
		durationOutput = fmt.Sprintf("%d year(s)", years)
	}

	fmt.Println("Duration: ", durationOutput)

	var newProject = project{
		Title:       name,
		Start:       start,
		End:         end,
		Duration:    durationOutput,
		Description: desc,
		Tech:        Tech,
	}

	dataProject = append(dataProject, newProject)

	fmt.Println(dataProject)

	return c.Redirect(http.StatusMovedPermanently, "/myProject")
}
func deleteProject(delete echo.Context) error {
	i, _ := strconv.Atoi(delete.Param("id"))

	fmt.Println("index : ", i)

	dataProject = append(dataProject[:i], dataProject[i+1:]...)

	return delete.Redirect(http.StatusMovedPermanently, "/myProject")
}

func editProject(edit echo.Context) error {
	i, _ := strconv.Atoi(edit.Param("id"))

	fmt.Println("index : ", i)

	dataProject = append(dataProject[:i], dataProject[i+1:]...)
	return edit.Redirect(http.StatusMovedPermanently, "/project")
}

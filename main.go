package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

var (
	counterMain             int
	counterStudentRules     int
	counterLyceumRegulation int
)

type Res struct {
	CounterMain             int
	CounterStudentRules     int
	CounterLyceumRegulation int
}

func main() {
	e := echo.New()
	e.GET("/lyceum-6/main", func(c echo.Context) error {
		counterMain++
		return c.Redirect(http.StatusFound, "https://licey-6.odinedu.ru")
	})
	e.GET("/lyceum-6/student-rules", func(c echo.Context) error {
		counterStudentRules++
		return c.Redirect(http.StatusFound, "https://licey-6.odinedu.ru/wp-content/uploads/2023/08/%D0%9F%D1%80%D0%B0%D0%B2%D0%B8%D0%BB%D0%B0-%D0%B2%D0%BD%D1%83%D1%82%D1%80%D0%B5%D0%BD%D0%BD%D0%B5%D0%B3%D0%BE-%D1%80%D0%B0%D1%81%D0%BF%D0%BE%D1%80%D1%8F%D0%B4%D0%BA%D0%B0-%D0%BE%D0%B1%D1%83%D1%87%D0%B0%D1%8E%D1%89%D0%B8%D1%85%D1%81%D1%8F-2024.pdf")
	})
	e.GET("/lyceum-6/lyceum-regulation", func(c echo.Context) error {
		counterLyceumRegulation++
		return c.Redirect(http.StatusFound, "https://licey-6.odinedu.ru/wp-content/uploads/2023/08/%D0%A3%D1%81%D1%82%D0%B0%D0%B2-%D0%9B%D0%B8%D1%86%D0%B5%D1%8F-6-15.09.2021.pdf")
	})

	e.GET("/counters", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Res{
			CounterMain:             counterMain,
			CounterStudentRules:     counterStudentRules,
			CounterLyceumRegulation: counterLyceumRegulation,
		})
	})
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:5173",
			"https://localhost:5173",
			"http://localhost:5174",
			"https://localhost:5174",
		},
		AllowHeaders: []string{"*"},
	}))

	log.Fatal(e.Start(":8080"))

}

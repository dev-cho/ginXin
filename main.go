package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

type Soccer struct {
	HomeTeam    string `json:"homeTeam"`
	HomeManager string `json:"homeManager"`
	HomeImg     string `json:"homeImg"`
	AwayTeam    string `json:"awayTeam"`
	AwayManager string `json:"awayManager"`
	AwayImg     string `json:"awayImg"`
	SoccerDate  string `json:"soccerDate"`
	Home        string `json:"home"`
	Drow        string `json:"drow"`
	Away        string `json:"away"`
}
type Soccers []Soccer

func main() {
	router := gin.Default()

	//new template engine
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "views/frontend",
		Extension: ".html",
		Master:    "layouts/master",
		Partials:  []string{"partials/ad"},
		Funcs: template.FuncMap{
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})
	router.Static("/public", "./public") //CSS

	router.GET("/", func(ctx *gin.Context) {
		// `HTML()` is a helper func to deal with multiple TemplateEngine's.
		// It detects the suitable TemplateEngine for each path automatically.
		gintemplate.HTML(ctx, http.StatusOK, "index", gin.H{
			"title": "Fontend title!",
		})
	})

	//=========== Backend ===========//

	//new middleware
	mw := gintemplate.NewMiddleware(gintemplate.TemplateConfig{
		Root:      "views/backend",
		Extension: ".html",
		Master:    "layouts/master",
		Partials:  []string{},
		Funcs: template.FuncMap{
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})

	// You should use helper func `Middleware()` to set the supplied
	// TemplateEngine and make `HTML()` work validly.
	backendGroup := router.Group("/admin", mw)

	backendGroup.GET("/", func(ctx *gin.Context) {
		// With the middleware, `HTML()` can detect the valid TemplateEngine.
		gintemplate.HTML(ctx, http.StatusOK, "index", gin.H{
			"title": "Backend title!",
		})
	})

	soccer := &Soccers{
		Soccer{
			HomeTeam: "Huddersfield", HomeManager: "David Wagner", HomeImg: "59",
			AwayTeam: "Chelsea", AwayManager: "Maurizio Sarri", AwayImg: "38",
			SoccerDate: "2018-08-11 23:00:00",
			Home:       "6.05", Away: "3.04", Drow: "1.04",
		},
		Soccer{
			HomeTeam: "Chelsea", HomeManager: "Maurizio Sarri", HomeImg: "38",
			AwayTeam: "Huddersfield", AwayManager: "David Wagner", AwayImg: "59",
			SoccerDate: "2018-08-12 23:00:00",
			Home:       "4.05", Away: "5.04", Drow: "2.04",
		},
		Soccer{
			HomeTeam: "Arsenal", HomeManager: "Unai Emery", HomeImg: "42",
			AwayTeam: "Man City", AwayManager: "Pep Guardiola", AwayImg: "17",
			SoccerDate: "2018-08-13 00:00:00",
			Home:       "2.05", Away: "1.04", Drow: "3.04",
		},
		Soccer{
			HomeTeam: "Man City", HomeManager: "Pep Guardiola", HomeImg: "17",
			AwayTeam: "Arsenal", AwayManager: "Unai Emery", AwayImg: "42",
			SoccerDate: "2018-08-14 00:00:00",
			Home:       "1.05", Away: "1.5", Drow: "2.04",
		},
	}
	router.GET("/soccerList", func(ctx *gin.Context) {
		ctx.JSON(200, soccer)
	})

	router.Run(":9090")
}

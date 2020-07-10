package routes

import (
	"youtube-manager-go/middlewares"
	"youtube-manager-go/web/api"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	g := e.Group("/api")
	{
		g.GET("/popular", api.FetchMostPopularVidoes())
		g.GET("/video/:id", api.GetVideo())
		g.GET("/video/:id/related", api.FetchRelatedVideos())
		g.GET("/search", api.SearchVideos())
	}

	fg := g.Group("/favorites", middlewares.FirebaseGuard())
	{
		fg.POST("/:id/toggle", api.ToggleFavoriteVideo())
	}
}

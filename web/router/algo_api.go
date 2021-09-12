package router

import (
	"github.com/Juminiy/my_go_lib/web/api_context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func ApiVersion_1(apiBase fiber.Router) {}
func ApiVersion_2(apiBase fiber.Router) {}
func AlgoBase(apiVersion fiber.Router)  {}
func ApiBase(app *fiber.App) {
	app.Get("/", api_context.OKContext)
	app.Get("/favicon", proxy.Forward("https://s.hualingnan.site/img/git-lookup.jpeg"))

	apiBase := app.Group("/api")
	apiVersion := apiBase.Group("/v1", func(ctx *fiber.Ctx) error {
		ctx.Set("Version", "v1")
		return ctx.Next()
	})
	apiVersion.Get("/info", api_context.ApiInfoContext)
	algo := apiVersion.Group("/algo")
	algoBase := algo.Group("/base")
	algoBase.Get("/intMin/:compValue/:compedValue", api_context.IntMinContext)
	algoBase.Post("/intQuickSort", api_context.QuickSortContext)
	algoBase.Post("/intMergeSort", api_context.MergeSortContext)
	algoBase.Post("/epsilonClosure", api_context.EpsilonClosureContext)
	// 404 Handler
	app.Use(api_context.NotFoundContext)
}

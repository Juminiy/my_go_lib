package api_context

import (
	"os"

	"github.com/Juminiy/my_go_lib/my_algo/algo_compile/input_struct"
	"github.com/Juminiy/my_go_lib/web/model"
	"github.com/Juminiy/my_go_lib/web/service"
	"github.com/Juminiy/my_go_lib/web/utils"
	"github.com/gofiber/fiber/v2"
)

var format utils.JsonFormat

func OKContext(ctx *fiber.Ctx) error {
	return ctx.JSON(format.Ok("Authored by Alan 2084TEAM"))
}
func NotFoundContext(ctx *fiber.Ctx) error {
	return ctx.JSON(format.Error(nil))
}

func ApiInfoContext(ctx *fiber.Ctx) error {
	return ctx.JSON(format.Ok(fiber.Map{
		"version":          ctx.Get("Version"),
		"gitRepositoryUrl": os.Getenv("gitRepositoryUrl"),
		"description":      "Algorithm Api by Golang code",
		"author":           "Alan Juminiy",
	}))
}

func IntMinContext(ctx *fiber.Ctx) error {
	return ctx.JSON(format.Ok(service.IntMinAlgoService(ctx.Params("compValue"), ctx.Params("compedValue"))))
}
func IntMaxContext(ctx *fiber.Ctx) error {
	return ctx.JSON(format.Ok(service.IntMaxAlgoService(ctx.Params("compValue"), ctx.Params("compedValue"))))
}
func QuickSortContext(ctx *fiber.Ctx) error {
	arr := new(model.Arr)
	if err := ctx.BodyParser(arr); err != nil {
		return err
	}
	return ctx.JSON(format.Ok(service.IntQuickSortService(arr.TArr)))
}
func MergeSortContext(ctx *fiber.Ctx) error {
	arr := new(model.Arr)
	if err := ctx.BodyParser(arr); err != nil {
		return nil
	}
	return ctx.JSON(format.Ok(service.IntMergeSortService(arr.TArr)))
}
func EpsilonClosureContext(ctx *fiber.Ctx) error {
	inputData := new(input_struct.ValuesInput)
	if err := ctx.BodyParser(inputData); err != nil {
		return nil
	}
	return ctx.JSON(format.Ok(service.EpsilonClosureService((*inputData).Edges, (*inputData).Nodes)))
}
func ConstructSubSetsContext(ctx *fiber.Ctx) error {
	inputData := new(input_struct.ValuesInput)
	if err := ctx.BodyParser(inputData); err != nil {
		return nil
	}
	return ctx.JSON(format.Ok(service.ConstructSubSetsService((*inputData).Edges, (*inputData).Nodes)))
}
func NFAToDFAContext(ctx *fiber.Ctx) error {
	inputData := new(input_struct.ValuesInput)
	if err := ctx.BodyParser(inputData); err != nil {
		return nil
	}
	return ctx.JSON(format.Ok(service.NFAToDFAService((*inputData).Edges, (*inputData).Nodes)))
}
func MinDFAContext(ctx *fiber.Ctx) error {
	inputData := new(input_struct.ValuesInput)
	if err := ctx.BodyParser(inputData); err != nil {
		return nil
	}
	return ctx.JSON(format.Ok(service.MinDFAService((*inputData).Edges, (*inputData).Nodes)))
}

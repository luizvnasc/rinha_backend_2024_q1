package extrato

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func RegistraHandlers(app *fiber.App) {
	app.Get("/clientes/:id/extrato", getextrato)
}

func getextrato(ctx *fiber.Ctx) error {
	clienteID := ctx.Params("id")
	id, err := strconv.ParseUint(clienteID, 10, 8)
	if err != nil {
		return ctx.SendStatus(404)
	}
	extrato, err := store.getExtrato(uint(id))
	if err != nil {
		return ctx.SendStatus(404)
	}
	return ctx.JSON(extrato)
}

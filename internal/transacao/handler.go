package transacao

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func RegistraHandlers(app *fiber.App) {
	app.Post("/clientes/:id/transacoes", criaTransacao)
}

func criaTransacao(ctx *fiber.Ctx) error {
	clienteID := ctx.Params("id")
	transacao := new(Transacao)

	if err := ctx.BodyParser(transacao); err != nil {
		return err
	}
	id, err := strconv.ParseUint(clienteID, 10, 8)
	if err != nil {
		return ctx.SendStatus(500)
	}
	transacao.ClienteID = uint(id)
	if err = store.CriaTransacao(transacao); err != nil {
		return ctx.SendStatus(404)
	}
	return ctx.SendStatus(200)
}

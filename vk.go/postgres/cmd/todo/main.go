package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"log"

	_ "github.com/jackc/pgx/v5"
)

func createHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Create")
}

func readHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Read")
}

func updateHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Update")
}

func deleteHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Delete")
}

func main() {

	dbConnectionStr := "postgresql://postgres:LemmeIn@127.0.0.1:5432/todo?sslmode=disable"

	conn, err := pgx.Connect(context.Background(), dbConnectionStr)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		_ = conn.Close(ctx)
	}(conn, context.Background())

	/*
		app := fiber.New()

		// TODO: config

		app.Get("/", readHandler)
		app.Post("/", createHandler)
		app.Put("/update", updateHandler)
		app.Delete("/delete", deleteHandler)

		port := ":3239"
		err = app.Listen(port)

		if err != nil {
			log.Fatalln(err)
		}
	*/
}

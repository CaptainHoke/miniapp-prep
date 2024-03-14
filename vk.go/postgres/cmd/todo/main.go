package main

// This garbage was written while following the worst fucking tutorial in existence

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/jackc/pgx/v5"
	"log"

	_ "github.com/jackc/pgx/v5"
)

type Todo struct {
	Item string
}

func createHandler(ctx *fiber.Ctx, conn *pgx.Conn) error {
	return ctx.SendString("Create")
}

func readHandler(ctx *fiber.Ctx, conn *pgx.Conn) error {
	var todos []string

	rows, err := conn.Query(context.Background(), "select * from items")
	defer rows.Close()

	if err != nil {
		log.Fatalln(err)
	}

	var row string
	for rows.Next() {
		err := rows.Scan(&row)
		log.Println(row)
		if err != nil {
			return err
		}
		todos = append(todos, row)
	}

	return ctx.Render("index", fiber.Map{"Todos": todos})
}

func updateHandler(ctx *fiber.Ctx, conn *pgx.Conn) error {
	return ctx.SendString("Update")
}

func deleteHandler(ctx *fiber.Ctx, conn *pgx.Conn) error {
	return ctx.SendString("Delete")
}

func main() {

	dbConnectionStr := "postgresql://postgres:LemmeIn@127.0.0.1:5432/todo"

	conn, err := pgx.Connect(context.Background(), dbConnectionStr)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		_ = conn.Close(ctx)
	}(conn, context.Background())

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	// TODO: config

	app.Get("/", func(c *fiber.Ctx) error {
		return readHandler(c, conn)
	})
	app.Post("/", func(c *fiber.Ctx) error {
		return createHandler(c, conn)
	})
	app.Put("/update", func(c *fiber.Ctx) error {
		return updateHandler(c, conn)
	})
	app.Delete("/delete", func(c *fiber.Ctx) error {
		return deleteHandler(c, conn)
	})

	app.Static("/", "./")

	port := ":3239"
	err = app.Listen(port)

	if err != nil {
		log.Fatalln(err)
	}
}

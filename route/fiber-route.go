package route

import (
	"github.com/JuliusIvander/case-02/entity"
	"github.com/JuliusIvander/case-02/repository"
	"github.com/JuliusIvander/case-02/service"
	"github.com/gofiber/fiber/v2"
)

type (
	// UserRoute struct
	userApp struct {
		service service.UserService
	}
)

func (m *userApp) getHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := m.service.GetUser(id)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	return c.Status(200).JSON(result)
}

func (m *userApp) addHandler(c *fiber.Ctx) error {
	u := new(entity.User)

	if err := c.BodyParser(u); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err := m.service.AddUser(u)
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.Status(201).JSON(&fiber.Map{
		"ok":      true,
		"message": "Data telah berhasil disimpan",
	})
}

// NewUserRoute function
func NewUserRoute() *fiber.App {
	route := fiber.New()
	userRepo := repository.NewMySQL("root:@tcp(127.0.0.1:3306)/case02")
	userService := service.NewUserService(userRepo)
	app := userApp{userService}

	route.Post("/add", app.addHandler)
	route.Get("/get/:id", app.getHandler)
	return route
}

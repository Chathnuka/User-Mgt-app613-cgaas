package apiHandlers

import (
	"User-Mgt/api"
	"strings"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Router(app *fiber.App) {
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	group := app.Group("/User-Mgt/api")
	defaultGroup := app.Group("/")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Static("/", "./docs/rapiDoc/build")
	DefaultMappings(defaultGroup)
	RouteMappings(group)
}

func RouteMappings(cg fiber.Router) {
cg.Post("/CreateUser", api.CreateUserApi)
cg.Put("/UpdateUser", api.UpdateUserApi)
cg.Delete("/DeleteUser", api.DeleteUserApi)
cg.Get("/FindUser", api.FindUserApi)
cg.Get("/FindallUser", api.FindallUserApi)
cg.Post("/UploadUser", api.UploadUserApi)
cg.Get("/DownloadUser", api.DownloadUserApi)
cg.Post("/CreateRole", api.CreateRoleApi)
cg.Put("/UpdateRole", api.UpdateRoleApi)
cg.Delete("/DeleteRole", api.DeleteRoleApi)
cg.Get("/FindRole", api.FindRoleApi)
cg.Get("/FindallRole", api.FindallRoleApi)
cg.Post("/UploadRole", api.UploadRoleApi)
cg.Get("/DownloadRole", api.DownloadRoleApi)
cg.Get("/UsermgtGetconfigRole", api.UsermgtGetconfigRoleApi)

cg.Get("/swagger", api.SwaggerHandler)

}

func DefaultMappings(cg fiber.Router) {
	cg.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "User-Mgt-APP613 service is up and running", "version": "1.0"})
	})
}
package main

import "github.com/gofiber/fiber"


func main(){
	app := fiber.New();

	app.Get("/foo", func (c *fiber.Ctx){
		 c.JSON(map[string]string{"msg":"ALl Ok"})
	})

	app.Listen(":5000")
}
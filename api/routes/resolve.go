/*
after creating the short url, we also need to redirect the user to the original site
so this class resolves the short url to the original URL
*/
package routes

import(
	"url-shortener/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"

)

func ResolveUrl(c *fiber.Ctx) error{
	//get req comes from the main file and then the url variable is created and value from url is put 
	//in it
	url := c.Params("url")

	r := database.CreateClient(0)
	//defer basically that execute it at the last of call stack
	defer r.Close()

	//running functions on db
	//redis is key value db, for every url(key) we get the information
	value, err := r.Get(database.Ctx, url).Result()
	//if we get value as nil, it means it does not exist in db
	if err == redis.Nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":"short not found in DB!"})
	}	else if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"can not connect to DB",})
	}

	//else we create a clien and then redirect it to the main link
	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	return c.Redirect(value, 301)
}

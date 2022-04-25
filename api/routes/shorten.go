package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

//custom data-type
//defining request and response as structs which will give it structure and defined properly which
//can be easily used in front-end
type request struct{
	//json:name is basically telling Go that when json format comes and the field is url
	//we convert/assign it to the URl field in the struct
	URL				string			`json:"url"`
	CustomShort		string			`json:"short"`
	Expiry			time.Duration	`json:expiry`
}

type response struct{
	URL					string				`json:"url"`
	CustomShort			string				`json:"short"`
	//below are added so the user cant make unlimited number of requests
	expiry				time.Duration		`json:"expiry"`
	XRateRemaining		int					`json:"rate_limit"`
	XRateLimitReset		time.Duration		`json:"rate_limit_reset"`
}

//creating the shorte function
//this funciton returns error in case something goes wrong
func ShortenUrl(c *fiber.Ctx) error{
	body := new(request)

	//body parser basically takes the json request and converts it into go struct which can be understood
	//by golang and operations can be performed on it
	if err := c.BodyParser(&body); err!=nil{
		//if there is error we return the fiber response as bad request and data in form of json
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"can not parse JSON"})
	}

	//implement rate limiting

	//checking if input is actual URl
	if !govalidator.isURL(body.URL){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"invalid URL"})
	}

	//checking for domain error
	if !helpers.RemoveDomainError(body.URL){
		c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map)
	}

	//enforce HTTP and SSL
	body.URL = helpers.EnforceHTTP(body.URL)
}
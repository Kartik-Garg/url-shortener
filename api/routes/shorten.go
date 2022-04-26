package routes

import (
	"os"
	"strconv"
	"time"
	"url-shortener/database"
	"url-shortner/helpers"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/internal/uuid"
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
	//creating new client
	r2 := database.CreateClient(1)
	defer r2.Close()
	//redis has get and set methods and we know it is key value pair db
	//so we get here and then check the value
	val, err := r2.Get(database.Ctx, c.IP()).Result()
	//if we do not find any value in the redis db
	if err == redis.Nil{
		//30*60*time.second is basically the time as 30 mins before it gets reset and eror in case
		//there are any error
		_ =r2.Set(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else{
		val, _ := r2.Get(database.Ctx, c.IP()).Result()
		//passing ip as key, getting the quota and then converting it to int and storing it in valInt
		valInt, _ := strconv.Atoi(val)
		if(valInt <= 0){
			limit, _ := r2.TTL(database.Ctx, c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error":"rate limit has exceeded",
				"rate_limit_reset": limit / time.Nanosecond / time.Minute,
			})
		}
	}


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

	//creating custom short url by user
	var id string
	if body.CustomShort ==""{
		id = uuid.New().String()[:6]
	} else{
		id = body.CustomShort
	}

	r := database.CreateClient(0)
	defer r.Close()

	val, _ = r.Get(database.Ctx, id).Result()
	if val != ""{
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":"URL custom short is already used",
		})
	} 

	//checking expiry
	if body.Expiry == 0 {
		//setting expiry for the shortened url
		body.Expiry = 24
	}

	//now we have to set the new url in our db
	/*
		for this particular id(key), in db we set the original url and the expiration of the
		newly created shorten url
		.Err is if we expect an error, we can take it in variable and then use it
	*/
	err = r.Set(database.Ctx, id, body.URL, body.Expiry*3600*time.Second).Err()
	//if err exists
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":"unable to connect db server",
		})
	}


	//for decrementing it, at last so that after its run on top, decremenet happens at last
	r2.Decr(database.Ctx, c.IP())
}
package routes

import (
	"time")

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
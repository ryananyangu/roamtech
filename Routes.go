package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ryananyangu/roamtech/controllers"
)

var basePath = "/api/v1"

// Routes Function to route mapping
var Routes = map[string]map[string]fiber.Handler{
	basePath + "/service/monitor": {
		"GET": controllers.ServiceMonitor,
	},
	basePath + "/mcc-mnc/scrapper": {
		"GET": controllers.MccMncScrap,
	},
	basePath + "/lookup/mcc-mnc": {
		"GET": controllers.Network,
	},
	basePath + "/lookup/country/networks": {
		"GET": controllers.CountryNetworks,
	},
}

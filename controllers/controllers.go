package controllers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ryananyangu/roamtech/database"
	"github.com/ryananyangu/roamtech/services"
)

func ServiceMonitor(ctx *fiber.Ctx) error {

	timestamp := time.Now().Unix()
	return ctx.Status(http.StatusOK).JSON(map[string]interface{}{
		"Timestamp": timestamp,
	})

}

func MccMncScrap(ctx *fiber.Ctx) error {

	list, err := services.ProcessMccMnc()

	if err != nil {
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	return ctx.Status(http.StatusOK).JSON(list)

}

func Network(ctx *fiber.Ctx) error {

	// query := map[string]string{}

	mcc := ctx.Query("mcc", "")
	mnc := ctx.Query("mnc", "")

	// ctx.

	networkcode := mcc + "-" + mnc

	if mcc == "" || mnc == "" {
		return ctx.Status(http.StatusBadRequest).SendString("Query string has to have both mnc & mcc")
	}

	db := *database.GetDB()
	network, found := db[networkcode]

	if !found {
		return ctx.Status(http.StatusNotFound).SendString("No matching info for provided mnc : " + mnc + " & mcc : " + mcc)

	}
	return ctx.Status(http.StatusOK).JSON(network)

}

func CountryNetworks(ctx *fiber.Ctx) error {

	country := ctx.Query("country", "")
	mcc := ctx.Query("mcc", "")

	if country != "" {

		return ctx.Status(http.StatusOK).JSON(services.GetByCountry(country))

	} else if mcc != "" {
		return ctx.Status(http.StatusOK).JSON(services.GetByMcc(mcc))
	} else {
		return ctx.Status(http.StatusBadRequest).SendString("Query string has to either have mcc code or country name")
	}

}

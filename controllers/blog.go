package controllers

import (
	"awesomeProject2/database"
	"awesomeProject2/model"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func BlogList(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "ok",
		"message":    "blog listed",
	}
	time.Sleep(time.Millisecond * 100)
	db := database.DBConn
	var records []model.Blog
	db.Find(&records)
	context["blog_records"] = records

	c.Status(200)
	return c.JSON(context)
}
func BlogRead(c *fiber.Ctx) error {
	c.Status(400)
	context := fiber.Map{
		"statusText": "",
		"message":    "",
	}
	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("record not found")
		context["msg"] = "record bulunamadı"

		c.Status(404)
		return c.JSON(context)
	}

	context["record"] = record
	context["statusText"] = "ok"
	context["msg"] = "blog detail"
	c.Status(200)
	return c.JSON(context)
}

func BlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "error",
		"msg":        "Something went wrong.",
	}

	record := new(model.Blog)

	if err := c.BodyParser(record); err != nil {
		log.Println("Error in parsing request:", err)
		context["msg"] = "Error in parsing request."
		return c.Status(fiber.StatusBadRequest).JSON(context)
	}

	result := database.DBConn.Create(record)
	if result.Error != nil {
		log.Println("Error in saving data:", result.Error)
		context["msg"] = "Error in saving data."
		return c.Status(fiber.StatusInternalServerError).JSON(context)
	}

	// Kayıt başarılı
	context["statusText"] = "ok"
	context["msg"] = "Record is saved successfully."
	context["data"] = record

	return c.Status(fiber.StatusCreated).JSON(context)
}

func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "ok",
		"message":    "blog updated",
	}
	id := c.Params("id")
	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("there is no man who has id=0.")
		context["statusText"] = ""
		context["msg"] = "record not found"
		c.Status(400)

		return c.JSON(context)
	}
	if err := c.BodyParser(&record); err != nil {
		log.Println("error in parisng reuqest")
		context["msg"] = "bir şeyler yanlış gitti..."
	}
	result := database.DBConn.Save(record)
	if result.Error != nil {
		log.Println("error in saving data")
		c.Status(400)
	}
	context["msg"] = "record updated successfully"
	context["data"] = record
	c.Status(200)
	return c.JSON(context)
}

func BlogDelete(c *fiber.Ctx) error {
	c.Status(400)
	context := fiber.Map{
		"statusText": "",
		"message":    ""}
	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("record not found")
		context["msg"] = "record bulunamadı"

		return c.JSON(context)
	}
	result := database.DBConn.Delete(record)
	if result.Error != nil {
		context["msg"] = "record bulunamadı"

		return c.JSON(context)
	}
	context["statusText"] = "ok"
	context["msg"] = "record deleted succesfully"
	c.Status(200)
	return c.JSON(context)
}

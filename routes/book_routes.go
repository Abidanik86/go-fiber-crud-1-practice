package routes

import (
	"go-fiber-crud/database"
	"go-fiber-crud/models"

	"github.com/gofiber/fiber/v2"
)

// সমস্ত বই লিস্ট করুন
func GetBooks(c *fiber.Ctx) error {
	var books []models.Book
	database.DB.Find(&books)
	return c.JSON(books)
}

// নির্দিষ্ট বই খুঁজে বের করুন
func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	if err := database.DB.First(&book, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "বই পাওয়া যায়নি!"})
	}

	return c.JSON(book)
}

// নতুন বই তৈরি করুন
func CreateBook(c *fiber.Ctx) error {
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ডাটা পার্স করা যায়নি!"})
	}

	database.DB.Create(&book)
	return c.Status(201).JSON(book)
}

// বই আপডেট করুন
func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	if err := database.DB.First(&book, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "বই পাওয়া যায়নি!"})
	}

	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ডাটা পার্স করা যায়নি!"})
	}

	database.DB.Save(&book)
	return c.JSON(book)
}

// বই ডিলিট করুন
func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	if err := database.DB.First(&book, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "বই পাওয়া যায়নি!"})
	}

	database.DB.Delete(&book)
	return c.JSON(fiber.Map{"message": "বই সফলভাবে ডিলিট করা হয়েছে!"})
}

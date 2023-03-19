package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/hzlnqodrey/golang-fiber-postgre-gorm/models"
	"github.com/hzlnqodrey/golang-fiber-postgre-gorm/storage"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// Book Models
type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

// Repo Struct
type Repository struct {
	DB *gorm.DB
}

// Book Controller
// Create BOok Controller
func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}

	err := context.BodyParser(&book)

	if err != nil {
		context.Status(fiber.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"message": "request failed",
			},
		)
		return err
	}

	err = r.DB.Create(&book).Error

	if err != nil {
		context.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"message": "Could not create book",
			},
		)
		return err
	}

	context.Status(fiber.StatusOK).JSON(
		&fiber.Map{
			"message": "Book has been added",
		},
	)

	return nil
}

func (r *Repository) DeleteBook(c *fiber.Ctx) error {
	booksModels := models.Books{}
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "id cannot be empty",
			})

		return nil
	}

	err := r.DB.Delete(booksModels, id)

	if err.Error != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"message": "could not delete book",
			})

		return err.Error
	}

	// "successfull message"
	c.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "books deleted successfully",
		},
	)

	return nil
}

// get books by id
func (r *Repository) GetBookByID(c *fiber.Ctx) error {
	booksModels := models.Books{}
	id := c.Params("id")

	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "id cannot be empty",
			},
		)

		return nil
	}

	fmt.Println("The ID is ", id)

	err := r.DB.Where("id = ?", id).First(booksModels).Error

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"message": "could not get the books ",
			},
		)

		return err
	}

	// everything when well
	c.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "book id fetched successfully/",
			"data":    booksModels,
		},
	)

	return nil
}

// Get Books Controller
func (r *Repository) GetBooks(c *fiber.Ctx) error {
	booksModels := &[]models.Books{}

	err := r.DB.Find(booksModels).Error

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"message": "Could not get books",
			},
		)
		return err
	}

	c.Status(fiber.StatusOK).JSON(
		&fiber.Map{
			"message": "Book fetched successfully",
			"data":    booksModels,
		},
	)

	return nil
}

// setup router method
func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Get("/get_book/:id", r.GetBookByID)
	api.Get("/books", r.GetBooks)
}

func main() {
	// LOAD ENV VAR
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBname:   os.Getenv("DB_NAME"),
	}

	// DB Config
	db, err := storage.NewConnection(config)

	if err != nil {
		panic("Could not load to database")
	}

	err = models.MigrateBooks(db)

	if err != nil {
		panic("could not migrate db")
	}

	// Fiber and App routing
	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hzlnqodrey/golang-fiber-postgre-gorm/models"
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

	// DB Config
	db, err := storage.NewConnection(config)

	if err != nil {
		panic("Could not load to database")
	}

	// Fiber and App routing
	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}

package controllers

import (
	"strconv"

	"github.com/cocoasterr/kalbe_test/models"
	"github.com/cocoasterr/kalbe_test/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ProductController struct {
	Service services.ProductService
}

func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{
		Service: service,
	}
}

func (c *ProductController) CreateController(ctx *fiber.Ctx) error {
	userName := ctx.Locals("userName").(string)

	var ProductInput models.ProductCreate
	if err := ctx.BodyParser(&ProductInput); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
		return err
	}

	payload := map[string]interface{}{
		"name":  ProductInput.Name,
		"harga": ProductInput.Harga,
		"stock": ProductInput.Stock,
	}

	trx, err := c.Service.Repo.BeginTransaction()
	if err != nil {
		logrus.WithError(err).Error("Error beginning transaction")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	err = c.Service.CreateService(ctx.Context(), trx, payload, userName)
	if err != nil {
		logrus.WithError(err).Error("Error creating Product")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	if commitErr := c.Service.Repo.CommitTransaction(trx); commitErr != nil {
		logrus.WithError(commitErr).Errorf("Error committing transaction")
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Product created successfully"})
}

func (c *ProductController) IndexController(ctx *fiber.Ctx) error {
	page, err := strconv.Atoi(ctx.Query("page", "1"))
	if err != nil || page <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid page parameter"})
	}

	limit, err := strconv.Atoi(ctx.Query("limit", "10"))
	if err != nil || limit <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid limit parameter"})
	}

	data, total, err := c.Service.IndexService(ctx.Context(), page, limit)
	if err != nil {
		logrus.WithError(err).Error("Error retrieving Product data")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": data, "total": total})
}

func (c *ProductController) FindByController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Product ID"})
	}

	data, err := c.Service.FindByService(ctx.Context(), "id", id)
	if err != nil {
		logrus.WithError(err).Error("Error retrieving Product data by ID")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": data})
}

func (c *ProductController) UpdateController(ctx *fiber.Ctx) error {
	userName := ctx.Locals("userName").(string)

	var ProductInput models.ProductCreate
	if err := ctx.BodyParser(&ProductInput); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
		return err
	}

	payload := map[string]interface{}{
		"name":  ProductInput.Name,
		"harga": ProductInput.Harga,
		"stock": ProductInput.Stock,
	}

	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Product ID"})
	}

	trx, err := c.Service.Repo.BeginTransaction()
	if err != nil {
		logrus.WithError(err).Error("Error beginning transaction")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	err = c.Service.UpdateService(ctx.Context(), trx, payload, id, userName)
	if err != nil {
		logrus.WithError(err).Error("Error updating Product")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if commitErr := c.Service.Repo.CommitTransaction(trx); commitErr != nil {
		logrus.WithError(commitErr).Errorf("Error committing transaction")
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Product updated successfully"})
}

func (c *ProductController) DeleteController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid attendance ID"})
	}
	trx, err := c.Service.Repo.BeginTransaction()
	if err != nil {
		logrus.WithError(err).Error("Error beginning transaction")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	err = c.Service.SoftDelService(ctx.Context(), trx, id)
	if err != nil {
		logrus.WithError(err).Error("Error deleting Product")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	if commitErr := c.Service.Repo.CommitTransaction(trx); commitErr != nil {
		logrus.WithError(commitErr).Errorf("Error committing transaction")
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Product deleted successfully"})
}

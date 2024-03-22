package controllers

import (
	"strconv"

	"github.com/cocoasterr/kalbe_test/models"
	"github.com/cocoasterr/kalbe_test/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type PesananController struct {
	Service services.PesananService
}

func NewPesananController(service services.PesananService) *PesananController {
	return &PesananController{
		Service: service,
	}
}

func (c *PesananController) CreateController(ctx *fiber.Ctx) error {
	userName := ctx.Locals("userName").(string)
	userID := ctx.Locals("userID").(string)

	var PesananInput models.PesananCreate
	PesananInput.UserId = userID
	if err := ctx.BodyParser(&PesananInput); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
		return err
	}

	data, err := c.Service.FindProduct(ctx.Context(), PesananInput.ProductId)
	if err != nil {
		logrus.WithError(err).Error("Error retrieving Pesanan data by ID")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	} else if data == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found!"})
	}

	if data["stock"].(int64) < int64(PesananInput.Qty) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Stock Product is not enough!"})
	}

	payload := map[string]interface{}{
		"user_id":    PesananInput.UserId,
		"status":     PesananInput.Status,
		"product_id": PesananInput.ProductId,
		"qty":        PesananInput.Qty,
	}

	payloadProduct := map[string]interface{}{
		"stock": data["stock"].(int64) - int64(PesananInput.Qty),
	}

	trx, err := c.Service.Repo.BeginTransaction()
	if err != nil {
		logrus.WithError(err).Error("Error beginning transaction")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	// if true or (paid) it will update product stock
	if payload["status"] == true {
		if err = c.Service.ProductService.UpdateService(ctx.Context(), trx, payloadProduct, payload["product_id"].(string), userName); err != nil {
			logrus.WithError(err).Error("Error updating Product")
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	if err = c.Service.CreateService(ctx.Context(), trx, payload, userName); err != nil {
		logrus.WithError(err).Error("Error creating Pesanan")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if commitErr := c.Service.Repo.CommitTransaction(trx); commitErr != nil {
		logrus.WithError(commitErr).Errorf("Error committing transaction")
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Pesanan created successfully"})
}

func (c *PesananController) IndexController(ctx *fiber.Ctx) error {
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
		logrus.WithError(err).Error("Error retrieving Pesanan data")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": data, "total": total})
}

func (c *PesananController) FindByController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Pesanan ID"})
	}

	data, err := c.Service.FindByService(ctx.Context(), "id", id)
	if err != nil {
		logrus.WithError(err).Error("Error retrieving Pesanan data by ID")
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": data})
}

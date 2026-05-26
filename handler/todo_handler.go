package handler

import (
	"golang-blog-api/helper"
	"golang-blog-api/model/web"
	"golang-blog-api/service"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type TodoHandler struct {
	Service service.TodoService
}

func NewTodoHandler(service service.TodoService) *TodoHandler {
	return &TodoHandler{
		Service: service,
	}
}

func (h *TodoHandler) Create(c fiber.Ctx) error {
	var request web.TodoCreateRequest

	err := c.Bind().Body(&request)
	if err != nil {
		return err
	}

	err = helper.Validate.Struct(request)
	if err != nil {
		return c.Status(400).JSON(helper.WebResponse{
			StatusCode: 400,
			Status:     "Error",
			Message:    err.Error(),
		})
	}

	userIdString := c.Locals("userId").(string)

	userId, err := uuid.Parse(userIdString)
	if err != nil {
		return err
	}

	result, err := h.Service.Create(c.Context(), userId, request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(helper.WebResponse{
		StatusCode: fiber.StatusCreated,
		Status:     "Success",
		Message:    "Success created todo",
		Data:       result,
	})
}
func (h *TodoHandler) Update(c fiber.Ctx) error {
	var request web.TodoUpdateRequest

	err := c.Bind().Body(&request)
	if err != nil {
		return err
	}

	err = helper.Validate.Struct(request)
	if err != nil {
		return c.Status(400).JSON(helper.WebResponse{
			StatusCode: 400,
			Status:     "Error",
			Message:    err.Error(),
		})
	}

	todoIdString := c.Params("id")

	id, err := uuid.Parse(todoIdString)
	if err != nil {
		return err
	}

	userIdString := c.Locals("userId").(string)

	userId, err := uuid.Parse(userIdString)
	if err != nil {
		return err
	}

	result, err := h.Service.Update(c.Context(), id, userId, request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(helper.WebResponse{
		StatusCode: fiber.StatusOK,
		Status:     "Success",
		Message:    "Success update todo",
		Data:       result,
	})
}
func (h *TodoHandler) Delete(c fiber.Ctx) error {
	todoIdString := c.Params("id")

	id, err := uuid.Parse(todoIdString)
	if err != nil {
		return err
	}

	userIdString := c.Locals("userId").(string)

	userId, err := uuid.Parse(userIdString)
	if err != nil {
		return err
	}

	err = h.Service.Delete(c.Context(), id, userId)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(helper.WebResponse{
		StatusCode: fiber.StatusOK,
		Status:     "Success",
		Message:    "Success delete todo",
	})
}
func (h *TodoHandler) FindByID(c fiber.Ctx) error {
	todoIdString := c.Params("id")

	id, err := uuid.Parse(todoIdString)
	if err != nil {
		return err
	}

	userIdString := c.Locals("userId").(string)

	userId, err := uuid.Parse(userIdString)
	if err != nil {
		return err
	}

	result, err := h.Service.FindByID(c.Context(), id, userId)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(helper.WebResponse{
		StatusCode: fiber.StatusOK,
		Status:     "Success",
		Message:    "Success get todo",
		Data:       result,
	})
}
func (h *TodoHandler) FindByUserID(c fiber.Ctx) error {
	userIdString := c.Locals("userId").(string)

	userId, err := uuid.Parse(userIdString)
	if err != nil {
		return err
	}

	result, err := h.Service.FindByUserID(c.Context(), userId)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(helper.WebResponse{
		StatusCode: fiber.StatusOK,
		Status:     "Success",
		Message:    "Success get todos",
		Data:       result,
	})
}

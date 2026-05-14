package handler

import (
	"fmt"
	"golang-blog-api/helper"
	"golang-blog-api/model/web"
	"golang-blog-api/service"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) Register(c fiber.Ctx) error {
	var request web.UserCreateRequest
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

	result, err := h.Service.Register(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(helper.WebResponse{
		StatusCode: 200,
		Status:     "Success",
		Message:    "Success created user",
		Data:       result,
	})
}

func (h *UserHandler) Update(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(helper.WebResponse{
			StatusCode: 400,
			Status:     "Error",
			Message:    "Invalid uuid",
		})
	}

	var request web.UserUpdateRequest
	err = c.Bind().Body(&request)
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

	result, err := h.Service.Update(c.Context(), id, request)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(helper.WebResponse{
		StatusCode: 200,
		Status:     "Success",
		Message:    "Success update user",
		Data:       result,
	})
}

func (h *UserHandler) Delete(c fiber.Ctx) error {
	id, _ := uuid.Parse(c.Params("id"))

	err := h.Service.Delete(c.Context(), id)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(helper.WebResponse{
		StatusCode: 200,
		Status:     "Success",
		Message:    "Success delete user",
	})
}

func (h *UserHandler) FindByID(c fiber.Ctx) error {
	id, _ := uuid.Parse(c.Params("id"))

	result, err := h.Service.FindByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(helper.WebResponse{
		StatusCode: 200,
		Status:     "Success",
		Message:    "Success get one user",
		Data:       result,
	})
}

func (h *UserHandler) FindAll(c fiber.Ctx) error {
	users, err := h.Service.FindAll(c.Context())
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(helper.WebResponse{
		StatusCode: 200,
		Status:     "Success",
		Message:    "Success get all users",
		Data:       users,
	})
}

func (h *UserHandler) Login(c fiber.Ctx) error {
	var request web.UserLoginRequest

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

	result, err := h.Service.Login(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(
		helper.WebResponse{
			StatusCode: fiber.StatusOK,
			Status:     "Success",
			Message:    "Login Success",
			Data:       result,
		},
	)
}

func (h *UserHandler) Me(c fiber.Ctx) error {
	userIdStr := c.Locals("userId").(string)
	fmt.Println(userIdStr)

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return c.Status(400).JSON(helper.WebResponse{
			StatusCode: 400,
			Status:     "Error",
			Message:    "Invalid user id",
		})
	}

	result, err := h.Service.FindByID(c.Context(), userId)
	if err != nil {
		return c.Status(400).JSON(helper.WebResponse{
			StatusCode: 400,
			Status:     "Error",
			Message:    "user not found 3",
		})
	}

	return c.Status(fiber.StatusOK).JSON(helper.WebResponse{
		StatusCode: 200,
		Status:     "Success",
		Message:    "Success got a user",
		Data:       result,
	})
}

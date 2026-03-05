package response

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Pagination struct {
	PreviousPage    string `json:"previousPage"`
	NextPage        string `json:"nextPage"`
	HasPreviousPage bool   `json:"hasPreviousPage"`
	HasNextPage     bool   `json:"hasNextPage"`
}

// PaginatedResponse represents paginated API response structure
type Response struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       any         `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	StatusCode int         `json:"statusCode"`
	Timestamp  string      `json:"timestamp,omitempty"`
}

// Common response builder
func buildResponse(c *fiber.Ctx, statusCode int, success bool, message string, data any) error {
	return c.Status(statusCode).JSON(Response{
		Success:    success,
		Message:    message,
		Data:       data,
		StatusCode: statusCode,
		Timestamp:  time.Now().String(),
	})
}

// Success Responses (2xx)
func Ok(c *fiber.Ctx, message string, data any) error {
	return buildResponse(c, fiber.StatusOK, true, message, data)
}

func Created(c *fiber.Ctx, message string, data any) error {
	return buildResponse(c, fiber.StatusCreated, true, message, data)
}

func Accepted(c *fiber.Ctx, message string, data any) error {
	return buildResponse(c, fiber.StatusAccepted, true, message, data)
}

func NoContent(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusNoContent).JSON(Response{
		Success:    true,
		Message:    message,
		Data:       nil,
		StatusCode: fiber.StatusNoContent,
		Timestamp:  time.Now().String(),
	})
}

// Client Error Responses (4xx)
func BadRequestException(c *fiber.Ctx, message string, data any) error {
	return buildResponse(c, fiber.StatusBadRequest, false, message, data)
}

func UnauthorizedException(c *fiber.Ctx, message string, data any) error {
	return buildResponse(c, fiber.StatusUnauthorized, false, message, data)
}

func ForbiddenException(c *fiber.Ctx, message string, data any) error {
	return buildResponse(c, fiber.StatusForbidden, false, message, data)
}

func NotFoundException(c *fiber.Ctx, message string, data any) error {
	return buildResponse(c, fiber.StatusNotFound, false, message, data)
}

func MethodNotAllowedException(c *fiber.Ctx, message string, data any) error {
	return buildResponse(c, fiber.StatusMethodNotAllowed, false, message, data)
}

func NotAcceptableException(c *fiber.Ctx, message string, data any) error {
	return buildResponse(c, fiber.StatusNotAcceptable, false, message, data)
}

func ConflictException(c *fiber.Ctx, message string, data any) error {
	return buildResponse(c, fiber.StatusConflict, false, message, data)
}

func UnprocessableEntityException(c *fiber.Ctx, message string, data any) error {
	return buildResponse(c, fiber.StatusUnprocessableEntity, false, message, data)
}

func TooManyRequestsException(c *fiber.Ctx, message string, data any) error {
	return buildResponse(c, fiber.StatusTooManyRequests, false, message, data)
}

// Server Error Responses (5xx)
func InternalServerErrorException(c *fiber.Ctx, message string, data any) error {
	return buildResponse(c, fiber.StatusInternalServerError, false, message, data)
}

func BadGatewayException(c *fiber.Ctx, message string, data any) error {
	return buildResponse(c, fiber.StatusBadGateway, false, message, data)
}

func ServiceUnavailableException(c *fiber.Ctx, message string, data any) error {
	return buildResponse(c, fiber.StatusServiceUnavailable, false, message, data)
}

func GatewayTimeoutException(c *fiber.Ctx, message string, data any) error {
	return buildResponse(c, fiber.StatusGatewayTimeout, false, message, data)
}

// Custom response with any status code
func Custom(c *fiber.Ctx, statusCode int, message string, data any) error {
	success := statusCode >= 200 && statusCode < 300
	return buildResponse(c, statusCode, success, message, data)
}

// Validation error response with detailed errors
func ValidationException(c *fiber.Ctx, message string, errors any) error {
	return c.Status(fiber.StatusBadRequest).JSON(Response{
		Success:    false,
		Message:    message,
		Data:       map[string]any{"errors": errors},
		StatusCode: fiber.StatusBadRequest,
		Timestamp:  time.Now().String(),
	})
}

// Paginated response for list endpoints
func PaginatedResponse(c *fiber.Ctx, message string, data any, pagination Pagination) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Success:    true,
		Message:    message,
		Data:       data,
		Pagination: &pagination,
		StatusCode: fiber.StatusOK,
		Timestamp:  time.Now().String(),
	})
}

// Generic error response that can infer status from error type
func ErrorResponse(c *fiber.Ctx, err error, data any) error {
	// You can extend this to handle different error types
	// and automatically determine the appropriate status code
	return buildResponse(c, fiber.StatusInternalServerError, false, err.Error(), data)
}

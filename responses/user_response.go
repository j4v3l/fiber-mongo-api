package responses

import "github.com/gofiber/fiber/v2"

// Struct tags allow us to attach meta-information to corresponding struct properties. 
// In other words, we use them to reformat the JSON response returned by the API.

type UserResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

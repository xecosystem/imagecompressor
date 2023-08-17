package main

import (
	_ "image/png"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	server := NewServer(ServerOptions{
		Path: "/",
		Handler: func(c *fiber.Ctx) error {
			imageUrl := c.Query("url")
			if imageUrl == "" {
				return c.Status(fiber.StatusBadRequest).SendString("Missing 'url' parameter")
			}

			cachedImage, ok := getImageCache(imageUrl)
			if ok {
				c.Append("Content-Type", "image/jpeg")
				return c.Send(cachedImage)
			}

			compressedImage, err := fetchImage(imageUrl)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}

			cacheImage(imageUrl, compressedImage)
			c.Append("Content-Type", "image/jpeg")
			return c.Send(compressedImage)
		},
	})

	log.Fatal(server.Listen(":8080"))
}

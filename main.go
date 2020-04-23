package main

import (
	"errors"
	"github.com/labstack/echo"
	"github.com/motoki317/wynn-guild-banner/api"
	"github.com/motoki317/wynn-guild-banner/banner"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// mkdir images
	_ = os.Mkdir("./images", os.ModePerm)

	e := echo.New()

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 3000
		log.Println("Failed to parse PORT env var, setting to default 3000")
	}

	e.GET("/", handleRoot)
	e.GET("/banners/:guildName", handleBannerGenerate)

	err = e.Start(":" + strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
}

func handleRoot(c echo.Context) error {
	return c.String(http.StatusOK, "Wynncraft Guild Banner Generator\n")
}

func handleBannerGenerate(c echo.Context) error {
	guildName := c.Param("guildName")
	if guildName == "" {
		return c.String(http.StatusBadRequest, "Specify a guild name")
	}

	imagePath := "./images/" + guildName + ".png"
	if _, err := os.Stat(imagePath); err == nil {
		// File exists
		return c.File(imagePath)
	} else if !os.IsNotExist(err) {
		// Something went wrong
		log.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	stats, err := api.GetGuildStats(guildName)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "Non existent guild?")
	}

	if stats.Name == "" {
		return c.String(http.StatusBadRequest, "Non existent guild")
	}

	err = generateBanner(stats)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.File(imagePath)
}

func generateBanner(stats *api.Guild) error {
	if stats.Banner.Base == "" {
		return errors.New("Guild " + stats.Name + " doesn't have a banner")
	}

	baseCol, ok := banner.GetColor(stats.Banner.Base)
	if !ok {
		return errors.New("Failed to get color for " + stats.Banner.Base)
	}
	layers := make([]banner.Layer, 0, len(stats.Banner.Layers))
	for _, l := range stats.Banner.Layers {
		c, ok := banner.GetColor(l.Colour)
		if !ok {
			return errors.New("Failed to get color for " + l.Colour)
		}
		layers = append(layers, banner.Layer{
			Color:   c,
			Pattern: l.Pattern,
		})
	}

	b := banner.NewBanner(baseCol, layers)
	b.Draw()
	err := b.Output("./images/" + stats.Name + ".png")
	if err != nil {
		return err
	}
	return nil
}

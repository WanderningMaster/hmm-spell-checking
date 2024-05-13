package cmd

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/WanderningMaster/hmm-spell-checking/services"
	"github.com/WanderningMaster/hmm-spell-checking/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type SpellCheckResponse struct {
	Corrections []services.Candidate `json:"corrections"`
	TotalErrors int                  `json:"totalErrors"`
}
type SpellCheckRequest struct {
	Text string `json:"text"`
}

func StartServer() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	spellChecker := services.NewSpellChecker(10)

	e.POST("api/spell-check", func(c echo.Context) error {
		req := SpellCheckRequest{}
		err := json.NewDecoder(c.Request().Body).Decode(&req)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
		}

		candidates, totalErrors, err := spellChecker.CorrectText(strings.ToLower(req.Text))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		res := SpellCheckResponse{
			Corrections: candidates,
			TotalErrors: totalErrors,
		}

		return c.JSON(http.StatusOK, res)
	})

	err := e.Start(":8080")
	utils.Require(err)
}

package api

import (
	"fmt"
	"html"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"gomorpher/internal/engine"
	"gomorpher/internal/exporter"
	"gomorpher/internal/models"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("hexcolor", func(fl validator.FieldLevel) bool {
		match, _ := regexp.MatchString(`^#(?:[0-9A-Fa-f]{3}){1,2}$`, fl.Field().String())
		return match
	})
}

type SCSSResponse struct {
	Scss string `json:"scss"`
}

// GenerateHandler
// @Summary      Generate SCSS animation
// @Description  Генерирует SCSS‑анимацию по ключевым кадрам
// @Tags         animation
// @Accept       json
// @Produce      json
// @Param        params  body      models.AnimationRequest  true  "Animation parameters"
// @Success      200     {object}  SCSSResponse
// @Failure      400     {object}  map[string]interface{}  "Validation error"
// @Failure      500     {object}  map[string]interface{}  "Internal error"
// @Router       /generate [post]
func GenerateHandler(c *gin.Context) {
	var req models.AnimationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	model, err := engine.Generate(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	selector := "morphing-" + req.Shape
	scss, err := exporter.ExportAdvancedSCSS(model, selector)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	format := c.DefaultQuery("format", "raw")

	switch format {
	case "json":
		c.JSON(http.StatusOK, SCSSResponse{Scss: scss})
	case "download":
		c.Header("Content-Disposition", `attachment; filename="animation.scss"`)
		c.Data(http.StatusOK, "text/x-scss; charset=utf-8", []byte(scss))
	case "html":
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(
			fmt.Sprintf("<pre>%s</pre>", html.EscapeString(scss)),
		))
	default:
		c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(scss))
	}
}

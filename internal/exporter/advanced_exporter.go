package exporter

import (
	"bytes"
	"fmt"
	"text/template"

	"gomorpher/internal/engine"
)

type scssContext struct {
	Selector   string
	Duration   float64
	Easing     string
	Iterations int
	Frames     []struct {
		Percent   float64
		Path      string
		Radius    string
		Color     string
		Filter    string
		Transform string
	}
}

var scssTmpl = template.Must(template.ParseFiles("internal/templates/style.scss.tmpl"))

func ExportAdvancedSCSS(model engine.AdvancedAnimationModel, selector string) (string, error) {
	ctx := scssContext{
		Selector:   selector,
		Duration:   model.Duration,
		Easing:     model.Easing,
		Iterations: model.Iterations,
	}

	for _, f := range model.Frames {
		ctx.Frames = append(ctx.Frames, struct {
			Percent   float64
			Path      string
			Radius    string
			Color     string
			Filter    string
			Transform string
		}{
			Percent:   f.T * 100,
			Path:      f.Path,
			Radius:    f.Radius,
			Color:     f.Color,
			Filter:    f.Filter,
			Transform: f.Transform,
		})
	}

	var buf bytes.Buffer
	if err := scssTmpl.Execute(&buf, ctx); err != nil {
		return "", fmt.Errorf("template execute: %w", err)
	}
	return buf.String(), nil
}

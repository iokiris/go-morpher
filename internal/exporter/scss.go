package exporter

import (
	"fmt"
	"gomorpher/internal/engine"
	"strings"
)

func ExportSCSS(model engine.AdvancedAnimationModel, selector string) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("$%s-duration: %.2fs;", selector, model.Duration))
	b.WriteString(fmt.Sprintf("$%s-easing: %s;", selector, model.Easing))
	b.WriteString(fmt.Sprintf("$%s-iterations: %s;",
		selector,
		iterationCount(model.Iterations),
	))

	b.WriteString(fmt.Sprintf("@mixin %s-animation() {", selector))
	b.WriteString(fmt.Sprintf("  @keyframes %s-morph {", selector))
	for _, frame := range model.Frames {
		percent := frame.T * 100.0
		b.WriteString(fmt.Sprintf("    %.2f%% {", percent))
		b.WriteString(fmt.Sprintf("      d: %s;", frame.Path))
		b.WriteString(fmt.Sprintf("      border-radius: %s;", frame.Radius))
		b.WriteString(fmt.Sprintf("      background: %s;", frame.Color))
		if frame.Filter != "" {
			b.WriteString(fmt.Sprintf("      filter: %s;", frame.Filter))
		}
		b.WriteString(fmt.Sprintf("      transform: %s;", frame.Transform))
		b.WriteString("    }")
	}
	b.WriteString("  }")
	b.WriteString(fmt.Sprintf("  animation: %s-morph $%s-duration $%s-easing $%s-iterations;",
		selector, selector, selector, selector))
	b.WriteString("}")

	b.WriteString(fmt.Sprintf(".%s {", selector))
	b.WriteString(fmt.Sprintf("  @include %s-animation();", selector))
	b.WriteString("}")

	return b.String()
}

func iterationCount(i int) string {
	if i <= 0 {
		return "infinite"
	}
	return fmt.Sprintf("%d", i)
}

package engine

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func interpolateBorderRadius(a, b string, u float64) string {
	partsA := strings.Split(a, "/")
	partsB := strings.Split(b, "/")

	interpolate := func(sa, sb string) string {
		ra := parsePercentList(sa)
		rb := parsePercentList(sb)
		n := len(ra)
		if len(rb) < n {
			n = len(rb)
		}
		out := make([]string, n)
		for i := 0; i < n; i++ {
			v := ra[i] + (rb[i]-ra[i])*u
			out[i] = fmt.Sprintf("%.2f%%", v)
		}
		return strings.Join(out, " ")
	}

	if len(partsA) == 2 && len(partsB) == 2 {
		h := interpolate(strings.TrimSpace(partsA[0]), strings.TrimSpace(partsB[0]))
		v := interpolate(strings.TrimSpace(partsA[1]), strings.TrimSpace(partsB[1]))
		return h + " / " + v
	}

	return interpolate(a, b)
}

func parsePercentList(s string) []float64 {
	fields := strings.Fields(s)
	out := make([]float64, len(fields))
	for i, f := range fields {
		f = strings.TrimSuffix(f, "%")
		if v, err := strconv.ParseFloat(f, 64); err == nil {
			out[i] = v
		} else {
			out[i] = 0
		}
	}
	return out
}

func interpolateColorHex(a, b string, u float64) string {
	ra, ga, ba := hexToRGB(a)
	rb, gb, bb := hexToRGB(b)
	r := int(math.Round(float64(ra) + (float64(rb)-float64(ra))*u))
	g := int(math.Round(float64(ga) + (float64(gb)-float64(ga))*u))
	bb2 := int(math.Round(float64(ba) + (float64(bb)-float64(ba))*u))
	return fmt.Sprintf("#%02X%02X%02X", clamp(r, 0, 255), clamp(g, 0, 255), clamp(bb2, 0, 255))
}

func hexToRGB(hex string) (r, g, b int) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) == 6 {
		if v, err := strconv.ParseUint(hex, 16, 32); err == nil {
			r = int((v >> 16) & 0xFF)
			g = int((v >> 8) & 0xFF)
			b = int(v & 0xFF)
		}
	}
	return
}

func clamp(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func interpolateFilter(a, b string, u float64) string {
	blurA, consA := parseFilter(a)
	blurB, consB := parseFilter(b)
	blur := blurA + (blurB-blurA)*u
	cons := consA + (consB-consA)*u

	parts := []string{}
	if blur > 0 {
		parts = append(parts, fmt.Sprintf("blur(%.2fpx)", blur))
	}
	if cons > 0 {
		parts = append(parts, fmt.Sprintf("contrast(%.2f%%)", cons))
	}
	return strings.Join(parts, " ")
}

func parseFilter(s string) (blur, contrast float64) {

	reBlur := regexp.MustCompile(`blur\(([\d.]+)px\)`)
	if m := reBlur.FindStringSubmatch(s); len(m) == 2 {
		blur, _ = strconv.ParseFloat(m[1], 64)
	}

	reCons := regexp.MustCompile(`contrast\(([\d.]+)%\)`)
	if m := reCons.FindStringSubmatch(s); len(m) == 2 {
		contrast, _ = strconv.ParseFloat(m[1], 64)
	}
	return
}

func parseRotateScale(s string) (rotate, scale float64) {
	reRot := regexp.MustCompile(`rotate\(\s*([-0-9.]+)deg\s*\)`)
	if m := reRot.FindStringSubmatch(s); len(m) == 2 {
		rotate, _ = strconv.ParseFloat(m[1], 64)
	}
	reSc := regexp.MustCompile(`scale\(\s*([-0-9.]+)\s*\)`)
	if m := reSc.FindStringSubmatch(s); len(m) == 2 {
		scale, _ = strconv.ParseFloat(m[1], 64)
	} else {
		scale = 1
	}
	return
}

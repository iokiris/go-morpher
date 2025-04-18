package engine

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"

	"gomorpher/internal/engine/timing"
	"gomorpher/internal/models"
)

type AdvancedAnimationModel struct {
	Duration   float64
	Easing     string
	Iterations int
	Frames     []Keyframe
}

type Keyframe struct {
	T         float64
	Path      string
	Radius    string
	Color     string
	Filter    string
	Transform string
}

func init() {

	rand.Seed(time.Now().UnixNano())
}

func Generate(req models.AnimationRequest) (AdvancedAnimationModel, error) {
	if len(req.Frames) < 2 {
		return AdvancedAnimationModel{}, fmt.Errorf("need at least 2 keyframes")
	}

	total := int(math.Ceil(req.Duration * float64(req.FPS)))

	ease := selectEasing(req.Easing)

	sort.Slice(req.Frames, func(i, j int) bool {
		return req.Frames[i].T < req.Frames[j].T
	})

	keys := make([]Keyframe, len(req.Frames))
	for i, fr := range req.Frames {
		keys[i] = Keyframe{
			T:         fr.T,
			Path:      fr.Path,
			Radius:    fr.Radius,
			Color:     fr.Color,
			Filter:    fr.Filter,
			Transform: fr.Transform,
		}
	}

	var out []Keyframe
	for i := 0; i <= total; i++ {
		tNorm := float64(i) / float64(total)

		eased := ease(tNorm)

		prev, next := findSegment(keys, tNorm)

		f := lerpFrame(prev, next, eased, req.NoiseAmp)
		f.T = tNorm
		out = append(out, f)
	}

	return AdvancedAnimationModel{
		Duration:   req.Duration,
		Easing:     req.Easing,
		Iterations: req.Iterations,
		Frames:     out,
	}, nil
}

func findSegment(keys []Keyframe, t float64) (Keyframe, Keyframe) {
	for i := 1; i < len(keys); i++ {
		if t <= keys[i].T {
			return keys[i-1], keys[i]
		}
	}

	n := len(keys)
	return keys[n-2], keys[n-1]
}

func lerpFrame(a, b Keyframe, u, noiseAmp float64) Keyframe {

	noise := (rand.Float64()*2 - 1) * noiseAmp

	return Keyframe{
		Path:      lerpPath(a.Path, b.Path, u),
		Radius:    lerpRadius(a.Radius, b.Radius, u),
		Color:     lerpColor(a.Color, b.Color, u),
		Filter:    lerpFilter(a.Filter, b.Filter, u),
		Transform: lerpTransform(a.Transform, b.Transform, u),

		T: a.T + (b.T-a.T)*(u+noise*0.01),
	}
}

func lerpPath(pa, pb string, u float64) string {
	return fmt.Sprintf("path(\"%s\")", pa) /* TODO: real morph */
}
func lerpRadius(ra, rb string, u float64) string { return interpolateBorderRadius(ra, rb, u) }
func lerpColor(ca, cb string, u float64) string  { return interpolateColorHex(ca, cb, u) }
func lerpFilter(fa, fb string, u float64) string { return interpolateFilter(fa, fb, u) }
func lerpTransform(ta, tb string, u float64) string {
	ra, sa := parseRotateScale(ta)
	rb, sb := parseRotateScale(tb)
	rot := ra + (rb-ra)*u
	sc := sa + (sb-sa)*u
	return fmt.Sprintf("rotate(%.2fdeg) scale(%.2f)", rot, sc)
}

func selectEasing(name string) func(float64) float64 {
	switch name {
	case "ease-in":
		return timing.CubicBezier(0.42, 0)
	case "ease-out":
		return timing.CubicBezier(0, 0.58)
	case "ease-in-out":
		return timing.EaseInOutQuad
	case "spring":
		return timing.Spring
	default:
		return timing.Linear
	}
}

package models

type AnimationRequest struct {
	Shape      string            `json:"shape" validate:"required,oneof=circle square blob"`
	Duration   float64           `json:"duration" validate:"required,gt=0"`
	Easing     string            `json:"easing"   validate:"required,oneof=linear ease-in ease-out ease-in-out spring"`
	Iterations int               `json:"iterations"`
	FPS        int               `json:"fps"       validate:"gt=0"`
	NoiseAmp   float64           `json:"noiseAmp"`
	Frames     []KeyframeRequest `json:"frames"   validate:"required,min=2,dive"`
}

type KeyframeRequest struct {
	T         float64 `json:"t" binding:"required" validate:"gte=0,lte=1"`
	Path      string  `json:"path" validate:"required"`
	Radius    string  `json:"radius" validate:"required"`
	Color     string  `json:"color"  validate:"required"`
	Filter    string  `json:"filter"`
	Transform string  `json:"transform" validate:"required"`
}

# 🎨 GoMorpher — Advanced SCSS Animation Generator

**GoMorpher** is a Go-based API that generates expressive and dynamic SCSS animations from structured keyframe data. You can animate shape morphing, gradients, filters, transforms, border-radius transitions, and more — all with precise frame control.

---

## 🚀 Example

**POST** `/generate`

```json
{
  "shape": "circle",
  "duration": 5.0,
  "easing": "ease-in-out",
  "iterations": 0,
  "fps": 30,
  "noiseAmp": 0.1,
  "frames": [
    {
      "t": 0.0,
      "path": "M150 0 L75 200 L225 200 Z",
      "radius": "50% 0% 50% 0%",
      "color": "linear-gradient(90deg, #ff0000, #0000ff)",
      "filter": "blur(5px)",
      "transform": "rotate(0deg) scale(1)"
    },
    {
      "t": 1.0,
      "path": "M150 0 L100 200 L200 200 Z",
      "radius": "25% 25% 75% 75%",
      "color": "#00ff00",
      "filter": "contrast(150%)",
      "transform": "rotate(180deg) scale(2)"
    }
  ]
}
```

**Response:**
```json
{
  "scss": "... $morphing-circle-duration: 4.20s;..."
}
```

---

## 🔧 Features

- 🔁 Morphing between SVG paths
- 🎨 Color transitions and gradients
- 🌈 Support for filters (blur, brightness, contrast, etc.)
- 📐 Animatable `border-radius` and `transform`
- ⚙️ Fully configurable keyframe timing and easing
- 💻 Generated output is ready to paste into any SCSS project

---

## 📦 Tech Stack

- **Backend**: Go + Gin
- **Validation**: go-playground/validator
- **Templating**: `text/template` for SCSS mixin generation
- **Frontend usage**: Inject via build pipeline

---

## 🛠 Local Development

```bash
git clone https://github.com/your-username/gomorpher.git
cd gomorpher
go run main.go
```

Visit: `http://localhost:8080/swagger/index.html` (if Swagger is set up)

---

## 📁 Project Structure

```
gomorpher/
├── cmd
│         └── server
│             └── main.go
├── docs
│         ├── docs.go
│         ├── swagger.json
│         └── swagger.yaml
├── go.mod
├── go.sum
├── internal
│         ├── api
│         │         └── handler.go
│         ├── engine
│         │         ├── core.go
│         │         ├── helpers.go
│         │         └── timing
│         │             └── easing.go
│         ├── exporter
│         │         ├── advanced_exporter.go
│         │         └── scss.go
│         ├── models
│         │         └── request.go
│         └── templates
│             └── style.scss.tmpl
├── README.md
├── tmp
│         ├── build-errors.log
│         └── main
└── web
    ├── index.html
    └── script.js

```

---




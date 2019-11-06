package main

import (
	"image/color"

	"github.com/oakmound/oak"
	"github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/entities"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/key"
	"github.com/oakmound/oak/physics"
	"github.com/oakmound/oak/render"
	"github.com/oakmound/oak/scene"
)

const (
	Ground collision.Label = 1
)

func main() {
	oak.Add(
		"Tank",
		func(string, interface{}) {
			char := entities.NewMoving(100, 100, 16, 32,
				render.NewColorBox(16, 32, color.RGBA{255, 0, 0, 255}),
				nil, 0, 0)
			render.Draw(char.R)

			char.Speed = physics.NewVector(3, 3)
			fallSpeed := .1
			char.Bind(
				func(id int, nothing interface{}) int {
					char := event.GetEntity(id).(*entities.Moving)
					if oak.IsDown(key.A) {
						char.ShiftX(-char.Speed.X())
					}
					if oak.IsDown(key.D) {
						char.ShiftX(char.Speed.X())
					}
					hit := char.HitLabel(Ground)
					if hit == nil {
						char.Delta.ShiftY(fallSpeed)
					} else {
						char.Delta.SetY(0)
					}
					char.ShiftY(char.Delta.Y())
					return 0
				}, event.Enter)

			ground := entities.NewSolid(0, 400, 500, 20,
				render.NewColorBox(500, 20, color.RGBA{0, 0, 255, 255}),
				nil, 0)
			ground.UpdateLabel(Ground)
			render.Draw(ground.R)
		},
		func() bool {
			return true
		},
		func() (string, *scene.Result) {
			return "Tank", nil
		},
	)
	oak.SetupConfig.Screen.X = 1
	oak.SetupConfig.Screen.X = 1
	oak.SetupConfig.Screen.Width = 670
	oak.SetupConfig.Screen.Height = 620
	oak.SetupConfig.LoadBuiltinCommands = true
	oak.Init("Tank")
}

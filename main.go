package main

import (
	char "github.com/M1ralai/raylib/characters"
	r "github.com/M1ralai/raylib/render"
	run "github.com/M1ralai/raylib/running"
	set "github.com/M1ralai/raylib/settings"
	"github.com/M1ralai/raylib/variables"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	run.Init()
	for variables.Runing {
		r.Render()
		char.Movement()
		Update()
	}
	run.Quit()
}
func Update() {
	variables.Runing = !rl.WindowShouldClose()
	set.Settings()
	variables.CamCharacter.Target = rl.NewVector2(float32(variables.CharacterDest.X-(variables.CharacterDest.Width/2)), float32(variables.CharacterDest.Y-(variables.CharacterDest.Height/2)))
	variables.MenuiconDest = rl.NewRectangle(variables.CharacterDest.X+400, variables.CharacterDest.Y+200, 32, 32)
}

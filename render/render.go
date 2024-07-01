package render

import (
	"github.com/M1ralai/raylib/variables"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Render() {
	rl.BeginDrawing()
	rl.BeginMode2D(variables.CamCharacter)
	DrawScene()
	rl.ClearBackground(variables.BackgroundColor)
	rl.EndDrawing()
	rl.EndMode2D()
}

func DrawScene() {
	rl.DrawTexturePro(variables.Ground, variables.GroundSrc, variables.GroundDest, rl.NewVector2(variables.GroundDest.Width, variables.GroundDest.Height), 0, rl.White)
	rl.DrawTexturePro(variables.Character_Sprite, variables.CharacterSrc, variables.CharacterDest, rl.NewVector2(variables.CharacterDest.Width, variables.CharacterSrc.Height), 0, rl.White)
	rl.DrawTexturePro(variables.Menuicon, variables.MenuiconSrc, variables.MenuiconDest, rl.NewVector2(variables.MenuiconDest.Width, variables.MenuiconDest.Height), 0, rl.White)
}

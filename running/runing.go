package runing

import (
	char "github.com/M1ralai/raylib/characters"
	"github.com/M1ralai/raylib/variables"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Quit() {
	rl.UnloadTexture(variables.Character_Sprite)
	rl.CloseWindow()
}
func Init() {
	rl.InitWindow(variables.ScreeWidth, variables.ScreenHeight, "Your Fucking Game")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
	variables.Ground = rl.LoadTexture("sprites/environment/ground.png")
	variables.GroundSrc = rl.NewRectangle(0, 0, 8, 8)
	variables.GroundDest = rl.NewRectangle(250, 300, 40, 40)
	variables.Character_Sprite = rl.LoadTexture("sprites/character/character.png")
	variables.CharacterSrc = rl.NewRectangle(0, 0, 8, 8)
	variables.CharacterDest = rl.NewRectangle(250, 320, 40, 40)
	variables.CamCharacter = rl.NewCamera2D(rl.NewVector2(float32(variables.ScreeWidth/2), float32(variables.ScreenHeight/2)), rl.NewVector2(float32(variables.CharacterDest.X-(variables.CharacterDest.Width/2)), float32(variables.CharacterDest.Y-(variables.CharacterDest.Height/2))), 0.0, 1.0)
	variables.Menuicon = rl.LoadTexture("sprites/ui/menu.png")
	variables.MenuiconDest = rl.NewRectangle(variables.CharacterDest.X+400, variables.CharacterDest.Y+200, 32, 32)
	variables.MenuiconSrc = rl.NewRectangle(0, 0, 16, 16)
}
func Imput() {
	char.Movement()
}

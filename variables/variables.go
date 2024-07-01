package variables

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	Runing                 = true
	ScreenHeight     int32 = 480
	ScreeWidth       int32 = 1000
	BackgroundColor        = rl.NewColor(28, 28, 132, 255)
	Character_Sprite rl.Texture2D
	CharacterDest    rl.Rectangle
	CharacterSrc     rl.Rectangle
	CharacterSpeed   float32 = 3
	CamCharacter     rl.Camera2D
	Ground           rl.Texture2D
	GroundSrc        rl.Rectangle
	GroundDest       rl.Rectangle
	Menuicon         rl.Texture2D
	MenuiconDest     rl.Rectangle
	MenuiconSrc      rl.Rectangle
)

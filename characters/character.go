package characters

import (
	"github.com/M1ralai/raylib/variables"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Movement() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		variables.CharacterDest.Y -= variables.CharacterSpeed
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		variables.CharacterDest.Y += variables.CharacterSpeed
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		variables.CharacterDest.X -= variables.CharacterSpeed
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		variables.CharacterDest.X += variables.CharacterSpeed
	}
}

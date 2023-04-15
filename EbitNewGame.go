package main

import (
	E "EbitNew"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	//ebiten.SetVsyncEnabled(false)
	//ebiten.SetTPS(ebiten.SyncWithFPS)
	//ebiten.SetFPSMode(ebiten.)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(&E.Game{}); err != nil {
		log.Fatal(err)
	}
}

func init() {

	var err error
	image1, _, err := ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}
	//hello.Image = &E.Image{image1}
	hello := E.NewEntity()
	hello.NewPosition(150, 150).
		NewImage(image1).
		NewWasd(7).
		NewGravity(3).NewRect(30, 30)
	print(hello.Position.X)

	dude1 := E.NewEntity()
	dude1.NewPosition(250, 300).
		NewImage(image1).
		NewFloor().NewRect(30, 30)

	hello2 := E.NewEntity()
	hello2.NewPosition(250, -250).
		NewImage(image1).
		NewRect(30, 30).
		NewGravity(1).NewFloor()

	for i := 0; i < 10; i++ {
		E.NewEntity().
			NewPosition(130*(float64(i)), 250).
			NewImage(image1).
			NewRect(130, 190).
			NewFloor()
	}

}

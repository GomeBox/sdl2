package sdl2

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/veandco/go-sdl2/sdl"
)

var _ graphics.Interface = (*GraphicsAdapter)(nil)

type GraphicsAdapter struct {
	window   *sdl.Window
	renderer *sdl.Renderer
}

func (g *GraphicsAdapter) Init() error {
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		return err
	}
	return nil
}

func (g *GraphicsAdapter) ShowWindow(windowSettings *graphics.WindowSettings) error {
	var screenPosX int32
	var screenPosY int32
	var resWidth, resHeight int
	var flags uint32

	flags |= sdl.WINDOW_OPENGL
	flags |= sdl.WINDOW_MOUSE_CAPTURE
	displayMode, err := sdl.GetDisplayMode(0, 0)
	if err != nil {
		return err
	}
	resWidth = int(displayMode.W)
	resHeight = int(displayMode.H)

	if windowSettings.Fullscreen {
		screenPosX = 0
		screenPosY = 0
		flags |= sdl.WINDOW_FULLSCREEN_DESKTOP
	} else {
		screenPosX = sdl.WINDOWPOS_UNDEFINED
		screenPosY = sdl.WINDOWPOS_UNDEFINED
		if windowSettings.Resolution.Width != 0 && windowSettings.Resolution.Height != 0 {
			resWidth = windowSettings.Resolution.Width
			resHeight = windowSettings.Resolution.Height
		}
		flags |= sdl.WINDOW_BORDERLESS
	}

	window, err := sdl.CreateWindow(
		windowSettings.Title,
		screenPosX, screenPosY,
		int32(resWidth), int32(resHeight),
		flags)
	if err != nil {
		return err
	}

	g.window = window

	renderer, err := sdl.CreateRenderer(
		window,
		-1,
		sdl.RENDERER_ACCELERATED)
	if err != nil {
		return err
	}
	g.renderer = renderer
	err = renderer.SetLogicalSize(int32(windowSettings.Resolution.Width), int32(windowSettings.Resolution.Height))
	if err != nil {
		return err
	}

	window.Show()
	renderer.Present()
	return nil
}

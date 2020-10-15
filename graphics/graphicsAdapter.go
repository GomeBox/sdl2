package graphics

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type GraphicsAdapter interface {
	graphics.TextureLoader
	graphics.WindowAdapter
	graphics.ScreenPresenter
	TextureLoader() graphics.TextureLoader
	WindowAdapter() graphics.WindowAdapter
	ScreenPresenter() graphics.ScreenPresenter
	FontLoader() graphics.FontLoader
	Init() error
}

func NewGraphicsAdapter() GraphicsAdapter {
	return new(graphicsAdapter)
}

type graphicsAdapter struct {
	window     *sdl.Window
	renderer   *sdl.Renderer
	fontLoader FontLoader
}

func (g *graphicsAdapter) Init() error {
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		return err
	}

	err = img.Init(img.INIT_JPG | img.INIT_PNG)
	if err != nil {
		return err
	}

	err = ttf.Init()
	if err != nil {
		return err
	}

	return nil
}

func (g *graphicsAdapter) TextureLoader() graphics.TextureLoader {
	return g
}

func (g *graphicsAdapter) WindowAdapter() graphics.WindowAdapter {
	return g
}

func (g *graphicsAdapter) ScreenPresenter() graphics.ScreenPresenter {
	return g
}

func (g *graphicsAdapter) FontLoader() graphics.FontLoader {
	return g.fontLoader
}

func (g *graphicsAdapter) ShowWindow(windowSettings *graphics.WindowSettings) error {
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

	//TODO: This should not be here
	g.fontLoader = NewFontLoader(renderer)

	window.Show()
	renderer.Present()
	return nil
}

func (g *graphicsAdapter) Load(fileName string) (graphics.TextureDrawer, error) {
	tex, err := img.LoadTexture(g.renderer, fileName)
	if err != nil {
		return nil, err
	}
	_, _, w, h, err := tex.Query()
	if err != nil {
		return nil, err
	}
	drawer := textureDrawer{renderer: g.renderer, sdlTexture: tex, width: int(w), height: int(h)}
	return drawer, nil
}

func (g *graphicsAdapter) Present() error {
	g.renderer.Present()
	return nil
}
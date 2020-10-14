package sdl2

import "github.com/veandco/go-sdl2/sdl"

type textureDrawer struct {
	sdlTexture    *sdl.Texture
	renderer      *sdl.Renderer
	width, height int32
}

func (drawer textureDrawer) Draw() error {
	destRect := sdl.Rect{X: 0, Y: 0, H: drawer.height, W: drawer.width}
	err := drawer.renderer.Copy(drawer.sdlTexture, nil, &destRect)
	//TODO move this to main loop
	drawer.renderer.Present()
	if err != nil {
		return err
	}
	return nil
}

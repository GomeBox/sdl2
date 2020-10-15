package mocks

import "github.com/GomeBox/gome/adapters/graphics"

type FontLoader struct {
}

func (fontLoader *FontLoader) Load(fileName string, size int) (graphics.TextCreator, error) {
	return nil, nil
}

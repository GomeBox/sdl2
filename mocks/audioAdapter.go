package mocks

import (
	"github.com/GomeBox/gome/adapters/audio"
)

type AudioAdapter struct {
	InitCnt int
}

func (adapter *AudioAdapter) Init() error {
	adapter.InitCnt++
	return nil
}

func (adapter *AudioAdapter) SongLoader() audio.SongLoader {
	return nil
}

func (adapter *AudioAdapter) SoundLoader() audio.SoundLoader {
	return nil
}

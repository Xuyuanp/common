package common

import "sync"

type MutexWrapper struct {
	sync.Mutex
}

func (wrapper *MutexWrapper) Wrap(fn func()) {
	wrapper.Lock()
	defer wrapper.Unlock()
	fn()
}

type RWMutexWrapper struct {
	sync.RWMutex
}

func (wrapper *RWMutexWrapper) Wrap(fn func()) {
	wrapper.Lock()
	defer wrapper.Unlock()
	fn()
}

func (wrapper *RWMutexWrapper) RWrap(fn func()) {
	wrapper.RLock()
	defer wrapper.RUnlock()
	fn()
}

type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (wrapper *WaitGroupWrapper) Wrap(fn func()) {
	wrapper.Add(1)
	go func() {
		defer wrapper.Done()
		fn()
	}()
}

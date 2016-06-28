// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"time"

	"github.com/concourse/atc/worker"
)

type FakeVolumeFactoryDB struct {
	GetVolumeTTLStub        func(volumeHandle string) (time.Duration, bool, error)
	getVolumeTTLMutex       sync.RWMutex
	getVolumeTTLArgsForCall []struct {
		volumeHandle string
	}
	getVolumeTTLReturns struct {
		result1 time.Duration
		result2 bool
		result3 error
	}
	ReapVolumeStub        func(handle string) error
	reapVolumeMutex       sync.RWMutex
	reapVolumeArgsForCall []struct {
		handle string
	}
	reapVolumeReturns struct {
		result1 error
	}
	SetVolumeTTLStub        func(string, time.Duration) error
	setVolumeTTLMutex       sync.RWMutex
	setVolumeTTLArgsForCall []struct {
		arg1 string
		arg2 time.Duration
	}
	setVolumeTTLReturns struct {
		result1 error
	}
	SetVolumeSizeInBytesStub        func(string, int64) error
	setVolumeSizeInBytesMutex       sync.RWMutex
	setVolumeSizeInBytesArgsForCall []struct {
		arg1 string
		arg2 int64
	}
	setVolumeSizeInBytesReturns struct {
		result1 error
	}
}

func (fake *FakeVolumeFactoryDB) GetVolumeTTL(volumeHandle string) (time.Duration, bool, error) {
	fake.getVolumeTTLMutex.Lock()
	fake.getVolumeTTLArgsForCall = append(fake.getVolumeTTLArgsForCall, struct {
		volumeHandle string
	}{volumeHandle})
	fake.getVolumeTTLMutex.Unlock()
	if fake.GetVolumeTTLStub != nil {
		return fake.GetVolumeTTLStub(volumeHandle)
	} else {
		return fake.getVolumeTTLReturns.result1, fake.getVolumeTTLReturns.result2, fake.getVolumeTTLReturns.result3
	}
}

func (fake *FakeVolumeFactoryDB) GetVolumeTTLCallCount() int {
	fake.getVolumeTTLMutex.RLock()
	defer fake.getVolumeTTLMutex.RUnlock()
	return len(fake.getVolumeTTLArgsForCall)
}

func (fake *FakeVolumeFactoryDB) GetVolumeTTLArgsForCall(i int) string {
	fake.getVolumeTTLMutex.RLock()
	defer fake.getVolumeTTLMutex.RUnlock()
	return fake.getVolumeTTLArgsForCall[i].volumeHandle
}

func (fake *FakeVolumeFactoryDB) GetVolumeTTLReturns(result1 time.Duration, result2 bool, result3 error) {
	fake.GetVolumeTTLStub = nil
	fake.getVolumeTTLReturns = struct {
		result1 time.Duration
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeFactoryDB) ReapVolume(handle string) error {
	fake.reapVolumeMutex.Lock()
	fake.reapVolumeArgsForCall = append(fake.reapVolumeArgsForCall, struct {
		handle string
	}{handle})
	fake.reapVolumeMutex.Unlock()
	if fake.ReapVolumeStub != nil {
		return fake.ReapVolumeStub(handle)
	} else {
		return fake.reapVolumeReturns.result1
	}
}

func (fake *FakeVolumeFactoryDB) ReapVolumeCallCount() int {
	fake.reapVolumeMutex.RLock()
	defer fake.reapVolumeMutex.RUnlock()
	return len(fake.reapVolumeArgsForCall)
}

func (fake *FakeVolumeFactoryDB) ReapVolumeArgsForCall(i int) string {
	fake.reapVolumeMutex.RLock()
	defer fake.reapVolumeMutex.RUnlock()
	return fake.reapVolumeArgsForCall[i].handle
}

func (fake *FakeVolumeFactoryDB) ReapVolumeReturns(result1 error) {
	fake.ReapVolumeStub = nil
	fake.reapVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolumeFactoryDB) SetVolumeTTL(arg1 string, arg2 time.Duration) error {
	fake.setVolumeTTLMutex.Lock()
	fake.setVolumeTTLArgsForCall = append(fake.setVolumeTTLArgsForCall, struct {
		arg1 string
		arg2 time.Duration
	}{arg1, arg2})
	fake.setVolumeTTLMutex.Unlock()
	if fake.SetVolumeTTLStub != nil {
		return fake.SetVolumeTTLStub(arg1, arg2)
	} else {
		return fake.setVolumeTTLReturns.result1
	}
}

func (fake *FakeVolumeFactoryDB) SetVolumeTTLCallCount() int {
	fake.setVolumeTTLMutex.RLock()
	defer fake.setVolumeTTLMutex.RUnlock()
	return len(fake.setVolumeTTLArgsForCall)
}

func (fake *FakeVolumeFactoryDB) SetVolumeTTLArgsForCall(i int) (string, time.Duration) {
	fake.setVolumeTTLMutex.RLock()
	defer fake.setVolumeTTLMutex.RUnlock()
	return fake.setVolumeTTLArgsForCall[i].arg1, fake.setVolumeTTLArgsForCall[i].arg2
}

func (fake *FakeVolumeFactoryDB) SetVolumeTTLReturns(result1 error) {
	fake.SetVolumeTTLStub = nil
	fake.setVolumeTTLReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolumeFactoryDB) SetVolumeSizeInBytes(arg1 string, arg2 int64) error {
	fake.setVolumeSizeInBytesMutex.Lock()
	fake.setVolumeSizeInBytesArgsForCall = append(fake.setVolumeSizeInBytesArgsForCall, struct {
		arg1 string
		arg2 int64
	}{arg1, arg2})
	fake.setVolumeSizeInBytesMutex.Unlock()
	if fake.SetVolumeSizeInBytesStub != nil {
		return fake.SetVolumeSizeInBytesStub(arg1, arg2)
	} else {
		return fake.setVolumeSizeInBytesReturns.result1
	}
}

func (fake *FakeVolumeFactoryDB) SetVolumeSizeInBytesCallCount() int {
	fake.setVolumeSizeInBytesMutex.RLock()
	defer fake.setVolumeSizeInBytesMutex.RUnlock()
	return len(fake.setVolumeSizeInBytesArgsForCall)
}

func (fake *FakeVolumeFactoryDB) SetVolumeSizeInBytesArgsForCall(i int) (string, int64) {
	fake.setVolumeSizeInBytesMutex.RLock()
	defer fake.setVolumeSizeInBytesMutex.RUnlock()
	return fake.setVolumeSizeInBytesArgsForCall[i].arg1, fake.setVolumeSizeInBytesArgsForCall[i].arg2
}

func (fake *FakeVolumeFactoryDB) SetVolumeSizeInBytesReturns(result1 error) {
	fake.SetVolumeSizeInBytesStub = nil
	fake.setVolumeSizeInBytesReturns = struct {
		result1 error
	}{result1}
}

var _ worker.VolumeFactoryDB = new(FakeVolumeFactoryDB)

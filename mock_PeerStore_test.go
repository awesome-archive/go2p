// Code generated by mockery v1.0.0. DO NOT EDIT.

package go2p

import mock "github.com/stretchr/testify/mock"

// MockPeerStore is an autogenerated mock type for the PeerStore type
type MockPeerStore struct {
	mock.Mock
}

// AddPeer provides a mock function with given fields: peer
func (_m *MockPeerStore) AddPeer(peer *Peer) error {
	ret := _m.Called(peer)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Peer) error); ok {
		r0 = rf(peer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IteratePeer provides a mock function with given fields: handler
func (_m *MockPeerStore) IteratePeer(handler func(*Peer)) {
	_m.Called(handler)
}

// LockPeer provides a mock function with given fields: addr, handler
func (_m *MockPeerStore) LockPeer(addr string, handler func(*Peer)) {
	_m.Called(addr, handler)
}

// OnPeerAdd provides a mock function with given fields: handler
func (_m *MockPeerStore) OnPeerAdd(handler func(*Peer)) {
	_m.Called(handler)
}

// OnPeerWantRemove provides a mock function with given fields: handler
func (_m *MockPeerStore) OnPeerWantRemove(handler func(*Peer)) {
	_m.Called(handler)
}

// RemovePeer provides a mock function with given fields: peer
func (_m *MockPeerStore) RemovePeer(peer *Peer) {
	_m.Called(peer)
}

// Start provides a mock function with given fields:
func (_m *MockPeerStore) Start() {
	_m.Called()
}

// Stop provides a mock function with given fields:
func (_m *MockPeerStore) Stop() {
	_m.Called()
}

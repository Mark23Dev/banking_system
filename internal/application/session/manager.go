package session

import "banking_system/internal/domain/user"

type Manager struct {
	currentUser *user.User
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) Login(u *user.User) {
	m.currentUser = u
}

func (m *Manager) Logout() {
	m.currentUser = nil
}

func (m *Manager) CurrentUser() *user.User {
	return m.currentUser
}

func (m *Manager) IsAuthenticated() bool {
	return m.currentUser != nil
}
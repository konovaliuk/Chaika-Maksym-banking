package model

import (
	"time"

	"github.com/google/uuid"
)

type ManagerRole string

const (
	ManagerRoleAdmin     ManagerRole = "admin"
	ManagerRoleModerator ManagerRole = "moderator"
)

type Manager struct {
	Id       uuid.UUID
	FullName string
	Role     ManagerRole
	JoinedAt time.Time
}

func NewManager(
	id uuid.UUID,
	fullName string,
	role ManagerRole,
	joinedAt time.Time,
) *Manager {
	return &Manager{
		Id:       id,
		FullName: fullName,
		Role:     role,
		JoinedAt: joinedAt,
	}
}

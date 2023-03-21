package model

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	id           uuid.UUID
	email        string
	passwordHash string
	createdAt    time.Time
	updatedAt    time.Time
}

func NewCustomer(
	id uuid.UUID,
	email string,
	passwordHash string,
	createdAt time.Time,
	updatedAt time.Time,
) *Customer {
	return &Customer{
		id:           id,
		email:        email,
		passwordHash: passwordHash,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
	}
}

func (c *Customer) Update(
	email string,
	passwordHash string,
) {
	c.email = email
	c.passwordHash = passwordHash
	c.updatedAt = time.Now()
}

func (c *Customer) Id() uuid.UUID {
	return c.id
}

func (c *Customer) SetId(id uuid.UUID) {
	c.id = id
}

func (c *Customer) Email() string {
	return c.email
}

func (c *Customer) SetEmail(email string) {
	c.email = email
}

func (c *Customer) PasswordHash() string {
	return c.passwordHash
}

func (c *Customer) SetPasswordHash(passwordHash string) {
	c.passwordHash = passwordHash
}

func (c *Customer) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Customer) SetCreatedAt(createdAt time.Time) {
	c.createdAt = createdAt
}

func (c *Customer) UpdatedAt() time.Time {
	return c.updatedAt
}

func (c *Customer) SetUpdatedAt(updatedAt time.Time) {
	c.updatedAt = updatedAt
}

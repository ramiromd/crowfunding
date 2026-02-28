package entity

import (
	sharedvalue "github.com/ramiromd/crowfunding/internal/shared/domain/value_object"
	identityvalue "github.com/ramiromd/crowfunding/internal/identity/domain/value_object"
)


type User struct {
	entityId	sharedvalue.EntityId
	nickname identityvalue.Nickname
	email identityvalue.Email
	passwordHash identityvalue.PasswordHash
	createdAt sharedvalue.CreationDate
	updatedAt sharedvalue.UpdateDate
}

func NewUser(entityId sharedvalue.EntityId, nickname identityvalue.Nickname, email identityvalue.Email, passwordHash identityvalue.PasswordHash, createdAt sharedvalue.CreationDate, updatedAt sharedvalue.UpdateDate) (*User, error) {
	
	// TODO: Inicializar createdAt implicitamente sin parametros.
	// TODO: UpdatedAt debe ser opcional. Y no debe tener un valor por defecto.
	if err := updatedAt.CheckGreaterOrEqualThan(createdAt); err != nil {
		return nil, err
	}
	return &User{
		entityId: entityId,
		nickname: nickname,
		email: email,
		passwordHash: passwordHash,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}

func (u *User) EntityId() sharedvalue.EntityId {
	return u.entityId
}

func (u *User) Nickname() identityvalue.Nickname {
	return u.nickname
}

func (u *User) Email() identityvalue.Email {
	return u.email
}

func (u *User) PasswordHash() identityvalue.PasswordHash {
	return u.passwordHash
}

func (u *User) ChangePassword(passwordHash identityvalue.PasswordHash) *User {
	u.passwordHash = passwordHash
	u.updatedAt = sharedvalue.NewUpdateDate()
	return u
}

func (u *User) CreatedAt() sharedvalue.CreationDate {
	return u.createdAt
}

func (u *User) UpdatedAt() sharedvalue.UpdateDate {
	return u.updatedAt
}
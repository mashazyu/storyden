package account

import (
	"context"

	"github.com/Southclaws/opt"
	"github.com/rs/xid"

	"github.com/Southclaws/storyden/internal/ent"
)

type Option func(*Account)

type Repository interface {
	Create(ctx context.Context, handle string, opts ...Option) (*Account, error)

	GetByID(ctx context.Context, id AccountID) (*Account, error)
	GetByHandle(ctx context.Context, handle string) (*Account, error)

	Update(ctx context.Context, id AccountID, opts ...Mutation) (*Account, error)
}

func WithID(id AccountID) Option {
	return func(a *Account) {
		a.ID = AccountID(id)
	}
}

func WithAdmin(admin bool) Option {
	return func(a *Account) {
		a.Admin = true
	}
}

func WithName(name string) Option {
	return func(a *Account) {
		a.Name = name
	}
}

func WithBio(bio string) Option {
	return func(a *Account) {
		a.Bio = opt.New(bio)
	}
}

type Mutation func(u *ent.AccountUpdateOne)

func SetHandle(handle string) Mutation {
	return func(u *ent.AccountUpdateOne) {
		u.SetHandle(handle)
	}
}

func SetName(name string) Mutation {
	return func(u *ent.AccountUpdateOne) {
		u.SetName(name)
	}
}

func SetBio(bio string) Mutation {
	return func(u *ent.AccountUpdateOne) {
		u.SetBio(bio)
	}
}

func SetAdmin(status bool) Mutation {
	return func(u *ent.AccountUpdateOne) {
		u.SetAdmin(status)
	}
}

func SetInterests(interests []xid.ID) Mutation {
	return func(u *ent.AccountUpdateOne) {
		u.ClearTags().AddTagIDs(interests...)
	}
}

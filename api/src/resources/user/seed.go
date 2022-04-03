package user

import (
	"context"

	"4d63.com/optional"
	"github.com/google/uuid"

	"github.com/Southclaws/storyden/api/src/infra/db/model"
)

var (
	SeedUser_01_Admin = User{
		ID:    UserID(uuid.MustParse("00000000-0000-0000-0000-000000000000")),
		Email: "tim@storyd.en",
		Name:  "TimManTheTinMan",
		Bio:   optional.Of("I run this place"),
		Admin: true,
	}

	SeedUser_02_User = User{
		ID:    UserID(uuid.MustParse("00000000-0000-0000-0000-000000000000")),
		Email: "tam@storyd.en",
		Name:  "IDontLikeTom",
		Bio:   optional.Of("I'm just called Tam"),
	}
)

func NewMockWithSeed() Repository {
	m := NewMock()
	Seed(m)
	return m
}

func NewWithSeed(db *model.Client) Repository {
	m := New(db)
	Seed(m)
	return m
}

func Seed(r Repository) {
	ctx := context.Background()

	r.CreateUser(ctx, SeedUser_01_Admin.Email, SeedUser_01_Admin.Name)
	r.CreateUser(ctx, SeedUser_02_User.Email, SeedUser_02_User.Name)
}
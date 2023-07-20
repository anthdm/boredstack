package migrations

import (
	"context"
	"github.com/anthdm/bs/data"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		db.RegisterModel((*data.Todo)(nil))

		//fixture := dbfixture.New(db, dbfixture.WithRecreateTables())
		//return fixture.Load(ctx, bunapp.FS(), "fixture/fixture.yml")
		return nil
	}, nil)
}

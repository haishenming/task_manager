package data

import (
	"context"

	"task_manager/ent"
	"task_manager/ent/migrate"
	"task_manager/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	// mysql driver.
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewClient, NewTaskRepo)

// Data .
type Data struct {
	client *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, client *ent.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		client: client,
	}, cleanup, nil
}

func (d *Data) Close() error {
	return d.client.Close()
}

func (d *Data) Client() *ent.Client {
	return d.client
}

func NewClient(c *conf.Data) (*ent.Client, error) {
	client, err := ent.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		return nil, err
	}

	err = client.Debug().Schema.Create(
		context.TODO(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}

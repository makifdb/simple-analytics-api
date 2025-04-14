package database

import (
	"context"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2"
	_ "github.com/ClickHouse/clickhouse-go/v2"
)

func NewClickHouseClient(ctx context.Context, host string, port int, username, password string) (clickhouse.Conn, error) {

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%d", host, port)},
		Auth: clickhouse.Auth{
			Database: "analytics",
			Username: username,
			Password: password,
		},
	})
	if err != nil {
		return nil, err
	}
	v, err := conn.ServerVersion()
	fmt.Println(v)

	if err != nil {
		return nil, err
	}
	return conn, nil
}

package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/makifdb/simple-analytics-api/ent"
)

func NewEntClickHouseClient(ctx context.Context, host string, port int, username, password string) (*ent.Client, error) {

	// clickhouse://username:password@host:port/database
	connString := fmt.Sprintf("clickhouse://%s:%s@%s:%d/%s", username, password, host, port, "analytics")
	db, err := sql.Open("clickhouse", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to open ClickHouse connection: %w", err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping ClickHouse: %w", err)
	}

	err = Migration(ctx, db)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate ClickHouse: %w", err)
	}

	drv := entsql.OpenDB("clickhouse", db)
	return ent.NewClient(ent.Driver(drv)), nil
}

func Migration(ctx context.Context, db *sql.DB) error {

	query := `
		CREATE TABLE IF NOT EXISTS events (
			event_id UUID NOT NULL,
			user_id UUID NOT NULL,
			event_type String NOT NULL,
			timestamp DateTime NOT NULL,
			event_data String,
			PRIMARY KEY (event_id)
		) ENGINE = MergeTree()
		ORDER BY (event_id);
	`

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to create events table: %w", err)
	}

	return nil
}

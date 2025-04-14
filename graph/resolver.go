package graph

//go:generate go run github.com/99designs/gqlgen generate
import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/makifdb/simple-analytics-api/ent"
)

type Resolver struct {
	EntClient *ent.Client
	CliClient driver.Conn
}

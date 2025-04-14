package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Event struct {
	ent.Schema
}

func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("event_id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable(),
		field.UUID("user_id", uuid.UUID{}).
			Default(uuid.New),
		field.String("event_type").
			NotEmpty(),
		field.Time("timestamp").
			Default(time.Now).
			Immutable(),
		field.JSON("event_data", map[string]any{}).
			Optional(),
	}
}

func (Event) Edges() []ent.Edge {
	return nil
}

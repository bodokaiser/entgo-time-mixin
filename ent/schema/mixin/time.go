package mixin

import (
	"log"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Time mixin implements created at and updated at fields.
type Time struct {
	mixin.Schema
}

func (Time) Fields() []ent.Field {
	log.Printf("using clock %T\n", clock)

	return []ent.Field{
		field.Time("created_at").
			StructTag(`json:"createdAt,omitempty"`).
			Immutable().
			Default(now),
		field.Time("updated_at").
			StructTag(`json:"updatedAt,omitempty"`).
			Default(now).
			UpdateDefault(now),
	}
}

// Default clock for timestamps.
var clock Clock = SecondClock{}

func now() time.Time {
	return clock.Now()
}

// SetClock changes the default clock for timestamps.
func SetClock(c Clock) {
	log.Printf("setting default clock to %T\n", c)
	clock = c
}

// Clock represents a clock.
type Clock interface {
	Now() time.Time
}

// SecondClock implements a Clock with second precision.
type SecondClock struct {
}

func (c SecondClock) Now() time.Time {
	t := time.Now()

	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, t.Location())
}

// FixedClock implements a Clock which always returns a fixed time.
type FixedClock struct {
	Time time.Time
}

func (c FixedClock) Now() time.Time {
	return c.Time
}

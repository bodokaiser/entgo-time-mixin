package hook

import (
	"context"
	"fmt"

	"github.com/gosimple/slug"

	"github.com/bodokaiser/entgo-time-mixin/ent"
)

type Slugable interface {
	Name() string
	SetSlug(string)
}

func GenerateSlug(next ent.Mutator) ent.Mutator {
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		value, ok := m.Field("name")
		if !ok {
			return nil, fmt.Errorf("mutation has no name")
		}
		name, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("mutation name is not a string")
		}

		if err := m.SetField("slug", slug.Make(name)); err != nil {
			return nil, fmt.Errorf("mutation cannot set slug: %v", err)
		}

		return next.Mutate(ctx, m)
	})
}

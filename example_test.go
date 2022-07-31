package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/bodokaiser/entgo-time-mixin/ent"
	_ "github.com/bodokaiser/entgo-time-mixin/ent/runtime"
	"github.com/bodokaiser/entgo-time-mixin/ent/schema/mixin"
)

func TestMain(m *testing.M) {
	log.Printf("bootstrapping tests")

	mixin.SetClock(mixin.FixedClock{
		Time: time.Date(2007, time.June, 29, 0, 0, 0, 0, time.UTC),
	})

	os.Exit(m.Run())
}

func Example() {
	ctx := context.Background()

	client := open(ctx)
	defer client.Close()

	pet, err := client.Pet.Create().SetName("Doby").Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pet)

	// Output:
	// Pet(id=1, created_at=, updated_at=, name=Doby, slug=doby)
}

func open(ctx context.Context) *ent.Client {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}

package domain

import "testing"

func TestEntity(t *testing.T) {
	t.Run("NewEntity", func(t *testing.T) {
		t.Run("should return an entity with the given ID", func(t *testing.T) {
			id := "123"
			entity := NewEntity(WithID(id))

			if entity.ID() != id {
				t.Errorf("expected ID to be %s, got %s", id, entity.ID())
			}
		})
	})
}
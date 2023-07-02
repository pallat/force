package force

import "testing"

func TestRid(t *testing.T) {
	t.Run("", func(t *testing.T) {
		seeding = "123"
		r := Validate("123")
		if !r {
			t.Error()
		}
	})
	t.Run("", func(t *testing.T) {
		seeding = "123"
		r := Validate("312")
		if !r {
			t.Error()
		}
	})
	t.Run("", func(t *testing.T) {
		seeding = "123"
		r := Validate("310")
		if r {
			t.Error()
		}
	})
}

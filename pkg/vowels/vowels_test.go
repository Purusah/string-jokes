package vowels

import "testing"

func TestEncode(t *testing.T) {
	t.Run("lower case", func(t *testing.T) {
		if Encode("hello") != "h2ll4" {
			t.Error("should be h2ll4")
		}
	})
	t.Run("mixed case", func(t *testing.T) {
		if Encode("hEllo LLA") != "h2ll4 LL1" {
			t.Error("should be h2ll4 LL1")
		}
	})
}

func TestDecode(t *testing.T) {
	t.Run("lower case", func(t *testing.T) {
		if Decode("h2ll4") != "hello" {
			t.Error("should be h2ll4")
		}
	})
	t.Run("mixed case", func(t *testing.T) {
		if Decode("h2ll4 LL1") != "hello LLa" {
			t.Error("should be hello LLa")
		}
	})
}

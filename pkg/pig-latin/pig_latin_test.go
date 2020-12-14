package latin

import "testing"

func TestEncode(t *testing.T) {
	t.Run("start consonant", func(t *testing.T) {
		if Encode("pig") != "igpay" {
			t.Error("got", Encode("pig"))
		}
		if Encode("banana") != "ananabay" {
			t.Error("got", Encode("ananabay"))
		}
		if Encode("me") != "emay" {
			t.Error("got", Encode("me"))
		}
	})
	t.Run("start consonant clusters", func(t *testing.T) {
		if Encode("smile") != "ilesmay" {
			t.Error("got", Encode("smile"))
		}
		if Encode("glove") != "oveglay" {
			t.Error("got", Encode("glove"))
		}
	})
	t.Run("start with vowel sounds", func(t *testing.T) {
		if Encode("omelet") != "omeletay" {
			t.Error("got", Encode("omelet"))
		}
		if Encode("I") != "Iay" {
			t.Error("got", Encode("I"))
		}
		if Encode("are") != "areay" {
			t.Error("got", Encode("are"))
		}
	})
}

package integer

import "testing"

func TestAdder(t *testing.T) {
	actual := Add(2, 3)
	expected := 5

	if actual != expected {
		t.Errorf("expected '%d' got '%d'", expected, actual)
	}
}

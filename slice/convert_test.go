package slice

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConvertStringToSlice(t *testing.T) {
	expected := []string{"item1", "item2", "item3"}

	res := ConvertStringToSlice(",", "item1,item2,item3")

	if diff := cmp.Diff(expected, res); diff != "" {
		t.Errorf("CompareSlices() mismatch (-want +got):\n%s", diff)
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestParseCourseYAMLTags(t *testing.T) {
	got := parseCourseYAML([]byte(`tags: ["robotics", "esp32", "servo"]`))
	want := []string{"robotics", "esp32", "servo"}
	if !reflect.DeepEqual(got.Tags, want) {
		t.Fatalf("tags = %#v, want %#v", got.Tags, want)
	}
}

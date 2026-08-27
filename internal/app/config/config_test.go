package config

import (
	"reflect"
	"testing"
)

func TestAppConfigContainsOnlyTemplateSections(t *testing.T) {
	t.Parallel()

	configType := reflect.TypeOf(AppConfig{})
	want := []string{"Server", "Logging"}

	if configType.NumField() != len(want) {
		t.Fatalf("AppConfig has %d sections, want %d", configType.NumField(), len(want))
	}

	for index, fieldName := range want {
		if got := configType.Field(index).Name; got != fieldName {
			t.Errorf("AppConfig section %d = %q, want %q", index, got, fieldName)
		}
	}
}

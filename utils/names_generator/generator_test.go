package names_generator_test

import (
	"strings"
	"testing"

	"github.com/jurelou/forensibus/utils/names_generator"
)

func TestNameFormat(t *testing.T) {
	name := names_generator.GetRandomName()
	if !strings.Contains(name, "_") {
		t.Fatalf("Generated name does not contain an underscore")
	}
}

func TestNameRandom(t *testing.T) {
	name := names_generator.GetRandomName()
	name2 := names_generator.GetRandomName()
	name3 := names_generator.GetRandomName()

	if name == name2 || name == name3 || name2 == name3 {
		t.Fatalf("Duplicated names ...")
	}
}

func BenchmarkGetRandomName(b *testing.B) {
	b.ReportAllocs()
	var out string
	for n := 0; n < b.N; n++ {
		out = names_generator.GetRandomName()
	}
	b.Log("Last result:", out)
}

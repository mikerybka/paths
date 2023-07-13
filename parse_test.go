package paths

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{"empty", "", []string{}},
		{"empty with slash", "/", []string{}},
		{"single", "foo", []string{"foo"}},
		{"single with slash", "/foo", []string{"foo"}},
		{"double", "foo/bar", []string{"foo", "bar"}},
		{"double with slash", "/foo/bar", []string{"foo", "bar"}},
		{"trailing slash", "foo/bar/", []string{"foo", "bar"}},
		{"trailing slash with slash", "/foo/bar/", []string{"foo", "bar"}},
		{"leading slash", "/foo/bar", []string{"foo", "bar"}},
		{"leading slash with trailing slash", "/foo/bar/", []string{"foo", "bar"}},
		{"leading and trailing slash", "/foo/bar/", []string{"foo", "bar"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.s); !equal(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, s := range a {
		if s != b[i] {
			return false
		}
	}
	return true
}

package stacks

import "testing"

func TestParseFeatureName(t *testing.T) {
	cases := []struct {
		input, wantPascal, wantCamel, wantSnake string
	}{
		{"user profile", "UserProfile", "userProfile", "user_profile"},
		{"UserProfile", "UserProfile", "userProfile", "user_profile"},
		{"user_profile", "UserProfile", "userProfile", "user_profile"},
		{"userProfile", "UserProfile", "userProfile", "user_profile"},
		{"my-feature", "MyFeature", "myFeature", "my_feature"},
		{"auth", "Auth", "auth", "auth"},
		{"User Profile", "UserProfile", "userProfile", "user_profile"},
	}
	for _, c := range cases {
		p, ca, s := ParseFeatureName(c.input)
		if p != c.wantPascal {
			t.Errorf("%q pascal: got %q want %q", c.input, p, c.wantPascal)
		}
		if ca != c.wantCamel {
			t.Errorf("%q camel: got %q want %q", c.input, ca, c.wantCamel)
		}
		if s != c.wantSnake {
			t.Errorf("%q snake: got %q want %q", c.input, s, c.wantSnake)
		}
	}
}

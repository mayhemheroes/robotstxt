package robotstxt

import "testing"

func FuzzParseBytes(f *testing.F) {
	f.Add([]byte("User-agent: *\nDisallow: /"))
	f.Fuzz(func(t *testing.T, data []byte) {
		r, err := FromBytes(data)
		if err != nil {
			if r != nil {
				t.Fatal("r != nil on error")
			}
			return
		}
		_ = r.FindGroup("googlebot")
		r.TestAgent("/path", "googlebot")
	})
}

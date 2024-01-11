package main

import (
	"testing"
)

func TestCreateHTML(t *testing.T) {
	// TODO: Write your test case here
	t.Run("simple text", func(t *testing.T) {
		f := fuckity{"this is simple text", "", nil}
		got, _ := Fuck(f)
		want := "<div>\nthis is simple text\n</div>"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("simple text with params", func(t *testing.T) {
		f := fuckity{"this is simple text", "hx-post=\"/submit\"", nil}
		got, _ := Fuck(f)
		want := "<div hx-post=\"/submit\" >\nthis is simple text\n</div>"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("fuck with asses", func(t *testing.T) {
		c1 := fuckity{"ass 1", "", nil}
		c2 := fuckity{"ass 2", "", nil}
		c3 := fuckity{"ass 3", "", nil}
		l := []fuckity{c1, c2, c3}
		f := fuckity{"", "hx-post=\"/submit\"", l}

		got, _ := Fuck(f)
		want := `<div hx-post="/submit" >
<div>
ass 1
</div>
<div>
ass 2
</div>
<div>
ass 3
</div>
</div>`
		if got != want {
			t.Errorf("got %q\nwant %q", got, want)
		}
	})
	t.Run("ass equals another fuck + Content", func(t *testing.T) {
		f_inside := fuckity{"inside content", "", nil}
		f := fuckity{"outside content", "", []fuckity{f_inside}}
		got, _ := Fuck(f)
		want :=
			`<div>
outside content
<div>
inside content
</div>
</div>`
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

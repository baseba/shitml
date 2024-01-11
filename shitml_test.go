package shitml

import (
	"testing"
)

func TestCreateHTML(t *testing.T) {
	// TODO: Write your test case here
	t.Run("simple text", func(t *testing.T) {
		f := Fuckity{"this is simple text", "", nil}
		got, _ := Fuck(f)
		want := "<div>\nthis is simple text\n</div>"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("simple text with params", func(t *testing.T) {
		f := Fuckity{"this is simple text", "hx-post=\"/submit\"", nil}
		got, _ := Fuck(f)
		want := "<div hx-post=\"/submit\" >\nthis is simple text\n</div>"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("fuck with asses", func(t *testing.T) {
		c1 := Fuckity{"ass 1", "", nil}
		c2 := Fuckity{"ass 2", "", nil}
		c3 := Fuckity{"ass 3", "", nil}
		l := []Fuckity{c1, c2, c3}
		f := Fuckity{"", "hx-post=\"/submit\"", l}

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
		f_inside := Fuckity{"inside content", "", nil}
		f := Fuckity{"outside content", "", []Fuckity{f_inside}}
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

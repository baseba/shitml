package main

import (
	"fmt"
	"strings"
)

type fuckity struct {
	Content string
	Params  string
	Ass     []fuckity
}

func deez(fuck fuckity) string {
	switch fuck.Params {
	case "":
		if fuck.Content != "" {

			return fmt.Sprintf("\n<div>\n%s\n</div>\n", fuck.Content)
		}
		return fmt.Sprintf("\n<div>\n%s</div>\n", fuck.Content)

	default:
		if fuck.Content != "" {

			return fmt.Sprintf("\n<div %s >\n%s\n</div>\n", fuck.Params, fuck.Content)
		}
		return fmt.Sprintf("\n<div %s >\n%s</div>\n", fuck.Params, fuck.Content)

	}
}

func fuck(shit fuckity) (string, error) {
	s := ""
	if shit.Ass == nil {
		d := deez(shit)
		d = strings.Replace(d, "\n\n", "\n", -1)
		d = strings.TrimSuffix(d, "\n")
		d = strings.TrimPrefix(d, "\n")

		return d, nil
	}

	for _, c := range shit.Ass {
		if c.Ass == nil {

			s += deez(c)
			s += "\n"

			continue
		}
		nested, _ := fuck(c)
		s += nested
		s += "\n"
	}
	s = strings.Replace(s, "\n\n", "\n", -1)
	// s += "\n"
	// fmt.Println(s)
	inter_fuck := fuckity{shit.Content + s, shit.Params, nil}
	return fuck(inter_fuck)

}

func main() {

	result, _ := fuck(fuckity{"content", "params", nil})
	fmt.Print(result)

}

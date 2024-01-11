package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type fuckity struct {
	Content string
	Params  string
	Ass     []fuckity
}

func Cheecks(butthole string) (string, error) {
	htmx := "<script src=\"https://unpkg.com/htmx.org@1.9.9\" type=\"text/javascript\"></script>"
	tailwind := "<script src=\"https://cdn.tailwindcss.com\"></script>"
	head := "<head>" + "\n" + htmx + "\n" + tailwind + "\n</head>"
	r := head + butthole
	return r, nil
}

func deez(fuck fuckity) string {

	switch fuck.Params {
	case "":
		if fuck.Content != "" {

			s := fmt.Sprintf("\n<div>\n%s\n</div>\n", fuck.Content)
			s = strings.Replace(s, "\n\n", "\n", -1)
			return s
		}
		s := fmt.Sprintf("\n<div>\n%s\n</div>\n", fuck.Content)
		s = strings.Replace(s, "\n\n", "\n", -1)
		return s
	default:
		if fuck.Content != "" {

			s := fmt.Sprintf("\n<div %s >\n%s\n</div>\n", fuck.Params, fuck.Content)
			s = strings.Replace(s, "\n\n", "\n", -1)
			return s
		}
		s := fmt.Sprintf("\n<div %s >\n%s\n</div>\n", fuck.Params, fuck.Content)
		s = strings.Replace(s, "\n\n", "\n", -1)
		return s
	}

}

func Fuck(shit fuckity) (string, error) {
	s := ""
	if shit.Ass == nil {
		d := deez(shit)
		d = strings.Replace(d, "\n\n", "\n", -1)
		// d = strings.TrimSuffix(d, "\n")
		// d = strings.TrimPrefix(d, "\n")

		return d, nil
	}

	for _, c := range shit.Ass {
		if c.Ass == nil {

			s += deez(c)
			s += "\n"

			continue
		}
		nested, _ := Fuck(c)
		s += nested
		s += "\n"
	}
	s += "\n"
	s = shit.Content + s
	s = strings.Replace(s, "\n\n", "\n", -1)
	s = strings.Replace(s, "\n \n", "\n", -1)

	// s = strings.TrimSuffix(s, "\n")
	// s = strings.TrimPrefix(s, "\n")
	inter_fuck := fuckity{s, shit.Params, nil}
	return Fuck(inter_fuck)

}

func main() {

	result, _ := Fuck(fuckity{"title", "class='bg-black text-white text-xl' hx-post='/clicked' hx-trigger='click'", []fuckity{{"subtitle 1", "class='text-lg'", []fuckity{{"sub-sub-title", "", nil}}}, {"subtitle2", "", nil}}})
	final_final_result, _ := Fheecks(result)
	f, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err2 := f.WriteString(final_final_result)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Print("done")

}

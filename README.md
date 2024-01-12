
# shitml
A superior way to write HTML.

## Installation
```bash
go get github.com/baseba/shitml
```

## Manifest

- Becauso of HTMX you only need `<div(s)>`.
- Functions to write elements.
- Cursewords are easier to remember.

### Syntax Inspiration
[Syntax Inspiration Link](https://github.com/baseba/shitml/assets/48297670/6c933438-7cdb-4f0d-a749-0316666e9629)

## How it Works
- SHiTML💩 produces only divs in the final HTML.
- It includes Tailwind and htmx for building a great UI.

## Functions and Types
There is only one type called `fuckity`:
```go
type fuckity struct {
	Content string
	Params  string
	Ass     []fuckity
}
```

### With this type, we can build a `fuck()` function that returns a `<div>`.

## Example
Here, we can see that the readable function:
```go
shitml.cheeks(shitml.fuck(shitml.fuckity{Content: "i am the content",Params:  "class='bg-black text-white text-xl' hx-post='/clicked' hx-trigger='click'",    Ass:     nil,}))
```
Generates this complex and messy HTML:
```html
<div class='bg-black text-white text-xl' hx-post='/clicked' hx-trigger='click' >
    i am the content
<div class='text-lg' >
```

**Note:** Always wrap the outer `shitml.fuck()` with `shitml.cheeks()` to have Tailwind and htmx support.
```

Feel free to make any additional adjustments or ask for further clarifications!
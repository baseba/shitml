# shitml
A superior way to write html


# Manifest 

- Only <code>\<div\></code> is necesary
- Function are superior
- Eric Cartman created the syntax

## Our syntax inspiration
https://github.com/baseba/shitml/assets/48297670/6c933438-7cdb-4f0d-a749-0316666e9629

# How it works

- sHiTML💩 produces only divs in the final HTML
- It includes tailwind and htmx so you can build a great experience

# functions and types
there is only one type called fuckity:
```
type fuckity struct {
	Content string
	Params  string
	Ass     []fuckity
}
```

### with this type we can build a `fuck()` function who returns a <code>\<div\></code>

- example
 ```
 fuck(fuckity{"i am the content", "class='bg-black text-white text-xl' hx-post='/clicked' hx-trigger='click'", nil})
 ```
 generates 
 ```
 <div class='bg-black text-white text-xl' hx-post='/clicked' hx-trigger='click' >
    i am the content
<div class='text-lg' >
 ```

### Be careful to always wrap the outer <code>fuck</code> with <code>cheeks()</code> to have tailwind and htmx support



# 5. Dynamic HTML templates

## 5.1. Displaying dynamic data

- Rendering multiple pieces of data
  - `html/template` allows you to pass in one -- and only one -- item of dynamic data when rendering a template
  - A lightweight and type-safe way to achieve this is to wrap your dynamic data in a struct which acs like a single "holding structure" for your data
- Dynamic content escaping
  - `html/template` automatically escapes anything between `{{ }}` tags
  - `htmp/template` is smart enough to make escaping context-dependent
- Nested templates
  - Dot needs to be epxlicitly passed or _pipelined_ to the template being invoked:
  - `{{template "main" .}}{{ block "sidebar" .}}{{end}}`
  - Get in the habit of always pipelining dot whenever you invoke a template with the `{{template}}` or `{{block}}` actions
- Calling methods
  - If the type you're yielding between `{{ }}` tags has methods defined against it, you can call these methods (so long as they are exported and they return only a single value -- or a single value and an error)
  - `<span>{{.Snippet.Created.Weekday}}</span>`
  - `<span>{{.Snippet.Created.AddDate 0 6 0}}</span>`
- HTML comments
  - `html/template` always strips out HTML comments you include in your templates, including conditional comments


## 5.2. Template actions and functions

| **Action** | **Description**|
| --- | --- |
| `{{if .Foo}} C1 {{else}} C2 {{end}}` | If `.Foo` is not empty then render the content C1, otherwise render the content C2. |
| `{{with .Foo}} C1 {{ else }} C2 {{end}} | If `.Foo` is not empty, then set dot to the value of `.Foo` and render the content C1, otherwise render the content C2. |
| `{{range .Foo}} C1 {{else}} C2 {{end}}` | If the length of `.Foo` is greater than zero then loop over each element, setting dot to the value of each element and rendering the content C1. If the length of `.Foo` is zero then render the content C2. The underlying type of `.Foo` must be an array, slice, map, or channel. |

- `{{else}}` is optional
- The _empty_ values are `false`, `0`, any nil pointer or interface value, and any array, slice, map, or string of length zero
- It's important to grasp that the `with` and `range` actions change the value of dot

- `html/template`:

| **Action** | **Description**|
| --- | --- |
| `{{eq .Foo .Bar}}` | Yields true if `.Foo` is equal to `.Bar` |
| `{{ne .Foo .Bar}}` | Yields true if `.Foo`. is not equal to `.Bar` |
| `{{not .Foo}}` | Yields the boolean negation of `.Foo` |
| `{{or .Foo .Bar}}` | Yields `.Foo`. if `.Foo` is not empty; otherwise yields `.Bar` |
| `{{index .Foo i}}` | Yields the value of `.Foo` at index `i`. The underlying type of `.Foo` must be a map, slice or array, and `i` must be an integer value. |
| `{{printf "%s-%s" .Foo .Bar}} | Yields a formatted string containing the `.Foo`. and `.Bar` values. Works in the same way as fmt.Sprintf() |
| `{{len .Foo}}` | Yields the length of `.Foo` as an integer. |
| `{{$bar := len .Foo}}` | Yields the length of `.Foo` as an integer. |
| `{{$bar := len .Foo}}` | Declare and assign the length of `.Foo` to the template variable `$bar` |

- Combining functions
  - Use parentheses `()` to surround the functions and their arguments as necessary
  - Ex: `{{if (and (eq .Foo 1) (le .Bar 20))}} C1 {{end}}`
- Controlling loop behavior
  - Within a `{{range}}` you can use `{{break}}` to end the loop early
  - `{{continue}}` to immediately start the next loop iteration


## 5.3. Caching templates

- Refactoring a few things in the code, things to fix:
  1. `template.ParseFiles()` is called for each render -- avoid this by parsing the files once, and storing templates in an in-memory cache
  2. Reduce code duplication in `home` and `snippetView`
  

## 5.4. Catching runtime errors

- Make template rendering a two-stage process:
  1. Make a _trial_ render by writing the template into a buffer
  2. If this fails, respond to the user with an error message, otherwise, write the contents of the buffer to the `http.ResponseWriter`


## 5.5. Common dynamic data

## 5.6. Custom template functions

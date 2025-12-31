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

## 5.3. Caching templates

## 5.4. Catching runtime errors

## 5.5. Common dynamic data

## 5.6. Custom template functions

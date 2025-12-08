# GSS - Go Style Sheets

GSS features a DSL and a program. The DSL can be used to specify CSS like layout and styling rules for individual elements of a canvas. The program can be used to render the an image file of the canvas out of individual elements and GSS rules. In a broad sense GSS is a very basic implementation of CSS. GSS enables Go developers to offer their users dynamic asset creation features using familiar interface without spanning browser like processes during execution.

## Writing GSS

User provide two files. One is for the hierarchy of elements:

```html
<div id="main">
  <div id="title">Do we really need to span a browser to render some HTML and CSS?</div>
  <img srcset="bugs@2x.png 2x, bugs@3x 3x" />
  <div id="author">Bugs Bunny</div>
</div>
```

The other is for styling and layout of elements:

```css
#main {
  width: 400px;
  height: 300px;
  padding: 20px;
}

.title {
  font-family: "Helvetica Neue", sans-serif;
  font-size: 12pt;
}

img {
}
```

GSS supports plenty of the fundamental CSS selector operators and CSS properties. GSS properties posses very similar behavior to CSS properties.

## Rendering GSS

```go
package main

import "go.ufukty.com/gss/pkg/render"

viewport = render.ViewportOptions{
  Width:   "400px",
  Height:  "300px",
  Density: "2x",
}
err := render.Gss(writer, "elements.html", "styles.css", "#main", viewport)
```

##

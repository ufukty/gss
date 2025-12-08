# Spec

## Tags

GSS supports `div`, `span` and `img` tags.

## Selectors

GSS supports the direct and indirect parenting relations in selectors:

```css
parent > child
ancestor successor
```

## Layout

### Display

**Outside**

GSS supports both the `inline` and `block` display modes. Defaults are `block` for `div` and `inline` for `img` and `span`.

```css
selector {
  display: [ inline | block ];
}
```

**Children positioning**

```css
selector {
  display: [ grid | flex];
}
```

Please note that when container has been set to either of `grid` and `flex` children's own choices between `block` and inline` is ignored.

### Dimensions

GSS calculates a element's width and height values according to the smallest\* bounding box than can house its children. Note that when the content is text smallest bounding box is calculated by assuming the text needs to be laid out in one line without folding into multiple lines.

Although when an element has calculated under inherited dimensional constrained (from its ancestors have `width` or `height` manually set)

## Styles

### Background

```css
selector {
  background: [ #RRGGBB | #RGB | #RRGGBBAA | #RGBA];
}
```

### Text

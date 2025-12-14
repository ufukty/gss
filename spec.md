# Spec

## Tags

GSS supports `div`, `span` and `img` tags.

## Selectors

GSS supports selecting set of elements based on the tag name, id, class name and direct or indirect container.

```
*
#an-id
.a-class
img
div
span
#main > .title
#main .preheading
```

## Layout

### Display

Display property accepts two values each for outside and inside positioning behavior. The outside is for the behavior of element inside its parent. The inside is for laying its children.

```css
selector {
  display: [ inline | block] [ flow | grid | flex];
}
```

GSS supports both the `inline` and `block` display modes. Defaults are `block` for `div` and `inline` for `img` and `span`. Subsequent inline elements share a line. Block elements sits alone in their lines. When siblings mix different outside positioning modes, `block` positioned ones break lines.

> Siblings with mixed modes such as `inline`, `inline`, `inline`, `block` and `inline` would be rendered across 3 lines. The first 3 would share the first line. Although, lack of available container space would make the first line fold too.

Outside positioning values `inline` and `block` are used only if parent inside positioning mode is `flow`.

### Dimensions

GSS calculates a element's width and height values according to the smallest\* bounding box than can house its children. Note that when the content is text smallest bounding box is calculated by assuming the text needs to be laid out in one line without folding into multiple lines.

Although when an element has calculated under inherited dimensional constrained (from its ancestors have `width` or `height` manually set)

## Styles

### Background

GSS accepts both the long and short form of 3-channel and 4-channel HEX based color values. Such as `#FF0000` for red and `#FFF8` for half-transperent white.

```css
selector {
  background: [ #RRGGBB | #RGB | #RRGGBBAA | #RGBA];
}
```

### Text

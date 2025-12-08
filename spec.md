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

Display property accepts two values each for outside and inside positioning behavior. The outside is for the behavior of element inside its parent which is considered as outside . The inside is for laying its children.

```css
selector {
  display: [ inline | block] [ flow | grid | flex];
}
```

**Outside**

GSS supports both the `inline` and `block` display modes. Defaults are `block` for `div` and `inline` for `img` and `span`. Subsequent inline elements share a line. Block elements sits alone in their lines.

When there are elements with both of the outside positioning values as in the siblings `inline`, `inline`, `inline`, `block`, `inline` GSS renders them in 3 lines by putting the first 3 at the same line. If the container element of those has an inside positioning value other than the `flow` (which is the default) those values will be ignored.

**Inside**

Please note that when container has been set to either of `grid` and `flex` children's own choices between `block` and `inline` is ignored.

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

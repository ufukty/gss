# Contributing

I hope this article will make you understand the design decisions made on internals quicker.

## How it works?

Outline:

1. Parsing HTML, GSSE and GSS.
1. Wrapping AST nodes within DOM nodes.
1. Calculating effective styles (DFS)
   1. Resolving GSSE to values in pixels, ms, angles etc. using contextual information.
1. Calculating content sizes (DFS postorder).
1. Calculating layout (BFS).

AST represents the user provided file contents in a format that is way closer to the how users write them in contrast to the DOM where the information is organized within structures that are designed to ease the internal processes. DOM nodes gather the information that is processed together into one place.

Calculating content size of a node requires its children to be set beforehand, so it is performed in DFS postorder. Content size of a node actually a range determined with `min-content` and `max-content`. Both is layout free. For example a grand-parent's `min-content` size can be the longest render size of a word one of its grand-children contains. Content sizes ignores display property for say.

Calculating content size of leaf nodes requires their affective styling to be decided because both the dimensions of textual and visual content get effected by styling rules (eg. `font-weight`, `border` etc.). This is way finalizing styles is before measuring content sizes.

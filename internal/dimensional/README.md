# Dimensional

This package provides the type `Dimensional` which allows storing values of GSS units.

## Features

- Canonicalization of any value into its core unit.
- Storage and processing of compound units like `em·pc·px²`.
- Arithmetical operations addition, subtraction, multiplication and division.
- Value comparison.
- Pretty printing.

## Limitations

Dimensional supports only the storage of values in contextual units like `em`, `rem`, `%`, `vw` etc. They need to be converted into `px` or other core unit before processed by package.

## Maintenance note

Dimensional is designed to be unaware of contextual information that is needed to resolve written length, angle etc. values to canvas values in order to avoid circular dependency issues. This way gss expressions can be parsed during gss parsing which is earlier than dom involve.

# Experiment: exposing internalized structs

This repo demonstrates that it's possible, in Go, to still access structs
defined within `internal` if they're returned by public functions defined
outside of `internal`.

- The `internalized` module has a single package `a` which exposes a function
  `MyFunc`, returning an internal type.
- The `uses-internalized` module has a `main` function which imports `a.MyFunc`,
  executes it and can access the internal type.

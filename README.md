# Traveller

Go library to help traverse a complex nest of struct/map/array/slice.

```go
package main

import (
	"fmt"

	"github.com/pwnedgod/traveller"
)

type swipe struct {
	Plain   string
	Meaning string
	Peace   int
}

type facade struct {
	Party      string
	Barrel     any
	Retirement float64
	Tiger      string
	unexported string
}

type offroad struct {
	Surprise []int
	Social   []float64
	Bake     string
	Deprive  string
}

type philosophy struct {
	Deliver string
	Thirsty string
	Job     any
	Peace   map[string]string
	Tire    offroad
}

func main() {
	x := philosophy{
		Deliver: "Something",
		Thirsty: "Juice",
		Job: facade{
			Party:      "Night",
			Barrel:     "Hidden",
			Retirement: 90,
			Tiger:      "Claws",
			unexported: "Inconspicuous",
		},
		Peace: map[string]string{
			"Unfathomable": "Indeed",
			"Joke":         "Batman",
		},
		Tire: offroad{
			Surprise: []int{1, 2, 3, 4},
			Social:   []float64{50.1, 60.2, 70.3},
			Bake:     "Cake",
			Deprive:  "Sleep",
		},
	}

	allStrings := traveller.GetAll[string](x, traveller.MustPath("**", false))
	fmt.Println(allStrings) // [Something Juice Night Hidden Claws Indeed Batman Cake Sleep]
}

```

## Getting

### Multiple Values
`traveller.GetAll[T]` will retrieve all value matching the given path and type.

```go
allPasswords := traveller.GetAll[string](val, traveller.MustPath("**.password", false))
```

### Single Value

`traveller.Get[T]` can be used along with the type that is desired. Will only return the first value matching the given path and type.


```go
password, ok := traveller.Get[string](val, traveller.MustPath("**.password", false))
```

`traveller.MustGet[T]` can also be used to obtain the desired value, but will panic if the value is not found.

```go
password := traveller.MustGet[string](val, traveller.MustPath("**.password", false))
```

## Setting

### Multiple Values
`traveller.SetAll` and `traveller.SetAllBy` will attempt to set all matching values.

If the type is unassignable to that type, then the attempt will be ignored.

```go
changeCount := traveller.SetAll(val, traveller.MustPath("**.password", false), "<hidden>")
```

```go
changeCount := traveller.SetAllBy(val, traveller.MustPath("**.password", false), func(oldVal any) {
	return MyHash(oldVal.(string))
})
```

### Single Value
`traveller.Set` and `traveller.SetBy` will attempt to set the **first successful matching** value.

If the assignment was unsuccessful, it will continue searching.

```go
hasChanged := traveller.Set(val, traveller.MustPath("**.password", false), "<hidden>")
```

```go
hasChanged := traveller.SetBy(val, traveller.MustPath("**.password", false), func(oldVal string) {
	return MyHash(oldVal)
})
```

### Caveat of Setting Values
Due to the nature of Go and some inaddressable values, if a value is deemed inaddressable, the traversed value will be reassigned as a copy on its parent. The resulting edit should still be the same, but please be aware of this little detail/hack.

Also be aware of pointers, especially if the same pointer to a value is unexpectedly used somewhere else.

## Matcher
This is what determines the matching behaviour. You can make your own Matcher by satisfying the following interface:

```go
type Matcher interface {
	// Return true to continue traversing.
	Match(reflect.Value, MatcherSegment) (keepSearching bool)
}
```

The included matchers are:
- `MatchExact`: Exact match along with its type for key (string for field name, int for array/slice index, etc.).
- `MatchPattern`: Match by wildcard pattern. Matching provided by [github.com/gertd/wild](github.com/gertd/wild).
- `MatchMulti`: Recursive matching. Allows free deep traversal.

`Path` and `MustPath` return a `[]traveller.Matcher` and it is the direct type to be used. You can also make your own `[]traveller.Matcher`.

```go
traveller.GetAll[string](val, []traveller.Matcher{traveller.MatchExact{Value: "something"}, traveller.MatchMulti{}})
```

## Options
There are several options that allows manipulation of the traversal behaviour.

```go
traveller.GetAll[string](val, traveller.MustPath("something.**.some*", false),
	traveller.WithIgnoreMaps(true),
	traveller.WithNoFlatEmbeds(true),
	// ...
)
```

- `WithNoFlatEmbeds`: If true, disallows "flattening" of embedded values. Like the following:
```go
// Cannot get by "Value" on this struct because NoFlatEmbeds is set.
// Only "Inner.Value" is allowed.
type Outer struct {
	Inner
}

type Inner struct {
	Value string
}
```
- `WithIgnoreStructs`: Ignores structs on traversal. If the main value is a struct, then it will not search anything.
- `WithIgnoreMaps`: Ignores maps on traversal. If the main value is a map, then it will not search anything.
- `WithIgnoreArrays`: Ignore arrays and slices on traversal. If the main value is an array or a slice, then it will not search anything.

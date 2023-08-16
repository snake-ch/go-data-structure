package list

// Compare returns an integer comparing two type 'T'.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
type Comparator[T any] func(i, j T) int

// Equaler reports whether then element equals the other or not.
type Equaler[T any] func(i, j T) bool

// Less reports whether the element should sort before the other.
type Less[T any] func(i, j T) bool

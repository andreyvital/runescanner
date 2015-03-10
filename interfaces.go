package runescanner

// RuneScanner scans runes in a `io.RuneScanner` and also provides the
// `LookAhead` method which allows you to get the `n-th` char ahead of the
// current position
type RuneScanner interface {
	// ReadRune reads a single UTF-8 encoded Unicode character and returns the
	// rune and its size in bytes. If the encoded rune is invalid, it consumes
	// one byte and returns unicode.ReplacementChar (U+FFFD) with a size of 1.
	ReadRune() (rune, int, error)

	// UnreadRune unreads the last rune. If the most recent read operation on
	// the buffer was not a ReadRune, UnreadRune returns an error. (In this
	// regard it is stricter than UnreadByte, which will unread the last byte
	// from any read operation.)
	UnreadRune() error

	// LookAhead reads the `n-th` rune ahead of the scanners position. It behaves
	// the same way as the `ReadRune` method. The value of `n` is the number of
	// `ReadRune` calls needed to get the rune.
	LookAhead(n int) (rune, int, error)
}

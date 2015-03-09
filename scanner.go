package runescanner

import "io"

func NewRuneScanner(r io.RuneScanner) RuneScanner {
	return &runeScanner{
		reader: r,
		queue:  newQueue(),
	}
}

type runeScanner struct {
	reader io.RuneScanner
	queue  *queue
	last   *result
}

func (s *runeScanner) ReadRune() (rune, int, error) {
	if r := s.queue.shift(); r != nil {
		s.last = r
		return r.r, r.n, r.e
	}

	r := newResult(s.reader.ReadRune())
	s.last = r
	return r.r, r.n, r.e
}

func (s *runeScanner) UnreadRune() error {
	return s.reader.UnreadRune()
}

func (s *runeScanner) LookAhead(n int) (rune, int, error) {
	if n <= 0 && s.last == nil {
		return 0, 1, nil
	}

	if n <= 0 && s.last != nil {
		return s.last.r, s.last.n, s.last.e
	}

	if n < s.queue.size() {
		r := s.queue.at(n - 1)
		return r.r, r.n, r.e
	}

	var r *result
	n -= s.queue.size()
	for n > 0 {
		n--
		r = newResult(s.reader.ReadRune())
		s.queue.push(r)
	}

	return r.r, r.n, r.e
}

func newResult(r rune, n int, e error) *result {
	return &result{
		r: r,
		n: n,
		e: e,
	}
}

type result struct {
	r rune
	n int
	e error
}

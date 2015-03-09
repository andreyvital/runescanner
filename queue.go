package runescanner

func newQueue() *queue {
	return &queue{
		results: make([]*result, 0),
	}
}

type queue struct {
	results []*result
}

func (q *queue) push(r *result) {
	q.results = append(q.results, r)
}

func (q *queue) shift() *result {
	if len(q.results) == 0 {
		return nil
	}

	r := q.results[0]
	l := len(q.results)
	for i := 1; i < l; i++ {
		q.results[i-1] = q.results[i]
	}
	q.results[l-1] = nil
	return r
}

func (q *queue) at(n int) *result {
	if n < 0 || n > len(q.results)-1 {
		return nil
	}

	return q.results[n]
}

func (q *queue) size() int {
	return len(q.results)
}

package runescanner_test

import (
	"strings"

	. "github.com/txgruppi/runescanner"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RuneScanner", func() {
	var scanner RuneScanner

	BeforeEach(func() {
		scanner = NewRuneScanner(strings.NewReader("Tarcísio"))
	})

	Describe("NewRuneScanner", func() {
		It("should return a valid RuneScanner instance", func() {
			Expect(scanner).ToNot(BeNil())
		})
	})

	Describe("ReadRune", func() {
		It("should read all runes in the reader", func() {
			r0 := []rune{'T', 'a', 'r', 'c', 'í', 's', 'i', 'o'}
			r1 := []int{1, 1, 1, 1, 2, 1, 1, 1}

			for index, expected := range r0 {
				r, n, e := scanner.ReadRune()
				Expect(r).To(Equal(expected))
				Expect(n).To(Equal(r1[index]))
				Expect(e).To(BeNil())
			}
		})

		It("should return the correct values even after several lookahead calls", func() {
			scanner.LookAhead(0)
			scanner.LookAhead(1)
			scanner.LookAhead(2)
			scanner.LookAhead(5)

			r0 := []rune{'T', 'a', 'r', 'c', 'í', 's', 'i', 'o'}
			r1 := []int{1, 1, 1, 1, 2, 1, 1, 1}

			for index, expected := range r0 {
				r, n, e := scanner.ReadRune()
				Expect(r).To(Equal(expected))
				Expect(n).To(Equal(r1[index]))
				Expect(e).To(BeNil())
			}
		})
	})

	Describe("UnreadRune", func() {
		It("should unread the last read rune", func() {
			r, n, e := scanner.ReadRune()
			Expect(r).To(Equal('T'))
			Expect(n).To(Equal(1))
			Expect(e).To(BeNil())

			e = scanner.UnreadRune()
			Expect(e).To(BeNil())

			r, n, e = scanner.ReadRune()
			Expect(r).To(Equal('T'))
			Expect(n).To(Equal(1))
			Expect(e).To(BeNil())
		})

		It("should return error if called twice", func() {
			r, n, e := scanner.ReadRune()
			Expect(r).To(Equal('T'))
			Expect(n).To(Equal(1))
			Expect(e).To(BeNil())

			e = scanner.UnreadRune()
			Expect(e).To(BeNil())

			e = scanner.UnreadRune()
			Expect(e).To(MatchError("strings.Reader.UnreadRune: previous operation was not ReadRune"))
		})
	})

	Describe("LookAhead", func() {
		It("should look ahead", func() {
			r, n, e := scanner.LookAhead(0)
			Expect(r).To(BeEquivalentTo(0))
			Expect(n).To(Equal(1))
			Expect(e).To(BeNil())

			r, n, e = scanner.LookAhead(1)
			Expect(r).To(Equal('T'))
			Expect(n).To(Equal(1))
			Expect(e).To(BeNil())

			r, n, e = scanner.LookAhead(2)
			Expect(r).To(Equal('a'))
			Expect(n).To(Equal(1))
			Expect(e).To(BeNil())

			r, n, e = scanner.LookAhead(5)
			Expect(r).To(Equal('í'))
			Expect(n).To(Equal(2))
			Expect(e).To(BeNil())
		})

		It("should return the correct result for n equals to 0", func() {
			r, n, e := scanner.ReadRune()
			Expect(r).To(Equal('T'))
			Expect(n).To(Equal(1))
			Expect(e).To(BeNil())

			r, n, e = scanner.LookAhead(0)
			Expect(r).To(Equal('T'))
			Expect(n).To(Equal(1))
			Expect(e).To(BeNil())

			r, n, e = scanner.ReadRune()
			Expect(r).To(Equal('a'))
			Expect(n).To(Equal(1))
			Expect(e).To(BeNil())

			r, n, e = scanner.LookAhead(0)
			Expect(r).To(Equal('a'))
			Expect(n).To(Equal(1))
			Expect(e).To(BeNil())

			r, n, e = scanner.LookAhead(3)
			Expect(r).To(Equal('í'))
			Expect(n).To(Equal(2))
			Expect(e).To(BeNil())
		})
	})
})

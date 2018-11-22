package slice

type IndexedStringSlice struct {
	s []string
	m map[string]struct{}
}

type IndexedIntSlice struct {
	s []int
	m map[int]struct{}
}

func NewStringSlice(s []string) *IndexedStringSlice {
	m := make(map[string]struct{}, len(s))
	for _, v := range s {
		m[v] = struct{}{}
	}
	return &IndexedStringSlice{
		s: s,
		m: m,
	}
}

func NewIntSlice(s []int) *IndexedIntSlice {
	m := make(map[int]struct{}, len(s))
	for _, v := range s {
		m[v] = struct{}{}
	}
	return &IndexedIntSlice{
		s: s,
		m: m,
	}
}

func (islice *IndexedStringSlice) Append(s string) {
	islice.s = append(islice.s, s)
	islice.m[s] = struct{}{}
}

func (islice *IndexedStringSlice) Has(s string) bool {
	_, ok := islice.m[s]
	return ok
}

func (islice *IndexedStringSlice) Values() []string {
	return islice.s
}

func (islice *IndexedIntSlice) Append(s int) {
	islice.s = append(islice.s, s)
	islice.m[s] = struct{}{}
}

func (islice *IndexedIntSlice) Has(s int) bool {
	_, ok := islice.m[s]
	return ok
}

func (islice *IndexedIntSlice) Values() []int {
	return islice.s
}

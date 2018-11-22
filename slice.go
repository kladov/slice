package slice

type IndexedStringSlice struct {
	s []string
	m map[string]struct{}
}

type IndexedFloat64Slice struct {
	s []float64
	m map[float64]struct{}
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

func NewFloatSlice(s []float64) *IndexedFloat64Slice {
	m := make(map[float64]struct{}, len(s))
	for _, v := range s {
		m[v] = struct{}{}
	}
	return &IndexedFloat64Slice{
		s: s,
		m: m,
	}
}

func (islice *IndexedStringSlice) Append(v string) {
	islice.s = append(islice.s, v)
	islice.m[v] = struct{}{}
}

func (islice *IndexedStringSlice) Has(v string) bool {
	_, ok := islice.m[v]
	return ok
}

func (islice *IndexedStringSlice) Values() []string {
	return islice.s
}

func (islice *IndexedIntSlice) Append(v int) {
	islice.s = append(islice.s, v)
	islice.m[v] = struct{}{}
}

func (islice *IndexedIntSlice) Has(v int) bool {
	_, ok := islice.m[v]
	return ok
}

func (islice *IndexedIntSlice) Values() []int {
	return islice.s
}

func (islice *IndexedFloat64Slice) Append(s float64) {
	islice.s = append(islice.s, s)
	islice.m[s] = struct{}{}
}

func (islice *IndexedFloat64Slice) Has(v float64) bool {
	_, ok := islice.m[v]
	return ok
}

func (islice *IndexedFloat64Slice) Values() []float64 {
	return islice.s
}
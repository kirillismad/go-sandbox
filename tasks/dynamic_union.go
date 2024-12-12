package tasks

type DynamicUnion interface {
	Union(p, q int)
	Find(p int) int
	Connected(p, q int) bool
	Count() int
}

type DynamicUnion1 struct {
	slice []int
	count int
}

func NewDynamicUnion1(n int) *DynamicUnion1 {
	slice := make([]int, n)
	for i := range slice {
		slice[i] = i
	}

	return &DynamicUnion1{
		slice: slice,
		count: n,
	}
}

func (f *DynamicUnion1) Union(p int, q int) {
	pRoot := f.Find(p)
	qRoot := f.Find(q)

	if pRoot == qRoot {
		return
	}

	for i := 0; i < len(f.slice); i++ {
		if f.slice[i] == qRoot {
			f.slice[i] = pRoot
		}
	}
	f.count--
}

func (f *DynamicUnion1) Find(p int) int {
	return f.slice[p]
}

func (f *DynamicUnion1) Connected(p int, q int) bool {
	return f.Find(p) == f.Find(q)
}

func (f *DynamicUnion1) Count() int {
	return f.count
}

type DynamicUnion2 struct {
	slice []int
	count int
}

func NewDynamicUnion2(n int) *DynamicUnion2 {
	slice := make([]int, n)
	for i := range slice {
		slice[i] = i
	}

	return &DynamicUnion2{
		slice: slice,
		count: n,
	}
}

func (f *DynamicUnion2) Union(p int, q int) {
	pRoot := f.Find(p)
	qRoot := f.Find(q)

	if pRoot == qRoot {
		return
	}

	f.slice[qRoot] = pRoot

	f.count--
}

func (f *DynamicUnion2) Find(p int) int {
	for f.slice[p] != p {
		p = f.slice[p]
	}
	return p
}

func (f *DynamicUnion2) Connected(p int, q int) bool {
	return f.Find(p) == f.Find(q)
}

func (f *DynamicUnion2) Count() int {
	return f.count
}

type DynamicUnion3 struct {
	slice []int
	sizes []int
	count int
}

func NewDynamicUnion3(n int) *DynamicUnion3 {
	slice := make([]int, n)
	for i := range slice {
		slice[i] = i
	}
	sizes := make([]int, n)
	for i := range sizes {
		sizes[i] = 1
	}

	return &DynamicUnion3{
		slice: slice,
		sizes: sizes,
		count: n,
	}
}

func (f *DynamicUnion3) Union(p int, q int) {
	pRoot := f.Find(p)
	qRoot := f.Find(q)

	if pRoot == qRoot {
		return
	}

	if f.sizes[pRoot] > f.sizes[qRoot] {
		f.slice[qRoot] = pRoot
		f.sizes[pRoot] += f.sizes[qRoot]
	} else {
		f.slice[pRoot] = qRoot
		f.sizes[qRoot] += f.sizes[pRoot]
	}

	f.count--
}

func (f *DynamicUnion3) Find(p int) int {
	for f.slice[p] != p {
		p = f.slice[p]
	}
	return p
}

func (f *DynamicUnion3) Connected(p int, q int) bool {
	return f.Find(p) == f.Find(q)
}

func (f *DynamicUnion3) Count() int {
	return f.count
}

type DynamicUnion4 struct {
	units map[int]int
	sizes map[int]int
	count int
}

func NewDynamicUnion4(n int) *DynamicUnion4 {
	units := make(map[int]int, n)
	sizes := make(map[int]int, n)

	return &DynamicUnion4{
		units: units,
		sizes: sizes,
		count: n,
	}
}

func (f *DynamicUnion4) ensureExists(p, q int) {
	if _, ok := f.units[p]; !ok {
		f.units[p] = p
		f.sizes[p] = 1
	}
	if _, ok := f.units[q]; !ok {
		f.units[q] = q
		f.sizes[q] = 1
	}
}

func (f *DynamicUnion4) Union(p int, q int) {
	f.ensureExists(p, q)
	pRoot := f.Find(p)
	qRoot := f.Find(q)

	if pRoot == qRoot {
		return
	}

	if f.sizes[pRoot] > f.sizes[qRoot] {
		f.units[qRoot] = pRoot
		f.sizes[pRoot] += f.sizes[qRoot]
	} else {
		f.units[pRoot] = qRoot
		f.sizes[qRoot] += f.sizes[pRoot]
	}

	f.count--
}

func (f *DynamicUnion4) Find(p int) int {
	for f.units[p] != p {
		p = f.units[p]
	}
	return p
}

func (f *DynamicUnion4) Connected(p int, q int) bool {
	if _, ok := f.units[p]; !ok {
		return false
	}
	if _, ok := f.units[q]; !ok {
		return false
	}
	return f.Find(p) == f.Find(q)
}

func (f *DynamicUnion4) Count() int {
	return f.count
}

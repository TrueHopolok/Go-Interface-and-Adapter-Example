package sortable

type Sortable interface {
	before(Sortable) bool
}

func MergeSort(arr []Sortable) {
	var t int
	if t = len(arr); t > 1 {
		h := t / 2
		L, R := make([]Sortable, h), make([]Sortable, t-h)
		_, _ = copy(L, arr[:h]), copy(R, arr[h:])
		MergeSort(L)
		MergeSort(R)
		i, j := 0, 0
		for i < h && j < t-h {
			if L[i].before(R[j]) {
				arr[i+j] = L[i]
				i++
			} else {
				arr[i+j] = R[j]
				j++
			}
		}
		for ; i < h; i++ {
			arr[i+j] = L[i]
		}
		for ; j < t-h; j++ {
			arr[i+j] = R[j]
		}
	}
}

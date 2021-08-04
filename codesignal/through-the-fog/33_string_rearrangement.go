package main

func main() {

}

var perms [][]string

func stringsRearrangement(inputArray []string) bool {

	perms = make([][]string, 0)
	n := len(inputArray)

	permute(inputArray, 0, n-1)
	for _, perm := range perms {
		fail := false
		prev := perm[0]
		for _, str := range perm[1:] {
			if !oneDiff(prev, str) {
				fail = true
			}
			prev = str
		}
		if !fail {
			return true
		}
	}
	return false
}

func permute(arr []string, l, r int) {
	if l == r {
		cp := make([]string, len(arr))
		copy(cp, arr)
		perms = append(perms, cp)
		return
	}
	for i := l; i <= r; i++ {
		swap(arr, l, i)
		permute(arr, l+1, r)
		swap(arr, l, i)
	}
}

func swap(arr []string, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func oneDiff(a, b string) bool {
	dif := 0
	for n, _ := range a {
		if a[n] != b[n] {
			dif++
		}
	}
	return dif == 1
}

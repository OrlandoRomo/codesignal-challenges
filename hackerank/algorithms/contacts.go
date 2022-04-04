package main

func main() {
	queries := [][]string{
		{"add", "hack"},
		{"add", "hackerrank"},
		{"find", "hac"},
		{"find", "hak"},
	}
	contacts(queries)
}

func contacts(queries [][]string) []int32 {
	// Write your code here
	// O(N^2)
	store := make(map[string]int, 0)
	counters := make([]int32, 0)
	for _, query := range queries {
		if query[0] == "add" {
			for i := 1; i <= len(query[1]); i++ {
				contact := query[1]
				_, ok := store[contact[0:i]]
				if ok {
					store[contact[0:i]]++
				}
				if !ok {
					store[contact[0:i]] = 1
				}
			}
		}

		if query[0] == "find" {
			count := store[query[1]]
			counters = append(counters, int32(count))
		}
	}
	return counters
}

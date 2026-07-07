package main

func ForEach(f func(int), a []int) {
	for _, n := range a {
		f(n)
	}
}

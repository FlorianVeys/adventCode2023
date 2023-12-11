package main

func main() {
	time := []int{48938595}
	distance := []int{296192812361391}

	possibility := make([]int, len(time))

	for index, t := range time {
		for i := 0; i < t; i++ {
			if i*(t-i) > distance[index] {
				possibility[index]++
			}
		}
	}

	result := 1

	for _, p := range possibility {
		result = result * p
	}

	println(result)
}
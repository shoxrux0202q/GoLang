package homework



func IsPerfectNumber(num int) bool {
	sum := 0
	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	return sum == num
}

/*func printPerfectNumbers(n int) {
	fmt.Printf("Mukammal sonlar %d gacha:\n", n)
	for i := 1; i <= n; i++ {
		if isPerfectNumber(i) {
			fmt.Println(i)
		}
	}
}*/
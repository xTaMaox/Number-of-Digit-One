func countDigitOne(n int) int {
	var result, k, a, b, c int
	for a, k = n, 1; k <= n; k *= 10 {
		a, b, c = a/10, a%10, n%k
		switch b {
		case 0:
			result += a * k
		case 1:
			result += a*k + c + 1
		default:
			result += a*k + k
		}
	}
	return result
}
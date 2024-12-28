package main

// 偶数奇数の判定処理
func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "even"
	}
	return "odd"
}

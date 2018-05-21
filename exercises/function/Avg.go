package main

func main() {
	//编写一个函数用于计算一个 float64 类型的 slice 的平均值。
}
func average(xs []float64) (avg float64) {

	sum := 0.0
	switch len(xs) {
	case 0:
		avg = 0
	default:
		for _, v := range xs {
			sum += v
		}
		avg = sum / float64(len(xs))
	}
	return .4
}

package mlgo

import "math"

type KnnClassifier struct {
	neighbors int
	X [][]float64
	Y []interface{}
}

func NewKnnClassifier(neighbors int) KnnClassifier {
	var ret KnnClassifier

	ret.neighbors = neighbors

	return ret
}

func (class *KnnClassifier) Fit(trainX [][]float64, trainY []interface{}) {
	if len(trainX) != len(trainY) {
		panic("The length of training data and targets is different")
	}

	class.X = trainX
	class.Y = trainY
}

func getDistance(dist Distance, points1[]float64, points2 []float64) float64 {
	switch dist {
	case Euclidean:
		var total float64 = 0
		for i := 0; i < len(points1);  i++ {
			total += math.Pow(points1[i] - points2[i], 2)
		}
		return math.Sqrt(total)
	default:
		panic("That is not a valid distance choice")
		return 0
	}
}
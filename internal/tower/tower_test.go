package tower

import "testing"

func TestTowerLocationDetection(t *testing.T) {
	var tableTest = []map[string][]int{
		{
			"input":     {1, 2},
			"placement": {2, 4},
			"output":    {-1, -1},
		},
		{
			"input":     {1, 1},
			"placement": {1, 1},
			"output":    {1, 1},
		},
	}

	for _, test := range tableTest {
		inputValue := test["input"]
		allowedTowerLocation := [][]int{test["placement"]}

		x, y := AllowedToPlaceTower(inputValue[0], inputValue[1], allowedTowerLocation)

		outputValue := test["output"]

		if x != outputValue[0] {
			t.Errorf("x value is %d, expected %d", x, outputValue[0])
		}

		if y != outputValue[1] {
			t.Errorf("y value is %d, expected %d", y, outputValue[1])
		}
	}
}

func TestUnitDistance(t *testing.T) {
	var tableTest = []map[string]any{
		{
			"input":  []float64{30, 12, 30, 8},
			"output": 4,
		},
		{
			"input":  []float64{10, 12, 30, 8},
			"output": 20,
		},
	}

	for _, test := range tableTest {
		point := test["input"].([]float64)
		expectResult := test["output"].(int)

		actualResult := euclideanFormula(point[0], point[1], point[2], point[3])

		if expectResult != actualResult {
			t.Errorf("Actual Result is %d, Which System Expect %d", actualResult, expectResult)
		}
	}
}

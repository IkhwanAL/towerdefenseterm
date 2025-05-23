package tower

import "testing"

func TestTowerLocationDetection(t *testing.T) {
	var tableTest = []map[string]interface{}{
		{
			"input":     []int{1, 2},
			"placement": []int{2, 4},
			"output":    false,
		},
		{
			"input":     []int{1, 1},
			"placement": []int{1, 1},
			"output":    true,
		},
	}

	for _, test := range tableTest {
		inputValue := test["input"].([]int)
		allowedTowerLocation := [][]int{test["placement"].([]int)}

		allowed, _ := AllowedToPlaceTower(inputValue[0], inputValue[1], allowedTowerLocation)

		outputValue := test["output"].(bool)

		if allowed != outputValue {
			t.Errorf("value is %v, expected %v", allowed, outputValue)
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

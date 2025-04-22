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

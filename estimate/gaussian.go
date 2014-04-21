package estimate

import (
	"../util"
	"fmt"
)

func main() {
	params := EstimateGaussian("test.csv")
	fmt.Println(params)
}

func EstimateGaussian(filename string) map[string]map[string]float64 {
	result := util.LoadCSV(filename)
	params := make(map[string]map[string]float64, len(result))
	for var_name, values := range result {
		params[var_name] = make(map[string]float64, 2)
		mean, std := CalcMeanAndStd(values)
		params[var_name]["mean"] = mean
		params[var_name]["std"] = std
	}
	return params
}

func CalcMeanAndStd(values []float64) (mean, std float64) {
	mean = util.CalcMean(values)
	std = util.CalcStd(values, mean)
	return
}

package estimate

import (
	"../util"
)

func EstimateAR(filename string, order int) map[string]map[string]float64 {
	result := util.LoadCSV(filename)
	params := make(map[string]map[string]float64, len(result))
	/*	for var_name, values := range result {
			params[var_name] = make(map[string]float64, 2)
			hoge := util.CalcCoefficients(values, order)
		}
	*/
	return params
}

package util

import (
	"encoding/csv"
	"fmt"
	//	matrix "github.com/skelterjohn/go.matrix"
	"io"
	"math"
	"os"
	"strconv"
)

func LoadCSV(filename string) map[string][]float64 {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer file.Close()
	reader := csv.NewReader(file)

	// Read fist line (variable names)
	variable_names, err := reader.Read()
	num_variable := len(variable_names)
	variable_values := make([][]float64, num_variable)

	for i := 0; i < num_variable; i++ {
		variable_values[i] = nil
	}

	for {
		record, read_err := reader.Read()
		if read_err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		for k := 0; k < len(record); k++ {
			value, value_err := strconv.ParseFloat(record[k], 64)
			if value_err == io.EOF {
				break
			}
			variable_values[k] = append(variable_values[k], value)
		}
	}

	variable_values_map := make(map[string][]float64)
	for i := 0; i < len(variable_names); i++ {
		var_name := variable_names[i]
		variable_values_map[var_name] = variable_values[i]
	}
	return variable_values_map
}

func CalcMean(values []float64) (mean float64) {
	sum := 0.0
	for _, value := range values {
		sum += value
	}
	mean = sum / float64(len(values))
	return
}

func CalcStd(values []float64, optional ...float64) (std float64) {
	mean := 0.0
	if len(optional) == 1 {
		mean = optional[0]
	} else {
		mean = CalcMean(values)
	}
	sum := 0.0
	for _, value := range values {
		sum += math.Pow(math.Abs(value-mean), 2)
	}
	std = math.Sqrt(sum / float64(len(values)))
	return
}

func CalcCoefficients(values []float64, order int) (coef []float64) {
	//	n := len(values) - order
	//	m := order
	return nil
}

func CalcCovariance(values []float64, order, k, index int) (cov float64) {
	cov = 0.0
	for s := order; s < 2*order; s++ {
		cov += values[s-index] * values[s-k]
	}
	return cov
}

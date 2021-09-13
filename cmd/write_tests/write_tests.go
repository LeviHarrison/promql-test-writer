package main

import (
	"fmt"
	"math"
)

type Metric struct {
	name   string
	series []Series
}

type Series struct {
	labels string
	value  float64
}

type Tests struct {
	functionNames []string
	functions     []func(float64) float64
	modifierNames []string
	modifiers     []float64
}

var metric = Metric{name: "trig", series: []Series{{labels: `l="x"`, value: 10}, {labels: `l="y"`, value: 20}, {labels: `l="NaN"`, value: math.NaN()}}}

var tests = Tests{
	functionNames: []string{"asinh", "acosh", "atanh"},
	functions:     []func(float64) float64{math.Asinh, math.Acosh, math.Atanh},
	modifierNames: []string{"- 10.1"},
	modifiers:     []float64{-10.1},
}

func main() {
	for i := 0; i < len(tests.functionNames); i++ {
		for j := 0; j < len(tests.modifiers); j++ {
			fmt.Printf("eval instant at 5m %s(%s%s)\n", tests.functionNames[i], metric.name, tests.modifierNames[j])

			for k := 0; k < len(metric.series); k++ {
				fmt.Println(fmt.Sprintf("	{%s}", metric.series[k].labels), tests.functions[i](metric.series[k].value+tests.modifiers[j]))
			}

			fmt.Print("\n")
		}
	}
}

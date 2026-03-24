
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Generator struct {
	Seed int64
}

func NewGenerator() *Generator {
	return &Generator{Seed: time.Now().UnixNano()}
}

func (g *Generator) GenerateTimeSeries(length int) []float64 {
	rand.Seed(g.Seed)
	data := make([]float64, length)
	current := 10.0
	for i := 0; i < length; i++ {
		current += rand.Float64()*2 - 1
		data[i] = current
	}
	return data
}

func main() {
	gen := NewGenerator()
	series := gen.GenerateTimeSeries(10)
	fmt.Println("Generated Synthetic Time Series:")
	for _, v := range series {
		fmt.Printf("%.2f ", v)
	}
	fmt.Println()
}

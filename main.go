package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	gene         = "trovami"
	genomaSet    = "abcdefghijklmnopqrstuvwxyz"
	geneLength   = len(gene)
	popNum       = 15
	mutationRate = 0.04
)

func initPopulation(popNum int, genomaSet string, geneLength int) [][]byte {
	population := make([][]byte, popNum)
	for i := 0; i < popNum; i++ {
		population[i] = createIndividual(genomaSet, geneLength)
	}
	return population
}

func createIndividual(genomaSet string, geneLength int) []byte {
	individual := make([]byte, geneLength)
	for i := 0; i < geneLength; i++ {
		individual[i] = genomaSet[rand.Intn(len(genomaSet))]
	}
	return individual
}

func computeFitness(gene string, population [][]byte) []int {
	fitness := make([]int, len(population))
	for i, individual := range population {
		for j, c := range gene {
			if byte(c) == individual[j] {
				fitness[i]++
			}
		}
	}
	return fitness
}

func maxFitness(fitness []int) (int, int) {
	max := -1
	index := -1
	for i, fit := range fitness {
		if fit >= max {
			max = fit
			index = i
		}
	}
	return index, max
}

func secondFitness(maxId int, fitness []int) (int, int) {
	max := -1
	index := -1
	for i, fit := range fitness {
		if maxId != i && fit >= max {
			max = fit
			index = i
		}
	}
	return index, max
}

func leastFitness(maxId, secondId int, fitness []int) int {
	min := geneLength
	index := -1
	for i, fit := range fitness {
		if maxId != i && secondId != i && fit <= min {
			min = fit
			index = i
		}
	}
	return index
}

func crossover(parent1, parent2 []byte) ([]byte, []byte) {
	crossoverPoint := rand.Intn(len(parent1))
	child1 := append(parent1[:crossoverPoint], parent2[crossoverPoint:]...)
	child2 := append(parent2[:crossoverPoint], parent1[crossoverPoint:]...)
	return child1, child2
}

func mutation(mutationRate float64, offspring1, offspring2 []byte) ([]byte, []byte) {
	for i := range offspring1 {
		if rand.Float64() < mutationRate {
			offspring1[i] = genomaSet[rand.Intn(len(genomaSet))]
			offspring2[i] = genomaSet[rand.Intn(len(genomaSet))]
		}
	}
	return offspring1, offspring2
}

func main() {
	rand.Seed(time.Now().UnixNano())

	startTime := time.Now()

	generation := 0
	population := initPopulation(popNum, genomaSet, geneLength)
	fitness := computeFitness(gene, population)

	for {
		generation++
		parent1Id, _ := maxFitness(fitness)
		parent2Id, _ := secondFitness(parent1Id, fitness)

		offspring1, offspring2 := crossover(population[parent1Id], population[parent2Id])
		offspring1, offspring2 = mutation(mutationRate, offspring1, offspring2)

		fitness = computeFitness(gene, population)

		grandfatherToKillIfIsntDeadBySeniorityId := leastFitness(parent1Id, parent2Id, fitness)

		best1Id, _ := maxFitness(fitness)
		best2Id, _ := secondFitness(parent1Id, fitness)

		var children []byte
		if fitness[best1Id] >= fitness[best2Id] {
			children = offspring1
		} else {
			children = offspring2
		}

		for i := range children {
			population[grandfatherToKillIfIsntDeadBySeniorityId][i] = children[i]
		}

		fitness = computeFitness(gene, population)
		maxFitnessId, _ := maxFitness(fitness)

		fmt.Printf("%sttGeneration n° %d\n", string(population[maxFitnessId]), generation)

		if fitness[maxFitnessId] == geneLength {
			break
		}
	}

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	fmt.Printf("Il codice è stato eseguito in %s\n", executionTime)
}

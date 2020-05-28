package main

import (
	"fmt"
	"math/rand"
	"math"
	"github.com/jszwec/csvutil"
	"os"
	"encoding/csv"
	"bufio"
	"github.com/dustin/go-humanize"
	"randomwalk"
)
func main() {
	var popN int
	fmt.Println("Enter size of population: ")
	fmt.Scanln(&popN)
	population := createPop(popN)
	fmt.Println(population)
	fmt.Println("")
	fmt.Println("The BMI of the current population is", 
		bmiPop(population))
	fmt.Println("The average age of the population is", 
		averageAge(population))
	fmt.Printf("The average salary of the population is $%s \n", 
		humanize.Commaf(averageSalary(population)))
	popWalk(population)
	population[1].takeDamage(5.2)
	fmt.Println(population[1].Health)
	createPopCSV(population)
}

type person struct{
	Name string
	Gender string
	Age int
	Height int
	Weight float32
	Occupation string
	X int
	Y int
	Health float32
}

func (p *person) talk() string {
	return string ("Hi, my name is " + p.Name)
}
func (p *person) salary() int {
	if p.Occupation == "doctor" {
		return 100000
	}
	if p.Occupation == "student" {
		return 0
	}
	if p.Occupation == "finance" {
		return 60000
	}
	if p.Occupation == "engineer" {
		return 50000
	} 
	return 25000
}

func (p *person) takeDamage(hit float32) {
	p.Health -= hit
}

func (p *person) Walk() (int, int) {
	dx, dy := randomwalk.Walk(10)
	p.X += dx
	p.Y += dy
	// fmt.Printf("x = %v, y = %v \n",p.x,p.y)
	return p.X,p.Y
}

func newPerson() person {
	uh := 190 //upper height in cm
	lh := 130 //lower height in cm
	uw := 100 //upper weight in kg
	lw := 45 //lower weight in kg
	genders := []string{"Male","Female","Other"}
	occupations := []string{"student","doctor","nurse","barista","finance",
	"teacher","engineer","marketing"}
	mNames := [] string{"Alex","Eddy","Mo","Fahim","Alexio","Roberto", 
	"Daniel", "Ishaq"}
	fNames := [] string{"Tham", "Alexus", "Ariana", "Rebecca", "Maryam","Jo"}
	chosenGender := genders[rand.Intn(3)]
	var chosenName string
	if chosenGender == "Male" {
		chosenName = mNames[rand.Intn(len(mNames))]
	} else if chosenGender == "Female" {
		chosenName = fNames[rand.Intn(len(fNames))]
	} else {
		allNames := []string{}
		allNames = append(mNames,fNames...)
		chosenName = allNames[rand.Intn(len(allNames))]
	}
	p := person{Name:chosenName,
	Gender: chosenGender,
	Age: rand.Intn(70), 
	Height: rand.Intn(uh - lh) + lh, 
	Weight: rand.Float32()*(float32(uw-lw)) + float32(lw),
	Occupation: occupations[rand.Intn(len(occupations))],
	X: 0, Y: 0, Health: 100.0}
	return p

}
func createPop(size int) []person {
	pop := [] person {}

	me := person{Name:"Sayed", Gender: "Male",Age: 21,Height: 161, Weight: 61.4,
	Occupation: "student", X: 0, Y: 0, Health: 100.0}
	pop = append(pop,me)

	for i:=1; i<size; i++{
		pop = append(pop, newPerson())
	}
	return pop
}

func bmiPop(population []person) float64 {
	bmi := 0.0
	for i := range population {
		w := float64(population[i].Weight)
		h := float64(population[i].Height)/100
		bmi += w/(math.Pow(2,h))
	}
	return bmi/float64(len(population))
}

func averageAge(population []person) float64 {
	run := 0.0
	for i := range population {
		run += float64(population[i].Age)
	}
	return run/float64(len(population))
}

func averageSalary(population []person) float64 {
	run := 0.0
	for i := range population {
		run += float64(population[i].salary())
	}
	return run/float64(len(population))
}

func popWalk(population []person) {
	for i := range population {
		population[i].Walk()
	}
}
func filterNamePop(population[] person, name string) [] person{
	var retPop []person
	for i := range(population) {
		if population[i].Name == name {
			retPop = append(retPop, population[i])
		}
	} 
	return retPop
}

func createPopCSV(population []person) {
	rows, err := csvutil.Marshal(population)
	if err != nil {
		fmt.Println("Error writing csv:", err)
	}
	csvFile, err := os.Create("population.csv")
	w := bufio.NewWriter(csvFile)
	_, err = w.Write(rows)
	w.Flush()
	// w := csv.NewWriter(csvFile)
	// w.Write(rows)
	// w.Flush()

}
func createCSV() {
	csvFile, err := os.Create("testing.csv")
	w := csv.NewWriter(csvFile)
	if err != nil {
		fmt.Println("Error writing csv:", err)
	}
	testString := []string {"fads","2","fasdfdsaf","fasdfakf","da"}
	w.Write(testString)
	w.Flush()

}

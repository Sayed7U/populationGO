package main

import (
	"fmt"
	"math/rand"
	"math"
	"github.com/jszwec/csvutil"
	"os"
	"encoding/csv"
	"strings"
)

func main() {
	population := createPop(50)
	fmt.Println(population)
	fmt.Println("")
	fmt.Println("The BMI of the current population is", 
		bmiPop(population))
}

type person struct{
	name string
	gender string
	age int
	height int
	weight float32
	occupation string
}

func (p *person) talk() string {
	return string ("Hi, my name is " + p.name)
}


func newPerson() person {
	uh := 190 //upper height in cm
	lh := 130 //lower height in cm
	uw := 100 //upper weight in kg
	lw := 45 //lower weight in kg
	genders := []string{"Male","Female","Other"}
	occupations := []string{"student","doctor","nurse","barista","finance",
	"teacher","engineer","marketing"}
	mNames := [] string{"Alex","Eddy","Mo","Fahim","Alexio","Roberto", "Daniel"}
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
	p := person{name:chosenName,
	gender: chosenGender,
	age: rand.Intn(70), 
	height: rand.Intn(uh - lh) + lh, 
	weight: rand.Float32()*(float32(uw-lw)) + float32(lw),
	occupation: occupations[rand.Intn(len(occupations))]}
	return p

}
func createPop(size int) []person {
	pop := [] person {}

	me := person{name:"Sayed", gender: "Male",age: 21, height: 161, weight: 61.4,
	occupation: "student"}
	pop = append(pop,me)

	for i:=1; i<size; i++{
		pop = append(pop, newPerson())
	}
	return pop
}

func bmiPop(population []person) float64 {
	bmi := 0.0
	for i := range population {
		w := float64(population[i].weight)
		h := float64(population[i].height)/100
		bmi += w/(math.Pow(2,h))
	}
	return bmi/float64(len(population))
}

func createPopCSV(population []person) {
	rows, err := csvutil.Marshal(population)
	if err != nil {
		fmt.Println("Error writing csv:", err)
	}
	// fmt.Println(rows)
	csvFile, err := os.Create("population.csv")
	w := csv.NewWriter(csvFile)
	w.Write(strings.Split(string(rows)," "))
	w.Flush()

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
// func toCSVFile(population []person) {
// 	w := csv.NewWriter(os.Stdout)

// 	for _, person := range population {
// 			headers := []string {"name",
// 	"gender",
// 	"age",
// 	"height",
// 	"weight",
// 	"occupation"}
// 		if err := w.Write(headers); err != nil{
// 			log.Fatalln("Error writing to csv:", err)
// 		}
// 	}

// 	w.Flush()

// 	if err := w.Error(); err != nil {
// 		log.Fatal(err)
// 	}
// }

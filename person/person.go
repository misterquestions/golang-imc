package person

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Gender byte

const (
	Female Gender = 'M'
	Male Gender = 'H'
)

type Person struct {
	name string
	age uint8
	nss string
	gender Gender
}

func NewPerson(name string, age uint8, gender Gender) *Person {
	var nss = generateRandomNSS()
	return &Person{ name, age, nss, gender }
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) GetAge() uint8 {
	return p.age
}

func (p *Person) GetNSS() string {
	return p.nss
}

func (p *Person) GetGender() Gender {
	return p.gender
}

func (p *Person) CalculateBMI(weight float64, height float64) int8 {
	var bmi = weight / math.Pow(height, 2)

	if (p.gender == Female && bmi < 19) || (p.gender == Male && bmi < 20) {
		return -1
	}

	if (p.gender == Female && bmi > 24) || (p.gender == Male && bmi > 25) {
		return 1
	}

	return 0
}

func (p *Person) IsAdult() bool {
	if p.age >= 18 {
		return true
	}

	return false
}

// Helps checking if the person's gender matches with given one
func (p *Person) checkGender(gender Gender) bool {
	return p.gender == gender
}

// Prints all the object information
func (p *Person) ToString() string {
	return fmt.Sprintf("La persona %s (%s) con el NSS %s tiene %s años", p.name, p.gender, p.nss, p.age)
}

// Generates a random string given the following characters
var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randomString(length int) string {
	// Ensure some randomness
	rand.Seed(time.Now().UnixNano())

	// Generate random string
	builder := make([]rune, length)

	for count := range builder {
		builder[count] = letters[rand.Intn(len(letters))]
	}

	return string(builder)
}

// Generates a random NSS using the above string generator
func generateRandomNSS() string {
	return randomString(6)
}

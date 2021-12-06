package main

import "fmt"

type Person struct {
	name string
	id   int64
}
type Cat struct {
	name  string
	id    int64
	owner *Person
}

//Pet is an interface to define all the common attributes of Pets
type Pet interface {
	walk()
	run()
	play()
	bite(interface{})
	feed()
	setOwner(*Person)
	getOwner() Person
}

func (person *Person) feedPet(pet Pet) {
	fmt.Printf("%v is feeding %v\n", person.name, pet)
	pet.feed()
}

func (c *Cat) walk() {
	fmt.Printf("%v is walking\n", c.name)

}
func (c *Cat) run() {
	fmt.Printf("%v is running\n", c.name)

}
func (c *Cat) feed() {
	fmt.Printf("%v was fed\n", c.name)

}
func (c *Cat) play() {
	fmt.Printf("%v was fed\n", c.name)

}
func (c *Cat) bite(interface{}) {
	fmt.Printf("%v was fed\n", c.name)

}
func (c *Cat) setOwner(p *Person) {
	c.owner = p
	fmt.Printf("'%v' owns '%v' now\n", c.owner.name, c.name)

}

func (c *Cat) getOwner() Person {
	return *c.owner
}

func String()

func main() {
	p1 := Person{
		name: "moda", id: 11,
	}
	p2 := Person{
		name: "Noone", id: 12,
	}
	c1 := Cat{name: "Mishou", id: 1, owner: &p1}
	c2 := &c1
	c1.walk()
	c1.run()
	c1.bite(p1)
	p1.feedPet(c2)
	c1.setOwner(&p2)

	fmt.Println("This is a message")
}

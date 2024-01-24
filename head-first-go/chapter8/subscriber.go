package main

import "fmt"

type Subscriber struct {
	name    string
	rate    float64
	active  bool
	Address // anonymous field
}

type Address struct {
	city    string
	street  string
	country string
}

func defaultSubscriber(name string) *Subscriber {
	var newSub Subscriber
	newSub.name = name
	newSub.rate = 9.9
	newSub.active = true
	newSub.city = "Athens"

	return &newSub
}

func printInfo(sub *Subscriber) {
	fmt.Println("Sub name", sub.name, ", rate", sub.rate, ", isActive", sub.active)
	fmt.Printf("Sub address %#v\n", sub.Address)
}

func applyDiscount(sub *Subscriber) {
	sub.rate = 4.99
}

func main() {
	sub := defaultSubscriber("Nikolas")
	printInfo(sub)
	applyDiscount(sub)
	printInfo(sub)

	var sub2 Subscriber = Subscriber{name: "Nick", rate: 10, active: true}

	printInfo(&sub2)
	applyDiscount(&sub2)
	printInfo(&sub2)

}

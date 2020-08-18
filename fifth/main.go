package main //Change it to main

import "fmt"

// Using Struct
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

// Using Method For User
func (user User) display() string {
	result := fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
	return result
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

// Using Method For Group
func (group Group) displayGroup() {
	fmt.Printf("Name Group : %s", group.Name)

	fmt.Printf("\n")

	fmt.Printf("Member Group : %d", len(group.Users))

	fmt.Printf("\n")

	for _, user := range group.Users {
		fmt.Println("-", user.FirstName)
	}
}
func main() {

	// Method 1 Use Struct
	user := User{
		ID:        1,
		FirstName: "Rudolf",
		LastName:  "Darwin",
		Email:     "rudolf@darwin.com",
		IsActive:  true,
	}
	result := user.display()
	fmt.Println("With Method : ")
	fmt.Println(result)

	// Method 2 Use Struct
	user2 := User{}
	user2.ID = 2
	user2.FirstName = "Emilia"
	user2.LastName = "Parker"
	user2.Email = "emilia@parker.com"
	user2.IsActive = true

	// How to Output Struct With Method
	result = user2.display()
	fmt.Println("With Method : ")
	fmt.Println(result)

	//How to Output Display User With Function
	display1 := displayUser(user)
	display2 := displayUser(user2)
	fmt.Println("With Function : ")
	fmt.Println(display1)
	fmt.Println(display2)

	//How to Output Display Group With Method
	users := []User{user, user2}
	group := Group{"Gamer PlayGround", user, users, true}
	fmt.Println("With Method : ")
	group.displayGroup()

	// How to Output Display Group With Function
	users1 := []User{user, user2}
	group1 := Group{"Gamer PlayGround 2", user, users1, true}
	fmt.Println("With Function : ")
	displayGroup(group1)

}

// function For Struct Group
func displayGroup(group Group) {
	fmt.Printf("Name : %s", group.Name)
	fmt.Printf("\n")
	fmt.Printf("Member Group : %d", len(group.Users))
	fmt.Printf("\n")
	for _, user := range group.Users {
		fmt.Println("-", user.FirstName)
	}
}

// Function For Struct User
func displayUser(user User) string {
	result := fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
	return result
}

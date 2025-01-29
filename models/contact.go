package models

type Contact struct {
	ID    int
	Name  string
	Email string
	Role  string 
}

var Contacts = []Contact{
	{ID: 1, Name: "John Wick", Email: "John@gmail.com", Role: "Admin"},
	{ID: 2, Name: "Kratos", Email: "Kratos@gmail.com", Role: "User"},
}

func GetAllContacts() []Contact {
	return Contacts
}

func AddContact(contact Contact) {
	contact.ID = len(Contacts) + 1
	Contacts = append(Contacts, contact)
}
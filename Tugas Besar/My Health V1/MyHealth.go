package main

const (
	MAX_PATIENT  int = 100
	MAX_DOCTOR   int = 100
	MAX_CHOICE   int = 10
	MAX_RESPONSE int = 100
	MAX_QUESTION int = 1000
)

type User struct {
	Name           string
	Age            int
	Email          string
	Password       string
	Role           string
	Specialization string
}

type ChoiceElement struct {
	Caption string
	Key     string
}

type Menu struct {
	Title   string
	Choices [MAX_CHOICE]ChoiceElement
	Prompt  string
}

type Response struct {
	Message string
	User    User
}
type Consultation struct {
	Question    string
	AskedBy     User
	Tags        string
	Responses   [MAX_RESPONSE]Response
	NumResponse int
}

var PatientList [MAX_PATIENT]User
var DoctorList [MAX_DOCTOR]User
var ConsultationList [MAX_QUESTION]*Consultation

var loggedInUser User
var registeredPatient int
var registeredDoctor int
var postedConsulation int

func main() {

	showMenu("main")
}

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type CV struct {
	PersonInfo struct {
		Name     string `json:"name"`
		Surname  string `json:"surname"`
		Birthday string `json:"birthday"`
		City     string `json:"city"`
		Profile  string `json:"profile"`
	} `json:"person_info"`

	Position   string   `json:"position"`
	SoftSkills []string `json:"soft_skills"`
	HardSkills []string `json:"hard_skills"`
	Gaols      string   `json:"gaols"`
	Hobby      []string `json:"hobby"`

	ContactInformation struct {
		PhoneNumber string `json:"phone_number"`
		Email       string `json:"email"`
		Github      string `json:"github"`
		Linkedin    string `json:"linkedin"`
	} `json:"contact_information"`

	WorkExperience []struct {
		Company        string   `json:"company"`
		Position       string   `json:"position"`
		City           string   `json:"city"`
		StartDate      string   `json:"start_date"`
		EndDate        string   `json:"end_date"`
		Description    string   `json:"description"`
		Responsibility []string `json:"responsibility"`
	} `json:"work_experience"`

	Education []struct {
		University string `json:"university"`
		Faculty    string `json:"faculty"`
		Specialty  string `json:"specialty"`
		StartDate  string `json:"start_date"`
		EndDate    string `json:"end_date"`
		Degree     string `json:"degree"`
	} `json:"education"`

	AdditionalEducation []string `json:"additional_education"`

	Language []struct {
		Language   string `json:"language"`
		Nativeness string `json:"nativeness"`
		Level      string `json:"level"`
	} `json:"language"`
}

var (
	cvData       CV
	buffer       bytes.Buffer
	env          *string
	webDirectory = "./web"
	goDirectory  = "./swp-generator"
)

func main() {
	env = flag.String("environment", "prod", "define environment (dev or prod)")
	flag.Parse()

	if *env == "dev" {
		webDirectory = "../.././web"
		goDirectory = "../."
	}

	cvJSON, err := ioutil.ReadFile(goDirectory + "/cv.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	err = json.Unmarshal(cvJSON, &cvData)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	tmpl, err := template.ParseFiles(goDirectory + "/template.html")
	if err != nil {
		log.Fatal("Error during template parsing: ", err)
	}

	err = tmpl.Execute(&buffer, cvData)
	if err != nil {
		log.Fatal("Error during template execution: ", err)
	}

	err = os.WriteFile(webDirectory+"/index.html", buffer.Bytes(), 0644)
	if err != nil {
		log.Fatal("Error during opening index.html: ", err)
	}

	log.Println("Successfully created index.html")
}

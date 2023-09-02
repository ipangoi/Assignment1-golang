package model

type Person struct {
	ID         int    `json:"id,string"`
	Code       string `json:"student_code"`
	Name       string `json:"student_name"`
	Address    string `json:"student_address"`
	Occupation string `json:"student_occupation"`
	Reason     string `json:"joining_reason"`
}

type Participant struct {
	Participant []Person `json:"participants"`
}

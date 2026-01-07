package models

type Certificate struct {
	StudentName string `json:"student_name"`
	RollNumber  string `json:"roll_number"`
	Course      string `json:"course"`
	University  string `json:"university"`
	Year        string `json:"year"`
}

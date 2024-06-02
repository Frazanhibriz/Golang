package main

import (
	"fmt"
)

func postNewQuestion() {
	if postedConsulation == MAX_QUESTION {
		fmt.Println("Pertanyaan baru tidak dapat dikirim. Penuh")
		return
	}
	p := &Consultation{
		AskedBy: loggedInUser,
	}

	formTitle("KIRIM PERTANYAAN BARU")
	formInput("Pertanyaan", &p.Question)
	formInput("Tag", &p.Tags)
	ConsultationList[postedConsulation] = p

	fmt.Printf("%v\n", ConsultationList[postedConsulation])

	postedConsulation++
}

func questionList() {
	if postedConsulation == 0 {
		fmt.Println("Belum ada pertanyaan yang terkirim")
		return
	}

	for i := 0; i < postedConsulation && ConsultationList[i].Question != ""; i++ {
		fmt.Printf("%d. Pertanyaan: %s ; Penanya: %s\n", (i + 1), ConsultationList[i].Question, ConsultationList[i].AskedBy.Name)
	}

	pilih := 1
	for pilih > 0 {
		fmt.Print("Pilih pertanyaan utk lebih detil atau ketik '0' utk kembali ke menu sebelumnya : ")
		fmt.Scan(&pilih)

		if pilih > 0 {
			questionDetail(ConsultationList[pilih-1])
		}
	}
}

func sortQuestionByTag(direction string) {
	var minTag, maxTag string
	var idx int

	//selection sort
	for i := 0; i < postedConsulation; i++ {
		minTag = ConsultationList[i].Tags
		maxTag = ConsultationList[i].Tags
		idx = i

		for j := i + 1; j < postedConsulation; j++ {
			if direction == "asc" {
				if minTag > ConsultationList[j].Tags {
					minTag = ConsultationList[j].Tags
					idx = j
				}
			} else if direction == "desc" {
				if maxTag < ConsultationList[j].Tags {
					maxTag = ConsultationList[j].Tags
					idx = j
				}
			}
		}

		if idx != i {
			ConsultationList[i], ConsultationList[idx] = ConsultationList[idx], ConsultationList[i]
		}
	}

	questionList()
}

func questionDetail(c *Consultation) {
	var pilih string

	printBorder("=")
	fmt.Printf("Pertanyaan: %s\n", c.Question)
	fmt.Printf("Nama Penanya: %s\n", c.AskedBy.Name)
	fmt.Printf("Umur Penanya: %d\n", c.AskedBy.Age)
	fmt.Printf("Tag: %s\n", c.Tags)
	for i := 0; i < MAX_RESPONSE && c.Responses[i].Message != ""; i++ {
		printBorder("-")
		if c.Responses[i].User.Role == "Dokter" {
			fmt.Printf("Tanggapan dr. %s :\n%s\n", c.Responses[i].User.Name, c.Responses[i].Message)
		} else {
			fmt.Printf("Tanggapan Tn/Ny. %s :\n%s\n", c.Responses[i].User.Name, c.Responses[i].Message)
		}
	}
	formInput("Apakah Anda ingin memberi tanggapan baru (y/n) ?", &pilih)
	if pilih == "y" || pilih == "yes" {
		postNewResponse(c)
	}
}

func postNewResponse(c *Consultation) {
	r := Response{
		User: loggedInUser,
	}

	formTitle("KIRIM TANGGAPAN BARU")
	formInput("Tanggapan", &r.Message)
	c.Responses[c.NumResponse] = r
	c.NumResponse++

	fmt.Println("Tanggapan Anda sudah terkirim")
}

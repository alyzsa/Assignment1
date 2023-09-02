package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Peserta struct {
	ID            string `json:"id"`
	StudentCode   string `json:"student_code"`
	Nama          string `json:"student_name"`
	Alamat        string `json:"student_address"`
	Pekerjaan     string `json:"student_occupation"`
	AlasanGabung  string `json:"joining_reason"`
}

type DataPeserta struct {
	Peserta []Peserta `json:"participants"`
}

func main() {
	kodePeserta := os.Args[1]
	dataFile, err := ioutil.ReadFile("participant.json")
	

	var data DataPeserta
	err = json.Unmarshal(dataFile, &data)
	if err != nil {
		fmt.Println("Gagal, coba lagi", err)
		return
	}

	var pesertaDitemukan Peserta
	for _, peserta := range data.Peserta {
		if peserta.StudentCode == kodePeserta {
			pesertaDitemukan = peserta
			break
		}
	}

	if pesertaDitemukan.ID != "" {
		fmt.Printf("Data peserta dengan kode peserta %s:\n", kodePeserta)
		fmt.Printf("ID: %s\n", pesertaDitemukan.ID)
		fmt.Printf("Kode Peserta: %s\n", pesertaDitemukan.StudentCode)
		fmt.Printf("Nama: %s\n", pesertaDitemukan.Nama)
		fmt.Printf("Alamat: %s\n", pesertaDitemukan.Alamat)
		fmt.Printf("Pekerjaan: %s\n", pesertaDitemukan.Pekerjaan)
		fmt.Printf("Alasan Gabung: %s\n", pesertaDitemukan.AlasanGabung)
	} else {
		fmt.Printf("Tidak ada peserta dengan kode peserta %s\n", kodePeserta)
	}
}

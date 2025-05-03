package main

import (
	"fmt"
	"time"
)

func main() {
	jalaninLirik()
}

func jalaninLirik() {
	lirik := []struct {
		baris            string
		delayPerKarakter time.Duration
		delayPerBaris    time.Duration
	}{
		{"Jika tidak hari ini", 140 * time.Millisecond, 550 * time.Millisecond},
		{"Mungkin minggu depan", 140 * time.Millisecond, 550 * time.Millisecond},
		{"Jika tidak minggu ini", 140 * time.Millisecond, 550 * time.Millisecond},
		{"Mungkin bulan depan", 140 * time.Millisecond, 550 * time.Millisecond},
		{"Jika tidak bulan ini", 140 * time.Millisecond, 600 * time.Millisecond},
		{"Mungkin tahun depan", 140 * time.Millisecond, 550 * time.Millisecond},
		{"Segala harapan kan datang", 120 * time.Millisecond, 450 * time.Millisecond},
		{"Yang kita impikan", 140 * time.Millisecond, 700 * time.Millisecond},

		// Chorus
		{"Janganlah menyerah dulu", 130 * time.Millisecond, 450 * time.Millisecond},
		{"Waktu masih panjang", 130 * time.Millisecond, 450 * time.Millisecond},
		{"Ingat doa kita selalu", 130 * time.Millisecond, 450 * time.Millisecond},
		{"Yang tak pernah usang", 140 * time.Millisecond, 450 * time.Millisecond},
		{"Kita usahakan lagi sayang", 150 * time.Millisecond, 800 * time.Millisecond},
		{"Yakin waktunya kan datang", 200 * time.Millisecond, 2000 * time.Millisecond},
	}

	fmt.Print("\033[1;97m")

	fmt.Println("\n🎵 == Kita Usahakan Lagi - Batas Senja == 🎵\n")

	for i, baris := range lirik {
		if i == 8 {
			fmt.Println()
		}

		for _, karakter := range baris.baris {
			fmt.Print(string(karakter))
			time.Sleep(baris.delayPerKarakter)
		}
		fmt.Println()
		time.Sleep(baris.delayPerBaris)
	}
	fmt.Print("\033[0m")

	fmt.Println("\n// Code by Digminds✨")
}

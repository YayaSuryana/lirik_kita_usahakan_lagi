# 🎶 Terminal Lirik: *Kita Usahakan Lagi* by Batas Senja

Sebuah proyek kecil berbasis **Go** yang menyajikan **animasi teks lirik lagu** _"Kita Usahakan Lagi"_ di terminal dengan efek ketikan dan jeda waktu yang sinematik. Proyek ini dibuat untuk menunjukkan bagaimana terminal bisa menjadi ruang ekspresi yang sederhana namun berkesan.

---

## ✨ Fitur

- ⌨️ Efek ketikan per karakter seperti mesin ketik
- ⏱️ Jeda baris yang disesuaikan agar terasa ritmis
- 🎨 Warna putih terang untuk tampilan elegan di terminal
- 🔄 Pemisah visual antara bait dan chorus
- 💻 Kompatibel dengan hampir semua terminal modern

---

## 🧪 Contoh Output

🎵 == Kita Usahakan Lagi - Batas Senja == 🎵

Jika tidak hari ini
Mungkin minggu depan
Jika tidak minggu ini
Mungkin bulan depan
Jika tidak bulan ini
Mungkin tahun depan
Segala harapan kan datang
Yang kita impikan

Janganlah menyerah dulu
Waktu masih panjang
Ingat doa kita selalu
Yang tak pernah usang
Kita usahakan lagi sayang
Yakin waktunya kan datang


---

## 🚀 Cara Menjalankan

### Clone repo ini
```bash
git clone https://github.com/username/nama-repo.git
cd nama-repo
go run main.go
```

## 📁 Struktur File
```bash
├── main.go        # File utama program untuk menampilkan lirik lagu
└── README.md      # Dokumentasi ini
```

## 🧠 Insight Teknis
Proyek ini memanfaatkan:
- time.Sleep() untuk mengatur jeda per karakter dan antar baris
- ANSI escape codes (\033[1;97m) untuk memberikan efek warna terang pada teks di terminal
- Slice of struct untuk menyusun lirik beserta timing-nya secara fleksibel


<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL II</strong></h2>
<h2 align="center"><strong> REVIEW STRUKTUR KONTROL </strong></h2>

<br>

<p align="center">

  <img src="https://github.com/user-attachments/assets/0a03461e-7740-4661-9e83-9925031bd72c" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  LAILATUR RAHMAH / 2311102177<br>
  S1 IF 11 06
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Abednego Dwi Septiadi
  
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong><br>
  <strong>FAKULTAS INFORMATIKA</strong><br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong><br>
  <strong>2024</strong>
</p>

------

## Features

- [DASAR TEORI](#dasar-teori)
- [GUIDED](#guided)
- [UNGUIDED](#unguided)

## DASAR TEORI

### 2.1 Struktur Program Go

Dalam kerangka program yang ditulis dalam bahasa pemograman Go, program utama selalu mempunyai dua komponen berikut:

* **package main** merupakan penanda bahwa file ini berisi program utama.

* **func main()** berisi kode utama dari sebuah program Go

Komentar bukan bagian dari kode program, dan dapat ditulis dimana saja di dalam program:

* satu baris teks yang diawali dengan garis miring gada ('//') s.d akhir baris.
* beberapa baris teks yang dimulai dengan pasangan karakter '/*' dan diakhiri dengan 
'*/'

Review struktur kontrol dalam bahasa Go (Golang) berkaitan dengan berbagai elemen yang digunakan untuk mengontrol aliran eksekusi program. 
Struktur kontrol mengatur bagaimana blok kode tertentu dieksekusi berdasarkan kondisi atau perulangan. [1]
Berikut beberapa dasar teori dari struktur kontrol dalam bahasa Go:

**A. Pernyataan Kondisional (`if, else if, else`)**

- Digunakan untuk mengeksekusi blok kode tertentu jika kondisi tertentu terpenuhi.[1]
- Bentuk dasar:
```go
if kondisi {
    // blok kode dieksekusi jika kondisi true
} else if kondisi_lain {
    // blok kode dieksekusi jika kondisi_lain true
} else {
    // blok kode dieksekusi jika tidak ada kondisi yang terpenuhi
}
```
- Go juga mendukung pendeklarasian variabel dalam pernyataan if :
```go
if nilai := hitungNilai(); nilai > 10 {
    fmt.Println("Nilai lebih besar dari 10")
}
```

**2. Pernyataan Switch (`Switch`)**

- Alternatif dari beberapa pernyataan if else, mempermudah pembacaan kode.[1]
- Tidak seperti bahasa lain, switch Go secara otomatis tidak membutuhkan break untuk menghentikan eksekusi.[1]
- Bentuk dasar
```go
switch ekspresi {
case nilai1:
    // blok kode jika ekspresi == nilai1
case nilai2:
    // blok kode jika ekspresi == nilai2
default:
    // blok kode jika tidak ada nilai yang cocok
}
```
- `switch` juga bisa digunakan tanpa ekspresi untuk mengevaluasi kondisi :
```go
switch {
case x > 10:
    fmt.Println("x lebih besar dari 10")
case x < 5:
    fmt.Println("x lebih kecil dari 5")
default:
    fmt.Println("x berada di antara 5 dan 10")
}
```
**3. Perulangan (`for`)**
   
- Bahasa Go hanya memiliki satu jenis perulangan, yaitu for. Namun, for ini fleksibel dan dapat digunakan untuk berbagai bentuk perulangan.[1]
- Bentuk dasar for
```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```
- Perulangan tanpa kondisi (`infinite loop`)
```go
for {
    fmt.Println("Perulangan tanpa akhir")
}
```
- Digunakan untuk mengiterasi elemen dalam slice, array, map, atau channel :
```go
for index, value := range collection {
    fmt.Println(index, value)
}
```

**4. Pernyataan `break` dan `continue`**

- `break` digunakan untuk menghentikan eksekusi dari perulangan atau `switch`[1]
```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break // keluar dari loop saat i == 5
    }
}
```
- `continue` digunakan untuk melewati eksekusi kode yang tersisa dalam iterasi dan langsung lanjut ke iterasi berikutnya.
```go
for i := 0; i < 10; i++ {
    if i % 2 == 0 {
        continue // loncat ke iterasi berikutnya jika i genap
    }
    fmt.Println(i) // hanya akan mencetak bilangan ganjil
}
```
**5. Pernyataan `goto`**
- Go mendukung `goto`, meskipun jarang digunakan karena dapat membuat kode sulit diikuti. `goto` memindahkan eksekusi program ke label yang ditentukan.[1]
```go
func main() {
    i := 0
loop:
    if i < 5 {
        fmt.Println(i)
        i++
        goto loop
    }
}
```

**6. Pernyataan `Defer`**
- `defer` digunakan untuk menunda eksekusi fungsi sampai fungsi yang menampung defer selesai.[1]
```go
func contoh() {
    defer fmt.Println("Ini dieksekusi terakhir")
    fmt.Println("Ini dieksekusi pertama")
}
```
- `defer`sering digunakan untuk menutup resource seperti file atau koneksi jaringan, memastikan resource tersebut ditutup meskipun terjadi error.

**7.  `Panic` dan `Recover`**
- `panic` digunakan untuk menghentikan eksekusi program secara tiba-tiba. Biasanya digunakan untuk menangani kesalahan yang tidak bisa ditangani di runtime.[1]
```go
panic("Terjadi kesalahan!")
```
- `recover` digunakan untuk mengontrol dan mengatasi panic, sehingga program bisa kembali berjalan.
```go
func contohPanic() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recover dari panic:", r)
        }
    }()
    panic("Terjadi panic")
}
```

**B. Catatan**

Semua proses terkait bahasa Go dilakukan melalui utilitas Go. Beberapa opsi dengan utilitas go:

* **go build** mengkompilasi program sumber yang ada dalam folder menjadi sebuah program.
* **go build file.go** mengkompilasi program sumber file.go saja.
* **go fmt** membaca semua program sumber dalam folder dan mereformat penulisannya agar sesuai dengan standar penulisan program sumber Go.
* **go clean** membersihkan file-file dalam folder sehingga tersisa program sumbernya saja.

### 2.2 Tipe Data dan Intruksi Dasar

#### A. Tipe Data

#### 1. Integer (Bilangan Bulat)

Go menyediakan beberapa tipe integer:
- `int`: ukuran tergantung platform (32 atau 64 bit)
- `int8`: -128 hingga 127
- `int16`: -32768 hingga 32767
- `int32`: -2147483648 hingga 2147483647
- `int64`: -9223372036854775808 hingga 9223372036854775807
- `uint`: unsigned integer, tergantung platform
- `uint8`: 0 hingga 255
- `uint16`: 0 hingga 65535
- `uint32`: 0 hingga 4294967295
- `uint64`: 0 hingga 18446744073709551615

Contoh penggunaan:
```go
var angka int = 42
var umur uint8 = 25
```

#### 2. Float (Bilangan Desimal)
Go memiliki dua tipe float:
- `float32`: ±1.18×10⁻³⁸ hingga ±3.4×10³⁸
- `float64`: ±2.23×10⁻³⁰⁸ hingga ±1.80×10³⁰⁸

Contoh penggunaan:
```go
var pi float64 = 3.14159265359
var suhu float32 = 27.5
```

#### 3. String
String di Go adalah urutan karakter UTF-8 yang tidak dapat diubah (immutable).

Fitur string di Go:
1. Dapat menggunakan tanda kutip ganda (`"`) untuk string biasa
2. Dapat menggunakan tanda kutip balik (`` ` ``) untuk raw string

Contoh penggunaan:
```go
var nama string = "Budi"
var alamat string = `Jalan Merdeka No. 17
Kota Jakarta`  // Raw string bisa multi-baris
```

### 4. Boolean
Tipe bool memiliki dua nilai: `true` atau `false`

Contoh penggunaan:
```go
var lulus bool = true
var menikah bool = false
```

### B. Percabangan

#### 1. If-Else
Struktur if-else di Go memiliki beberapa fitur unik:
1. Tidak memerlukan tanda kurung `()` untuk kondisi
2. Kurung kurawal `{}` wajib digunakan
3. Dapat memiliki statement singkat sebelum kondisi

Contoh:
```go
// If-else biasa
if nilai >= 75 {
    fmt.Println("Lulus")
} else {
    fmt.Println("Tidak lulus")
}

// If dengan statement singkat
if skor := hitungSkor(); skor > 100 {
    fmt.Println("Skor maksimum!")
}
```

#### 2. Switch
Switch di Go lebih fleksibel dibanding bahasa lain:
1. Tidak perlu `break` di setiap case
2. Case bisa memiliki multiple values
3. Bisa menggunakan kondisi di setiap case

Contoh:
```go
// Switch biasa
switch buah {
case "apel":
    fmt.Println("Ini apel")
case "jeruk", "lemon":  // Multiple values
    fmt.Println("Ini jeruk atau lemon")
default:
    fmt.Println("Buah tidak dikenal")
}

// Switch dengan kondisi
switch {
case nilai >= 90:
    fmt.Println("A")
case nilai >= 80:
    fmt.Println("B")
default:
    fmt.Println("C")
}
```

### C. Perulangan

Go hanya memiliki satu konstruksi perulangan: `for`. Namun, ini bisa digunakan dalam beberapa cara:

#### 1. For Tradisional
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```
Komponen:
1. Statement inisialisasi: `i := 0`
2. Kondisi: `i < 5`
3. Statement post: `i++`

#### 2. For sebagai While
```go
counter := 0
for counter < 5 {
    fmt.Println(counter)
    counter++
}
```
Hanya menggunakan kondisi, mirip `while` di bahasa lain.

#### 3. For Range
Digunakan untuk iterasi melalui struktur data:
```go
slice := []string{"apel", "jeruk", "mangga"}
for index, nilai := range slice {
    fmt.Printf("Index: %d, Nilai: %s\n", index, nilai)
}
```
Dapat digunakan pada:
1. Array atau slice
2. String
3. Map
4. Channel

#### 4. Infinite Loop
```go
for {
    fmt.Println("Loop tanpa henti")
    // Gunakan 'break' untuk keluar dari loop
}
```

#### 5. Break
1. Gunakan `break` untuk keluar dari loop
2. Gunakan `continue` untuk melompat ke iterasi berikutnya
3. Label bisa digunakan dengan `break` dan `continue` untuk kontrol yang lebih spesifik

Contoh label:
```go
outerLoop:
for i := 0; i < 5; i++ {
    for j := 0; j < 5; j++ {
        if i*j > 10 {
            break outerLoop
        }
    }
}
```

**Data dan Variabel**

Variabel adalah nama dari suatu lokasi di memori yang data dengan tipe tertentu dapat disimpan.

* Nama variabel dimulai dengan huruf dan dapat diikuti dengan sejumlah huruf, angka, atau garis bawah.

Contoh: ketemu, found, rerata, mhs1, data_2, ...

* Tipe data yang umum tersedia adalah integer, real, boolean, karakter, dan string. 

* Nilai data yang tersimpan dalam variabel dapat diperoleh dengan menyebutkan langsung nama variabelnya.

Contoh: Menyebutkan nama found akan mengambil nilai tersimpan dalam memori untuk variabel found pastinya.

* Informasi alamat atau lokasi dari variabel dapat diperoleh dengan menambahkan prefiks & di depan nama variabel tersebut.

Contoh: &found akan mendapatkan alamat memori untuk menyimpan data pada found.

* Jika variabel berisi alamat memori, prefiksn * pada variabel tersebut akan memberikan nilai yang tersimpan dalam memori yang lokasinya disimpan dalam variabel tersebut.

Contoh: *mem akan mendapatkan data di memori yang alamatnya tersimpan di mem. Karenanya *(&found) akan mendapatkan data dari lokasi memori variabel found berada, alias sama saja dengan menyebutkan langsung found 8=).


## GUIDED

### 1. Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silahkan masukkan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang diperoleh. Coba terangkan apa sebenarnya yang di lakukan program tersebut.

#### Source Code

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Println("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Println("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Println("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/6eba38b9-15f0-4eae-b514-6d1c99f7771d)

#### Deskripsi Program

Program ini memberikan contoh sederhana tentang bagaimana sebuah program komputer dapat memanipulasi data dengan cara yang terstruktur. Konsep pertukaran nilai yang diimplementasikan dalam program ini adalah salah satu konsep dasar dalam pemrograman yang memiliki banyak aplikasi dalam berbagai bidang.

* **Inisialisasi Variabel:** Program memulai dengan mendeklarasikan tiga variabel string (misalnya, satu, dua, dan tiga) untuk menyimpan input dari pengguna. Selain itu, dideklarasikan pula variabel temp sebagai variabel sementara untuk membantu proses pertukaran.<br/>
* **Input Pengguna:** Program meminta pengguna untuk memasukkan tiga buah string. String-string ini kemudian disimpan ke dalam ketiga variabel yang telah dideklarasikan sebelumnya. <br/>
* **Pertukaran Nilai:** Proses inti dari program ini adalah pertukaran nilai antara ketiga variabel. Ini dilakukan dengan langkah-langkah berikut: <br/>
  + Nilai variabel satu disimpan sementara ke variabel temp.<br/>
  + Nilai variabel dua dipindahkan ke variabel satu.<br/>
  + Nilai variabel tiga dipindahkan ke variabel dua.<br/>
  + Nilai yang tersimpan di variabel temp (nilai awal satu) dipindahkan ke variabel tiga.<br/>
* **Output:**  Setelah proses pertukaran selesai, program menampilkan nilai akhir dari ketiga variabel string tersebut.<br/>

### 2. Tahun kabisat adalah tahun yang habis dibagi 400 atau habis dibagi 4 tetapi tidak habis dibagi 100. Buatlah sebuah program yang menerima input sebuah bilangan bulat dan memeriksa apakah bilangan tersebut merupakan tahun kabisat (true) atau bukan (false).

#### Source Code

```go
package main

import "fmt"

func main() {
	var tahun int
	fmt.Println("Masukkan sebuah tahun: ")
	fmt.Scanln(&tahun)

	if (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0) {
		fmt.Println(tahun, "adalah tahun kabiat: true")
	} else {
		fmt.Println(tahun, "bukan tahun kabisat: false")
	}
}
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/d488caad-409a-481a-ac17-f8fc28fa569e)

#### Deskripsi Program

Program ini adalah aplikasi sederhana yang ditulis dalam bahasa Go untuk menentukan apakah sebuah tahun adalah tahun kabisat atau bukan. Dengan meminta input tahun dari pengguna, program menggunakan logika berbasis kondisi untuk memeriksa kriteria tahun kabisat—yaitu dapat dibagi 400 atau dapat dibagi 4 tetapi tidak 100. Hasilnya kemudian ditampilkan kepada pengguna, memberikan informasi yang jelas tentang status tahun yang dimasukkan. Program ini dapat digunakan dalam konteks perhitungan tanggal atau perencanaan yang membutuhkan pemahaman tentang tahun kabisat.

* **Variabel:** tahun digunakan untuk menyimpan nilai tahun yang dimasukkan pengguna. <br/>
* **Operator Modulus (%):** Operator ini digunakan untuk mencari sisa pembagian. Dalam kasus ini, digunakan untuk memeriksa apakah suatu bilangan habis dibagi dengan bilangan lain. <br/>
* **Kondisi if-else:** Struktur kontrol ini digunakan untuk membuat keputusan berdasarkan suatu kondisi.<br/>
* **Operator Logika:** Operator && (AND) digunakan untuk menggabungkan dua kondisi. <br/>

### 3. Buat program Bola yang menerima input jari-jari suatu bola (bilangan bulat). Tampilkan Volume dan Luas kulit bola. volumebola = r³ dan luasbola = 4πr² (π = 3.1415926535).

#### Source Code

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var jariJari int
	fmt.Print("Jejari = ")
	fmt.Scan(&jariJari)

	// Hitung volume dan luas permukaan
	volume := (4.0 / 3.0) * math.Pi * float64(jariJari*jariJari*jariJari)
	luas := 4 * math.Pi * float64(jariJari*jariJari)

	// Tampilkan hasil
	fmt.Printf("Bola dengan jari-jari %d memiliki volume %.4f dan luas kulit %.4f\n", jariJari, volume, luas)
}
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/f97bd147-04fc-4c2f-b957-29d611b03a5a)

#### Deskripsi Program

Program ini adalah aplikasi sederhana yang ditulis dalam bahasa Go untuk menghitung volume dan luas permukaan bola berdasarkan jari-jari yang dimasukkan oleh pengguna. Setelah meminta input jari-jari dari pengguna, program menggunakan rumus matematis—yaitu volume bola \( V = \frac{4}{3} \pi r^3 \) dan luas permukaan bola \( A = 4 \pi r^2 \)—untuk menghitung nilai volume dan luas permukaan. Hasil perhitungan ditampilkan dengan format yang jelas, menampilkan nilai volume dan luas kulit bola dengan presisi empat angka desimal. Program ini berguna dalam konteks pendidikan atau aplikasi yang memerlukan penghitungan geometri tiga dimensi.

**Cara Kerja:**
* **Deklarasi Paket dan Import:**
  + **package main:** Menentukan bahwa ini adalah program utama.
  + **import "fmt":** Mengimpor paket fmt untuk melakukan input/output (misalnya, mencetak ke layar atau membaca input dari pengguna).
  + **import "math":** Mengimpor paket math untuk menggunakan fungsi-fungsi matematika seperti math.Pi.

**Deklarasi Variabel:**
  + **var jariJari int:** Mendeklarasikan variabel jariJari dengan tipe data integer untuk menyimpan nilai jari-jari bola yang dimasukkan pengguna.

**Input Jari-Jari:**
  + **fmt.Print("Jejari = "):** Mencetak pesan ke layar untuk meminta pengguna memasukkan nilai jari-jari.
  + **fmt.Scan(&jariJari):** Membaca nilai yang dimasukkan pengguna dan menyimpannya ke dalam variabel jariJari.

Perhitungan Volume dan Luas Permukaan:
  + **volume := (4.0 / 3.0) * math.Pi * float64(jariJari*jariJari*jariJari):** Menghitung volume bola menggunakan rumus yang sudah diketahui. Perhatikan bahwa kita menggunakan float64 untuk memastikan hasil perhitungan berupa bilangan desimal dengan presisi yang cukup.
  + **luas := 4 * math.Pi * float64(jariJari*jariJari):** Menghitung luas permukaan bola menggunakan rumus yang sesuai.

## Unguided

### 1. Di baca nilai temperatur dalam derajat Celsius. Nyatakan temperatur tersebut dalam Fahrenheit. <br/>
Fahrenheit = ( Celsius 9/5) + 32 <br/>
Reamur = Celsius 4/5 <br/>
Kelvin = (Fahrenheit +459.67) 5/9 <br/>
Lanjutkan program di atas, sehingga temperatur dinyatakan juga dalam derajat Reamur dan Kelvin. <br/>

#### Source Code

```go
package main

import "fmt"

func main() {
	var celsius float64

	fmt.Print("Masukkan suhu dalam derajat Celsius: ")
	fmt.Scanln(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32
	reamur := celsius * 4 / 5
	kelvin := celsius + 273.15

	fmt.Printf("Suhu dalam Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Suhu dalam Reamur: %.2f\n", reamur)
	fmt.Printf("Suhu dalam Kelvin: %.2f\n", kelvin)
}
```
#### Screenshoot Output

![image](https://github.com/user-attachments/assets/3aefeaf1-1f77-417f-ab6e-7e495a4c5c7a)

#### Deskripsi Program

Program ini adalah aplikasi yang ditulis dalam bahasa Go untuk mengonversi suhu dari derajat Celsius ke beberapa skala suhu lainnya, yaitu Fahrenheit, Reamur, dan Kelvin. Setelah pengguna memasukkan nilai suhu dalam derajat Celsius, program menghitung nilai suhu dalam skala Fahrenheit menggunakan rumus \( F = \frac{9}{5}C + 32 \), Reamur dengan rumus \( R = \frac{4}{5}C \), dan Kelvin dengan rumus \( K = C + 273.15 \). Hasil konversi suhu ditampilkan dengan format yang jelas, mencakup nilai dalam masing-masing skala suhu dengan presisi dua angka desimal. Program ini berguna untuk aplikasi sehari-hari yang memerlukan konversi suhu dan pemahaman tentang berbagai skala suhu.

**Deklarasi Paket dan Import:**

* **package main:** Menentukan bahwa ini adalah program utama.<br/>
* **import "fmt":** Mengimpor paket fmt untuk melakukan input/output (misalnya, mencetak ke layar atau membaca input dari pengguna).<br/>

**Deklarasi Variabel:**

* **var celsius float64:** Mendeklarasikan variabel celsius dengan tipe data float64 untuk menyimpan nilai suhu dalam derajat Celsius yang akan diinputkan oleh pengguna. Tipe data float64 dipilih karena suhu bisa berupa bilangan desimal.

**Input Suhu:**

* **fmt.Print("Masukkan suhu dalam derajat Celsius: "):** Mencetak pesan ke layar untuk meminta pengguna memasukkan nilai suhu dalam derajat Celsius.<br/>
* **fmt.Scanln(&celsius):** Membaca nilai yang diinputkan oleh pengguna dari input standar (biasanya keyboard) dan menyimpannya ke dalam variabel celsius.<br/>

**Konversi Suhu:**

* **fahrenheit := (celsius * 9 / 5) + 32:** Menghitung suhu dalam derajat Fahrenheit menggunakan rumus konversi yang sesuai.<br/>
* **reamur := celsius * 4 / 5:** Menghitung suhu dalam derajat Reamur menggunakan rumus konversi yang sesuai.<br/>
* **kelvin := celsius + 273.15:** Menghitung suhu dalam Kelvin menggunakan rumus konversi yang sesuai.<br/>

### 2. Buat program ASCII yang akan membaca 5 buah data integer dan mencetaknya dalam format karakter. Kemudian membaca 3 buah data karakter dan mencetak 3 buah karakter setelah karakter tersebut (menurut tabel ASCII) <br/>
Masukan terdiri dari dua baris. Baris pertarna berisi 5 buah data integer. Data integer mempunyai nilai antara 32 s.d. 127. Baris kedua berisi 3 buah karakter yang berdampingan satu dengan yang lain (tanpa dipisahkan spasi). <br/>
Keluaran juga terdiri dari dua baris. Baris pertama berisi 5 buah representasi karakter dari data yang diberikan, yang berdampingan satu dengan lain, tanpa dipisahkan spasi. Baris kedua berisi 3 buah karakter (juga tidak dipisahkan oleh spasi).<br/>

#### Source Code

```go
package main

import "fmt"

func main() {
	var numbers [5]int
	var chars [3]byte

	fmt.Println("Masukkan 5 buah data integer (32-127):")
	for i := 0; i < 5; i++ {
		fmt.Scan(&numbers[i])
	}

	fmt.Println("Masukkan 3 buah karakter:")
	for i := 0; i < 3; i++ {
		fmt.Scanf("%c", &chars[i])
	}

	fmt.Print("Karakter dari nilai integer: ")
	for _, num := range numbers {
		fmt.Printf("%c", num)
	}
	fmt.Println()

	fmt.Print("Karakter setelahnya: ")
	for _, char := range chars {
		fmt.Printf("%c", char+1)
	}
	fmt.Println()
}
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/f7c427ea-930e-4a6a-9e6e-c7c779aaa4dc)

#### Deskripsi Program

Program ini dirancang untuk melakukan konversi antara nilai integer dan karakter ASCII. Pengguna diminta untuk memasukkan sejumlah bilangan bulat dalam rentang tertentu. Setiap bilangan bulat ini kemudian diinterpretasikan sebagai kode ASCII dan dikonversi menjadi karakter yang sesuai. Selain itu, program juga meminta pengguna untuk memasukkan beberapa karakter. Karakter-karakter ini kemudian akan dikonversi menjadi nilai ASCII-nya, lalu dinaikkan satu angka, dan dikonversi kembali menjadi karakter. Dengan kata lain, program ini menunjukkan hubungan antara representasi numerik (ASCII) dan karakter yang kita lihat di layar. Hasil akhir dari program adalah menampilkan karakter yang sesuai dengan nilai integer yang dimasukkan, serta karakter setelahnya dari karakter yang diinputkan. pengguna untuk memasukkan lima nilai integer dalam rentang 32 hingga 127, yang kemudian dicetak sebagai karakter ASCII. Selanjutnya, program juga meminta tiga karakter untuk diinput, dan mencetak karakter-karakter tersebut setelah ditambahkan satu ke nilai ASCII-nya. Namun, program ini bisa ditingkatkan dengan menangani whitespace saat membaca karakter dan memvalidasi input integer agar sesuai dengan rentang yang ditentukan, sehingga memberikan feedback yang lebih baik jika data yang dimasukkan tidak valid.



### 3. Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah 'merah', 'kuning', 'hijau', dan 'ungu' selama 5 kali percobaan berulang.<br/>
Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya.<br/>

#### Source Code

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Urutan warna yang benar
	correctOrder := []string{"merah", "kuning", "hijau", "ungu"}

	// Membaca input untuk 5 percobaan
	reader := bufio.NewReader(os.Stdin)
	var hadError bool

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)

		// Membaca input dari pengguna
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Memisahkan input berdasarkan spasi
		colors := strings.Split(input, " ")

		// Mengecek apakah urutan warna sesuai
		for j := 0; j < 4; j++ {
			if colors[j] != correctOrder[j] {
				hadError = true
				break
			}
		}
	}

	// Menampilkan hasil
	if !hadError {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}
}
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/70b2bc7b-5ff4-4189-a205-8a0777ca0249)

#### Deskripsi Program

Program ini adalah aplikasi dalam bahasa Go yang meminta pengguna untuk memasukkan urutan warna dalam lima percobaan dan memeriksa apakah urutan tersebut sesuai dengan urutan yang benar, yaitu "merah", "kuning", "hijau", dan "ungu". Setelah pengguna memasukkan input, program membandingkan setiap warna dengan urutan yang benar; jika ada kesalahan, program menandai kesalahan tersebut. Di akhir percobaan, program mencetak "BERHASIL: true" jika semua urutan benar, atau "BERHASIL: false" jika terdapat kesalahan. Program ini berguna sebagai latihan pengenalan warna atau permainan memori untuk meningkatkan daya ingat pengguna.

**Cara Kerja:**

* **Deklarasi Paket dan Import:**

  + **package main:** Menentukan bahwa ini adalah program utama.<br/>
  + **import "bufio", "fmt", "os", "strings":** Mengimpor paket-paket yang diperlukan untuk membaca input dari pengguna, melakukan format input/output, berinteraksi dengan sistem operasi, dan memanipulasi string.<br/>

* **Inisialisasi Urutan Warna Benar:**

  + **correctOrder := []string{"merah", "kuning", "hijau", "ungu"}:** Membuat slice (array yang dinamis) yang berisi urutan warna yang benar. Slice ini akan digunakan sebagai acuan untuk membandingkan dengan input pengguna.

* **Membaca Input Pengguna:**

 + **reader := bufio.NewReader(os.Stdin):** Membuat objek reader untuk membaca input dari pengguna secara baris per baris.<br/>
 + **Loop for i := 1; i <= 5; i++:** Melakukan perulangan sebanyak 5 kali untuk 5 percobaan.<br/>
 + **fmt.Printf("Percobaan %d: ", i):** Mencetak nomor percobaan ke layar.<br/>
 + **input, _ := reader.ReadString('\n'):** Membaca satu baris input dari pengguna dan menyimpannya dalam variabel input.<br/>
 + **input = strings.TrimSpace(input):** Menghapus spasi kosong di awal dan akhir string input.<br/>
 + **colors := strings.Split(input, " "):** Membagi string input menjadi slice of string berdasarkan spasi, sehingga setiap kata (warna) menjadi elemen dalam slice colors.<br/>

* **Membandingkan Urutan Warna:**

  + **Loop for j := 0; j < 4; j++:** Melakukan perulangan untuk membandingkan setiap elemen dalam slice colors dengan elemen yang sesuai dalam correctOrder.<br/>
  + **if colors[j] != correctOrder[j]:** Jika warna pada posisi j tidak sama, maka hadError diubah menjadi true dan loop dihentikan menggunakan break.<br/>

* **Menampilkan Hasil:**

  + Setelah semua percobaan selesai, program akan mencetak "BERHASIL: true" jika tidak ada kesalahan (semua urutan benar), atau "BERHASIL: false" jika ada setidaknya satu urutan yang salah.

### 4.	Buatlah sebuah program yang menerima input sebuah bilangan bulat positif (dan tidak nol) N, kemudian program, akan meminta input berupa nama bunga secara berulang sebanyak N kali dan nama tersebut disimpan ke dalam pita.<br/>
Modifikasi program sebelumnya, proses input akan berhenti apabila user mengetikkan 'selesai'. Kemudian tampilkan isi pita beserta banyaknya bunga yang ada di dalam pita <br/>

#### Source Code

```go
package main

import "fmt"

func main() {
	// Deklarasi variabel untuk menyimpan pita dan banyaknya bunga
	var pita string     // Variabel 'pita' digunakan untuk menggabungkan nama bunga
	var banyakBunga int // Variabel 'banyakBunga' untuk menghitung jumlah bunga yang dimasukkan

	// Menggunakan loop untuk meminta input bunga dari pengguna secara berulang
	for {
		var bunga string
		// Meminta input nama bunga
		fmt.Print("Bunga (ketik 'SELESAI' untuk berhenti): ")
		fmt.Scanln(&bunga) // Membaca input dari pengguna

		// Jika pengguna mengetik "SELESAI", maka keluar dari loop
		if bunga == "SELESAI" {
			break
		}

		// Jika pita sudah memiliki isi, tambahkan pemisah " – "
		if pita != "" {
			pita += " – "
		}
		// Menambahkan nama bunga ke pita
		pita += bunga

		// Meningkatkan jumlah bunga yang dimasukkan
		banyakBunga++
	}

	// Menampilkan hasil akhir
	fmt.Println("Pita:", pita)                   // Menampilkan pita yang berisi nama-nama bunga
	fmt.Println("Banyaknya bunga:", banyakBunga) // Menampilkan jumlah bunga yang dimasukkan
}
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/c9197737-fac7-4555-adc1-5d2baf9e15a6)

#### Deskripsi Program

Program ini meminta pengguna untuk memasukkan nama-nama bunga secara berulang hingga mereka mengetik "SELESAI". Setiap nama bunga yang dimasukkan akan digabungkan ke dalam satu string bernama `pita`, dengan pemisah " – " di antara setiap nama. Selain itu, program juga menghitung jumlah bunga yang dimasukkan dan menyimpannya dalam variabel `banyakBunga`. Setelah pengguna selesai, program menampilkan daftar bunga yang telah dimasukkan dan total jumlah bunga tersebut.

**Cara Kerja:**

* **Inisialisasi Variabel:**

  + **pita:** Variabel jenis string digunakan untuk menyimpan nama-nama bunga yang diinputkan. Awalnya, variabel ini kosong.<br/>
  + **banyakBunga:** Variabel bertipe integer digunakan untuk menghitung jumlah bunga yang telah dimasukkan. Awalnya, nilai variabel ini adalah 0.<br/>

* **Looping Input:**

  + Program menggunakan struktur for untuk melakukan perulangan. Perulangan akan terus berjalan hingga kondisi break terpenuhi.<br/>
  + Dalam setiap iterasi, program meminta pengguna untuk memasukkan nama bunga.<br/>
  + Jika pengguna memasukkan "SELESAI", maka program akan keluar dari perulangan.<br/>

* **Membangun Kalimat:**

  + Setiap kali pengguna memasukkan nama bunga, nama bunga tersebut akan ditambahkan ke variabel pita.<br/>
  + Untuk memisahkan setiap nama bunga, digunakan tanda hubung ("–").<br/>
  + Jika pita sudah berisi nama bunga sebelumnya, maka tanda hubung akan ditambahkan terlebih dahulu sebelum menambahkan nama bunga yang baru.<br/>

* **Menghitung Jumlah Bunga:**

  + Setiap kali ada nama bunga yang ditambahkan ke pita, nilai banyakBunga akan ditambah 1.

* **Menampilkan Hasil:**

  + Setelah perulangan selesai, program akan menampilkan nilai dari pita (yang berisi semua nama bunga) dan banyakBunga (yang menunjukkan jumlah bunga).

### 5.	Buatlah program Pak Andi yang menerima input dua buah bilangan real positif yang menyatakan berat total masing-masing isi kantong terpal. Program akan terus meminta input bilangan tersebüt hingga salah satu kantong terpal berisi 9 lig atau lebih. <br/>
Pada modifikasi program tersebut, program akan menampilkan true jika selisih kedua isi kantong lebih dari atau sama dengan 9kg, program akan berhenti memproses apabila total berat isi kedua kantong melebihi 150kg atau salah satu kantong beratnya negatif <br/>

#### Source Code

```go
package main

import "fmt"

func main() {
	for {
		var berat1, berat2 float64

		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		_, err := fmt.Scan(&berat1, &berat2)
		if err != nil {
			fmt.Println("Input tidak valid.")
			return
		}

		if berat1+berat2 > 150 || berat1 < 0 || berat2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		selisih := berat1 - berat2
		if selisih < 0 {
			selisih = -selisih
		}
		akanOleng := selisih >= 9

		fmt.Printf("Sepeda motor pak Andi akan oleng: %t\n", akanOleng)
	}
}
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/83464312-7955-4cff-8590-6e9c12b113a3)

#### Deskripsi Program

Program ini dirancang untuk mensimulasikan kondisi keseimbangan beban pada sepeda motor Pak Andi. Program akan terus meminta input berat belanjaan di dua kantong hingga kondisi tertentu terpenuhi.

**Cara Kerja:**

* **Looping Berulang:** Program menggunakan for untuk membuat perulangan tak terbatas, sehingga akan terus meminta input dari pengguna hingga kondisi tertentu terpenuhi.<br/>
Input Berat Belanjaan: Pada setiap iterasi, program meminta pengguna untuk memasukkan berat belanjaan di dua kantong. Nilai-nilai ini disimpan dalam variabel berat1 dan berat2.<br/>
* **Validasi Input:** Program memeriksa apakah input yang diberikan oleh pengguna valid. Jika terjadi kesalahan saat membaca input, program akan menampilkan pesan kesalahan dan berhenti.<br/>
* **Program akan berhenti jika:** <br/>
  + Total berat kedua kantong melebihi 150 kg. <br/>
  + Salah satu atau kedua kantong memiliki berat negatif.<br/>
* **Perhitungan Selisih:** Program menghitung selisih berat antara kedua kantong. Nilai selisih selalu dibuat positif dengan menggunakan fungsi abs.<br/>
* **Menentukan Kondisi Oleng:** Jika selisih berat antara kedua kantong lebih besar atau sama dengan 9 kg, maka dianggap sepeda motor akan oleng. Hasilnya akan ditampilkan ke layar.<br/>

Intinya, program ini akan terus berjalan hingga total berat belanjaan melebihi batas atau salah satu kantong memiliki berat negatif. Pada setiap iterasi, program akan memeriksa apakah selisih berat antara kedua kantong cukup besar untuk membuat sepeda motor oleng.

### 6.	Diberikan sebuah persamaan sebagai berikut ini. <br/>
f(k) = (4k+2)² /(4k+1)(4k+3) <br/>
Buatlah sebuah program yang menerima input sebuah bilangan sebagai K, hemudian menghitüng dan menampilkan nilai f(K) sesuai persamaan di atas. <br/>

√2 merupakan bilangan irasional. Meskipun demikian, nilai tersebut dapat dihampin dengan rumus berikut: <br/>
√Σ-Π П (4+2)² (4k+1)(4k+3) <br/>

Modifikası program sebelumnya yang menerima input integer K dan menghitung √2 untuk K tersebut. Hampiran v2 dituliskan dalam ketelitian 10 angka di belakang koma.

#### Source Code

**Sebelum dimodifikasi**

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var k float64

	fmt.Print("Masukkan nilai K: ")
	fmt.Scan(&k)

	fk := math.Pow(4*k+2, 2) / ((4*k + 1) * (4*k + 3))

	fmt.Printf("Nilai f(k) = %.10f\n",fk)
}
```

**Sesudah dimodifikasi**

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var k int

	for {
		fmt.Print("Masukkan nilai K: ")
		_, err := fmt.Scan(&k)
		if err != nil {
			fmt.Println("Input tidak valid.")
			return
		}

		if k < 0 {
			fmt.Println("Nilai K harus positif.")
			continue
		}

		akarDua := 0.0
		for i := 0; i <= k; i++ {
			pembilang := math.Pow(4*float64(i)+2, 2)
			penyebut := (4*float64(i) + 1) * (4*float64(i) + 3)
			akarDua += pembilang / penyebut
		}

		fmt.Printf("Nilai akar 2 = %.10f\n", akarDua)
	}
}
```

#### Screenshoot Output

**Sebelum dimodifikasi**

![image](https://github.com/user-attachments/assets/3ec2185a-2809-4dbb-a8c5-71e0f178a1bf)

**Sesudah dimodifikasi**

![image](https://github.com/user-attachments/assets/441efdc2-6d61-46c6-8072-6ef195a11350)

#### Deskripsi Program

Program ini adalah aplikasi dalam bahasa Go yang menghitung nilai fungsi \( f(k) \) berdasarkan input nilai \( k \) yang diberikan oleh pengguna. Setelah pengguna memasukkan nilai \( k \), program menggunakan rumus \( f(k) = \frac{(4k + 2)^2}{(4k + 1)(4k + 3)} \) untuk melakukan perhitungan. Hasil dari perhitungan fungsi ditampilkan dengan presisi sepuluh angka desimal. Program ini berguna untuk aplikasi matematis yang memerlukan evaluasi fungsi atau pengujian nilai berdasarkan rumus tertentu.

Lalu setelah dimodifikasi program ini akan menghitung estimasi nilai akar dua menggunakan metode penjumlahan berdasarkan input nilai \( k \) yang diberikan oleh pengguna. Dengan memastikan bahwa nilai \( k \) valid dan positif, program menghitung estimasi akar dua menggunakan rumus matematis tertentu, yaitu \( \text{akarDua} = \sum_{i=0}^{k} \frac{(4i + 2)^2}{(4i + 1)(4i + 3)} \). Hasil perhitungan ditampilkan dengan presisi sepuluh angka desimal, dan program terus meminta input hingga pengguna memilih untuk berhenti. Kesimpulannya, program ini tidak hanya berfungsi sebagai alat untuk menghitung estimasi akar dua, tetapi juga memberikan wawasan tentang bagaimana variasi dalam nilai \( k \) dapat memengaruhi hasil estimasi, serta mengedukasi pengguna tentang konsep estimasi numerik.

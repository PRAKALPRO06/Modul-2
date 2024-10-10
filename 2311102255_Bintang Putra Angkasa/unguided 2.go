package main
import "fmt"

func main() {
    const jumlahPercobaan = 5
    urutanWarnaDiharapkan := [4]string{"merah", "kuning", "hijau", "ungu"}
    var hasilAkhir bool = true

    // Melakukan iterasi sebanyak 5 kali percobaan
    for i := 0; i < jumlahPercobaan; i++ {
        var warnaInput [4]string
        fmt.Printf("Masukkan urutan warna untuk percobaan ke-%d (pisahkan dengan spasi):\n", i+1)
        
        // Meminta input 4 warna dari pengguna
        for j := 0; j < 4; j++ {
            fmt.Scan(&warnaInput[j])
        }
        
        // Memeriksa apakah urutan warna yang dimasukkan sesuai dengan urutan yang diharapkan
        for j := 0; j < 4; j++ {
            if warnaInput[j] != urutanWarnaDiharapkan[j] {
                hasilAkhir = false
            }
        }
    }

    // Menampilkan hasil akhir: true jika semua percobaan berhasil, false jika ada yang gagal
    fmt.Println(hasilAkhir)
}

# INTERVIEW QUESTION

## 1. COIN

### SOAL

Jika ada deret angka sejumlah "K" ingin dijumlahkan hingga mendapat angka "N", berapa banyak kombinasi yang bisa dibentuk.

### CONTOH

#### Input

- N = 4
- K = 5

kombinasi yang bisa dibentuk dari angka 1 hingga 4 (1,2,3,4) hingga hasilnya 5 adalah 6 kombinasi:

```
[1,4]
[2,3]
[1,1,3]
[1,2,2]
[1,1,1,2]
[1,1,1,1,1]
```

#### Output

- 6

---

## 2. CONTAIN

### SOAL

apakah sebuah kalimat mengandung kata tertentu

### CONTOH

#### Input

- kalimat = "jumps"
- kata = "quick brown fox jumps over the lazy dog"

#### Output

- true

---

## 3. FACTORIAL

### SOAL

mencari hasil faktorial dari angka tertentu

### CONTOH

#### Input

- input = 10!

#### Output

- 3628800

---

## 4. INTEREST

### SOAL

mencari total uang hasil bunga tertentu dalam kurun waktu (tahun) tertentu

### CONTOH

#### Input

- amount = 1_000_000
- year = 5 years
- interest = 2 %

#### Output

- 1_104_080.8032

---

## 5. JUMP

### SOAL

mencari deret angka dari 2 angka pertama dengan panjang tertentu

### CONTOH

#### Input

- first = 3
- second = 6
- length = 5

#### Output

- [3, 6, 9, 12, 15]

---

## 6. LONGEST TIME

### SOAL

mencari jarak waktu tempuh terjauh dari jarak tombol yang dituju dengan tombol sebelumnya

### CONTOH

#### Input

- input = [[0,2],[1,9],[0,12],[2,15]]

- sub array index pertama dalam input merepresentasikan tombol mana yang ditekan, 0 = "a", 1 = "b" dan seterusnya hingga 25 = "z"
- sub array index kedua merepresentasikan waktu kapan tombol di tekan

- jarak waktu tombol pertama (0) adalah 2 karna 2-0 = 2
- jarak waktu tombol kedua (1) adalah 7 karna 9-2 = 7
- jarak waktu tombol ketiga (0) adalah 3 karna 12-9 = 3
- jarak waktu tombol keempat (2) adalah 3 karna 15-12 = 3

karna yang memiliki waktu tempuh terjauh adalah tombol kedua (1) maka outputnya adalah "b"

#### Output

- "b"

---

## 7. MINIMUM GROUP

### SOAL

cari berapa kelompok minimum yang bisa di buat dari daftar hewan yang memiliki pemangsanya di daftar tersebut dengan ketentuan:

- 1 hewan hanya memiliki 1 pemangsa langsung
- hewan juga tidak bisa digabung denga pemangsa tidak langsungnya

### CONTOH

#### Input

- input = [-1, 8, 6, 0, 7, 3, 8, 9, -1, 6]

- index dari input adalah nomor hewan
- value dari input adalah nomor index pemangsa hewan tersebut
- jika value -1 berarti hewan tersebut tidak memiliki pemangsa

maka akan menghasilkan graph berikut:

- ![alt](https://media.cheggcdn.com/media/a7e/a7e3e8da-c574-4f14-8fba-9ab6cbe15685/image.png)

#### Output

- 5

---

## 8. MOCK 1

### SOAL

cari 2 angka pertama dalam sebuah array yang akan menghasilkan nilai tertentu, jika tidak ada return (0, 0)

### CONTOH

#### Input

- arr = [1, 2, 3, 2, 3, 4, 5]
- target = 6

#### Output

- [3,3]

---

## 9. MOCK 2

### SOAL

buat deret sebagai berikut ["1", "\*", "3", "4", "\*", "6", "\*", "8", "9", "\*"]

### CONTOH

#### Input

- max = 10

#### Output

- ["1", "\*", "3", "4", "\*", "6", "\*", "8", "9", "\*"]

---

## 10. MOCK 3

### SOAL

tentukan kondisi dari semua bracket "(" "{" "[" "]" "}" ")" menutup sempurna atau tidak

### CONTOH

#### Input

- input = "\[\{\}\]\(\{\[\]\}\)"

#### Output

- true

---

## 11. MOCK 4

### SOAL

membalik integer

### CONTOH

#### Input

- input = 1234

#### Output

- 4321

---

## 12. ODD NUMBER

### SOAL

membuat deret bilangan ganjil dengan panjang tyertentu

### CONTOH

#### Input

- total = 10

#### Output

- [1, 3, 5, 7, 9, 11, 13, 15, 17, 19]

---

## 13. PALINDROME

### SOAL

tentukan kondisi kalimat input adalahs ebuah palindrome atau bukan

### CONTOH

#### Input

- input = "ABBA"

#### Output

- true

---

## 14. PRIME

### SOAL

buat deret bilangan prima dengan pajang tertentu

### CONTOH

#### Input

- total = 10

#### Output

- [2, 3, 5, 7, 11, 13, 17, 19, 23, 29]

---

## 15. SQUARE

### SOAL

buat kotak menggunakan hash "#" dengan panjang sisi tertentu

### CONTOH

#### Input

- t = 5

#### Output

```
# # # # #
# # # # #
# # # # #
# # # # #
# # # # #
```

---

## 16. SUM OF SERIES

### SOAL

buat fungsi untuk menghitung baris berikut
$$ sum = {1 \over 1 + \sqrt{2}} + {1 \over \sqrt{2} + \sqrt{3}} ... + {1 \over \sqrt{n} + \sqrt{n+1}}$$

### CONTOH

#### Input

- t = 10

#### Output

- 2.3166247903554

---

## 17. TRIANGLE

### SOAL

buat pattern segitiga

### CONTOH

#### Input

- l = 5

#### Output

```
# # # # #
# # # #
# # #
# #
#
```

---

## 18. TRIANGLE HOLE

### SOAL

buat pattern segitiga 2

### CONTOH

#### Input

- l = 5

#### Output

```

# # # # #
#     #
#   #
# #
#
```

---

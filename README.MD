# 🚀 Go 30 Days Journey

> Learning Golang from **Syntax → Memory → Engineer Mindset**

## 📅 Day 1 & Day 2

---

# 🎯 Tujuan Belajar

Target belajar ini bukan hanya bisa membuat REST API menggunakan Go, tetapi memahami:

- Cara kerja compiler
- Memory model Go
- Filosofi desain Go
- Best practice penulisan kode
- Cara berpikir seorang Go Backend Engineer

---

# 📖 Filosofi Go

Go dibuat dengan prinsip:

- **Simple**
- **Readable**
- **Maintainable**
- **Fast Compilation**
- **Explicit over Magic**

Go tidak berusaha menjadi bahasa yang memiliki banyak fitur, tetapi menjadi bahasa yang mudah dipelajari dan mudah dipelihara untuk project besar.

## Prinsip penting

> **Clear is better than clever.**

Kode yang sederhana dan mudah dibaca lebih baik daripada kode yang terlalu pintar tetapi sulit dipahami.

---

# ⚙️ Compiled Language

Alur kerja Go:

```text
Source Code (.go)

        ↓

    Compiler

        ↓

 Binary / Executable

        ↓

 Operating System

        ↓

       CPU

        ↓

      Output
```

Berbeda dengan PHP yang menggunakan interpreter pada setiap request.

Keuntungan:

- Startup cepat
- Performa tinggi
- Mudah didistribusikan (single binary)

---

# 📝 Zero Value

Semua variabel di Go selalu mempunyai nilai default.

| Type      | Zero Value |
| --------- | ---------- |
| int       | 0          |
| float64   | 0          |
| string    | ""         |
| bool      | false      |
| pointer   | nil        |
| slice     | nil        |
| map       | nil        |
| interface | nil        |

Filosofi:

> **Never leave memory uninitialized.**

---

# 📦 Value Semantics

```go
x := 10
y := x

y = 20
```

Output:

```
10
20
```

Yang terjadi adalah:

```
COPY VALUE
```

Memory:

```
x = 10

y = 20
```

Value Semantics berlaku untuk:

- int
- float
- bool
- string
- array
- struct (default)

---

# 🔗 Reference-like Semantics

```go
a := []int{1,2,3}

b := a

b[0] = 100
```

Output:

```
a = [100 2 3]

b = [100 2 3]
```

Yang di-copy bukan isi array, tetapi Slice Header.

```
ptr

len

cap
```

Kedua slice tetap menunjuk underlying array yang sama.

---

# 📚 Slice Internal

Slice bukan Array.

Slice terdiri dari:

```
ptr

len

cap
```

Visualisasi:

```
users

ptr ---> underlying array

len = jumlah data

cap = kapasitas yang dialokasikan
```

---

# 🧠 Stack vs Heap

## Stack

- Cepat
- Otomatis dibersihkan
- Digunakan untuk object sederhana

## Heap

- Dikelola Garbage Collector
- Digunakan ketika object perlu hidup lebih lama

> Catatan: Tidak semua pointer berada di Heap.

---

# 🔍 Escape Analysis

Compiler Go akan menentukan sendiri apakah object ditempatkan di Stack atau Heap.

Konsep:

```
Apakah object masih dibutuhkan setelah function selesai?

YES

↓

Heap

NO

↓

Stack
```

---

# 🎭 Variable Shadowing

```go
a := 10

{
    a := 20
}
```

Output:

```
20
10
```

Karena variable di dalam block adalah variable baru.

Sedangkan:

```go
a := 10

{
    a = 20
}
```

Output:

```
20
```

Karena mengubah variable yang sama.

---

# 👉 Pointer

Pointer menyimpan alamat memory.

```go
x := 10

change(&x)

func change(value *int){
    *value = 100
}
```

Output:

```
100
```

Yang di-copy tetap VALUE.

Tetapi value tersebut adalah alamat memory.

---

# ✅ Go Always Pass By Value

Go selalu menggunakan Pass By Value.

Perbedaannya hanya pada value yang di-copy.

| Type    | Yang di-copy   |
| ------- | -------------- |
| int     | angka          |
| string  | value          |
| pointer | alamat memory  |
| slice   | slice header   |
| map     | runtime header |
| channel | runtime header |

---

# 🆕 new()

Fungsi:

```
Allocate Memory

↓

Zero Value

↓

Return Pointer
```

Contoh:

```go
x := new(int)
```

Return:

```
*int
```

---

# 🛠 make()

Digunakan khusus untuk:

- slice
- map
- channel

Fungsi:

```
Initialize Runtime Structure

↓

Return Value
```

Contoh:

```go
users := make([]User, 0, 10)

m := make(map[string]string)

ch := make(chan int)
```

---

# 📌 Nil vs Empty Slice

```go
var users []User
```

JSON:

```json
{
  "data": null
}
```

Sedangkan

```go
users := []User{}
```

JSON:

```json
{
  "data": []
}
```

Untuk REST API biasanya lebih disukai:

```
[]
```

daripada

```
null
```

---

# ⚠️ Error Handling

Go tidak menggunakan try-catch.

Go menggunakan:

```go
if err != nil {
    return err
}
```

Filosofi:

> **Errors are Values**

Error adalah value biasa yang bisa:

- dikembalikan
- dibandingkan
- dibungkus
- diteruskan

---

# 🏗 Composition over Inheritance

Go tidak menggunakan inheritance seperti Java.

Java:

```
Car IS A Vehicle
```

Go:

```
Car HAS AN Engine
```

Contoh:

```go
type Engine struct{}

type Car struct {
    Engine
}
```

Keuntungan:

- Lebih sederhana
- Lebih mudah dibaca
- Lebih fleksibel
- Tidak memiliki inheritance bertingkat

---

# 🧹 Clean Code Philosophy

Jangan membuat sesuatu sampai benar-benar diperlukan.

Kurang baik:

```
CalculatorService

CalculatorRepository

CalculatorFactory

CalculatorDTO
```

untuk program yang hanya mempunyai empat function.

Lebih baik:

```go
func Add()

func Subtract()

func Multiply()

func Divide()
```

---

# 💬 Comment Philosophy

Komentar harus menjelaskan:

**KENAPA**

bukan

**APA**

Kurang baik:

```go
// function untuk menambahkan dua angka
```

Lebih baik:

```go
// Pre-allocate slice untuk mengurangi reallocation.
```

---

# 🧠 Go Engineer Mindset

Setiap melihat satu baris Go, tanyakan:

- Apa yang di-copy?
- Apa yang terjadi di memory?
- Apa yang ada di Stack?
- Apa yang mungkin ada di Heap?
- Kenapa Go didesain seperti ini?
- Apa trade-off dari pendekatan ini?

---

# 📚 Istilah Penting

## Zero Value

Nilai default setiap tipe data.

## Value Semantics

Copy nilai secara langsung.

## Reference-like Semantics

Copy header/alamat sehingga masih menunjuk data yang sama.

## Slice Header

```
ptr

len

cap
```

## Variable Shadowing

Variable baru dengan nama yang sama pada scope berbeda.

## Pointer

Variable yang menyimpan alamat memory.

## Escape Analysis

Analisis compiler untuk menentukan Stack atau Heap.

## Composition over Inheritance

Menggabungkan object daripada membuat inheritance bertingkat.

## Errors are Values

Error diperlakukan sebagai value biasa.

## Clear is Better than Clever

Kode yang sederhana lebih baik daripada kode yang terlalu kompleks.

---

# 🎯 Rule Selama 30 Hari

Sebelum menulis kode, tanyakan:

1. Apa yang di-copy?
2. Apa yang terjadi di memory?
3. Kenapa Go menyediakan cara ini?
4. Apa trade-off dari pendekatan ini?
5. Apakah ada cara yang lebih sederhana?

---

# 🚀 Target Akhir

Menjadi Backend Engineer Go yang mampu memahami:

- Compiler
- Memory Model
- Pointer
- Interface
- Concurrency
- Clean Architecture
- Performance
- Production Best Practice

> **Goal:** Bukan hanya bisa menulis Go, tetapi memahami alasan di balik setiap baris kode yang ditulis.

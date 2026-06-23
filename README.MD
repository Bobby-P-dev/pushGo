# 📚 Go 30 Days - Day 3

# Pointer Receiver, Value Receiver & Interface Philosophy

---

# 🎯 Tujuan Pembelajaran

Hari ini fokus bukan mempelajari syntax baru, tetapi memahami **cara berpikir Go sebagai sebuah bahasa**.

Target:

- Memahami Pointer Receiver vs Value Receiver
- Memahami Copy Cost
- Memahami Interface Philosophy
- Memahami Implicit Interface
- Memahami Trade-off dalam memilih Pointer atau Value

---

# 1. Struct sebagai Value

Struct di Go secara default adalah Value Type.

Contoh:

```go
type User struct {
	Name string
}

func change(user User) {
	user.Name = "Alice"
}
```

Ketika function dipanggil:

```go
change(user)
```

Yang terjadi adalah:

```text
COPY STRUCT
```

Visualisasi:

```text
Main Stack

user
+----------------+
Name = Bobby
+----------------+

↓

change(user)

↓

Function Stack

user
+----------------+
Name = Bobby
+----------------+
```

Perubahan pada function tidak akan mengubah object asli.

---

# 2. Pointer Receiver

Contoh:

```go
func change(user *User) {
	user.Name = "Alice"
}
```

Pemanggilan:

```go
change(&user)
```

Yang dikirim bukan isi object, tetapi alamat memory.

Visualisasi:

```text
Main Stack

user

↓

0x1000

↓

+----------------+
Name = Bobby
+----------------+

↓

change(&user)

↓

Pointer

↓

0x1000

↓

Object yang sama

↓

Name berubah menjadi Alice
```

Output:

```
Alice
```

---

# 3. Value Receiver

Contoh:

```go
func (u User) GetName() string {
	return u.Name
}
```

Karakteristik:

- Tidak mengubah object asli
- Melakukan copy value
- Cocok untuk object kecil
- Aman digunakan untuk operasi read-only

---

# 4. Pointer Receiver

Contoh:

```go
func (u *User) ChangeName(name string) {
	u.Name = name
}
```

Karakteristik:

- Mengubah object asli
- Menghindari copy struct besar
- Lebih efisien untuk object besar
- Digunakan ketika state perlu diubah

---

# 5. Jangan Gunakan Pointer Hanya Karena Update

Pertanyaan yang harus ditanyakan:

```
Apakah object perlu diubah?

ATAU

Apakah copy object terlalu mahal?
```

Jika jawabannya tidak, maka Value Receiver sering kali lebih sederhana.

---

# 6. Copy Cost

Struct kecil:

```go
type Point struct {
	X int
	Y int
}
```

Copy:

```
X
Y
```

Sangat murah.

---

Struct besar:

```go
type Report struct {
	Data [100000]int
}
```

Value Receiver:

```
COPY

100000 integer
```

Pointer Receiver:

```
COPY

1 alamat memory (pointer)
```

Lebih efisien.

---

# 7. Go Always Pass By Value

Ini adalah konsep yang sangat penting.

Go SELALU menggunakan Pass By Value.

Perbedaannya hanya value yang di-copy.

Contoh:

```
int

↓

copy angka

pointer

↓

copy alamat

slice

↓

copy slice header

map

↓

copy runtime header
```

Pointer bukan Pass By Reference.

Pointer tetap Pass By Value, hanya value yang di-copy berupa alamat memory.

---

# 8. Interface Philosophy

Go mempunyai filosofi yang berbeda dari Java.

Java:

```java
class Dog implements Animal {

}
```

Go:

```go
type Dog struct{}

func (Dog) Speak() {}
```

Tidak ada keyword:

```
implements
```

---

# 9. Implicit Interface

Go menggunakan konsep:

> "If it walks like a duck and quacks like a duck, then it is a duck."

Artinya:

Jika sebuah type mempunyai method yang dibutuhkan oleh interface, maka otomatis dianggap memenuhi interface tersebut.

Contoh:

```go
type Speaker interface {
	Speak()
}

type Dog struct{}

func (Dog) Speak() {}
```

Dog otomatis menjadi Speaker tanpa deklarasi tambahan.

---

# 10. Kenapa Tidak Menggunakan implements?

Java:

```
Dog

↓

implements Animal
```

Jika interface berubah:

```
Animal

Speak()

Walk()
```

Semua implementasi akan error.

---

Go:

```
Dog

↓

Speak()
```

Selama consumer hanya membutuhkan Speak(), maka Dog tetap valid.

Dependency menjadi lebih longgar (Loose Coupling).

---

# 11. Interface Dimiliki Oleh Consumer

Go menganjurkan:

```
Producer

↓

Concrete Struct

↓

Consumer

↓

Interface kecil
```

Bukan:

```
Struct

↓

Memiliki Interface besar
```

---

# 12. Composition over Inheritance

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
- Tidak ada inheritance bertingkat
- Lebih mudah dipelihara
- Lebih fleksibel

---

# 13. Filosofi Penting Go

## Clear is Better than Clever

Kode yang mudah dibaca lebih baik daripada kode yang terlalu pintar.

---

## Errors are Values

Error adalah value biasa yang dapat:

- dikembalikan
- dibandingkan
- dibungkus
- diteruskan

---

## Composition over Inheritance

Menggabungkan object lebih baik daripada inheritance bertingkat.

---

## Implicit Interface

Type tidak perlu mengetahui interface apa yang dia implementasikan.

---

# 14. Engineer Mindset

Jangan bertanya:

```
Apa outputnya?
```

Tetapi bertanya:

```
Apa yang di-copy?

Apa yang terjadi di memory?

Apa trade-off dari pendekatan ini?

Kenapa Go didesain seperti ini?

Apa dampaknya terhadap performa?
```

---

# 15. Catatan Pribadi

## Insight Hari Ini

- Go selalu Pass By Value.
- Pointer tetap melakukan copy, tetapi yang di-copy adalah alamat memory.
- Interface di Go bersifat implicit.
- Struct adalah Value Type secara default.
- Pointer Receiver tidak selalu lebih baik daripada Value Receiver.

## Hal yang Akan Selalu Diingat

> "Framework mengajarkan cara memakai sesuatu."

> "Fundamental mengajarkan kenapa sesuatu dibuat seperti itu."

Target belajar:

**Menjadi Go Engineer yang memahami compiler, memory, runtime, dan design philosophy, bukan hanya programmer yang dapat membuat fitur berjalan.**

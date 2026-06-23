# Day2

Apa yang terjadi di code day 2 ini?

- di user.go saya mempunyai struct user yang di gunakan langsung oleh changename dan getname, di sini kedua nya menggunakan implitic interface yang dimana interface yang mendeskripsikan perilaku yang di butuhkan

- tetapi di change name menggunakan \*user karena untuk mengupdate nilai yang ada di struct user melalui memorry nya langsung contoh seperti ini:

Main Stack

user

↓

0x1000

+----------------+

Name : Bobby

Age : 20

+----------------+

↓

ChangeName()

↓

Pointer

↓

0x1000

↓

Name berubah menjadi Alice

- untuk yang get name menggunakan user saja karena untuk mengambil nilai yang ada di struct user tidak untuk mengubah nilai yang ada di struct user contoh seperti ini hanya copy value nya saja:

Main Stack

user

↓

0x1000

+----------------+

Name : Bobby

Age : 20

+----------------+

↓

GetName()

↓

Value

↓

0x1000

↓

Name = "Bobby"

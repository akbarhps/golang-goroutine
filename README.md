# Golang Goroutine

Sumber Tutorial:
[Udemy](https://www.udemy.com/course/pemrograman-go-lang-pemula-sampai-mahir/learn/lecture/24456458#overview)
[Slide](https://docs.google.com/presentation/d/1A78dn_g6HfxfRor9XBUAGPQM9vT6_SnrQGrQ2z0myOo/edit#slide=id.p)


## Pengenalan Concurrency dan Parallel Programming
---


### Pengenalan Parallel Programming

- Saat ini kita hidup dimana jarang sekali kita menggunakan prosesor yang single core
- Semakin canggih pengangkat keras, maka software pun akan mengikuti, dimana sekarang kita bisa dengan mudah membuat proses parallel dia aplikasi
- Parallel programming sederhananya adalah memecahkan suatu masalah dengan membaginya menjadi yang lebih kecil, dan dijalankan secara bersamaan pada waktu yang bersamaan pula


### Contoh Parallel 

- Menjalankan beberapa aplikasi sekaligus di sistem operasi kita (office, editor, browser, dan lain-lain)
- Beberapa koki menyiapkan makanan di restoran, dimana tiap koki membuat makanan masing-masing
- Antrian di Bank, dimana tiap teller melayani nasabah nya masing-masing


### Process vs Thread

| Process                                          | Thread                                                      |
| ------------------------------------------------ | ----------------------------------------------------------- |
| Process adalah sebuah eksekusi program           | Thread adalah segmen dari process                           |
| Process mengkonsumsi memory besar                | Thread menggunakan memory kecil                             |
| Process saling terisolasi dengan process lainnya | Thread bisa saling berhubungan jika dalam process yang sama |
| Process lama untuk dijalankan dan dihentikan     | Thread cepat untuk dijalakan dan dihentikan                 |


### Parallel vs Concurrency

- Berbeda dengan paralel (menjalankan beberapa pekerjaan secara bersamaan), concurrency adalah menjalankan beberapa pekerjaan secara bergantian
- Dalam parallel kita biasanya membutuhkan banyak Thread, sedangkan dalam concurrency, kita hanya membutuhkan sedikit Thread

![Diagram Parallel](https://user-images.githubusercontent.com/69947442/140457850-6fb5a9b5-7597-4edb-baed-f8ea3fa6ce6a.png)

![Diagram Concurrency](https://user-images.githubusercontent.com/69947442/140457843-d1978601-9b74-4a91-ad33-6646b2657f6e.png)


### Contoh Concurrency

- Saat kita makan di cafe, kita bisa makan, lalu ngobrol, lalu minum, makan lagi, ngobrol lagi, minum lagi, dan seterusnya. Tetapi kita tidak bisa pada saat yang bersamaan minum, makan dan ngobrol, hanya bisa melakukan satu hal pada satu waktu, namun bisa berganti kapanpun kita mau.


### CPU-Bound

- Banyak algoritma dibuat yang hanya membutuhkan CPU untuk menjalankannya. Algoritma jenis ini biasanya sangat tergantung dengan kecepatan CPU.
- Contoh yang paling populer adalah Machine Learning, oleh karena itu sekarang banyak sekali teknologi Machine Learning yang banyak menggunakan GPU karena memiliki core yang lebih banyak dibanding CPU biasanya.
- Jenis algoritma seperti ini tidak ada benefitnya menggunakan Concurrency Programming, namun bisa dibantu dengan implementasi Parallel Programming.


### I/O Bound

- I/O-bound adalah kebalikan dari sebelumnya, dimana biasanya algoritma atau aplikasinya sangat tergantung dengan kecepatan input output devices yang digunakan. 
- Contohnya aplikasi seperti membaca data dari file, database, dan lain-lain.
- Kebanyakan saat ini, biasanya kita akan membuat aplikasi jenis seperti ini.
- Aplikasi jenis I/O-bound, walaupun bisa terbantu dengan implementasi Parallel Programming, tapi benefitnya akan lebih baik jika menggunakan Concurrency Programming.
- Bayangkan kita membaca data dari database, dan Thread harus menunggu 1 detik untuk mendapat balasan dari database, padahal waktu 1 detik itu jika menggunakan Concurrency Programming, bisa digunakan untuk melakukan hal lain lagi


## Pengenalan Goroutine
---

- Goroutine adalah sebuah thread ringan yang dikelola oleh Go Runtime
- Ukuran Goroutine sangat kecil, sekitar 2kb, jauh lebih kecil dibandingkan Thread yang bisa sampai 1mb atau 1000kb
- Namun tidak seperti thread yang berjalan parallel, goroutine berjalan secara concurrent


### Cara Kerja Goroutine

- Sebenarnya, Goroutine dijalankan oleh Go Scheduler dalam thread, dimana jumlah thread nya sebanyak GOMAXPROCS (biasanya sejumlah core CPU)
- Jadi sebenarnya tidak bisa dibilang Goroutine itu pengganti Thread, karena Goroutine sendiri berjalan di atas Thread
- Namun yang mempermudah kita adalah, kita tidak perlu melakukan manajemen Thread secara manual, semua sudah diatur oleh Go Scheduler


### Cara Kerja Scheduler

Dalam Go-Scheduler, kita akan mengenal beberapa terminologi

- G : Goroutine
- M : Thread (Machine)
- P : Processor

![Cara Kerja Go Scheduler](https://user-images.githubusercontent.com/69947442/140458213-e39da53a-81e3-4c04-8ff4-7cef12dcb710.png)


## Membuat Project
---

- Buat folder belajar-golang-goroutine
- Buat module : 

```bash
go mod init belajar-golang-goroutine
```


## Membuat Goroutine
---

- Untuk membuat goroutine di Golang sangatlah sederhana
- Kita hanya cukup menambahkan perintah go sebelum memanggil function yang akan kita jalankan dalam goroutine
- Saat sebuah function kita jalankan dalam goroutine, function tersebut akan berjalan secara asynchronous, artinya tidak akan ditunggu sampai function tersebut selesai
- Aplikasi akan lanjut berjalan ke kode program selanjutnya tanpa menunggu goroutine yang kita buat selesai

```go
func RunHelloWorld() {
    fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
    go RunHelloWorld()
    fmt.Println("Ups")

    time.Sleep(1 * time.Second)
}
```

### Menjalankan Test 

![Output Test](https://user-images.githubusercontent.com/69947442/140458730-36123bc7-7420-4149-849d-4376d7de491b.png)


## Goroutine Sangat Ringan
---

- Seperti yang sebelumnya dijelaskan, bahwa goroutine itu sangat ringan
- Kita bisa membuat ribuan, bahkan sampai jutaan goroutine tanpa takut boros memory
- Tidak seperti thread yang ukurannya berat, goroutine sangatlah ringan

### Kode: Membuat Banyak Goroutine

```go
func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}
```

### Output

![Output Code](https://user-images.githubusercontent.com/69947442/140459012-f5e9634f-50f6-4cee-8187-249a4535a782.png)


## Pengenalan Channel
---

- Channel adalah tempat komunikasi secara synchronous yang bisa dilakukan oleh goroutine
- Di Channel terdapat pengirim dan penerima, biasanya pengirim dan penerima adalah goroutine yang berbeda
- Saat melakukan pengiriman data ke Channel, goroutine akan ter-block, sampai ada yang menerima data tersebut
- Maka dari itu, channel disebut sebagai alat komunikasi synchronous (blocking)
- Channel cocok sekali sebagai alternatif seperti mekanisme async await yang terdapat di beberapa bahasa pemrograman lain


![Diagram Channel](https://user-images.githubusercontent.com/69947442/140459264-84bef74d-acca-4a44-a6ad-0636f378772b.png)


### Karakteristik Channel

- Secara default channel hanya bisa menampung satu data, jika kita ingin menambahkan data lagi, harus menunggu data yang ada di channel diambil
- Channel hanya bisa menerima satu jenis data
- Channel bisa diambil dari lebih dari satu goroutine
- Channel harus di close jika tidak digunakan, atau bisa menyebabkan memory leak


## Membuat Channel
---

- Channel di Go-Lang direpresentasikan dengan tipe data chan
- Untuk membuat channel sangat mudah, kita bisa menggunakan make(), mirip ketika membuat map
- Namun saat pembuatan channel, kita harus tentukan tipe data apa yang bisa dimasukkan kedalam channel tersebut


### Kode: Membuat Channel

```go
channel := make(chan string)
```


### Mengirim dan Menerima Data dari Channel

- Seperti yang sudah dibahas sebelumnya, channel bisa digunakan untuk mengirim dan menerima data
- Untuk mengirim data, kita bisa gunakan kode : channel <- data
- Sedangkan untuk menerima data, bisa gunakan kode : data <- channel
- Jika selesai, jangan lupa untuk menutup channel menggunakan function close()


### Kode: Channel

```go
func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hello World"
	}()

	data := <-channel
	fmt.Println(data)
	// fmt.Println(<-channel) this work to

	time.Sleep(5 * time.Second)
}
```

### Output

![Output](https://user-images.githubusercontent.com/69947442/140465665-05568ccc-7b12-4bbc-89a6-59baef3d5211.png)


## Channel Sebagai Parameter
---

- Dalam kenyataan pembuatan aplikasi, seringnya kita akan mengirim channel ke function lain via parameter
- Sebelumnya kita tahu bahkan di Go-Lang by default, parameter adalah pass by value, artinya value akan diduplikasi lalu dikirim ke function parameter, sehingga jika kita ingin mengirim data asli, kita biasa gunakan pointer (agar pass by reference). 
- Berbeda dengan Channel, kita tidak perlu melakukan hal tersebut
 

### Kode: Channel Sebagai Parameter

```go
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello World"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	go GiveMeResponse(channel)
	data := <-channel
	fmt.Println(data)
}
```


### Output

![Output](https://user-images.githubusercontent.com/69947442/140466282-dac31882-c337-4d4b-ae84-9db661425998.png)


## Channel In dan Out
---

- Saat kita mengirim channel sebagai parameter, isi function tersebut bisa mengirim dan menerima data dari channel tersebut
- Kadang kita ingin memberi tahu terhadap function, misal bahwa channel tersebut hanya digunakan untuk mengirim data, atau hanya dapat digunakan untuk menerima data
- Hal ini bisa kita lakukan di parameter dengan cara menandai apakah channel ini digunakan untuk in (mengirim data) atau out (menerima data)


### Kode: Channel In dan Out

```go
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello World"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyOut(channel)
	go OnlyIn(channel)

	time.Sleep(3 * time.Second)
}
```


### Output

![Output](https://user-images.githubusercontent.com/69947442/140466755-768ce16c-23ff-4c10-84ea-71a347f1c48b.png)



## Buffered Channel
---

- Seperti yang dijelaskan sebelumnya, bahwa secara default channel itu hanya bisa menerima 1 data
- Artinya jika kita menambah data ke-2, maka kita akan diminta menunggu sampai data ke-1 ada yang mengambil
- Kadang-kadang ada kasus dimana pengirim lebih cepat dibanding penerima, dalam hal ini jika kita menggunakan channel, maka otomatis pengirim akan ikut lambat juga
- Untuknya ada Buffered Channel, yaitu buffer yang bisa digunakan untuk menampung data antrian di Channel


### Buffer Capacity

- Kita bebas memasukkan berapa jumlah kapasitas antrian di dalam buffer
- Jika kita set misal 5, artinya kita bisa menerima 5 data di buffer.
- Jika kita mengirim data ke 6, maka kita diminta untuk menunggu sampai buffer ada yang kosong
- Ini cocok sekali ketika memang goroutine yang menerima data lebih lambat dari yang mengirim data


![Diagram Channel Buffer](https://user-images.githubusercontent.com/69947442/140467106-cc79f104-ca1b-431c-b4e3-b02b2874cb7b.png)


### Kode: Membuat Buffered Channel

```go
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Hello"
	channel <- "World"
	channel <- "!"

	fmt.Println(<-channel) // Hello
	fmt.Println(<-channel) // World
	fmt.Println(<-channel) // !

	fmt.Println(cap(channel)) // 3
	fmt.Println(len(channel)) // 0
}
```


## Range Channel
---

- Kadang-kadang ada kasus sebuah channel dikirim data secara terus menerus oleh pengirim
- Dan kadang tidak jelas kapan channel tersebut akan berhenti menerima data
- Salah satu yang bisa kita lakukan adalah dengan menggunakan perulangan range ketika menerima data dari channel
- Ketika sebuah channel di close(), maka secara otomatis perulangan tersebut akan berhenti
- Ini lebih sederhana dari pada kita melakukan pengecekan channel secara manual


### Kode: Range Channel

```go
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- fmt.Sprintf("Perulangan ke %d", i)
		}
		defer close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}
}
```


### Output

![Output](https://user-images.githubusercontent.com/69947442/140469421-37a44532-ffc0-4c12-9ac0-ef8cae2a83e5.png)


## Select Channel
---

- Kadang ada kasus dimana kita membuat beberapa channel, dan menjalankan beberapa goroutine
- Lalu kita ingin mendapatkan data dari semua channel tersebut
- Untuk melakukan hal tersebut, kita bisa menggunakan select channel di Go-Lang
- Dengan select channel, kita bisa memilih data tercepat dari beberapa channel, jika data datang secara bersamaan di beberapa channel, maka akan dipilih secara random


### Kode: Select Channel

```go
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	for counter := 0; counter < 2; {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel1: ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel2: ", data)
			counter++
		}
	}
}
```


### Output

![Output](https://user-images.githubusercontent.com/69947442/140470892-78aecbbb-4e78-4a7e-be21-1efb6443428f.png)


## Default Select
---

- Apa yang terjadi jika kita melakukan select terhadap channel yang ternyata tidak ada datanya?
- Maka kita akan menunggu sampai data ada
- Kadang mungkin kita ingin melakukan sesuatu jika misal semua channel tidak ada datanya ketika kita melakukan select channel
- Dalam select, kita bisa menambahkan default, dimana ini akan dieksekusi jika memang di semua channel yang kita select tidak ada datanya


### Kode: Default Select Channel

```go
func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	for counter := 0; counter < 2; {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel1: ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel2: ", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}
	}
}
```

### Output

![Output](https://user-images.githubusercontent.com/69947442/140472155-58225a32-7f96-4057-94a4-b1c1becc3e5b.png)


## Race Condition
---

### Masalah Dengan Goroutine

- Saat kita menggunakan goroutine, dia tidak hanya berjalan secara concurrent, tapi bisa parallel juga, karena bisa ada beberapa thread yang berjalan secara parallel
- Hal ini sangat berbahaya ketika kita melakukan manipulasi data variable yang sama oleh beberapa goroutine secara bersamaan
- Hal ini bisa menyebabkan masalah yang namanya Race Condition


### Kode: Race Condition

```go
func TestRaceCondition(t *testing.T) {
	counter := 0

	for i := 0; i <= 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				counter++
			}
		}()
	}

	fmt.Println("Counter: ", counter)
}
```


### Output

![Output](https://user-images.githubusercontent.com/69947442/140495868-9ea46806-047b-40dd-8d50-cb4e92eb39a1.png)


## sync.Mutex
---

### Mutex (Mutual Exclusion)

- Untuk mengatasi masalah race condition tersebut, di Go-Lang terdapat sebuah struct bernama sync.Mutex
- Mutex bisa digunakan untuk melakukan locking dan unlocking, dimana ketika kita melakukan locking terhadap mutex, maka tidak ada yang bisa melakukan locking lagi sampai kita melakukan unlock
- Dengan demikian, jika ada beberapa goroutine melakukan lock terhadap Mutex, maka hanya 1 goroutine yang diperbolehkan, setelah goroutine tersebut melakukan unlock, baru goroutine selanjutnya diperbolehkan melakukan lock lagi
- Ini sangat cocok sebagai solusi ketika ada masalah race condition yang sebelumnya kita hadapi


### Kode: Race Condition With Mutex

```go
func TestRaceConditionWithMutex(t *testing.T) {
	counter := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter: ", counter)
}
```


### Output

![Output](https://user-images.githubusercontent.com/69947442/140497383-b21ac7ae-b1b8-4fd9-8100-1307ebb8c542.png)


## sync.RWMutex
---

### RWMutex (Read Write Mutex)

- Kadang ada kasus dimana kita ingin melakukan locking tidak hanya pada proses mengubah data, tapi juga membaca data
- Kita sebenarnya bisa menggunakan Mutex saja, namun masalahnya nanti akan rebutan antara proses membaca dan mengubah
- Di Go-Lang telah disediakan struct RWMutex (Read Write Mutex) untuk menangani hal ini, dimana Mutex jenis ini memiliki dua lock, lock untuk Read dan lock untuk Write


### Kode: RWMutex

```go
type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.RWMutex.Lock()
	defer b.RWMutex.Unlock()

	b.Balance += amount
}

func (b *BankAccount) GetBalance() int {
	b.RWMutex.RLock()
	defer b.RWMutex.RUnlock()

	return b.Balance
}
```


### Output

![Output](https://user-images.githubusercontent.com/69947442/140498513-e335bbfe-592b-4ba8-aa28-516b90c586a3.png)


## Deadlock
---

- Hati-hati saat membuat aplikasi yang parallel atau concurrent, masalah yang sering kita hadapi adalah Deadlock
- Deadlock adalah keadaan dimana sebuah proses goroutine saling menunggu lock sehingga tidak ada satupun goroutine yang bisa jalan
- Sekarang kita coba simulasikan proses deadlock


### Kode: Simulasi Deadlock

```go
type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (u *UserBalance) Lock() {
	u.Mutex.Lock()
}

func (u *UserBalance) Unlock() {
	u.Mutex.Unlock()
}

func Transfer(from, to *UserBalance, amount int) {
	from.Lock()
	defer from.Unlock()

	fmt.Printf("%s Locked\n", from.Name)
	from.Balance -= amount

	time.Sleep(time.Second)

	to.Lock()
	defer to.Unlock()

	fmt.Printf("%s Locked\n", to.Name)
	to.Balance += amount

	time.Sleep(time.Second)
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "User1",
		Balance: 10000,
	}
	user2 := UserBalance{
		Name:    "User2",
		Balance: 10000,
	}

	go Transfer(&user1, &user2, 100)
	go Transfer(&user2, &user1, 200)

	time.Sleep(3 * time.Second)

	fmt.Printf("User %s Balance %d\n", user1.Name, user1.Balance) // expect 10100
	fmt.Printf("User %s Balance %d\n", user2.Name, user2.Balance) // expect 9900
}
```


### Output

![Output](https://user-images.githubusercontent.com/69947442/140501104-91e27001-d3bc-437d-8930-84c716bd73dc.png)


## sync.WaitGroup
---

### WaitGroup

- WaitGroup adalah fitur yang bisa digunakan untuk menunggu sebuah proses selesai dilakukan
- Hal ini kadang diperlukan, misal kita ingin menjalankan beberapa proses menggunakan goroutine, tapi kita ingin semua proses selesai terlebih dahulu sebelum aplikasi kita selesai
- Kasus seperti ini bisa menggunakan WaitGroup
- Untuk menandai bahwa ada proses goroutine, kita bisa menggunakan method Add(int), setelah proses goroutine selesai, kita bisa gunakan method Done()
- Untuk menunggu semua proses selesai, kita bisa menggunakan method Wait()


### Kode: sync.WaitGroup

```go
func RunAsynchronous(group *sync.WaitGroup) {
	group.Add(1)
	defer group.Done()

	fmt.Println("Hello")
	time.Sleep(time.Second)
}

func TestWaitGroup(t *testing.T) {
	var group sync.WaitGroup

	for i := 0; i < 100; i++ {
		go RunAsynchronous(&group)
	}

	group.Wait()
	fmt.Println("Complete")
}
```


### Output

![Output](https://user-images.githubusercontent.com/69947442/140505929-6e4718a1-200e-46c2-94ae-5b89c7c1062f.png)


## sync.Once
---


### Once

- Once adalah fitur di Go-Lang yang bisa kita gunakan untuk memastikan bahsa sebuah function di eksekusi hanya sekali
- Jadi berapa banyak pun goroutine yang mengakses, bisa dipastikan bahwa goroutine yang pertama yang bisa mengeksekusi function nya
- Goroutine yang lain akan di hiraukan, artinya function tidak akan dieksekusi lagi


### Kode: sync.Once

```go
func TestOnce(t *testing.T) {
	var once sync.Once
	var group sync.WaitGroup
	counter := 0

	for i := 0; i < 100; i++ {
		group.Add(1)

		go func() {
			once.Do(func() {
				counter++
			})
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println(counter)
}
```


### Output

![Output](https://user-images.githubusercontent.com/69947442/140507199-87f6cabe-b6c9-4540-a816-2dc88d5fb825.png)


## sync.Pool
---

- Pool adalah implementasi design pattern bernama object pool pattern. 
- Sederhananya, design pattern Pool ini digunakan untuk menyimpan data, selanjutnya untuk menggunakan datanya, kita bisa mengambil dari Pool, dan setelah selesai menggunakan datanya, kita bisa menyimpan kembali ke Pool nya
- Implementasi Pool di Go-Lang ini sudah aman dari problem race condition


### Kode: Membuat Pool

```go
func TestPool(t *testing.T) {
	var pool = sync.Pool{ // override the default sync.Pool
		New: func() interface{} {
			return "Busy"
		},
	}

	pool.Put("Hello")
	pool.Put("World")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Done")
}
```


### Output

![Output](https://user-images.githubusercontent.com/69947442/140508355-f84f43b9-7a96-42dc-b882-c2addc22c504.png)


## sync.Map
---

### Map

- Go-Lang memiliki sebuah struct beranama sync.Map
- Map ini mirip Go-Lang map, namun yang membedakan, Map ini aman untuk menggunaan concurrent menggunakan goroutine
- Ada beberapa function yang bisa kita gunakan di Map :
  - Store(key, value) untuk menyimpan data ke Map
  - Load(key) untuk mengambil data dari Map menggunakan key
  - Delete(key) untuk menghapus data di Map menggunakan key
  - Range(function(key, value)) digunakan untuk melakukan iterasi seluruh data di Map


### Kode: sync.Map

```go
func StoreToMap(group *sync.WaitGroup, data *sync.Map, value int) {
	group.Add(1)
	defer group.Done()

	data.Store(value, value)
}

func TestStoreToMap(t *testing.T) {
	data := sync.Map{}
	group := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go StoreToMap(&group, &data, i)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Printf("key: %v, value: %v\n", key, value)
		return true
	})
}
```


### Output

![Output](https://user-images.githubusercontent.com/69947442/140509289-0e9512b9-0f8d-4418-8694-97f9522c5c04.png)


## sync.Cond
---


### Cond

- Cond adalah adalah implementasi locking berbasis kondisi. 
- Cond membutuhkan Locker (bisa menggunakan Mutex atau RWMutex) untuk implementasi locking nya, namun berbeda dengan Locker biasanya, di Cond terdapat function Wait() untuk menunggu apakah perlu menunggu atau tidak
- Function Signal() bisa digunakan untuk memberi tahu sebuah goroutine agar tidak perlu menunggu lagi, sedangkan function Broadcast() digunakan untuk memberi tahu semua goroutine agar tidak perlu menunggu lagi
- Untuk membuat Cond, kita bisa menggunakan function sync.NewCond(Locker)


### Kode: sync.Cond

```go
var cond = sync.NewCond(&sync.Mutex{})
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()

	cond.L.Lock()
	cond.Wait() // wait for signal

	fmt.Println("Done", value)

	cond.L.Unlock()
}

func TestWaitCondition(t *testing.T) {
	for i := 0; i < 10; i++ {
		group.Add(1)
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			cond.Signal() // send signal to wait condition
		}

		// time.Sleep(time.Second)
		// cond.Broadcast() // send signal to all wait condition
	}()

	group.Wait()
}
```


### Output

![Output](https://user-images.githubusercontent.com/69947442/140510856-f303e4a6-1257-47a0-b333-1e76a6d55e96.png)


## Atomic
---

- Go-Lang memiliki package yang bernama sync/atomic
- Atomic merupakan package yang digunakan untuk menggunakan data primitive secara aman pada proses concurrent
- Contohnya sebelumnya kita telah menggunakan Mutex untuk melakukan locking ketika ingin menaikkan angka di counter. Hal ini sebenarnya bisa digunakan menggunakan Atomic package
- Ada banyak sekali function di atomic package, kita bisa eksplore sendiri di halaman dokumentasinya
- https://golang.org/pkg/sync/atomic/ 


### Kode: Atomic

```go
func TestAtomic(t *testing.T) {
	var counter int64 = 0
	group := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		group.Add(1)

		go func() {
			defer group.Done()

			for j := 0; j < 100; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	group.Wait()
	fmt.Println(counter)
}
```


### Output

![Output](https://user-images.githubusercontent.com/69947442/140511766-4a1e18cf-a1ad-4af5-8631-2ea79f6143e2.png)


## time.Timer
---

- Timer adalah representasi satu kejadian
- Ketika waktu timer sudah expire, maka event akan dikirim ke dalam channel
- Untuk membuat Timer kita bisa menggunakan time.NewTimer(duration)


### Kode: time.Timer

```go
func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}
```


### Output

![Output](https://user-images.githubusercontent.com/69947442/140512377-8e711079-f81a-4284-b4b7-127a222e17df.png)



### time.After()

- Kadang kita hanya butuh channel nya saja, tidak membutuhkan data Timer nya
- Untuk melakukan hal itu kita bisa menggunakan function time.After(duration)


### Kode: time.After()

```go
func TestTimeAfter(t *testing.T) {
	channel := time.After(time.Second)
	time := <-channel
	fmt.Println(time)
}
```


### Output

![Output](https://user-images.githubusercontent.com/69947442/140512810-a460dae3-cfc1-4d35-9266-c13c8abac45a.png)


### time.AfterFunc()

- Kadang ada kebutuhan kita ingin menjalankan sebuah function dengan delay waktu tertentu
- Kita bisa memanfaatkan Timer dengan menggunakan function time.AfterFunc()
- Kita tidak perlu lagi menggunakan channel nya, cukup kirim kan function yang akan dipanggil ketika Timer mengirim kejadiannya


### Kode: time.AfterFunc()

```go
func TestTimeAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(time.Second, func() {
		defer group.Done()
		fmt.Println(time.Now())
	})

	fmt.Println(time.Now())
	group.Wait()
}
```


### Output

![Output](https://user-images.githubusercontent.com/69947442/140513551-7ac65f01-7ddc-436b-af70-59ed8784ed09.png)


## time.Ticker
---

- Ticker adalah representasi kejadian yang berulang
- Ketika waktu ticker sudah expire, maka event akan dikirim ke dalam channel
- Untuk membuat ticker, kita bisa menggunakan time.NewTicker(duration)
- Untuk menghentikan ticker, kita bisa menggunakan Ticker.Stop()


### Kode: time.Ticker

```go
func TestTimeTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second)

	go func() {
		time.Sleep(time.Second * 5)
		ticker.Stop()
	}()

	for time := range ticker.C {
		fmt.Println(time)
	}
}
```


### Output

![Output](https://user-images.githubusercontent.com/69947442/140514483-89290a95-2acf-4e50-8731-5b5fd160c7bb.png)



### time.Tick()

- Kadang kita tidak butuh data Ticker nya, kita hanya butuh channel nya saja
- Jika demikian, kita bisa menggunakan function timer.Tick(duration), function ini tidak akan mengembalikan Ticker, hanya mengembalikan channel timer nya saja


### Kode: time.Tick()

```go
func TestTimeTick(t *testing.T) {
	ticker := time.Tick(time.Second)

	for time := range ticker {
		fmt.Println(time)
	}
}
```


### Output

![Output](https://user-images.githubusercontent.com/69947442/140514838-69dbac8e-6541-4b2a-ade1-d305ffba2bae.png)


## GOMAXPROCS
---

- Sebelumnya diawal kita sudah bahas bahwa goroutine itu sebenarnya dijalankan di dalam Thread
- Pertanyaannya, seberapa banyak Thread yang ada di Go-Lang ketika aplikasi kita berjalan?
- Untuk mengetahui berapa jumlah Thread, kita bisa menggunakan GOMAXPROCS, yaitu sebuah function di package runtime yang bisa kita gunakan untuk mengubah jumlah thread atau mengambil jumlah thread
- Secara default, jumlah thread di Go-Lang itu sebanyak jumlah CPU di komputer kita. 
- Kita juga bisa melihat berapa jumlah CPU kita dengan menggunakan function runtime.NumCpu()


### Kode: Melihat Jumlah Thread

```go
func TestGetGoMaxProcs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("totalCpu:", totalCpu)

	maxProcs := runtime.GOMAXPROCS(-1)
	fmt.Println("maxProcs:", maxProcs)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("totalGoroutine:", totalGoroutine)
}
```


### Output

![Output](https://user-images.githubusercontent.com/69947442/140515871-868a69bd-a9c9-4067-a019-47a5022f09ab.png)



### Kode: Mengubah Jumlah Thread

```go
runtime.GOMAXPROCS(numOfThread)
```


### Output

![Output](https://user-images.githubusercontent.com/69947442/140516070-7bdc0f1b-eccd-4e40-94d8-12eda991623e.png)



### Peringatan

- Menambah jumlah thread tidak berarti membuat aplikasi kita menjadi lebih cepat
- Karena pada saat yang sama, 1 CPU hanya akan menjalankan  1 goroutine dengan 1 thread
- Oleh karena ini, jika ingin menambah throughput aplikasi, disarankan lakukan vertical scaling (dengan menambah jumlah CPU) atau horizontal scaling (menambah node baru)
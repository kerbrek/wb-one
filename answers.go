// Устные вопросы

// 1. Какой самый эффективный способ конкатенации строк?

// Тип Builder из пакета strings и его методы WriteString и String.
// https://pkg.go.dev/strings#Builder

// 2. Что такое интерфейсы, как они применяются в Go?

// Механизм абстракции. Интерфейсы описывают поведение (сигнатуры методов),
// которое конкретные типы должны реализовать. Это позволяет использовать
// разные конкретные типы там, где требуется интерфейсный тип (при условии
// что они реализуют поведение интерфейса). И все это проверяется компилятором.
// https://go.dev/doc/effective_go#interfaces_and_types
// Вот несколько встроенных интерфейсов:

// error https://pkg.go.dev/builtin#error
type error interface {
	Error() string
}

// Stringer https://pkg.go.dev/fmt#Stringer
type Stringer interface {
	String() string
}

// Reader https://pkg.go.dev/io#Reader
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Writer https://pkg.go.dev/io#Writer
type Writer interface {
	Write(p []byte) (n int, err error)
}

// 3. Чем отличаются RWMutex от Mutex?

// Блокировка Mutex может удерживаться только одним читателем/писателем.
// А блокировка RWMutex может удерживаться произвольным числом читателей или
// одним писателем.

// 4. Чем отличаются буферизированные и не буферизированные каналы?

// С не буферизированным каналом отправка и получение в/из канала блокируется
// пока отправитель и получатель не готовы. С буферизированным каналом
// отправка проходит успешно без блокировки пока буфер не заполнен, а получение
// проходит успешно без блокировки пока буфер не пуст.
// https://go.dev/ref/spec#Channel_types

// 5. Какой размер у структуры struct{}{}?

// 0 байт.
// https://dave.cheney.net/2014/03/25/the-empty-struct

// 6. Есть ли в Go перегрузка методов или операторов?

// Нет.
// https://go.dev/doc/faq#overloading

// 7. В какой последовательности будут выведены элементы map[int]int?
// Пример:
// m[0]=1
// m[1]=124
// m[2]=281

// Рандомно. Порядок обхода мапы случаен, и не гарантируется, что он будет
// одинаковым от одного обхода к другому.
// https://go.dev/ref/spec#RangeClause

// 8. В чем разница make и new?

// new выделяет память для переменной указанного типа и возвращает указатель.
// Переменные инициализируются нулевыми значениями.
// https://go.dev/ref/spec#Allocation
// https://go.dev/doc/effective_go#allocation_new

// make создает слайсы, мапы и каналы, и возвращает значение (не указатель).
// https://go.dev/ref/spec#Making_slices_maps_and_channels
// https://go.dev/doc/effective_go#allocation_make

// 9. Сколько существует способов задать переменную типа slice или map?

var s []int
var s = []int{}
var s = make([]int, 0, 7)
s := []int{}
s := make([]int, 0, 7)

// 10. Что выведет данная программа и почему?

func update(p *int) {
	b := 2
	p = &b
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(p)
	fmt.Println(*p)
}

// Выведет:
//   1
//   1
// В функции update мы всего лишь изменяем значение локальной переменной p.

// 11. Что выведет данная программа и почему?

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}

// Выведет:
//   4
//   2
//   3
//   0
//   1
//   fatal error: all goroutines are asleep - deadlock!
// В аргументах горутин передаются копии переменной wg. Нужно либо использовать
// указатели, либо напрямую использовать переменную wg, так как горутины замыкают ее.

// 12. Что выведет данная программа и почему?

func main() {
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n)
}

// Выведет:
//   0
// В блоке if мы создали новую локальную переменную блока if, затенив внешнюю переменную.

// 13. Что выведет данная программа и почему?

func someAction(v []int8, b int8) {
	v[0] = 100
	v = append(v, b)
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
}

// Выведет:
//   [100 2 3 4 5]
// В функцию someAction мы передаем копию дескриптора слайса. При изменени
// нулевого элемента слайса мы изменям элемент нижележащего массива. Так как
// длина и вместимость слайса равны, при добавлении элемента в слайс в локальную
// переменную функции v запишется новый дескриптор слайса с указателем на новый
// увеличенный нижележащий массив (в него скопируются элементы из старого и добавится
// новый элемент).
// https://go.dev/blog/slices-intro
// https://go.dev/ref/spec#Appending_and_copying_slices

// 14. Что выведет данная программа и почему?

func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}

// Выведет:
//   [b b a][a a]
// В функцию мы передаем копию дескриптора слайса. Так как длина и вместимость
// слайса равны, при добавлении элемента в слайс в локальную переменную функции
// slice запишется новый дескриптор слайса с указателем на новый увеличенный
// нижележащий массив (в него скопируются элементы из старого и добавится новый
// элемент). И далее мы изменяем уже новый слайс (точнее его новый нижележащий
// массив).

# go语言基础

## day01

### **1.变量定义**

先赋值在定义，两步完成

```go
var name1 string          //先声明
name1 = "这是使用先定义再赋值再使用的值" //再赋值
fmt.Println(name1)
```

定义的时候声明变量类型并赋值，一行完成

```go
var name2 string = "声明并赋值"
fmt.Println(name2)
```

使用`:=`简短使用，不需要声明变量的类型

```
name3 := "简短声明使用:="
fmt.Println(name3)
```

局部变量和全局变量

```go
func hello() {
	//函数内的局部变量作用域为当前的函数内
	age1 := 13
	fmt.Println("这是hello函数内的 ", age1, sex)
}

	// 这是全局变量，全局变量定义可以不使用；必须使用var
var sex string = "男"
```

定义常量，使用const定义，必须定义时赋值，且无法修改

```go
const version string = "1.0.0"
```

同时定义多个变量

```go
//使用逗号隔开变量，一一对应
var a1, a2 = 1, 2
fmt.Println(a1, a2)
//使用括号的方式
var (
b1 string = "b1"
b2 string = "b2"
)
fmt.Println("b1的值为", b1, "b2的值为", b2)
```

访问其他包的变量

```go
package book

//变量名需要首字母大写，使用时小写会报错
var Version string = "2.0"//✅️行为
var name string = "十万个冷笑话"//❌️行为
const EditorBy string = "alen" //✅️行为
```

```go
func main() {
	fmt.Println(book.Version) 
	fmt.Println(book.name) 
    fmt.Println(book.EditorBy)
}
```

## day02

### 1.输入输出

**常用三种输出**

println，自带换行

```go
fmt.Println("println 自带换行")
```

print，不会换行

```go
fmt.Print("没有换行")
```

printf，格式化输出

```go
fmt.Printf("%s格式化输出%s\n", name)       //该输出语句会有一个missing,因为第二个%s没有值
fmt.Printf("%s格式化输出%s\n", name, "你好") //"你好"这个字符串传给了第二个%s
```

**格式化输出**

%s：格式化为字符串输出

%d：格式化为整数类型输出

%f：格式化为浮点数输出，%.2f表示保留两位小数

%v：输出值，对控制不友好，%#v会加上双引号，输出控制较为友好

```go
fmt.Printf("%s格式化输出%s\n", name)       //该输出语句会有一个missing,因为第二个%s没有值
fmt.Printf("%s格式化输出%s\n", name, "你好") //"你好"这个字符串传给了第二个%s
fmt.Printf("这是整数类型%d \n", 3)
fmt.Printf("这是浮点数类型%.4f \n", 3.1415926)
fmt.Printf("这是字符串类型%s \n", "hello")
fmt.Printf("打印变量的类型%T %T \n", "hello", 1)
fmt.Printf("打印值：%v \n", "")
fmt.Printf("打印值：%#v \n", "")
```

**输入**

从控制台输出并传入给变量

```go
var name1 string
fmt.Print("请输入你的名称：")
fmt.Scan(&name1)
```

输入的值和变量类型不一致时，会有错误信息，例如下方，age需要输入int类型，在控制台输入alen时，会输出

```go
var age int //age为int类型
fmt.Print("请输入你的年龄：")
n, err := fmt.Scan(&age)   //n err是Scan自带的
fmt.Println(n, err, age)
==================
//age需要输入int类型，在控制台输入alen时，会输出以下内容
请输入你的年龄：alen
0 expected integer 0	//此处的两个0为int类型的默认值
```

### 2.基本数据类型

**整数类型**

```go
//uint8表示无符号位，存正整数
var u8 uint8 = 255 //存无符号位的正整数
// 0 0 0 0 0 0 0 0 = 2^8-1 = 255
```

```go
//int8表示第一位为符号位，存整数
var a1 int8 = -127
// 0   0 0 0 0 0 0 0
//第一位表示符号
```

**浮点**

float32和float64表示，多少位的小数，无声明时，默认为float64

```go
var f1 float32 = 1.1
var f2 float64 = 1.2
fmt.Println(f1, f2)
```

**字符类型**

byte：值范围与uint8一致，ASCII码中的值

```go
var c1 byte = 'a'
var c2 int8 = 97
fmt.Printf("%d,%c", c1, c2)
//'a'在ASCII码中的值为97，%c将c2的97转为了'a'
```

rune：取值范围和uint32,中日韩等字符

```go
var c3 rune = '傻'
var c4 int32 = 20667
fmt.Printf("%d,%c\n", c3, c4)
```

字符串，双引号括起来

**转义字符**

```go
fmt.Println("制表符1\t制表符2")
fmt.Println("\"转义\"")
fmt.Println(`hello
word
,,,

`)
==========================
制表符1	制表符2
"转义"
hello
word
,,,
```

**布尔类型**

值为true和false，默认值为false，无法做类型转换

```go
var b1 bool = true
var b2 bool = false
fmt.Println(b1, b2)
```

**零值问题**

基本变量类型 定义时不赋值时，默认值为0、空、false……

```go
var q int
var w float32
var e string
var r rune
var t bool
fmt.Println(w, q, w, e, r, t)
//   0 0 0  0 false
```

## day03

### 1.数组

数组的长度在定义时就已经固定无法增加长度，但是可以修改某个元素的值

定义一个长度为3，字符类型为string的数组，并给每个元素赋值

```go
var nameList [3]string = [3]string{"alen", "jack", "tom"}
```

也可以先定义再赋值

```go
var nameList2 [3]int
nameList2 = [3]int{21, 22, 23}
```

修改某个元素的值

```go
nameList[0] = "tim"
fmt.Println(nameList)
```

数组的第一个元素的索引为0，从左往右数，查看数组中某个元素的值时，使用变量名[索引]几个查看到第‘索引+1’位置元素的值，例如查看第1个元素的值为

```go
fmt.Println(nameList[0])
```

不论数组有多长，查看最后一个元素的值

```go
//len(nameList)，会输出nameList的长度
fmt.Println(nameList[len(nameList)-1])
```

### 2.切片

理解为数组plus版本，可以增加元素

定义一个长度为空的数组，即为切片，定义时赋值。

```go
var nameList []string = []string{"tom", "tim", "alen"}
```

先定义再赋值

```go
var ageList []int
ageList = []int{10, 20, 30}
```

直接打印空切片会报错

```go
var sexList []string
fmt.Println(sexList)	//打印整个数组会输出[]
fmt.Println(sexList[0])	//打印第1个元素时，会报错
fmt.Println(sexList == nil)	//结果为true
```

通常在定义一个切片的时候需要初始化一下，有以下几种方式

定义切片时，初始化一下

```go
var nameList1 []string = []string{}
fmt.Println(nameList1 == nil)

=====================
false//表示该切片不为空
```

使用**make函数**，make(元素类型,长度)

```go
nameList2 := make([]string, 0)
fmt.Println(nameList2 == nil)
=====================
false
```

创建一个全是0的切片

```go
ageList2 := make([]int,3)
fmt.Println(ageList2)
```

切片增加一个元素使用append方法

```go
ageList3 := make([]int ,0)
ageList3 = append(ageList3,7)
```

定义一个数组，有五个元素，切出从第几个到第几个的元素

```go
array := [5]string{"tom", "tim", "alen", "jack", "jim"}
slices := array[3:5]
fmt.Println(slices)
//array可以看成{🔪"tom", 🔪"tim", 🔪"alen",🔪"jack", 🔪"jim"🔪}，从第一个元素的左边开始切一刀，直到切到最后一个元素的右边。取某一串值可以在[开始🔪:结束🔪]，第一个位置的🔪索引为0。
//所以，我想要得到alen到jim的值为：alen左边的🔪到jim右边的🔪的索引号。
```

**切片排序**

```go
ints := []int{231, 234, 25, 12}
fmt.Println(ints) //原封不动输出
sort.Ints(ints)   //升序
fmt.Println(ints)
sort.Sort(sort.Reverse(sort.IntSlice(ints))) //降序
fmt.Println(ints)
```

## day04

### 1.map

map，定义后需要初始化

```go
//定义并初始化
var userMap map[int]string = map[int]string{
	1: "alen",
	2: "jack",
	3: "",
}
fmt.Println(userMap) //打印map
```

使用make函数初始化

```go
var ageMap map[string]int = make(map[string]int)
ageMap["alen"] = 18
fmt.Println(ageMap)
```

增加一个key、value

```
userMap[4] = "tom"
fmt.Println(userMap)
```

将key为1的value修改为hello

```go
userMap[1] = "hello"
```

当没有key为5的时候，可以使用该方法查看

```go
fmt.Printf("%#v\n", userMap[5])
value, ok := userMap[5]
fmt.Println(value, ok)
=======================================
 false
```

当key为3的value为空值时使用上述方法返回为true

```go
value1, ok1 := userMap[3]
fmt.Println(value1, ok1)
```

### 2.if判断

中断式（推荐），只会走一个if判断

```go
//中断式
var age int
fmt.Println("请输入你的年龄：")
fmt.Scan(&age)
if age <= 0 {
	fmt.Println("未出生")
	return
}
if age <= 18 {
	fmt.Println("未成年")
	return
}
if age <= 35 {
	fmt.Println("中年")
	return
}
```

嵌入式，先用一个判断条件将几个条件分为两部分，在从某一部分中进行逐一判断

```go
if age <= 18 {
	if age <= 0 {
		fmt.Println("未出生")
	} else {
		fmt.Println("未成年")
	}
} else {
	if age <= 35 {
		fmt.Println("青年")
	} else {
		fmt.Println("中年")
	}
}
```

多条件判断式

```go
//多条件判断式
var age int
fmt.Println("请输入你的年龄")
fmt.Scan(&age)

if age <= 0 {
	fmt.Println("未出生")
}
if age > 0 && age <= 18 {
	fmt.Println("未成年")
}
if age > 18 && age < 35 {
	fmt.Println("青年")
}
if age >= 35 {
	fmt.Println("中年")
}
```

## day05

### 1.switch

多选一，遇到匹配到一个case后就不会往下走了

```go
switch {
case age <= 0:
	fmt.Println("未出生")
case age <= 18:
	fmt.Println("未成年")
case age <= 35:
	fmt.Println("青年")
default:
	fmt.Println("中年")
}
===============
请输入你的年龄：15
未成年
```

当匹配一个值还想继续往下走时，在那个case最后一个语句下加上`fallthrough`

```go
switch {
case age <= 0:
	fmt.Println("未出生")
case age <= 18:
	fmt.Println("未成年")
	fallthrough
case age <= 35:
	fmt.Println("青年")
default:
	fmt.Println("中年")
}
===============
请输入你的年龄：15
未成年
青年
```

第二种写法

```go
switch week {
case 1, 2, 3, 4:
	fmt.Println("😭")
case 5:
	fmt.Println("😊")
case 6, 7:
	fmt.Println("非常😊")
}
```

## day06

### 1.for循环

**for 初始化;判断条件;迭代{}**

```go
//从1加到100
var sum int
for i := 1; i <= 100; i++ {
	sum = sum + i
	//sum += i
}
fmt.Println(sum)
```

for死循环三种写法

```go
//判断条件直接定义为true，i不参与判断
for i := 1; true; i++ {
	fmt.Println(time.Now())
	time.Sleep(time.Second)
}
```

```go
//i不参与条件判断，所以不需要定义i
for true {
	fmt.Println(time.Now())
	time.Sleep(time.Second)

```

```go
//最后简化版本
for {
	fmt.Println(time.Now())
	time.Sleep(time.Second)
}
```

**for循环实现while逻辑**，先判断再执行

```go
var sum int
var i int = 1
for i <= 100 {
	sum += i
	i++
}
fmt.Println(sum)
```

do while模式，先执行再判断

```go
var sum int
var i int = 1
for{
	sum += i
	i++
	if i>100 {
	break
	}
}
fmp.Println(sum)
```

**遍历数组切片Map**

遍历数组

```go
var List [3]string = [3]string{"hello", "world", "!"}
for i := 0; i < len(List); i++ {
	fmt.Println(i, List[i])
}
```

遍历切片使用

```go
//常规for遍历
var ageList []int = []int{12,13,14,15}
for i:=0;i<len(ageList);i++{
    fmt.Println(i,ageList[i])
}
//for range
var nameList []string = []string{"hello","world","!"}
for index,item := range nameList{
	fmt.Println(index,item)
}
```

使用for range的方式**遍历Map**

```go
var nameMap map[int]string = map[int]string{1001:"alen",2003:"tom",3:"jack"}
for key,value := range nameMap{
    fmt.Println(key,value)
}
```

break:跳出循环体

```go
//打印1-10
for i := 1; i <= 10; i++ {
	if i == 5 {
		break
	}
	fmt.Printf("第%d次循环\t", i)
}
================================
第1次循环	第2次循环	第3次循环	第4次循环	
```

continue：跳出本次循环，不再执行本次循环的剩余部分内容

```go
//打印1-10
for i := 1; i <= 10; i++ {
	if i == 5 {
		//break
		continue
	}
	fmt.Printf("第%d次循环\t", i)
}
```

### 2.函数

简单的函数，没有不需要传参

```go
func sayHello() {
	fmt.Printf("Hello World")
}
func main() {
	sayHello()
}
```

传递参数给函数，传参类型为int型

```go
func param1(id int) {
	fmt.Println(id)
}
func main() {
	param1(3)
}
```

传递两个参数，并且两个参数的变量类型都为int

```go
func param2(id, age int) {
	fmt.Println(id, age)
}
func main() {
	param2(3, 4)
```

传递两个不同的参数，分别为int和string

```go
func param3(id int, name string) {
	fmt.Println(id, name)
}
func main() {
	param3(14, "alen")
}
```

定义一个函数，将传参的值相加求和

```go
func addNum(numberList ...int){
    var sum int
    for _,item := range numberList{
        sum += item
    }
    fmt.Println(sum)
}
func main() {
	addNum(1, 2, 3, 4, 5)
}
```

返回值

一个返回值

```go
func r1(getName string) string{
    return getName
}
```

两个返回值

```go
func r2(nameList ...int)(ium int,sum int){
    for ium,sum := range(nameList){
        ium +=1
        sum += sum  
    }
    return ium,sum
}
```

定义一个函数，输入一个切片，返回元素的个数和切片的和

```go
func r2(getList []int) (indexSum, itemSum int) {
	for index, item := range getList {
		indexSum = index + 1
		itemSum += item
	}
	return indexSum, itemSum
}
func main() {
	var nameList []int = []int{1, 2, 3, 4}
	fmt.Println(r2(nameList))
    //	
    index, sum := r2(nameList)
	fmt.Println(index, sum)
}
=================================
4 10
```

return后不加返回值和变量时，默认返回定义函数时返回的变量

```go
func r4() (val string, ok bool) {
	if 1 < 2 {
		val = "11"
	}
	return
}
func main() {
	fmt.Println(r4())
}
```

## day07

### 1.匿名函数

在一个函数内定义一个函数，直接定义是不允许的，可以使用定义一个变量的方式定义一个函数，变量的类型是函数类型

定义一个匿名函数，使用return返回值

```go
func main(){
    var addSum = func () string{
        return "alen1"
    }
    fmt.Println(addSum())
}
```

定义一个函数，传一个name值传入函数

```go
func main(){
    var SetName = func (name string){
        fmt.Println(name)
    }
    SetName("alen2")
}
```

### 2.高阶函数

定义一个Map存放每个选项，1 登录、2 注册、3 用户中心，每个Map的key是选项，value是对于操作的函数

```go
import "fmt"

func main() {
	fmt.Println("请输入你要进行的操作")
	fmt.Println("1.登录")
	fmt.Println("2.注册")
	fmt.Println("3.个人中心")
	//定义一个变量接受用户输入的值
	var index int
	fmt.Scan(&index)

	//定义一个map，key为选项==index，value为所对应的函数操作
	var menu = map[int]func(){
		1: login,
		2: register,
		3: userCenter,
	}
	//将menu这个map的value赋值给fun
	fun, ok := menu[index]
	//如果这个key不存在ok为false，如果值存在，ok为true
	if ok {
		fun()
	}

}

func login() {
	fmt.Println("登录")
}
func register() {
	fmt.Println("注册")
}
func userCenter() {
	fmt.Println("用户中心")
}

```

**闭包**

一个函数的返回值也是一个函数，在内层的函数用到了外层的函数称为**闭包**

定义一个函数，func(2)(1,2,3)，实现效果为等待2s后将后面的括号内的参数求和

定义的函数addAwait的返回值也是一个函数，

```go
func addAwait(sec int) func() int{
    return func(nameList ...int) int{
        time.Sleep(time.Duration(sec) * time.Second)
        var sum int
        for _,i := range nameList{
            sum += i
        }
        return sum
    }
    
}
func main(){
    
}
```

### 3.指针

**值传递**：直接传递值给函数会重新开辟一个新的地址

```go
func Copy(fname string) {
	fmt.Printf("%p\n", &fname)

}

func main() {

	var name string = "alen"
	Copy(name)
	fmt.Printf("%p\n", &name)
}
```

**引用传递**

```go
func copy(fname *string) { //定义一个函数，他的参数为引用类型 *表示，表示他要接受的值是一个内存地址
	fmt.Println(fname) //此时fname存的值是name的内存地址
	*fname = "tom"     //修改 fname 指针指向的内存中的值，因为fname和name指向同一个内存地址的值，所以修改这个值为tom，name的值也会修改为tom
}

func main() {
	var name string = "alen"
	fmt.Printf("%p,%v\n", &name, name) //将name的内存地址打印
	copy(&name)                        //将变量name的内存地址的值传给函数
	fmt.Printf("%p,%v\n", &name, name) //将name的内存地址打印

}
```

## day08

### 1.init函数

init函数在main函数之前执行，没有参数的输入和返回值

```go
func init() {
	fmt.Println("init1")
}
func init() {
	fmt.Println("init2")
}
func main() {
	fmt.Println("main")
}
```

### 2.defer函数

在return前执行的函数，先进后出，离return越近的defer函数越先执行

```go
func main() {
	defer func() {
		fmt.Println("defer1")
	}()
	defer fmt.Println("defer2") 	//另外一种定义方式
	return
}
===============================
defer2
defer1
```

其他代码执行完之后才会执行defer函数

先定义name的值为alen，在defer函数后修改值为tom

```go
func main() {
	var name string = "alen"
	defer func() {
		fmt.Println(name)
	}()
	name = "tom"
	return
}
===============================
tom
```

### 3.结构体

定义一个结构体，并创建一个对象

```go
type Student struct {
	Name string		//Student的成员Name为string类型
	Age  int
}
func main() {
	s1 := Student{Name: "alen", Age: 18}		//创建一个对象
	fmt.Println(s1.Name)
}
```

给结构体绑定一个方法

```go
func (s Student) Study() {		//Study是Student的一个方法
	fmt.Printf("%s is studying", s.Name)
}
func main() {
	s1 := Student{Name: "alen", Age: 18}
	fmt.Println(s1.Name)
	s1.Study()
}
```

### 4.继承

再定义一个Class班级的结构体，Student的成员中有Class，再添加一个显示Student的Class信息

```go
type Class struct { 	//定义一个Class结构体，含有Name一个成员
	Name string
}
type Student struct {
	Name string		//Student结构体中也含有一个Name成员
	Age  int
	Class
}

func (s Student) Info() {
	fmt.Printf("%s今年%d岁班级在：%s\n", s.Name, s.Age, s.Class.Name)	//在此处同名需要先加上是哪个结构体
}
func (s Student) Study() {
	fmt.Printf("%s 正在学习", s.Name)
}

func main() {
	c1 := Class{Name: "三年级"}		//需要先给Class结构体创建一个对象，才能在下面赋值
	s1 := Student{Name: "alen", Age: 14, Class: c1}
	s1.Info()
}
```


# for 循环
**Go语言中只有for循环一个循环结构**

```
for 表达式1;表达式2;表达式3{
  //循环体
}
```

* for循环用的最多的地方就是遍历数组或切片等



* 经典for循环结构中 , for关键字后面有三个表达式,且每个表达式都可以省略

```go
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//等价于
	j := 0
	for ; j < 5; {
		fmt.Println(j)
		j++
	}
```

* for关键字后面也可以只有一个表达式,表示如果条件成立执行循环体代码

```go
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//等价于
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}
```

* 可以通过数组的长度判断循环结束条件遍历整个数组

```go
	arr := [3]string{"小猪", "胖", "胖猪"}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
```

* for循环遍历数组等时常与range结合使用
  * range函数返回两个值,第一个是脚标,第二个是内容

```go
	arr := [3]string{"小猪", "胖", "胖猪"}
	for i, n := range arr {
		//其中n=arr[i]
		fmt.Println(i, n)
	}
```

## 一. 双重for循环

* 可以在循环中执行循环,称为双重for循环
* 代码示例

```go
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(i, j)
		}
	}
```

* 上面代码中注意:
  * 只有当子循环完全执行结束才能再次执行外侧循环.因为循环体执行完才能执行表达式3

## 二.冒泡排序

* 排序就是把一组数据按照特定的顺序重新排列.可以是升序,降序等
* 冒泡排序利用双重for循环把最大(小)的值移动到一侧,每次可以判断出一个数据,如果有n个数组,执行n-1次循环就可以完成排序
* 排序代码(升序还是降序主要是看`if`判断是大于还是小于)

```go
	arr := [5]int{1, 7, 3, 6, 2}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
```


## break & continue
语义与其他语言一致 😀

### 一.continue

* continue关键字控制结束本次 **循环体 **结束,执行表达式三.

```go
	for i := 0; i < 5; i++ {
		fmt.Println("开始")
		if i == 2 || i == 3 {
			continue
		}
		fmt.Println("结束")
	}
```

* 在双重for循环中continue默认影响最内侧循环,与最外层循环无关

```
go
func main() {
	for k := 0; k < 2; k++ {
		for i := 0; i < 3; i++ {
			if i == 1 {
				continue
			}
			fmt.Println(k, i, "结束")
		}
	}
```

* Go语言执行标签写法,可以通过定义标签让continue控制影响哪个for循环

```
go
	myfor:for k := 0; k < 2; k++ {
		for i := 0; i < 3; i++ {
			if i == 1 {
				continue myfor
			}
			fmt.Println(k, i, "结束")
		}
	}
```

### 二. break

* break可以中断for循环,无论for循环还有几次执行,立即停止

```
go
	for i := 0; i < 5; i++ {
		if i == 2 {
			break
		}
		fmt.Println(i)
	}
```

* 在双重for循环中,break默认也影响到最近的for循环

```go
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if j == 1 {
				break
			}
			fmt.Println(i, j)
		}
	}
```

* **break也可以通过定义标签,控制break对哪个for循环生效**

```
go
	myfor:for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if j == 1 {
				break myfor
			}
			fmt.Println(i, j)
		}
	}
```


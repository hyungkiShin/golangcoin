## ch 1 
```go
변수 기초
---
package main

func main() {
	// var name string = "nico"
	// name := "nico" // 위 결과와 같다. : 는 create = 는 업데이트 + 함수 밖에선 사용 할 수 없다.
}

함수 기초
----
package main

import "fmt"

// tip ( 두개의 타입이 같다면 생략 가능 + 두개의 타입 명시 할시 타입 추가. 물론 사용하는 곳에서도. 받는 변수도 모두 추가.)
func plus(a, b int, name string) (int, string) {
	return a + b, name
}

func main() {
	result, name := plus(1, 2, "hi")
	fmt.Println(result, name)
}

변수 반복문
---

package main

import "fmt"

func plus(a ...int) int {
	var total int            // -> short cut total := 0
	for _, item := range a { // index 가 없을떈 _, 언더 스코어로 무시 가능
		total += item
	}
	return total
}

func main() {
	result := plus(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(result)
}

배열 반복문 
---
package main

import "fmt"

func main() {
	name := "Hi go Lang!!!"
	for _, letter := range name {
		fmt.Println(string(letter))
	}
}

진수 변환
---
package main

import "fmt"

func main() {
	x := 423523523235235235
	xAsBinary := fmt.SPrintf("%b\n", x) // binary to -> String 
    fmt.Println(x,xAsBinary)
}

array 를 생성후 -> iterator 하는 방법
---
package main

import "fmt"

func main() {
    // [] 괄호안의 숫자를 지워주면 slice ( 무한 으로 담을 수 있다.)
	foods := [3]string{"potato", "pizza", "pasta"}
// 1 번 유형
	for _, food := range foods {
		fmt.Println(food)
	}
// 2번 유형 
	for i := 0; i < len(foods); i++ {
		fmt.Println(foods[i])
	}
}

---
package main

import "fmt"

func main() {
	foods := []string{"potato", "pizza", "pasta"}
	fmt.Printf("%v", foods)

    // 1 
	foods = append(foods, "tomato") // slice 에 item 을 추가하는 방법
	fmt.Printf("%v\n", foods)

    // 2 
	foodsNew := append(foods, "tomato") // 복사본에 할당 하는 방법
	fmt.Printf("%v\n", foodsNew)
    
	fmt.Println(len(foodsNew)) // 길이.

}
```


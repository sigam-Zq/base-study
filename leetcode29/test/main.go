package main

import (
	"fmt"
)

/*
label test

(双层循环中才有效)

continue label
break label
*/

func main() {
	// log.Println("--")

label:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			// fmt.Printf("Inner i --- %d  j --- %d \n", i, j)
			// if j == 2 && i == 2 {
			// 上下一致-- 这里双层循环的 continue 等价于 break ??? 不一致其实
			// continue label
			// 	break
			// }
			if j == 2 && i == 2 {
				fmt.Printf("Inner i --- %d  j --- %d \n", i, j)
				break label
			}
		}
		fmt.Printf("Outer i --- %d  \n", i)
	}
}

/*

Outer i --- 0
Outer i --- 1
Inner i --- 2  j --- 2
*/

/*
问:




	outer:
		for _, s := range special {
			for i, need := range curNeeds {
				if s[i] > need {
					// break ?????
					continue outer
				}
				nextNeeds[i] = need - s[i]
			}



上面的循环

		for _, s := range special {
			for i, need := range curNeeds {
				if s[i] > need {
					break
				}
				nextNeeds[i] = need - s[i]
			}

和下面这个循环有什么区别么


答:




这两段代码有很重要的区别：

1. 使用 `continue outer` 的版本：
```go
outer:
    for _, s := range special {
        for i, need := range curNeeds {
            if s[i] > need {
                continue outer  // 直接跳到外层循环的下一次迭代
            }
            nextNeeds[i] = need - s[i]
        }
        // 这里的代码只有在内层循环完全完成时才会执行
    }
```

2. 使用 `break` 的版本：
```go
for _, s := range special {
    for i, need := range curNeeds {
        if s[i] > need {
            break  // 只跳出内层循环
        }
        nextNeeds[i] = need - s[i]
    }
    // 这里的代码会在break后执行
}
```

主要区别：
1. `continue outer` 会：
   - 立即跳出内层循环
   - 跳过外层循环的剩余代码
   - 直接开始外层循环的下一次迭代

2. `break` 只会：
   - 跳出内层循环
   - 继续执行外层循环中的剩余代码
   - 然后才开始外层循环的下一次迭代

举个例子：
```go
// 假设：
special = [[2,1], [1,2]]
curNeeds = [1,1]

// 使用 continue outer:
// 当 s=[2,1] 时，发现 2>1，立即跳到下一个 special
// 完全跳过当前 special 的所有处理

// 使用 break:
// 当 s=[2,1] 时，发现 2>1，跳出内层循环
// 但可能还会执行这个 special 的其他处理逻辑
```

在你的场景中：
- 如果发现任何一个 `s[i]` 超过了 `need`，这个 `special` 就完全不可用，应该使用 `continue outer`
- 如果只是要停止检查当前 `special` 的剩余元素，但这个 `special` 仍然可用，才使用 `break`

看起来你是在检查 `special` 是否可用，所以 `continue outer` 是更合适的选择。

*/

package quickMul

import (
	"log"
	"math"
	"testing"
	"time"
)

func TestPriority(t *testing.T) {
	t.Log(2 * 3 % 2)
}

func TestQuickMul(t *testing.T) {
	x, y, m := int64(2e12+33123), int64(3e12+2334), int64(1e9+7)

	var res, res2 int64
	timeCost(func() {
		res = quickMul(x, y, m)
	})
	timeCost(func() {
		// res2 = (x * y) % m
		res2 = int64(math.Pow(float64(x), float64(y))) % m
	})
	t.Log(res, res2, res == res2)

}
func timeCost(f func()) {
	s := time.Now()

	for i := 0; i < 10000; i++ {
		f()
	}

	tc := time.Since(s)

	log.Printf("time COST %v\n", tc)
}

// 这里的是 x^y mod m
func quickMul(x, y, m int64) int64 {
	res := int64(1)
	for y > 0 {
		if (y & 1) == 1 {
			res = (res * x) % m
		}
		y >>= 1
		x = (x * x) % m
	}
	return res
}

/*
是的，`quickMul` 函数在某些情况下会比直接计算 `res = (x * y) % m` 更高效，具体原因取决于乘法运算的规模和算法的应用场景。

### **`quickMul` 的意义**
`quickMul` 实现的是**快速幂算法**，主要用于计算 \(x^y \mod m\)，而不是单纯的 \(x \times y \mod m\)。快速幂的效率高是因为它通过**指数的二进制分解**来减少运算次数。

例如，普通计算 \(x^y\) 的时间复杂度是 \(O(y)\)，而快速幂算法的时间复杂度是 \(O(\log y)\)。

---

### **对比适用场景**
1. **直接计算 `res = (x * y) % m`**
   - 适合用于**单次乘法**的情况。
   - 时间复杂度是 \(O(1)\)。
   - 只需要一次常规的乘法操作。

2. **`quickMul`**
   - 适合用于计算 \(x^y \mod m\)（幂运算）时使用。
   - 时间复杂度是 \(O(\log y)\)，远低于直接幂运算的 \(O(y)\)。
   - 在需要反复使用幂运算的场景（比如密码学、数论问题、组合数学）中，`quickMul` 的性能优势非常明显。

---

### **总结**
- 如果你的需求是**简单的两个数相乘并取模**，直接使用 `res = (x * y) % m` 即可，效率最高。
- 如果你的需求是**幂运算（比如 \(x^y \mod m\)**），`quickMul` 明显更快，因为它将指数运算降为对数复杂度。

所以，选择哪种方法取决于具体的应用需求。如果你只是为了两数相乘，`quickMul` 的复杂性反而会拖慢性能。
*/

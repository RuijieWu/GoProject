// Fuck MD5 ro SHA
package main

import (
	"crypto/md5"
	"fmt"
	"math"
	"strings"
	"sync"
	"time"
)

/*
	func sha(str string, wg *sync.WaitGroup) string {
		h := sha256.New()
		h.Write([]byte(str))
		a := h.Sum(nil)
		hashCode2 := hex.EncodeToString(a[:]) //将数组转换成切片，转换成16进制，返回字符串
		return hashCode2
	}
*/
func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func f(s string, cnt int) {
	cnt--

	for i := 33; i <= 126; i++ {
		temp_s := strings.Replace(s, "?", string(i), 1)
		if cnt == 0 {
			count++
			str := MD5(temp_s)
			if str == target_string {
				flag = 1
				ans <- temp_s
				wg.Done()
			}
		} else {
			go f(temp_s, cnt)
		}
	}

}

func main() {
	startTime := time.Now()
	wg = sync.WaitGroup{}

	for i := 0; i < len(source_string); i++ {
		if source_string[i] == '?' {
			num++
		}
	}
	wg.Add(1)
	sum = int(math.Pow((126 - 33 + 1), float64(num)))

	go f(source_string, num)

	for flag != 1 {
		fmt.Println(time.Now(), " ", count, "/", sum, ":", flag)
		time.Sleep(1 * time.Second)
	}
	elapsedTime := time.Since(startTime)
	fmt.Println(time.Now(), " ", count, "/", sum, " ", "Cracked!")
	fmt.Println("Time:", elapsedTime)
	fmt.Println("The Answer is :", <-ans)
}

package util

import (
	"math/rand"
	"time"
	"unsafe"
)

//https://www.flysnow.org/2019/09/30/how-to-generate-a-random-string-of-a-fixed-length-in-go.html
//仅仅从rune到byte的改进，我们就获得了** 24%**的提升，内存占用降低了**三分之一** 。
//使用rand.Int63替换掉原来的rand.Intn，我们又获得了近**20%**的提升。
//单纯的使用掩码，因为重复获取可用索引的问题，性能下降了** -22%**。
//但是当我们对 Masking 掩码 进行改进，分为10部分缓存的时候，我们获得了3倍的提升。
//使用rand.Source 代替 rand.Rand, 我们再次获得了**21%**的提升。
//使用strings.Builder,速度提升虽然只有3.5%,但是内存分配降低了50% 。
//最后，通过unsafe包精简重写了strings.Builder的功能，我们又获得了**14%**的提升。
//最终，RandStringBytesMaskImprSrcUnsafe比RandStringRunes快6.3倍，并且只使用了六分之一的内存和一半的内存分配，我们就完成了任务

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

type Source interface {
	Int63() int64
	Seed(seed int64)
}

var src = rand.NewSource(time.Now().UnixNano())

func RandomString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	// 不同类型转换语法：<结果类型> := <目标类型> ( <表达式> )
	return *(*string)(unsafe.Pointer(&b))
}

# 一. GC

* GC英文全称 `garbage collector`
* Go语言GC是相对C/C++语言非常重要的改进
* 一些常用GC算法
  * 引用计算法.当对象被引用时计算器加一.不被引用计数器减一  与java一样，无法解决循环引用的问题
    * PHP和Object-C使用
    * 计数增加消耗
  * Mark And Sweep 标记和清除算法.停止程序运行,递归遍历对象,进行标记.标记完成后将所有没有引用的对象进行清除
    * 由于标记需要停止程序(Stop the world),当对象特别多时,标记和清除过程比较耗时(可能几百毫秒),很难接受
  * 三色标记法:是Mark And Sweep的改进版.从逻辑上分为白色区(未搜索),灰色区(正搜索),黑色区(已搜索).灰色区内容是子引用没有进行搜索,黑色区表示子引用存在
    - 从root根出发扫描所有根对象（类似于JVM的GCRoot，解决了循环引用的问题），将他们引用的对象标记为灰色（图中A，B）
  * 分代收集.一般情况都有三代,例如java中新生代,老年代,永久代.当新生代中带有阈值时会把对象放入到老年代,相同道理老年代内容达到阈值会放入到永久代

# 二.Go语言中的GC

* Go语言中采用`Stop The World`方式
* Golang每个版本基本上都会对GC进行优化,从Golang1.5开始支持并发(concurrent )收集,从1.8版本已经把STW时间优化到了100微妙,通常只需要10微妙以下.且在1.10版本时再次优化减少GC对CPU占用
* Go语言中GC是自动运行的,在下列情况下会触发GC
  * 超过内存大小阈值(当用户程序申请分配 32KB 以上的大对象时，一定会构建` runtime.gcTrigger `结构体尝试触发GC)
  * 达到定时时间,阈值是由一个`gcpercent`的变量控制的,当新分配的内存占已在使用中的内存的比例超过gcprecent时就会触发。比如一次回收完毕后，内存的使用量为5M，那么下次回收的时机则是内存分配达到10M的时候。如果达不到内存大小的阈值GC就会被定时时间触发,默认2min触发一次.

 **runtime.gcTrigger.test 方法决定是否需要触发垃圾收集**
```go
func (t gcTrigger) test() bool {
	if !memstats.enablegc || panicking != 0 || gcphase != _GCoff {
		return false
	}
	switch t.kind {
	case gcTriggerHeap:
		return memstats.heap_live >= memstats.gc_trigger
	case gcTriggerTime:
		if gcpercent < 0 {
			return false
		}
		lastgc := int64(atomic.Load64(&memstats.last_gc_nanotime))
		return lastgc != 0 && t.now-lastgc > forcegcperiod
	case gcTriggerCycle:
		return int32(t.n-work.cycles) > 0
	}
	return true
}
```
1.`gcTriggerHeap `— 堆内存的分配达到达`控制器计算`的触发堆大小；
2.`gcTriggerTime` — 如果一定时间内没有触发，就会触发新的循环，该条件由 runtime.forcegcperiod 变量控制，默认为` 2` 分钟；
3.`gcTriggerCycle` — 如果当前没有开启垃圾收集，则触发新的循环；




* GC调优
  * 小对象复用,局部变量尽量少声明,多个小对象可以放入到结构体,方便GC扫描
  * 少用string的”+”
* 在runtime包下mgc.go中明确的说明了Golang的GC的解释


package main

// HOI4RNG 模拟 HOI4 的随机数生成器 (线性同余生成器)
type HOI4RNG struct {
	state uint32 // 当前状态（种子）
}

// 模数、乘数、增量参数常量定义
const (
	modulus    uint64 = 2147483647 // m = 2^31 - 1
	multiplier uint64 = 16807      // a = 16807 (7^5)
	increment  uint64 = 0          // c = 0
)

// NewHOI4RNG 初始化一个 HOI4 RNG 实例
func NewHOI4RNG(seed uint32) *HOI4RNG {
	// 避免种子为0的情况，如果为0则用默认值1
	if seed == 0 {
		seed = 1
	}
	return &HOI4RNG{state: seed}
}

// Next 生成下一个随机数 (返回 uint32 类型)
func (r *HOI4RNG) Next() uint32 {
	// 计算下一状态，使用64位以防止乘法溢出，再取模
	nextState := (uint64(r.state)*multiplier + increment) % modulus
	r.state = uint32(nextState) // 更新内部状态
	return r.state
}

// 为了方便，也实现一个方法返回 [0,1) 之间的浮点随机数
func (r *HOI4RNG) NextFloat() float64 {
	// 将当前状态（1到m-1）映射为 0到1 之间的小数
	return float64(r.Next()) / float64(modulus)
}

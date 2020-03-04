package option

// Options 选项结构体
type Options struct {
	strOption1 string
	strOption2 string
	intOption1 int
	intOption2 int
}

// Option 定义一个函数类型的变量
type Option func(opts *Options)

// NewOptions 实例化选项结构体对象
func NewOptions(opts ...Option) (options *Options) {
	options = &Options{}
	// 遍历opts,得到每一个函数
	for _, opt := range opts {
		// 调用函数，在函数中给对象赋值
		opt(options)
	}
	return
}

// SetStrOption1 给strOption1属性设置值
func SetStrOption1(str string) Option {
	return func(opts *Options) {
		opts.strOption1 = str
	}
}

// SetStrOption2 给strOption2属性设置值
func SetStrOption2(str string) Option {
	return func(opts *Options) {
		opts.strOption2 = str
	}
}

// SetIntOption1 给intOption1属性设置值
func SetIntOption1(num int) Option {
	return func(opts *Options) {
		opts.intOption1 = num
	}
}

// SetIntOption2 给intOption2属性设置值
func SetIntOption2(num int) Option {
	return func(opts *Options) {
		opts.intOption2 = num
	}
}

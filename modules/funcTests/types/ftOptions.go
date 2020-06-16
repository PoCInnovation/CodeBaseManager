package ft_types

type ftOptions struct {
	Repeat     int `toml:"repeat"`
	Time       bool `toml:"time"`
	ShouldFail bool `toml:"shouldFail"`
	Timeout    float64 `toml:"timeout"`
	// TODO: BuildWith  string
}

func (opt *ftOptions) ApplyDefault(common ftOptions) {
	if opt.Repeat == 0 {
		opt.Repeat = common.Repeat
	}
	if opt.Time == false {
		opt.Time = common.Time
	}
	if opt.Timeout == 0 {
		opt.Timeout = common.Timeout
	}
}
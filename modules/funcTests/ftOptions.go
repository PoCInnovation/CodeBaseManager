package funcTests

type ftOptions struct {
    Repeat     int
    Time       bool
    ShouldFail bool
    Timeout    int
    // TODO: BuildWith  string
}

func (opt *ftOptions) SetCommon(common *ftOptions) {
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

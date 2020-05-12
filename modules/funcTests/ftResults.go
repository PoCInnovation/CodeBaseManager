package funcTests

const (
    OK = iota
    NOK
    IGNORED

)

type ftResult struct {
    statusFail int
    stdoutFail int
    stderrFail int
}

func compareOutput(out, exp string) int {
    switch exp {
    case "":
        return IGNORED
    case out:
        return OK
    default:
        return NOK
    }
}

func (r *ftResult) CompareToRef(ref, my *ftExecution) {
    switch {
    case ref.status != my.status:
        r.statusFail = NOK
    case ref.outBuf.String() != my.outBuf.String():
        r.stdoutFail = NOK
    case ref.errBuf.String() != my.errBuf.String():
        r.stderrFail = NOK
    }
}

func (r *ftResult) CompareToExp(exp *ftExpected, my *ftExecution) {
    if exp.Status != my.status {
        r.statusFail = NOK
    }
    r.stdoutFail = compareOutput(exp.getExpOut(), my.outBuf.String())
    r.stderrFail = compareOutput(exp.getExpErr(), my.errBuf.String())
}
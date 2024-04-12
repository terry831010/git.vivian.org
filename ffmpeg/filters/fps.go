package filters

//调整帧率
type fps struct {
	filter
	fps        string `args:"fps"`        //A string describing desired output framerate (default "25")
	start_time string `args:"start_time"` //double  Assume the first PTS should be this value.
	round      string `args:"round"`      //set rounding method for timestamps (from 0 to 5) (default near)
	eof_action string `args:"eof_action"` //action performed for last frame (from 0 to 1) (default round)
}

func Fps() *fps {
	c := &fps{}
	c.meta = c
	return c
}

func (me *fps) Fps(fps string) *fps {
	me.fps = fps
	return me
}

func (me *fps) StartTime(s string) *fps {
	me.start_time = s
	return me
}

func (me *fps) Round(round string) *fps {
	me.round = round
	return me
}

func (me *fps) EofAction(eof_action string) *fps {
	me.eof_action = eof_action
	return me
}

func (me *fps) Name() string {
	return "fps"
}

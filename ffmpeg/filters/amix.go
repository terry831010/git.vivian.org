package filters

//混合音频
type amix struct {
	filter
	inputs             string `args:"inputs"`             //Number of inputs from 1 to 32767  defalut 2
	duration           string `args:"duration"`           //longest:0  shortest:1  first:2
	dropout_transition string `args:"dropout_transition"` //Transition time
	weights            string `args:"weights"`            //Set weight for each input. (default "1 1")
	normalize          string `args:"normalize"`          //Scale inputs (default true)
}

func AMix() *amix {
	c := &amix{}
	c.meta = c
	return c
}

func (me *amix) Inputs(inputs string) *amix {
	me.inputs = inputs
	return me
}

func (me *amix) Duration(a string) *amix {
	me.duration = a
	return me
}

func (me *amix) DropoutTransition(d string) *amix {
	me.dropout_transition = d
	return me
}

func (me *amix) Weights(w string) *amix {
	me.weights = w
	return me
}

func (me *amix) Normalize(n string) *amix {
	me.normalize = n
	return me
}
func (me *amix) Name() string {
	return "amix"
}

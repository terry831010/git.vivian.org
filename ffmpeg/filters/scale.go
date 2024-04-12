package filters

//缩放
type scale struct {
	filter
	w                           string `args:"w"` //宽
	h                           string `args:"h"` //高
	flags                       string `args:"flags"`
	interl                      string `args:"interl"`
	in_color_matrix             string `args:"in_color_matrix"`
	out_color_matrix            string `args:"out_color_matrix"`
	out_range                   string `args:"out_range"`
	in_range                    string `args:"in_range"`
	force_original_aspect_ratio string `args:"force_original_aspect_ratio"`
}

func Scale() *scale {
	c := &scale{}
	c.meta = c
	return c
}

func (me *scale) Width(w string) *scale {
	me.w = w
	return me
}

func (me *scale) Height(h string) *scale {
	me.h = h
	return me
}

func (me *scale) ForceRatio(fr string) *scale {
	me.force_original_aspect_ratio = fr
	return me
}

func (me *scale) Flags(flags string) *scale {
	me.flags = flags
	return me
}

func (me *scale) Interl(interl string) *scale {
	me.interl = interl
	return me
}

func (me *scale) InColorMatrix(in_color_matrix string) *scale {
	me.in_color_matrix = in_color_matrix
	return me
}

func (me *scale) OutColorMatrix(out_color_matrix string) *scale {
	me.out_color_matrix = out_color_matrix
	return me
}

func (me *scale) OutRange(out_range string) *scale {
	me.out_range = out_range
	return me
}

func (me *scale) InRange(in_range string) *scale {
	me.in_range = in_range
	return me
}

func (me *scale) Name() string {
	return "scale"
}

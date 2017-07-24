package hocon

import (
	"github.com/jdevelop/go-hocon/parser"
)

func (r *hocon) ExitObject_data(ctx *parser.Object_dataContext) {
	current, _ := r.stack.Pop()
	parent, _ := r.stack.Peek()
	parent.setObject(ctx.Key().GetText(), current.(*ConfigObject))
}

func (r *hocon) EnterObject_data(ctx *parser.Object_dataContext) {
	r.stack.Push(NewConfigObject())
}

func (r *hocon) ExitString_data(ctx *parser.String_dataContext) {
	sd := ctx.String_value().GetText()
	if sd[0] == '"' || sd[0] == '\'' {
		sd = sd[1 : len(sd)-1]
	}
	if v, err := r.stack.Peek(); err == nil {
		v.setString(ctx.Key().GetText(), sd)
	}
}

func (r *hocon) ExitNumber_data(ctx *parser.Number_dataContext) {
	if v, err := r.stack.Peek(); err == nil {
		v.setInt(ctx.Key().GetText(), ctx.NUMBER().GetText())
	}
}

func (r *hocon) ExitReference_data(ctx *parser.Reference_dataContext) {
	if v, err := r.stack.Peek(); err == nil {
		v.setInt(ctx.Key().GetText(), ctx.REFERENCE().GetText())
	}
}

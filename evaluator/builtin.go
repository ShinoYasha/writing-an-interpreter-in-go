package evaluator

import "interpreter/object"

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments.got=%d, want=1",
					len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Element))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		}},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments.got=%d,want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s",
					args[0].Type())
			}
			arg := args[0].(*object.Array)
			if len(arg.Element) > 0 {
				return arg.Element[0]
			}
			return NULL
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments.got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("arguments to `last` must be ARRAY, got %s",
					args[0].Type())
			}
			arg := args[0].(*object.Array)
			length := len(arg.Element)
			if length > 0 {
				return arg.Element[length-1]
			}
			return NULL
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments.got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("arguments to `last` must be ARRAY, got %s",
					args[0].Type())
			}
			arg := args[0].(*object.Array)
			length := len(arg.Element)
			if length > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, arg.Element[1:length])
				return &object.Array{Element: newElements}
			}
			return NULL
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments.got=%d, want=2", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("arguments to `last` must be ARRAY, got %s",
					args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Element)
			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, arr.Element)
			newElements[length] = args[1]
			return &object.Array{
				Element: newElements,
			}
		},
	},
}

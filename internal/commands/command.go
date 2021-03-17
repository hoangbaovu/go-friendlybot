package commands

// ;;say ;;s ;;

type Command interface {
	Invokes() []string
	Description() string
	AdminRequired() bool
	Exec(ctx *Context) error
}

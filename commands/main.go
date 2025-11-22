package commands

type Client interface {
	FetchData(path string, dest any) error
}

type Command struct {
	client Client
}

func NewCommands(client Client) *Command {
	return &Command{
		client: client,
	}
}

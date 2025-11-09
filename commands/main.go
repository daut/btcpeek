package commands

type Command struct {
	FetchData FetchDataFunc
}

type FetchDataFunc func(endpoint string, dest any) error

func NewCommands(fetchData FetchDataFunc) *Command {
	return &Command{FetchData: fetchData}
}

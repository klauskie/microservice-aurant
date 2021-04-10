package models

type Client struct {
	Id, Name string
}

func NewClient(id, name string) Client {
	return Client{
		Id: id,
		Name: name,
	}
}

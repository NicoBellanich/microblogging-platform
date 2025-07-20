package repository

type IMessageRepository interface {
	Save()
	Load()
	LoadAll()
}

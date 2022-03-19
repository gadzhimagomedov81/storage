package file

import "github.com/google/uuid"

type File struct {
	ID   uuid.UUID
	Name string
	Data []byte
}

func NewFile(name string, blob []byte) (*File, error) {
	uuid, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	return &File{
		ID:   uuid,
		Name: name,
		Data: blob,
	}, nil
}

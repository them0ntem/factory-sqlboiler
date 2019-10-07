package factory

import "github.com/gofrs/uuid"

func UUID() string {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	return id.String()
}

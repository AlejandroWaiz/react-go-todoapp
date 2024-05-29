package database

type Database struct {
}

type DatabaseImplementation interface {
}

func New(database string) (DatabaseImplementation, error) {

	//TODO: implementar el devolver el tipo de base segun el string que te pasen al invocar la funcion
	switch database {

	case "firestore":
		return Database{}, nil

	case "local":
		return Database{}, nil

	default:
		return Database{}, nil

	}

}

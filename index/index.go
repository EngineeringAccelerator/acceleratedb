package index


type ADBIndex interface {
	//DB/Idx init
	Init(name string) error

	//Release file lock. Close session
	Close() error

	//Insert into file and also create a new index entry.
	Insert(key string, value string) error

	//Seek and update the value.
	Update(key string, value string) error

	//Get the value for the given key. 
	Fetch(key string)(string, error)

	//Fetch the whole index.
	FetchWhole() (map[string]string, error)
}



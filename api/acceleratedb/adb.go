package acceleratedb


import (
	"fmt"
	"os"
	idx "github.com/morpheyesh/AccelerateDB/index"
	
)

/**
Main class that handles the operations performed on the DB
**/
type AccelerateDB struct {
	name 	string
	index   idx.ADBIndex
}

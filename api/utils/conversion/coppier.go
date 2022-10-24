package conversion

import (
	"github.com/jinzhu/copier"
)

func Coppier(toVal, fromVal interface{}) {
	err := copier.Copy(toVal, fromVal)
	if err != nil {
		panic(err)
	}

}

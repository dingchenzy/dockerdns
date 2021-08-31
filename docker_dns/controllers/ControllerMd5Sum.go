package controllers

import (
	"crypto/md5"
	"fmt"
)

func Md5Run() bool {
	newstr := fmt.Sprintf("%#v", AllContainerIPName)
	oldstr := fmt.Sprintf("%#v", OldContainerIPName)

	return md5.Sum([]byte(newstr)) == md5.Sum([]byte(oldstr))
}

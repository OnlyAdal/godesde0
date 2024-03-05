package users

import (
	"fmt"
	"github/OnlyAdal/godesde0/modelos"
	"time"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUSer(10, "Pablo", time.Now(), true)
	fmt.Println(u)

}

package user_service_test

import (
	"testing"
	"time"

	model "../../models"
	userService "../user.service"
)

func TestCreate(t *testing.T) {
	user := model.User{
		Name:     "Juan",
		Email:    "Jerez@gmail.com",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	err := userService.Create(user)
	if err != nil {
		t.Error("la prueba de persistencia de datos del usuario ha fallado")
		t.Fail()
	} else {
		t.Log("test passed successfully")
	}

}

func TestRead(t *testing.T) {
	users, err := userService.Read()

	if err != nil {
		t.Error("error in the query")
		t.Fail()
	}

	if len(users) == 0 {
		t.Error("response whitout data")
		t.Fail()
	}
	t.Log("test success")

}
func TestUpdate(t *testing.T) {
	user := model.User{
		Name:  "Angela",
		Email: "angelical@gmail.com",
	}
	err := userService.Update(user, "602fe5464f6aa05220e55aed")

	if err != nil {
		t.Error("Error al trata de actualizar el usuario")
		t.Fail()
	}
	t.Log("La prueba de actualizacion funciono correctamente")

}
func TestDelete(t *testing.T) {
	err := userService.Delete("602f20f2d988877038dcde0e")
	if err != nil {
		t.Error("Error al tratar de eliminar el usuario")
		t.Fail()
	}
	t.Log("La prueba de eliminacion finalizo con exito")

}

package rest

import (
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("about to start test cases...")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	repository := usersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")

	fmt.Println(user)
	fmt.Println(err)

}

func TestLoginUserInvalidErrorInterface(t *testing.T) {

}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {

}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {

}

func TestLoginUserNoError(t *testing.T) {

}

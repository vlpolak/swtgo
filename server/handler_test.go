package server

import (
	"bytes"
	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestLoginHandlerFunct(t *testing.T) {
	defer gock.Off()
	gock.New("http://127.0.0.1:8080").
		Post("/login/").
		Persist().
		Reply(200).
		JSON(map[string]string{"user": "Jinzhu", "password": "122234234"})

	body := bytes.NewBuffer([]byte(`{"user": "Jinzhu","password": "122234234"}`))
	res, err := http.Post("http://127.0.0.1:8080/login/", "application/json", body)
	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode, 200)

	res_body, _ := ioutil.ReadAll(res.Body)
	st.Expect(t, string(res_body), `{"password":"122234234","user":"Jinzhu"}`)
	st.Expect(t, gock.IsDone(), true)
}

func TestHandleHandleHome(t *testing.T) {
	res, err := http.Get("http://127.0.0.1:8080/")
	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode, 200)
}

func TestHandleRegisterUser(t *testing.T) {
	defer gock.Off()
	gock.New("http://127.0.0.1:8080").
		Post("/register/").
		Persist().
		Reply(200).
		JSON(map[string]string{"name": "Testweeterere", "password": "$2a$14$2vsA3RSLj6icrDfGsprrQuKVf36sRE3GPgkwAcUpy/KshXu9ssuvS"})

	body := bytes.NewBuffer([]byte(`{"user":"Testweeterere","password":"$2a$14$2vsA3RSLj6icrDfGsprrQuKVf36sRE3GPgkwAcUpy/KshXu9ssuvS"}`))
	res, err := http.Post("http://127.0.0.1:8080/register/", "application/json", body)
	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode, 200)

	res_body, _ := ioutil.ReadAll(res.Body)
	st.Expect(t, string(res_body), string(`{"user":"Testweeterere","password":"$2a$14$2vsA3RSLj6icrDfGsprrQuKVf36sRE3GPgkwAcUpy/KshXu9ssuvS"}`))
	st.Expect(t, gock.IsDone(), true)
}

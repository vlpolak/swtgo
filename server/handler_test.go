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
	gock.New("localhost:8080").
		Post("/login").
		Reply(200).
		JSON(map[string]string{"user": "Jinzhu", "password": "122234234"})

	body := bytes.NewBuffer([]byte(`{"user": "Jinzhu","password": "122234234"}`))
	res, err := http.Post("http://127.0.0.1:8080/login/", "application/json", body)
	st.Expect(t, err, nil)
	st.Expect(t, res.StatusCode, 200)

	res_body, _ := ioutil.ReadAll(res.Body)
	st.Expect(t, string(res_body), "hello Jinzhu")
	st.Expect(t, gock.IsDone(), true)
}

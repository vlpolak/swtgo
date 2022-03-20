package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/vlpolak/swtgo/domain/entity"
	"github.com/vlpolak/swtgo/logger"
	"github.com/vlpolak/swtgo/pkg/hasher"
	"html/template"
	"image/png"
	"log"
	"net/http"
)

const webSocketUrl = "ws://fancy-chat.io/ws&token=one-time-token"

//TOTP для одного пользователя
var key *otp.Key

func (s *Server) HandleRegisterUser() func(writer http.ResponseWriter, request *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {

		writer.Header().Set("Content-Type", "application/json")

		r, err := CreateCreateUserRequest(*request)
		if err != nil {
			logger.ErrorLogger("Couldn't create user", err).Log()
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		err = validateCreateUserRequest(r)
		if err != nil {
			logger.ErrorLogger("Couldn't create user", err).Log()
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		userEntity, err := createUser(r.UserName, r.Password)
		if err != nil {
			logger.ErrorLogger("Couldn't create user", err).Log()
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		(*s.Users).us.SaveUser(&userEntity)

		logger.CommonLogger("User %s was created", &userEntity).Log()

		response := &CreateUserResponse{Id: entity.User{}.Uuid.String(), UserName: entity.User{}.UserName}
		err = json.NewEncoder(writer).Encode(response)

		if err != nil {
			logger.ErrorLogger("Couldn't convert user to json", err).Log()
			http.Error(writer, "Registration failed", http.StatusInternalServerError)
			return
		}
	}
}

//func (s *Server) HandleLogin() func(writer http.ResponseWriter, request *http.Request) {
//
//	return func(writer http.ResponseWriter, request *http.Request) {
//
//		writer.Header().Set("Content-Type", "application/json")
//
//		r, err := createLoginUserRequest(*request)
//		if err != nil {
//			logger.ErrorLogger("Couldn't login", err).Log()
//			http.Error(writer, err.Error(), http.StatusBadRequest)
//			return
//		}
//
//		err = validateLoginUserRequest(r)
//		if err != nil {
//			logger.ErrorLogger("Couldn't login", err).Log()
//			http.Error(writer, err.Error(), http.StatusBadRequest)
//			return
//		}
//
//		foundUser, _ := (*s.Users).us.FindUser(r.UserName)
//
//		(*s.Users).lc.SaveActiveUser(foundUser)
//		logger.CommonLogger("User was found", foundUser).Log()
//
//		hasher.CheckPasswordHash(r.Password, foundUser.HashedPassword)
//
//		response := &LoginUserResponse{"https://www.google.com/"}
//		err = json.NewEncoder(writer).Encode(response)
//
//		if err != nil {
//			logger.ErrorLogger("Couldn't convert user to json", err).Log()
//			http.Error(writer, "Registration failed", http.StatusInternalServerError)
//			return
//		}
//	}
//}

func (s *Server) HandleGetAvtiveUsers() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		au := (*s.Users).lc.Get()
		logger.CommonLogger("Active users", au).Log()
	}
}

func (s *Server) HandleHome(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, nil)
}

func (s *Server) LoginHandlerFunc(w http.ResponseWriter, r *http.Request) {
	//Обрабатываем только POST-запрос
	if r.Method != "POST" {
		http.NotFound(w, r)
	}

	var err error
	err = r.ParseForm()
	if err != nil {
		panic(err)
	}

	user := r.FormValue("user")
	password := r.FormValue("password")

	err = validateRequest(r)
	if err != nil {
		return
	}

	foundUser, err := (*s.Users).us.FindUser(user)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	logger.CommonLogger("User was found", foundUser).Log()
	checked := hasher.CheckPasswordHash(password, foundUser.HashedPassword)
	if !checked {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}
	(*s.Users).lc.Set(user, foundUser)
	w.Write([]byte("hello " + foundUser.UserName))
}

func validateRequest(r *http.Request) error {
	if len(r.FormValue("user")) < 4 {
		return errors.New("request was incorrect, user name should contain at least 4 characters")
	}
	if len(r.FormValue("password")) < 8 {
		return errors.New("request was incorrect, user password should contain at least 8 characters")
	}
	return nil
}

//Отображает страницу с QR-кодом
func (s *Server) Setup2FAHandlerFunc(w http.ResponseWriter, r *http.Request) {
	faTemplate.Execute(w, nil)
}

//Генерирует QR-код для добавления аккаунта в Яндекс.Ключ/Google.Authentificator
func (s *Server) GenQRCodeHandlerFunc(w http.ResponseWriter, r *http.Request) {
	//Настраиваем TOTP
	//для каждого пользователя TOTP ключ должен быть уникальным
	//В нашей программе ключ будет разный с каждым запуском (!)
	var err error
	key, err = totp.Generate(totp.GenerateOpts{
		Issuer:      "Example.com",
		AccountName: "zaz600@example.com",
	})
	if err != nil {
		panic(err)
	}
	// Convert TOTP key into a PNG
	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	//для простоты не обрабатываем ошибки
	png.Encode(&buf, img)
	w.Header().Set("Content-Type", "image/png")
	w.Write(buf.Bytes())
}

func (s *Server) Verifi2faHandlerFunc(w http.ResponseWriter, r *http.Request) {
	//Обрабатываем только POST-запрос
	if r.Method != "POST" {
		http.NotFound(w, r)
	}
	//для простоты не обрабатываем ошибки
	r.ParseForm()
	passcode := r.FormValue("passcode")
	valid := totp.Validate(passcode, key.Secret())
	if !valid {
		http.Error(w, "Неверный код", http.StatusForbidden)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`2ФА успешно настроена.`))
	//далее нам надо сохранить в базе key.Secret() пользователя
	//чтобы позднее верифицировать его одноразовые коды по этому секрету
}

func CreateCreateUserRequest(request http.Request) (CreateUserRequest, error) {
	var r CreateUserRequest
	err := json.NewDecoder(request.Body).Decode(&r)

	if err != nil {
		log.Printf("Couldn't convert json to object %v. Request was %v", err, request)
		return CreateUserRequest{}, errors.New("invalid json")
	}
	return r, nil
}

func createLoginUserRequest(request http.Request) (LoginUserRequest, error) {
	var r LoginUserRequest
	err := json.NewDecoder(request.Body).Decode(&r)

	if err != nil {
		log.Printf("Couldn't convert json to object %v. Request was %v", err, request)
		return LoginUserRequest{}, errors.New("invalid json")
	}
	return r, nil
}

func validateCreateUserRequest(r CreateUserRequest) error {
	if len(r.UserName) < 4 {
		return errors.New("request was incorrect, user name should contain at least 4 characters")
	}
	if len(r.Password) < 8 {
		return errors.New("request was incorrect, user password should contain at least 8 characters")
	}
	return nil
}

func validateLoginUserRequest(r LoginUserRequest) error {
	if len(r.UserName) < 4 {
		return errors.New("request was incorrect, user name should contain at least 4 characters")
	}
	if len(r.Password) < 8 {
		return errors.New("request was incorrect, user password should contain at least 8 characters")
	}
	return nil
}

func createUser(userName, password string) (entity.User, error) {
	hp, err := hasher.HashPassword(password)
	if err != nil {
		log.Printf("Couldn't create user")
		return entity.User{}, errors.New("couldn't create user")
	}
	return entity.User{Uuid: uuid.New(), UserName: userName, HashedPassword: hp}, nil
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
    <head>
        <title>Red Tea Pot 250ml</title>
    </head>
    <body>
    <h1>Please Sign in</h1>
    <form action="/login/" method="POST">
        <div>
            <label for="user">user:</label><br>
            <input type="input" name="user" id="user" value="zaz600">
        </div>
        <div>
            <label for="password">password:</label><br>
            <input type="password" name="password" id="password" value="">
        </div>
        <div><input type="submit" value="Sign in"></div>
    </form>
	<div>
	  <a href="/2fa/">Enable 2FA</a>
	</div>
    </body>
</html>
`))

var faTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
    <head>
        <title>Red Tea Pot 250ml</title>
    </head>
    <body>
	<h1>Настройка 2ФА</h1>
	<p>Отсканируйте код в приложении Яндекс.Ключ</p>
	<div>
	 <img src="/qr.png">
	</div>
	<p>Введите код из Яндекс.Ключ</p>
	<form action="/verify2fa/" method="POST">
		<div>
		 <input type="input" name="passcode" id="passcode">
		</div>
		<div><input type="submit" value="Verify"></div>
	</form>
	<div>
	 	<a href="/">На главную</a>
	</div>
    </body>
</html>
`))

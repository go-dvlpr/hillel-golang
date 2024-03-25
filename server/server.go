package server

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
)

func Start() error {
	http.HandleFunc("/", Handle) //localhost:3333

	logrus.Info("prepared to listen :3333 port")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		logrus.Error("failed to start listening :3333 port", err)
		return err
	}

	return nil
}

func Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		HandleGet(w, r)
		return
	}

	if r.Method == http.MethodPost {
		HandlePost(w, r)
		return
	}

	if r.Method == http.MethodPut {
		HandlePut(w, r)
		return
	}

	fmt.Println("method is not get ot post")
}

type HandleGetResponse struct {
	Key             string  `json:"Key"`
	Id              int     `json:"Id"`
	Something       *string `json:"Something,omitempty"`
	ShouldNotBeSend bool    `json:"-"`
}

func HandleGet(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	key := query.Get("key")
	idParam := query.Get("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		logrus.Error("failed to convert id")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("id should be int"))
		return
	}

	resp := HandleGetResponse{
		Key:             key,
		Id:              id,
		ShouldNotBeSend: false,
	}

	marshaledResponse, err := json.Marshal(resp)
	if err != nil {
		logrus.Error("failed to marshal response")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("id should be int"))
		return
	}

	fmt.Println("this is GET hendler")
	w.Write(marshaledResponse)
}

type HandlePostReq struct {
	Name  string `json:"name"`
	Age   int
	Email string
}

func HandlePut(w http.ResponseWriter, r *http.Request) {
	var req HandlePostReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logrus.Error("failed to unmarshal body")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("failed to unmarshal body"))
		return
	}

	fmt.Println(req)
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Error("failed to read body")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("failed to read body"))
		return
	}

	req := HandlePostReq{}

	err = json.Unmarshal(reqBytes, &req)
	if err != nil {
		logrus.Error("failed to unmarshal body")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("failed to unmarshal body"))
		return
	}

	fmt.Println(req)

	if req.Age < 18 {
		logrus.Error("ivan carevich too young")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ivan carevich too young"))
		return
	}

	reqBytesSecond, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Error("failed to read body")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("failed to read body"))
		return
	}

	fmt.Println("=====> ", string(reqBytesSecond))

	io.WriteString(w, "this is POST handler 123")
	//w.Write([]byte("this is POST hendler"))
}

//1: [GET, POST, PUT(full), DELETE, PATCH(part)]
//2: URI: http://localhost:3333/users?id=12345
//2.1:  http / https 		- protocol
//2.2:  localhost (0.0.0.0) - host
//2.3:  :3333 				- port
//2.4:  /users?id=12345		- URL
//2.5   ?id=12345 			- query parameters
// 3: POST, PUT, PATCH - Body
// 4: HEADER: Content-Length: 100

// 1xx:
// 2xx:	200 - OK, 201 - Created, 204 - No content
// 3xx: - Redirect
// 4xx: - Client error
// 5xx: - Server error

//http://localhost:3333/users?task_id=12345
//http://localhost:3333/task/12345

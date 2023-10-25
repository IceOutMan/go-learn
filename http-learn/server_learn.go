package main

import "net/http"

type CommonHandler struct{}

func (h *CommonHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello World\n"))
}

type UserHandler struct{}

func (h *UserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("User Handler\n"))
}

type OrderHandler struct{}

func (h *OrderHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Order	 Handler\n"))
}

func simple_server() {
	server := http.Server{
		Addr:    "localhost:8088",
		Handler: &CommonHandler{},
	}
	server.ListenAndServe()
}

func multi_handle_server() {
	server := http.Server{
		Addr: "localhost:8088",
	}

	http.Handle("/user", &UserHandler{})
	http.Handle("/order", &OrderHandler{})

	server.ListenAndServe()
}

func user_info(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user info"))
}

func order_info(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("order info"))
}

func multi_handle_func_server() {
	server := http.Server{
		Addr: "localhost:8088",
	}
	http.HandleFunc("/user", user_info)
	http.HandleFunc("/order", order_info)

	server.ListenAndServe()

}

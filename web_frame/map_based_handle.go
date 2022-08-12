package main

import (
	"net/http"
)

type HandlerBaseOnMap struct {
	handlers map[string]func(ctx *Context)
}

func (h *HandlerBaseOnMap) ServeHTTP(writer http.ResponseWriter,request *http.Request){
	key := h.key(request)
	if handler,ok := h.handlers[key];ok{
		handler(NewContext(writer,request))
	}else{
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("not found"))
	}
}

func (h *HandlerBaseOnMap) key(method string,pattern string) string{
	return  method + "#" + method
}
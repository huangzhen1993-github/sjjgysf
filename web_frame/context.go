package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (c *Context) ReadJson(req interface{}) error {
	body,_err := ioutil.ReadAll(c.R.Body)
	if _err != nil{
		return _err;
	}
	_err = json.Unmarshal(body,req)
	if _err != nil{
		return _err;
	}
	return nil
}

func (c *Context) WriteJson(code int,resp interface{}) error{
	c.W.WriteHeader(code)
	respJson,_err := json.Marshal(resp)
	if _err != nil{
		return _err
	}
	_,_err =  c.W.Write(respJson)
	return _err
}

type signUpReq struct {

}

type commonResponse struct {
	Data int
}

func SignUp(ctx *Context)  {
	req := &signUpReq{}
	_err := ctx.ReadJson(req)
	if _err != nil{
		ctx.BadRequestJson(_err)
		return
	}
	resp := &commonResponse{
		Data:123,
	}
	_err = ctx.WriteJson(http.StatusOK,resp)
	if _err != nil{
		ctx.BadRequestJson(_err)
		return
	}
}

func (c *Context) BadRequestJson(resp interface{}) error{
	return c.WriteJson(http.StatusBadRequest,resp)
}

func NewContext(writer http.ResponseWriter,request *http.Request) *Context  {
	return &Context{
		W:writer,
		R:request,
	}
}

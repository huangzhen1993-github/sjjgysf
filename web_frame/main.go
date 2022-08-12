package main

func main() {
	server := NewHttpServer("test-server")
	server.Route("/",SignUp)
	_err := server.Start(":8888")
	if _err != nil{
		panic(_err.Error())
	}
}

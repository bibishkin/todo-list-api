package main

func main() {

	go func() {
		infoLog.Println("HTTP server listen on", Server.Addr)
		err := Server.ListenAndServe()
		if err != nil {
			errorLog.Fatal(err)
		}
	}()

	infoLog.Println("HTTPS server listen on", ServerTLS.Addr)
	err := ServerTLS.ListenAndServeTLS(certPath, keyPath)
	if err != nil {
		errorLog.Fatal(err)
	}

}

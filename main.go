package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello, world!")
}

/*func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go Monitor(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(10 * time.Second)
}

func Monitor(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("monitor stopped:", ctx.Err())
			return
		default:
			fmt.Println("...")
			time.Sleep(500 * time.Millisecond) // prevent spamming
		}
	}
}*/

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Println("Starting server at http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}

}

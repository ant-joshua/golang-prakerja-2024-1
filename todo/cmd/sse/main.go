package main

import (
	"fmt"
	"log"
	"net/http"
)

type Event struct {
	Data any
}

func sendToChannel(eventChannel chan Event, data any) {
	eventChannel <- Event{Data: data}
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")                                                                                   // Allow any domain
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                                                    // Allow specific methods
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization") // Allow specific headers

		// If this is a preflight request, the method will be OPTIONS,
		// so call the real handler only if method is not OPTIONS
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func main() {
	// channel to hold event

	eventChannel := make(chan Event)

	// go func() {
	// 	for {
	// 		time.Sleep(5 * time.Second)
	// 		eventChannel <- Event{Data: fmt.Sprintf("Notification at %s", time.Now())}
	// 	}
	// }()

	http.HandleFunc("POST /orders/create", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		// create new order
		fmt.Println("Process Order")

		// Insert to database
		fmt.Println("Order created")

		// send event to channel
		go sendToChannel(eventChannel, "New Order Created")

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Order Created"))
	}))

	http.HandleFunc("POST /orders/{id}/pay", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")

		// pay order
		fmt.Println("Pay Order")

		// Update order status
		fmt.Println("Order paid")

		// send event to channel
		go sendToChannel(eventChannel, fmt.Sprintf("Order %s Paid", id))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Order Paid"))
	}))

	http.HandleFunc("/events", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		// set header to allow server sent event
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		closeNotifier := w.(http.CloseNotifier).CloseNotify()

		// write to response writer
		for {
			select {
			case <-closeNotifier:
				fmt.Println("Connection closed")
				return
			case event := <-eventChannel:
				fmt.Fprintf(w, "data: %s\n\n", event.Data)
				if f, ok := w.(http.Flusher); ok {
					f.Flush()
				}
			}
		}
	}))

	log.Println("server started at :5000")
	log.Fatalf("server error: %v", http.ListenAndServe(":5000", nil))
}

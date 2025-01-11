package main

import (
	"fmt"
	"net/http"

	"go_setad_saater/handlers"
)

func main() {

	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/forgot-password", handlers.ForgotPasswordHandler)
	http.HandleFunc("/update-user", handlers.UpdateUserHandler)

	http.HandleFunc("/register-order", handlers.RegisterOrderHandler)                    // روت جدید
	http.HandleFunc("/update-order-files", handlers.UpdateOrderFilesHandler)             // روت جدید
	http.HandleFunc("/update-order-description", handlers.UpdateOrderDescriptionHandler) // روت جدید
	http.HandleFunc("/delete-order", handlers.DeleteOrderHandler)                        // روت جدید
	http.HandleFunc("/get-order", handlers.GetOrderHandler)                              // روت جدید
	http.HandleFunc("/get-user-orders", handlers.GetUserOrdersHandler)                   // روت جدید
	http.HandleFunc("/update-payment-receipt", handlers.UpdatePaymentReceiptHandler)
	http.HandleFunc("/update-invoice-file", handlers.UpdateInvoiceFileHandler)

	http.HandleFunc("/upload-files", handlers.UploadFilesHandler)
	http.HandleFunc("/delete-file", handlers.DeleteFileHandler)
	http.HandleFunc("/serve-file", handlers.ServeFileHandler)

	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

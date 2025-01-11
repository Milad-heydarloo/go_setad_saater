// package handlers

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"go_setad_saater/utils"
// 	"io"
// 	"net/http"
// )

// func RegisterOrderHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var orderData struct {
// 		DateSh       string `json:"date_sh"`
// 		DateAd       string `json:"date_ad"`
// 		UserId       string `json:"user"` // ای دی کاربر ارسال شده
// 		OrderProcess string `json:"order_process"`
// 	}

// 	err := json.NewDecoder(r.Body).Decode(&orderData)
// 	if err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	// ارسال داده‌ها به PocketBase
// 	data := map[string]interface{}{
// 		"date_sh":       orderData.DateSh,
// 		"date_ad":       orderData.DateAd,
// 		"user":          orderData.UserId,
// 		"order_process": orderData.OrderProcess,
// 	}

// 	orderId, err := utils.CreateOrder(data)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to create order: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	// برگرداندن آی‌دی سفارش به کلاینت
// 	response := map[string]string{
// 		"order_id": orderId,
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

// func UpdateOrderFilesHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPatch {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var requestData struct {
// 		OrderID string   `json:"order_id"` // ID سفارش
// 		FileIDs []string `json:"file_ids"` // لیست ID فایل‌ها
// 	}

// 	err := json.NewDecoder(r.Body).Decode(&requestData)
// 	if err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	// اعتبارسنجی ورودی‌ها
// 	if requestData.OrderID == "" || len(requestData.FileIDs) == 0 {
// 		http.Error(w, "Missing order_id or file_ids", http.StatusBadRequest)
// 		return
// 	}

// 	// آپدیت سفارش در PocketBase
// 	err = utils.UpdateOrderFiles(requestData.OrderID, requestData.FileIDs)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to update order files: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	// پاسخ موفقیت‌آمیز
// 	response := map[string]string{
// 		"status":  "ok",
// 		"message": "Order files updated successfully",
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

// func UpdateOrderDescriptionHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPatch {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var requestData struct {
// 		OrderID     string `json:"order_id"`    // ID سفارش
// 		Description string `json:"description"` // توضیحات جدید
// 	}

// 	err := json.NewDecoder(r.Body).Decode(&requestData)
// 	if err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	// اعتبارسنجی ورودی‌ها
// 	if requestData.OrderID == "" || requestData.Description == "" {
// 		http.Error(w, "Missing order_id or description", http.StatusBadRequest)
// 		return
// 	}

// 	// آپدیت فیلد description در PocketBase
// 	err = utils.UpdateOrderDescription(requestData.OrderID, requestData.Description)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to update order description: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	// پاسخ موفقیت‌آمیز
// 	response := map[string]string{
// 		"status":  "ok",
// 		"message": "Order description updated successfully",
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

// func DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var requestData struct {
// 		OrderID string   `json:"order_id"` // ID سفارش
// 		FileIDs []string `json:"file_ids"` // لیست ID فایل‌ها
// 	}

// 	err := json.NewDecoder(r.Body).Decode(&requestData)
// 	if err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	// اعتبارسنجی ورودی‌ها
// 	if requestData.OrderID == "" {
// 		http.Error(w, "Missing order_id", http.StatusBadRequest)
// 		return
// 	}

// 	// اگر لیست فایل خالی بود، مستقیماً سفارش را حذف کن
// 	if len(requestData.FileIDs) == 0 {
// 		err := utils.DeleteOrder(requestData.OrderID)
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("Failed to delete order: %v", err), http.StatusInternalServerError)
// 			return
// 		}

// 		// پاسخ موفقیت‌آمیز
// 		response := map[string]string{
// 			"status":  "ok",
// 			"message": "Order deleted successfully",
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	// حذف فایل‌ها یکی‌یکی
// 	for _, fileID := range requestData.FileIDs {
// 		fileDeleteRequest := map[string]string{
// 			"id": fileID,
// 		}

// 		fileDeleteBody, _ := json.Marshal(fileDeleteRequest)
// 		req, err := http.NewRequest("POST", "http://localhost:8080/delete-file", bytes.NewBuffer(fileDeleteBody))
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("Failed to create delete file request: %v", err), http.StatusInternalServerError)
// 			return
// 		}
// 		req.Header.Set("Content-Type", "application/json")

// 		client := &http.Client{}
// 		resp, err := client.Do(req)
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("Failed to delete file: %v", err), http.StatusInternalServerError)
// 			return
// 		}
// 		defer resp.Body.Close()

// 		if resp.StatusCode != http.StatusOK {
// 			body, _ := io.ReadAll(resp.Body)
// 			http.Error(w, fmt.Sprintf("Failed to delete file: %s", string(body)), http.StatusInternalServerError)
// 			return
// 		}
// 	}

// 	// وقتی فایل‌ها حذف شدند، سفارش را حذف کن
// 	err = utils.DeleteOrder(requestData.OrderID)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to delete order: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	// پاسخ موفقیت‌آمیز
// 	response := map[string]string{
// 		"status":  "ok",
// 		"message": "Order and all associated files deleted successfully",
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

// func GetOrderHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost { // استفاده از متد POST
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// خواندن Body درخواست
// 	var requestData struct {
// 		OrderID string `json:"order_id"` // ID سفارش
// 	}

// 	err := json.NewDecoder(r.Body).Decode(&requestData)
// 	if err != nil || requestData.OrderID == "" {
// 		http.Error(w, "Invalid or missing order_id", http.StatusBadRequest)
// 		return
// 	}

// 	// دریافت اطلاعات سفارش

// 	order, err := utils.GetOrder(requestData.OrderID)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to fetch order: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	// استخراج فیلدهای مورد نظر
// 	description := ""
// 	if desc, ok := order["description"].(string); ok {
// 		description = desc
// 	}

// 	files := []interface{}{}
// 	if expand, ok := order["expand"].(map[string]interface{}); ok {
// 		if fileData, ok := expand["file"].([]interface{}); ok {
// 			files = fileData
// 		}
// 	}

// 	// ساخت لیست فایل‌ها
// 	fileList := []map[string]string{}
// 	for _, file := range files {
// 		if fileData, ok := file.(map[string]interface{}); ok {
// 			fileID := ""
// 			fileName := ""
// 			if id, ok := fileData["id"].(string); ok {
// 				fileID = id
// 			}
// 			if fields, ok := fileData["field"].([]interface{}); ok && len(fields) > 0 {
// 				if name, ok := fields[0].(string); ok {
// 					fileName = name
// 				}
// 			}
// 			fileList = append(fileList, map[string]string{
// 				"id":   fileID,
// 				"name": fileName,
// 			})
// 		}
// 	}

// 	// پاسخ نهایی
// 	response := map[string]interface{}{
// 		"order_id":    requestData.OrderID,
// 		"description": description,
// 		"files":       fileList,
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

// func GetUserOrdersHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// خواندن Body درخواست
// 	var requestData struct {
// 		UserID string `json:"user_id"` // ID کاربر
// 	}

// 	err := json.NewDecoder(r.Body).Decode(&requestData)
// 	if err != nil || requestData.UserID == "" {
// 		http.Error(w, "Invalid or missing user_id", http.StatusBadRequest)
// 		return
// 	}

// 	// دریافت لیست سفارشات
// 	orders, err := utils.GetOrdersByUser(requestData.UserID)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Failed to fetch orders: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	// ساختاردهی خروجی
// 	var response []map[string]interface{}
// 	for _, order := range orders {
// 		// description
// 		description := ""
// 		if desc, ok := order["description"].(string); ok {
// 			description = desc
// 		}

// 		// date_sh
// 		dateSh := ""
// 		if ds, ok := order["date_sh"].(string); ok {
// 			dateSh = ds
// 		}

// 		// date_ad
// 		dateAd := ""
// 		if da, ok := order["date_ad"].(string); ok {
// 			dateAd = da
// 		}

// 		// user
// 		user := ""
// 		if u, ok := order["user"].(string); ok {
// 			user = u
// 		}

// 		// order_process
// 		orderProcess := ""
// 		if op, ok := order["order_process"].(string); ok {
// 			orderProcess = op
// 		}

// 		// فایل‌ها
// 		files := []map[string]string{}
// 		if expand, ok := order["expand"].(map[string]interface{}); ok {
// 			if fileData, ok := expand["file"].([]interface{}); ok {
// 				for _, file := range fileData {
// 					if fileMap, ok := file.(map[string]interface{}); ok {
// 						fileID := ""
// 						fileName := ""
// 						if id, ok := fileMap["id"].(string); ok {
// 							fileID = id
// 						}
// 						if fields, ok := fileMap["field"].([]interface{}); ok && len(fields) > 0 {
// 							if name, ok := fields[0].(string); ok {
// 								fileName = name
// 							}
// 						}
// 						files = append(files, map[string]string{
// 							"id":   fileID,
// 							"name": fileName,
// 						})
// 					}
// 				}
// 			}
// 		}

// 		// اکسپند invoice_file
// 		invoiceFile := map[string]string{
// 			"id":   "",
// 			"name": "",
// 		}
// 		if expand, ok := order["expand"].(map[string]interface{}); ok {
// 			if invoiceData, ok := expand["invoice_file"].(map[string]interface{}); ok {
// 				if id, ok := invoiceData["id"].(string); ok {
// 					invoiceFile["id"] = id
// 				}
// 				if fields, ok := invoiceData["field"].([]interface{}); ok && len(fields) > 0 {
// 					if name, ok := fields[0].(string); ok {
// 						invoiceFile["name"] = name
// 					}
// 				}
// 			}
// 		}

// 		// اکسپند payment_receipt_file
// 		paymentReceiptFile := map[string]string{
// 			"id":   "",
// 			"name": "",
// 		}
// 		if expand, ok := order["expand"].(map[string]interface{}); ok {
// 			if receiptData, ok := expand["payment_receipt_file"].(map[string]interface{}); ok {
// 				if id, ok := receiptData["id"].(string); ok {
// 					paymentReceiptFile["id"] = id
// 				}
// 				if fields, ok := receiptData["field"].([]interface{}); ok && len(fields) > 0 {
// 					if name, ok := fields[0].(string); ok {
// 						paymentReceiptFile["name"] = name
// 					}
// 				}
// 			}
// 		}

// 		// ساخت خروجی برای هر سفارش
// 		response = append(response, map[string]interface{}{
// 			"order_id":             order["id"],
// 			"description":          description,
// 			"date_sh":              dateSh,
// 			"date_ad":              dateAd,
// 			"user":                 user,
// 			"order_process":        orderProcess,
// 			"files":                files,
// 			"invoice_file":         invoiceFile,
// 			"payment_receipt_file": paymentReceiptFile,
// 		})
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_setad_saater/utils"
	"io"
	"net/http"
)

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "https://setad.saaterco.com")
	w.Header().Set("Access-Control-Allow-Methods", "POST, PATCH, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func RegisterOrderHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	// پاسخ به OPTIONS
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var orderData struct {
		DateSh       string `json:"date_sh"`
		DateAd       string `json:"date_ad"`
		UserId       string `json:"user"`
		OrderProcess string `json:"order_process"`
	}

	err := json.NewDecoder(r.Body).Decode(&orderData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	data := map[string]interface{}{
		"date_sh":       orderData.DateSh,
		"date_ad":       orderData.DateAd,
		"user":          orderData.UserId,
		"order_process": orderData.OrderProcess,
	}

	orderId, err := utils.CreateOrder(data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create order: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"order_id": orderId,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateOrderFilesHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	// پاسخ به OPTIONS
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodPatch {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		OrderID string   `json:"order_id"`
		FileIDs []string `json:"file_ids"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if requestData.OrderID == "" || len(requestData.FileIDs) == 0 {
		http.Error(w, "Missing order_id or file_ids", http.StatusBadRequest)
		return
	}

	err = utils.UpdateOrderFiles(requestData.OrderID, requestData.FileIDs)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update order files: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"status":  "ok",
		"message": "Order files updated successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateOrderDescriptionHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	// پاسخ به OPTIONS
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodPatch {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		OrderID     string `json:"order_id"`
		Description string `json:"description"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if requestData.OrderID == "" || requestData.Description == "" {
		http.Error(w, "Missing order_id or description", http.StatusBadRequest)
		return
	}

	err = utils.UpdateOrderDescription(requestData.OrderID, requestData.Description)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update order description: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"status":  "ok",
		"message": "Order description updated successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	// پاسخ به OPTIONS
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		OrderID string   `json:"order_id"`
		FileIDs []string `json:"file_ids"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if requestData.OrderID == "" {
		http.Error(w, "Missing order_id", http.StatusBadRequest)
		return
	}

	for _, fileID := range requestData.FileIDs {
		fileDeleteRequest := map[string]string{
			"id": fileID,
		}

		fileDeleteBody, _ := json.Marshal(fileDeleteRequest)
		req, err := http.NewRequest("POST", "http://localhost:8080/delete-file", bytes.NewBuffer(fileDeleteBody))
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to create delete file request: %v", err), http.StatusInternalServerError)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to delete file: %v", err), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			http.Error(w, fmt.Sprintf("Failed to delete file: %s", string(body)), http.StatusInternalServerError)
			return
		}
	}

	err = utils.DeleteOrder(requestData.OrderID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete order: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"status":  "ok",
		"message": "Order and all associated files deleted successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	// پاسخ به OPTIONS
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		OrderID string `json:"order_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil || requestData.OrderID == "" {
		http.Error(w, "Invalid or missing order_id", http.StatusBadRequest)
		return
	}

	order, err := utils.GetOrder(requestData.OrderID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch order: %v", err), http.StatusInternalServerError)
		return
	}

	description := ""
	if desc, ok := order["description"].(string); ok {
		description = desc
	}

	files := []interface{}{}
	if expand, ok := order["expand"].(map[string]interface{}); ok {
		if fileData, ok := expand["file"].([]interface{}); ok {
			files = fileData
		}
	}

	fileList := []map[string]string{}
	for _, file := range files {
		if fileData, ok := file.(map[string]interface{}); ok {
			fileID := ""
			fileName := ""
			if id, ok := fileData["id"].(string); ok {
				fileID = id
			}
			if fields, ok := fileData["field"].([]interface{}); ok && len(fields) > 0 {
				if name, ok := fields[0].(string); ok {
					fileName = name
				}
			}
			fileList = append(fileList, map[string]string{
				"id":   fileID,
				"name": fileName,
			})
		}
	}

	response := map[string]interface{}{
		"order_id":    requestData.OrderID,
		"description": description,
		"files":       fileList,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetUserOrdersHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	// پاسخ به OPTIONS
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		UserID string `json:"user_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil || requestData.UserID == "" {
		http.Error(w, "Invalid or missing user_id", http.StatusBadRequest)
		return
	}

	orders, err := utils.GetOrdersByUser(requestData.UserID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch orders: %v", err), http.StatusInternalServerError)
		return
	}

	var response []map[string]interface{}
	for _, order := range orders {
		description := ""
		if desc, ok := order["description"].(string); ok {
			description = desc
		}

		dateSh := ""
		if ds, ok := order["date_sh"].(string); ok {
			dateSh = ds
		}

		dateAd := ""
		if da, ok := order["date_ad"].(string); ok {
			dateAd = da
		}

		user := ""
		if u, ok := order["user"].(string); ok {
			user = u
		}

		orderProcess := ""
		if op, ok := order["order_process"].(string); ok {
			orderProcess = op
		}

		files := []map[string]string{}
		if expand, ok := order["expand"].(map[string]interface{}); ok {
			if fileData, ok := expand["file"].([]interface{}); ok {
				for _, file := range fileData {
					if fileMap, ok := file.(map[string]interface{}); ok {
						fileID := ""
						fileName := ""
						if id, ok := fileMap["id"].(string); ok {
							fileID = id
						}
						if fields, ok := fileMap["field"].([]interface{}); ok && len(fields) > 0 {
							if name, ok := fields[0].(string); ok {
								fileName = name
							}
						}
						files = append(files, map[string]string{
							"id":   fileID,
							"name": fileName,
						})
					}
				}
			}
		}

		invoiceFile := map[string]string{
			"id":   "",
			"name": "",
		}
		if expand, ok := order["expand"].(map[string]interface{}); ok {
			if invoiceData, ok := expand["invoice_file"].(map[string]interface{}); ok {
				if id, ok := invoiceData["id"].(string); ok {
					invoiceFile["id"] = id
				}
				if fields, ok := invoiceData["field"].([]interface{}); ok && len(fields) > 0 {
					if name, ok := fields[0].(string); ok {
						invoiceFile["name"] = name
					}
				}
			}
		}

		paymentReceiptFile := map[string]string{
			"id":   "",
			"name": "",
		}
		if expand, ok := order["expand"].(map[string]interface{}); ok {
			if receiptData, ok := expand["payment_receipt_file"].(map[string]interface{}); ok {
				if id, ok := receiptData["id"].(string); ok {
					paymentReceiptFile["id"] = id
				}
				if fields, ok := receiptData["field"].([]interface{}); ok && len(fields) > 0 {
					if name, ok := fields[0].(string); ok {
						paymentReceiptFile["name"] = name
					}
				}
			}
		}

		response = append(response, map[string]interface{}{
			"order_id":             order["id"],
			"description":          description,
			"date_sh":              dateSh,
			"date_ad":              dateAd,
			"user":                 user,
			"order_process":        orderProcess,
			"files":                files,
			"invoice_file":         invoiceFile,
			"payment_receipt_file": paymentReceiptFile,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

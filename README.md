{
	"info": {
		"_postman_id": "c1ed5ae6-30ef-4752-a096-1a13a3b8c906",
		"name": "saater_setad",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "16273864"
	},
	"item": [
		{
			"name": "order_manager",
			"item": [
				{
					"name": "https://s-sater.liara.run/register-order",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "POST",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\n  \"date_sh\": \"1402/10/19\",\n  \"date_ad\": \"2025-01-09\",\n  \"user\": \"kvbm4a40in6wvc2\",\n  \"order_process\": \"0.2\"\n}\n",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "https://s-sater.liara.run/register-order",
							"protocol": "https",
							"host": [
								"s-sater",
								"liara",
								"run"
							],
							"path": [
								"register-order"
							]
						}
					},
					"response": []
				},
				{
					"name": "https://s-sater.liara.run/update-order-files",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "PATCH",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\n  \"order_id\": \"zj1l0dtj95p1jfw\",\n  \"file_ids\": [\"8118zbau4end9th\", \"una8deywd4lio48\"]\n}\n",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "https://s-sater.liara.run/update-order-files",
							"protocol": "https",
							"host": [
								"s-sater",
								"liara",
								"run"
							],
							"path": [
								"update-order-files"
							],
							"query": [
								{
									"key": "",
									"value": null,
									"disabled": true
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "https://s-sater.liara.run/update-order-description",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "PATCH",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\n  \"order_id\": \"zj1l0dtj95p1jfw\",\n  \"description\": \"This is a new description for sthe order.\"\n}\n",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "https://s-sater.liara.run/update-order-description",
							"protocol": "https",
							"host": [
								"s-sater",
								"liara",
								"run"
							],
							"path": [
								"update-order-description"
							]
						}
					},
					"response": []
				},
				{
					"name": "https://s-sater.liara.run/delete-order",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "POST",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "// {\n//   \"order_id\": \"zj1l0dtj95p1jfw\",\n//   \"file_ids\": []\n// }\n\n{\n  \"order_id\": \"zj1l0dtj95p1jfw\",\n  \"file_ids\": [\"pfqlig6tcqd21x9\", \"7zs1bgmygh36e3m\"]\n}\n",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "https://s-sater.liara.run/delete-order",
							"protocol": "https",
							"host": [
								"s-sater",
								"liara",
								"run"
							],
							"path": [
								"delete-order"
							]
						}
					},
					"response": []
				},
				{
					"name": "https://s-sater.liara.run/get-order",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "POST",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "// {\n//   \"order_id\": \"123456789\"\n// }\n\n\n\n{\n  \"order_id\": \"zj1l0dtj95p1jfw\"\n}\n",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "https://s-sater.liara.run/get-order",
							"protocol": "https",
							"host": [
								"s-sater",
								"liara",
								"run"
							],
							"path": [
								"get-order"
							]
						}
					},
					"response": []
				},
				{
					"name": "https://s-sater.liara.run/get-user-orders",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "POST",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\n  \"user_id\": \"kvbm4a40in6wvc2\"\n}\n",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "https://s-sater.liara.run/get-user-orders",
							"protocol": "https",
							"host": [
								"s-sater",
								"liara",
								"run"
							],
							"path": [
								"get-user-orders"
							]
						}
					},
					"response": []
				},
				{
					"name": "https://s-sater.liara.run/update-payment-receipt",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "POST",
						"header": [
							{
								"key": "Content-Type",
								"value": "multipart/form-data",
								"type": "text"
							}
						],
						"body": {
							"mode": "formdata",
							"formdata": [
								{
									"key": "order_id",
									"value": "8knue0ocu9s83jr",
									"type": "text"
								},
								{
									"key": "file",
									"type": "file",
									"src": "postman-cloud:///1efce681-adf0-49c0-8ae4-81d11db10f9c"
								}
							]
						},
						"url": {
							"raw": "https://s-sater.liara.run/update-payment-receipt",
							"protocol": "https",
							"host": [
								"s-sater",
								"liara",
								"run"
							],
							"path": [
								"update-payment-receipt"
							]
						}
					},
					"response": []
				},
				{
					"name": "https://s-sater.liara.run/update-invoice-file",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "POST",
						"header": [],
						"body": {
							"mode": "formdata",
							"formdata": [
								{
									"key": "order_id",
									"value": "zj1l0dtj95p1jfw",
									"type": "text"
								},
								{
									"key": "file",
									"type": "file",
									"src": "postman-cloud:///1efce67d-3d0b-4b00-b690-4445ca6bc53b"
								}
							]
						},
						"url": {
							"raw": "https://s-sater.liara.run/update-invoice-file",
							"protocol": "https",
							"host": [
								"s-sater",
								"liara",
								"run"
							],
							"path": [
								"update-invoice-file"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "user",
			"item": [
				{
					"name": "https://s-sater.liara.run/login",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "// {\n//   \"organization_code\": \"1323456\",\n//   \"password\": \"password123\"\n// }\n\n\n{\n  \"organization_code\": \"132s3456\",\n  \"password\": \"passworxd123\"\n}\n",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "https://s-sater.liara.run/login",
							"protocol": "https",
							"host": [
								"s-sater",
								"liara",
								"run"
							],
							"path": [
								"login"
							]
						}
					},
					"response": []
				},
				{
					"name": "https://s-sater.liara.run/forgot-password",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "POST",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\n  \"mobile_number\": \"09013757395\"\n}\n",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "https://s-sater.liara.run/forgot-password",
							"protocol": "https",
							"host": [
								"s-sater",
								"liara",
								"run"
							],
							"path": [
								"forgot-password"
							]
						}
					},
					"response": []
				},
				{
					"name": "https://s-sater.liara.run/update-user",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "POST",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "text"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\n  \"id\": \"biyl3sjp14znujo\",\n  \"organizational_address\": \" \",\n  \"landline_number\": \" \"\n}\n",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "https://s-sater.liara.run/update-user",
							"protocol": "https",
							"host": [
								"s-sater",
								"liara",
								"run"
							],
							"path": [
								"update-user"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "file",
			"item": [
				{
					"name": "https://s-sater.liara.run/upload-files",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "POST",
						"header": [],
						"body": {
							"mode": "formdata",
							"formdata": [
								{
									"key": "field",
									"type": "file",
									"src": [
										"postman-cloud:///1efce67d-3d0b-4b00-b690-4445ca6bc53b",
										"postman-cloud:///1efce681-adf0-49c0-8ae4-81d11db10f9c"
									]
								}
							]
						},
						"url": {
							"raw": "https://s-sater.liara.run/upload-files",
							"protocol": "https",
							"host": [
								"s-sater",
								"liara",
								"run"
							],
							"path": [
								"upload-files"
							]
						}
					},
					"response": []
				},
				{
					"name": "https://s-sater.liara.run/delete-file",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\n  \"id\": \"8vq8l09lfqwhe9f\"\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "https://s-sater.liara.run/delete-file",
							"protocol": "https",
							"host": [
								"s-sater",
								"liara",
								"run"
							],
							"path": [
								"delete-file"
							]
						}
					},
					"response": []
				},
				{
					"name": "https://s-sater.liara.run/serve-file?file_id=4sc9jf6c7lt24re&action=view",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "GET",
						"header": [],
						"url": {
							"raw": "https://s-sater.liara.run/serve-file?file_id=4sc9jf6c7lt24re&action=view",
							"protocol": "https",
							"host": [
								"s-sater",
								"liara",
								"run"
							],
							"path": [
								"serve-file"
							],
							"query": [
								{
									"key": "file_id",
									"value": "4sc9jf6c7lt24re"
								},
								{
									"key": "action",
									"value": "view"
								}
							]
						}
					},
					"response": []
				}
			]
		}
	]
}

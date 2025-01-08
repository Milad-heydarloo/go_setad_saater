curl -X POST https://s-sater.liara.run/register \
-H "Content-Type: application/json" \
-d '{
  "organization_code": "123456",
  "landline_number": "02112345678",
  "email": "example@test.com",
  "password": "password123",
  "full_name": "John Doe",
  "organizational_address": "Tehran, Iran",
  "mobile_number": "09121234567"
}'




curl -X POST https://s-sater.liara.run/login \
-H "Content-Type: application/json" \
-d '{
  "organization_code": "123456",
  "password": "password123"
}'


curl -X POST https://s-sater.liara.run/forgot-password \
-H "Content-Type: application/json" \
-d '{
  "mobile_number": "09121234567"
}'

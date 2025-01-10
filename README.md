# API Documentation for Saater Setad

A detailed description of APIs used in the Saater Setad project, including the purpose of each request and the relevant parameters.

---

## Table of Contents

1. [User APIs](#user-apis)
2. [Order Manager APIs](#order-manager-apis)
3. [File APIs](#file-apis)

---

## User APIs

### Login

#### Endpoint
```http
POST https://s-sater.liara.run/login
```

#### Body
```json
{
  "organization_code": "132s3456",
  "password": "passworxd123"
}
```

#### Description
این API برای ورود کاربر با استفاده از کد سازمان و رمز عبور استفاده می‌شود.

---

### Forgot Password

#### Endpoint
```http
POST https://s-sater.liara.run/forgot-password
```

#### Headers
| Key           | Value          | Description               |
|---------------|----------------|---------------------------|
| Content-Type  | application/json | Specifies the content type of the request. |

#### Body
```json
{
  "mobile_number": "09013757395"
}
```

#### Description
این API برای بازیابی رمز عبور از طریق شماره موبایل کاربر استفاده می‌شود.

---

### Update User

#### Endpoint
```http
POST https://s-sater.liara.run/update-user
```

#### Headers
| Key           | Value          | Description               |
|---------------|----------------|---------------------------|
| Content-Type  | application/json | Specifies the content type of the request. |

#### Body
```json
{
  "id": "biyl3sjp14znujo",
  "organizational_address": " ",
  "landline_number": " "
}
```

#### Description
این API برای به‌روزرسانی اطلاعات کاربر استفاده می‌شود.

---

### Register

#### Endpoint
```http
POST https://s-sater.liara.run/register
```

#### Headers
| Key           | Value          | Description               |
|---------------|----------------|---------------------------|
| Content-Type  | application/json | Specifies the content type of the request. |

#### Body
```json
{
  "organization_code": "123443345",
  "landline_number": "02112345678",
  "email": "exam3dple@example.com",
  "password": "securepassword",
  "full_name": "John Doe",
  "organizational_address": "123 Street, City",
  "mobile_number": "09123456789"
}
```

#### Description
این API برای ثبت‌نام کاربر جدید استفاده می‌شود.

---

## Order Manager APIs

### Register Order

#### Endpoint
```http
POST https://s-sater.liara.run/register-order
```

#### Headers
| Key           | Value          | Description               |
|---------------|----------------|---------------------------|
| Content-Type  | application/json | Specifies the content type of the request. |

#### Body
```json
{
  "date_sh": "1402/10/19",
  "date_ad": "2025-01-09",
  "user": "kvbm4a40in6wvc2",
  "order_process": "0.2"
}
```

#### Description
این API برای ثبت سفارش جدید استفاده می‌شود. تاریخ شمسی و میلادی، شناسه کاربر و وضعیت فرایند سفارش ارسال می‌شود.

---

### Update Order Files

#### Endpoint
```http
PATCH https://s-sater.liara.run/update-order-files
```

#### Headers
| Key           | Value          | Description               |
|---------------|----------------|---------------------------|
| Content-Type  | application/json | Specifies the content type of the request. |

#### Body
```json
{
  "order_id": "zj1l0dtj95p1jfw",
  "file_ids": ["8118zbau4end9th", "una8deywd4lio48"]
}
```

#### Description
این API برای به‌روزرسانی فایل‌های مربوط به یک سفارش خاص استفاده می‌شود. فایل‌ها باید با شناسه‌های خاص خود ارسال شوند.

---

### Update Order Description

#### Endpoint
```http
PATCH https://s-sater.liara.run/update-order-description
```

#### Headers
| Key           | Value          | Description               |
|---------------|----------------|---------------------------|
| Content-Type  | application/json | Specifies the content type of the request. |

#### Body
```json
{
  "order_id": "zj1l0dtj95p1jfw",
  "description": "This is a new description for the order."
}
```

#### Description
این API برای تغییر توضیحات سفارش خاص استفاده می‌شود. شناسه سفارش و توضیحات جدید ارسال می‌شوند.

---

### Delete Order

#### Endpoint
```http
POST https://s-sater.liara.run/delete-order
```

#### Headers
| Key           | Value          | Description               |
|---------------|----------------|---------------------------|
| Content-Type  | application/json | Specifies the content type of the request. |

#### Body
```json
{
  "order_id": "zj1l0dtj95p1jfw",
  "file_ids": ["pfqlig6tcqd21x9", "7zs1bgmygh36e3m"]
}
```

#### Description
این API برای حذف سفارش و فایل‌های مرتبط با آن استفاده می‌شود. شناسه سفارش و فایل‌های مرتبط باید ارسال شوند.

---

### Get Order

#### Endpoint
```http
POST https://s-sater.liara.run/get-order
```

#### Headers
| Key           | Value          | Description               |
|---------------|----------------|---------------------------|
| Content-Type  | application/json | Specifies the content type of the request. |

#### Body
```json
{
  "order_id": "zj1l0dtj95p1jfw"
}
```

#### Description
این API برای دریافت اطلاعات یک سفارش خاص با استفاده از شناسه سفارش استفاده می‌شود.

---

### Get User Orders

#### Endpoint
```http
POST https://s-sater.liara.run/get-user-orders
```

#### Headers
| Key           | Value          | Description               |
|---------------|----------------|---------------------------|
| Content-Type  | application/json | Specifies the content type of the request. |

#### Body
```json
{
  "user_id": "kvbm4a40in6wvc2"
}
```

#### Description
این API برای دریافت لیست سفارش‌های کاربر استفاده می‌شود.

---

### Update Payment Receipt

#### Endpoint
```http
POST https://s-sater.liara.run/update-payment-receipt
```

#### Headers
| Key           | Value          | Description               |
|---------------|----------------|---------------------------|
| Content-Type  | multipart/form-data | Specifies the content type of the request. |

#### Body
Form Data:
| Key      | Value                     | Type   |
|----------|---------------------------|--------|
| order_id | 8knue0ocu9s83jr           | text   |
| file     | postman-cloud:///file_url | file   |

#### Description
این API برای به‌روزرسانی رسید پرداخت مربوط به یک سفارش استفاده می‌شود.

---

### Update Invoice File

#### Endpoint
```http
POST https://s-sater.liara.run/update-invoice-file
```

#### Body
Form Data:
| Key      | Value                     | Type   |
|----------|---------------------------|--------|
| order_id | zj1l0dtj95p1jfw           | text   |
| file     | postman-cloud:///file_url | file   |

#### Description
این API برای به‌روزرسانی فایل فاکتور مربوط به سفارش استفاده می‌شود.

---

## File APIs

### Upload Files

#### Endpoint
```http
POST https://s-sater.liara.run/upload-files
```

#### Body
Form Data:
| Key    | Type   | Description        |
|--------|--------|--------------------|
| field  | file   | فایل‌هایی که باید آپلود شوند. |

#### Description
این API برای آپلود فایل‌ها به سرور استفاده می‌شود.

---

### Delete File

#### Endpoint
```http
POST https://s-sater.liara.run/delete-file
```

#### Body
```json
{
  "id": "8vq8l09lfqwhe9f"
}
```

#### Description
این API برای حذف یک فایل خاص از سرور استفاده می‌شود.

---

### Serve File

#### Endpoint
```http
GET https://s-sater.liara.run/serve-file?file_id=4sc9jf6c7lt24re&action=view
```

#### Query Parameters
| Key       | Value         | Description                     |
|-----------|---------------|---------------------------------|
| file_id   | 4sc9jf6c7lt24re | شناسه فایل                     |
| action    | view          | عملیات مورد نظر (نمایش فایل)   |

#### Description
این API برای نمایش فایل از سرور استفاده می‌شود.

---


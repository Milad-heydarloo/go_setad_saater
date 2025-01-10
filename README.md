<!DOCTYPE html>
<html lang="fa">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>مستندات پروژه</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.8;
            direction: rtl;
            text-align: right;
            margin: 0;
            padding: 0;
            background-color: #f9f9f9;
            color: #333;
        }
        header {
            background-color: #007bff;
            color: white;
            padding: 1rem;
            text-align: center;
        }
        main {
            max-width: 800px;
            margin: 2rem auto;
            background: white;
            padding: 2rem;
            border-radius: 10px;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        }
        h1, h2, h3 {
            color: #007bff;
        }
        code {
            background: #f1f1f1;
            padding: 0.2rem 0.4rem;
            border-radius: 5px;
            font-size: 0.9rem;
        }
        pre {
            background: #f1f1f1;
            padding: 1rem;
            border-radius: 10px;
            overflow-x: auto;
            margin: 1rem 0;
        }
        section {
            margin-bottom: 2rem;
        }
        footer {
            text-align: center;
            padding: 1rem;
            background-color: #f1f1f1;
            color: #777;
            font-size: 0.9rem;
        }
    </style>
</head>
<body>
    <header>
        <h1>مستندات پروژه</h1>
        <p>راهنمای استفاده از APIهای پروژه</p>
    </header>
    <main>
        <section>
            <h2>ساختار پروژه</h2>
            <p>این پروژه شامل سه بخش اصلی است:</p>
            <ul>
                <li>مدیریت سفارشات</li>
                <li>مدیریت کاربران</li>
                <li>مدیریت فایل‌ها</li>
            </ul>
        </section>
        <section>
            <h2>مدیریت سفارشات</h2>
            <h3>ثبت سفارش</h3>
            <p><strong>آدرس:</strong> <code>https://s-sater.liara.run/register-order</code></p>
            <p><strong>روش:</strong> <code>POST</code></p>
            <p><strong>نمونه داده ورودی:</strong></p>
            <pre>
{
  "date_sh": "1402/10/19",
  "date_ad": "2025-01-09",
  "user": "kvbm4a40in6wvc2",
  "order_process": "0.2"
}
            </pre>
        </section>
        <section>
            <h2>مدیریت کاربران</h2>
            <h3>ورود کاربر</h3>
            <p><strong>آدرس:</strong> <code>https://s-sater.liara.run/login</code></p>
            <p><strong>روش:</strong> <code>POST</code></p>
            <p><strong>نمونه داده ورودی:</strong></p>
            <pre>
{
  "organization_code": "132s3456",
  "password": "passworxd123"
}
            </pre>
        </section>
        <section>
            <h2>مدیریت فایل‌ها</h2>
            <h3>آپلود فایل</h3>
            <p><strong>آدرس:</strong> <code>https://s-sater.liara.run/upload-files</code></p>
            <p><strong>روش:</strong> <code>POST</code></p>
            <p><strong>هدر:</strong> <code>Content-Type: multipart/form-data</code></p>
        </section>
    </main>
    <footer>
        <p>این پروژه تحت لایسنس MIT منتشر شده است.</p>
    </footer>
</body>
</html>

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
            max-width: 900px;
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
        ul {
            padding-right: 1rem;
            list-style: none;
        }
        ul li {
            margin: 0.5rem 0;
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

        <!-- مدیریت سفارشات -->
        <section>
            <h2>مدیریت سفارشات</h2>

            <!-- ثبت سفارش -->
            <h3>۱. ثبت سفارش</h3>
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

            <!-- بروزرسانی فایل‌های سفارش -->
            <h3>۲. بروزرسانی فایل‌های سفارش</h3>
            <p><strong>آدرس:</strong> <code>https://s-sater.liara.run/update-order-files</code></p>
            <p><strong>روش:</strong> <code>PATCH</code></p>
            <p><strong>نمونه داده ورودی:</strong></p>
            <pre>
{
  "order_id": "zj1l0dtj95p1jfw",
  "file_ids": ["8118zbau4end9th", "una8deywd4lio48"]
}
            </pre>

            <!-- بروزرسانی توضیحات سفارش -->
            <h3>۳. بروزرسانی توضیحات سفارش</h3>
            <p><strong>آدرس:</strong> <code>https://s-sater.liara.run/update-order-description</code></p>
            <p><strong>روش:</strong> <code>PATCH</code></p>
            <p><strong>نمونه داده ورودی:</strong></p>
            <pre>
{
  "order_id": "zj1l0dtj95p1jfw",
  "description": "This is a new description for the order."
}
            </pre>

            <!-- حذف سفارش -->
            <h3>۴. حذف سفارش</h3>
            <p><strong>آدرس:</strong> <code>https://s-sater.liara.run/delete-order</code></p>
            <p><strong>روش:</strong> <code>POST</code></p>
            <p><strong>نمونه داده ورودی:</strong></p>
            <pre>
{
  "order_id": "zj1l0dtj95p1jfw",
  "file_ids": ["pfqlig6tcqd21x9", "7zs1bgmygh36e3m"]
}
            </pre>

            <!-- دریافت اطلاعات سفارش -->
            <h3>۵. دریافت اطلاعات سفارش</h3>
            <p><strong>آدرس:</strong> <code>https://s-sater.liara.run/get-order</code></p>
            <p><strong>روش:</strong> <code>POST</code></p>
            <p><strong>نمونه داده ورودی:</strong></p>
            <pre>
{
  "order_id": "zj1l0dtj95p1jfw"
}
            </pre>

            <!-- دریافت سفارش‌های کاربر -->
            <h3>۶. دریافت سفارش‌های کاربر</h3>
            <p><strong>آدرس:</strong> <code>https://s-sater.liara.run/get-user-orders</code></p>
            <p><strong>روش:</strong> <code>POST</code></p>
            <p><strong>نمونه داده ورودی:</strong></p>
            <pre>
{
  "user_id": "kvbm4a40in6wvc2"
}
            </pre>
        </section>

        <!-- مدیریت کاربران -->
        <section>
            <h2>مدیریت کاربران</h2>

            <!-- ورود کاربر -->
            <h3>۱. ورود کاربر</h3>
            <p><strong>آدرس:</strong> <code>https://s-sater.liara.run/login</code></p>
            <p><strong>روش:</strong> <code>POST</code></p>
            <p><strong>نمونه داده ورودی:</strong></p>
            <pre>
{
  "organization_code": "132s3456",
  "password": "passworxd123"
}
            </pre>

            <!-- بازیابی رمز عبور -->
            <h3>۲. بازیابی رمز عبور</h3>
            <p><strong>آدرس:</strong> <code>https://s-sater.liara.run/forgot-password</code></p>
            <p><strong>روش:</strong> <code>POST</code></p>
            <p><strong>نمونه داده ورودی:</strong></p>
            <pre>
{
  "mobile_number": "09013757395"
}
            </pre>
        </section>

        <!-- مدیریت فایل‌ها -->
        <section>
            <h2>مدیریت فایل‌ها</h2>

            <!-- آپلود فایل -->
            <h3>۱. آپلود فایل</h3>
            <p><strong>آدرس:</strong> <code>https://s-sater.liara.run/upload-files</code></p>
            <p><strong>روش:</strong> <code>POST</code></p>
            <p><strong>نوع ارسال:</strong> <code>multipart/form-data</code></p>

            <!-- حذف فایل -->
            <h3>۲. حذف فایل</h3>
            <p><strong>آدرس:</strong> <code>https://s-sater.liara.run/delete-file</code></p>
            <p><strong>روش:</strong> <code>POST</code></p>
            <p><strong>نمونه داده ورودی:</strong></p>
            <pre>
{
  "id": "8vq8l09lfqwhe9f"
}
            </pre>

            <!-- مشاهده فایل -->
            <h3>۳. مشاهده فایل</h3>
            <p><strong>آدرس:</strong> <code>https://s-sater.liara.run/serve-file?file_id=4sc9jf6c7lt24re&action=view</code></p>
            <p><strong>روش:</strong> <code>GET</code></p>
        </section>
    </main>
    <footer>
        <p>این پروژه تحت لایسنس MIT منتشر شده است.</p>
    </footer>
</body>
</html>

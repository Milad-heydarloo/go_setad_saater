<!DOCTYPE html>
<html lang="fa">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>مستندات پروژه</title>
    
       
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

            <h3>۵. دریافت اطلاعات سفارش</h3>
            <p><strong>آدرس:</strong> <code>https://s-sater.liara.run/get-order</code></p>
            <p><strong>روش:</strong> <code>POST</code></p>
            <p><strong>نمونه داده ورودی:</strong></p>
            <pre>
{
  "order_id": "zj1l0dtj95p1jfw"
}
            </pre>
        </section>

        <section>
            <h2>مدیریت کاربران</h2>
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

        <section>
            <h2>مدیریت فایل‌ها</h2>
            <h3>۱. آپلود فایل</h3>
            <p><strong>آدرس:</strong> <code>https://s-sater.liara.run/upload-files</code></p>
            <p><strong>روش:</strong> <code>POST</code></p>
            <p><strong>نوع ارسال:</strong> <code>multipart/form-data</code></p>

            <h3>۲. مشاهده فایل</h3>
            <p><strong>آدرس:</strong> <code>https://s-sater.liara.run/serve-file?file_id=4sc9jf6c7lt24re&action=view</code></p>
            <p><strong>روش:</strong> <code>GET</code></p>
        </section>
    </main>
    <footer>
        <p>این پروژه تحت لایسنس MIT منتشر شده است.</p>
    </footer>
</body>
</html>

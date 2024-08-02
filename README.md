# Uy vazifasi: Repository Pattern-dan foydalangan holda RESTful API-ni amalga oshirish

## Maqsad
`Repository` pattern-dan foydalangan holda kitoblar ro'yxatini boshqarish uchun asosiy `RESTful API`-ni amalga oshiring. Ilova `PostgreSQL`-dan ma'lumotlar ombori sifatida foydalanishi va `Go` ilovalarini tuzish uchun eng yaxshi amaliyotlarga rioya qilishi kerak.

## Talablar
1. Go ilovasini yarating:
    - `Repository` pattern-dan foydalananing.
    - Kitoblarni boshqarish uchun asosiy `CRUD` (Yaratish, O'qish, Yangilash, O'chirish) operatsiyalariga ega bo'ladi.
    - `PostgreSQL`-dan ma'lumotlar ombori sifatida foydalaning.
    - Oson sozlash va joylashtirish uchun dockerizatsiya qiling.

2. `Book` modeli quyidagi maydonlarga ega bo'lishi kerak:
    - `ID` (int, primary key)
    - `Title` (string)
    - `Author` (string)
    - `PublishedYear` (int)

3. Quyidagi endpointlarni amalga oshiring:

    - `GET /books`: Barcha kitoblarni olish.
    - `GET /books/{id}`: ID bo'yicha kitobni olish.
    - `POST /books`: Yangi kitob yaratish.
    - `PUT /books/{id}`: Mavjud kitobni yangilash.
    - `DELETE /books/{id}`: ID bo'yicha kitobni o'chirish.

4. Ma'lumotlar bazasi konfiguratsiyasi uchun `environment variables` foydalaning.

5. Ilovani va `PostgreSQL` ma'lumotlar bazasini ishga tushirish uchun `Dockerfile` va `Docker Compose` faylini taqdim eting.

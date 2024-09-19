-- Lihat Detail Order berdasarkan ID (order_id = 1)
SELECT * FROM "order" WHERE order_id = 1;

-- Lihat Daftar Order
SELECT * FROM "order";

-- Selesaikan Order (set completion_date = sekarang untuk order_id = 1)
UPDATE \"order\" SET completion_date = '2024-09-20' WHERE order_id = 1;

-- Cek keberadaan Order (order_id = 1)
SELECT order_id FROM "order" WHERE order_id = 1;

-- Buat Order Detail (order_id = 1, service_id = 2, qty = 3)
INSERT INTO order_detail (order_id, service_id, qty) VALUES (1, 2, 3);

-- Buat Order (order_id = 1, customer_id = 1, order_date = sekarang, received_by = 'John')
INSERT INTO "order" (order_id, customer_id, order_date, received_by) VALUES (1, 1, '2024-09-20', 'John');

-- Cek keberadaan Customer (customer_id = 1)
SELECT customer_id FROM customer WHERE customer_id = 1;

-- Hapus Service (service_id = 1)
DELETE FROM service WHERE service_id = 1;

-- Cek Service di Order Detail (service_id = 1)
SELECT service_id FROM order_detail WHERE service_id = 1;

-- Cek keberadaan Service (service_id = 1)
SELECT service_id FROM service WHERE service_id = 1;

-- Update Service (service_id = 1, nama = 'Laundry', unit = 'kg', harga = 5000)
UPDATE service SET service_name = 'Laundy', unit = 'kg', price = 5000 WHERE service_id = 1;

-- Pilih Service berdasarkan ID (service_id = 1)
SELECT * FROM service WHERE service_id = 1;

-- Pilih Semua Service
SELECT * FROM service;

-- Tambah Service (service_id = 1, nama = 'Laundry', unit = 'kg', harga = 5000)
INSERT INTO service (service_id, service_name, unit, price) VALUES (1, 'Laundry', 'kg', 5000);

-- Hapus Customer (customer_id = 1)
DELETE FROM customer WHERE customer_id = 1;

-- Cek keberadaan Customer (customer_id = 1)
SELECT customer_id FROM customer WHERE customer_id = 1;

-- Update Customer (customer_id = 1, nama = 'Alice', telepon = '12345', alamat = 'Jl. Utama')
UPDATE customer SET name = 'Alice', phone = '12345', address = 'Jl. Utama' WHERE customer_id = 1;

-- Pilih Customer berdasarkan ID (customer_id = 1)
SELECT * FROM customer WHERE customer_id = 1;

-- Pilih Semua Customer
SELECT * FROM customer;

-- Tambah Customer (customer_id = 1, nama = 'Alice', telepon = '12345', alamat = 'Jl. Utama')
INSERT INTO customer (customer_id, name, phone, address) VALUES (1, 'Alice', '12345', 'Jl. Utama');


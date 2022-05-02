CREATE TABLE employee (
	emp_id varchar(12) NOT null PRIMARY KEY ,
	emp_name varchar(255) NULL,
	phone_no varchar(32) NULL,
	address varchar(255) NULL,
	join_date date NULL
);

INSERT INTO public.employee (emp_id,emp_name,phone_no,address,join_date) VALUES
	 ('E0013','Marrie Moore','0836453245','Jl. Jarih No 43','2021-01-02'),
	 ('E0014','Naylendra Nilam','0847442472','Jl. Daniel Sambodo No 43','2021-02-03'),
	 ('E0015','Olivia Oscar','0858544573','Jl. Kranggan Wetan No 3','2021-03-04'),
	 ('E0016','Padme Parker','0863623436','Jl. Jambu Air No 15','2021-04-05'),
	 ('E0017','Quincy Queen','0873912362','Jl. Ramah Air No 35','2021-05-06'),
	 ('E0001','Alex Andrian','0812313421','Jl. Ragunan No 1, RT 06/01, Jakarta Selatan','2021-01-02'),
	 ('E0002','Brad Binder','0825424523','Jl. Tanah Abang, No 45 Jakarta Pusat','2021-02-03'),
	 ('E0003','Clay Calloway','0835242533','Jl. Slipi Raya No 87, RT 3/4 Jakarta Selatan','2021-03-04'),
	 ('E0004','Dana Donovan','0845645345','Jl. Abdul Manaf No 22, RT 4/5 Jakarta Timur','2021-04-05'),
	 ('E0005','Ewan Enoch','0857656576','Jl. Jaka Sembung No 83, Jakarta Barat','2021-05-06');
INSERT INTO public.employee (emp_id,emp_name,phone_no,address,join_date) VALUES
	 ('E0006','Farah Fauchet','0869234234','Jl. Abadi Selamanya No 321 Depok','2021-06-07'),
	 ('E0007','George Graham','0867324376','Kampung Rambutan Gg Senih No 43','2021-07-08'),
	 ('E0008','Herbert Hawk ','0883121345','Kampung Durian Runtuh no 332','2021-08-09'),
	 ('E0009','Iman Indiana','0892543254','Jl. Sari Asih no 92','2021-09-10'),
	 ('E0010','Jack Jones','0812634343','Jl. Kelapa Sambit No 22','2021-10-11'),
	 ('E0011','Key Karuna','0811765433','Jl. Kemang Raya No 38','2021-11-12'),
	 ('E0012','Luna Lawless','0823736357','Jl. Kesadaran No 87','2021-12-13'),
	 ('E0019','Stuart Smith','0885125633','Jl. Angin Mamiri No 82','2021-07-08'),
	 ('E0018','Roy Ryder','0873912362','Jl. Kebahagiaan No 88','2021-06-07'),
	 ('E0020','Theo Taylor','0804241253','Jl. Proklamasi Utara No 32','2021-08-09');
	

create table payment_type (
	id serial primary key,
	payment_name varchar(255)
);

insert into payment_type (payment_name) values ('Cash');
insert into payment_type (payment_name) values ('Debit');

create table trx (
	id serial primary key,
	trx_time timestamp,
	total int,
	payment_type int,
	cashier_id varchar(12)
); 

insert into trx (trx_time, total, payment_type, cashier_id) values (now(), 100000, 1, 'E0003');
insert into trx (trx_time, total, payment_type, cashier_id) values (now(), 230000, 2, 'E0011');
insert into trx (trx_time, total, payment_type, cashier_id) values (now(), 120000, 1, 'E0007');

create table goods (
	id serial primary key,
	good_name varchar(255),
	price int,
	stock int
);

insert into goods (good_name, price, stock) values ('Teh Pucuk', 4000, 2000);
insert into goods (good_name, price, stock) values ('Kopi Kapal Api', 2500, 5000);
insert into goods (good_name, price, stock) values ('Sari Roti Coklat', 8000, 1000);


create table trx_items (
	id serial primary key,
	trx_id int,
	good_id int,
	sell_price int,
	amount int
);

insert into trx_items (trx_id, good_id, sell_price, amount) values (1, 1, 4000, 5);
insert into trx_items (trx_id, good_id, sell_price, amount) values (1, 2, 2500, 10);
insert into trx_items (trx_id, good_id, sell_price, amount) values (1, 1, 8000, 2);

insert into trx_items (trx_id, good_id, sell_price, amount) values (2, 1, 4000, 10);
insert into trx_items (trx_id, good_id, sell_price, amount) values (2, 1, 8000, 4);

insert into trx_items (trx_id, good_id, sell_price, amount) values (3, 2, 2500, 8);
insert into trx_items (trx_id, good_id, sell_price, amount) values (3, 1, 8000, 4);

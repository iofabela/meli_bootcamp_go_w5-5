create table products(
    `id` int not null primary key auto_increment,
    `description` text not null,
    expiration_rate float not null,
    freezing_rate float not null,
    height float not null,
    `length` float not null,
    netweight float not null,
    product_code text not null,
    recommended_freezing_temperature float not null,
    width float not null,
    id_product_type int not null,
    id_seller int not null
);
create table products_types(
    `id` int not null primary key auto_increment,
    `description` text not null
);
create table employees(
    `id` int not null primary key auto_increment,
    card_number_id text not null,
    first_name text not null,
    last_name text not null,
    warehouse_id int not null
);
create table warehouses(
    `id` int not null primary key auto_increment,
    `address` text null,
    telephone text null,
    warehouse_code text null,
    minimum_capacity int null,
    minimum_temperature int null
);

create table sections(
    `id` int not null primary key auto_increment,
    section_number int not null,
    current_temperature int not null,
    minimum_temperature int not null,
    current_capacity int not null,
    minimum_capacity int not null,
    maximum_capacity int not null,
    warehouse_id int not null,
    id_product_type int not null
);
create table sellers(
    `id` int not null primary key auto_increment,
    cid int not null,
    company_name text not null,
    `address` text not null,
    telephone varchar(15) not null
);
create table buyers(
    `id` int not null primary key auto_increment,
    card_number_id text not null,
    first_name text not null,
    last_name text not null
);


CREATE TABLE localities (
	`id` INT PRIMARY KEY AUTO_INCREMENT,
    locality_name VARCHAR(255),
    province_id INT
);

CREATE TABLE order_status (
	`id` INT PRIMARY KEY AUTO_INCREMENT,
    `description` VARCHAR(255)
);

CREATE TABLE product_types (
	`id` INT PRIMARY KEY AUTO_INCREMENT,
    `description` VARCHAR(255)
);

create table countries(
    `id` int not null primary key auto_increment,
    country_name VARCHAR(255)
);

create table provinces(
    `id` int not null primary key auto_increment,
    province_name VARCHAR(255),
    id_country int
);

create table inbound_orders(
    id INT NOT NULL PRIMARY KEY auto_increment,
    order_date DATETIME(6),
    order_number VARCHAR(255),
    employe_id INT,
    product_batch_id INT,
    wareHouse_id INT
);

create table product_batches(
    id INT NOT NULL PRIMARY KEY auto_increment,
    batch_number VARCHAR(255),
    current_quantity INT,
    current_temperature DECIMAL(19,2),
    due_date DATETIME(6),
    initial_quantity INT,
    manufacturing_date DATETIME(6),
    manufacturing_hour DATETIME(6),
    minimum_temperature DECIMAL(19,2),
    product_id INT,
    section_id INT
);

create table users(
    id INT NOT NULL PRIMARY KEY auto_increment,
    password VARCHAR(255),
    username VARCHAR(255)
);

create table user_rol(
    usuario_id INT,
    rol_id INT
);

create table rol(
    id INT NOT NULL PRIMARY KEY auto_increment,
    description VARCHAR(255),
    rol_name VARCHAR(255)
);

alter table warehouses
add locality_id INT;

create table purchase_orders(
    `id` int not null primary key auto_increment,
    order_number varchar(255) not null,
    order_date datetime(6) not null,
    tracking_code varchar(255) not null,
    buyer_id int not null,
    order_status_id int not null,
    wareHouse_id int,
    carrier_id int 
);

create table carries(
    `id` int not null primary key auto_increment,
    cid varchar(255),
    company_name varchar(255),
    `address` varchar(255),
    telephone varchar(255),
    locality_id int
); 

create table product_records(
    `id` int not null primary key auto_increment,
    last_update_date datetime(6),
    purchase_price decimal(19,2),
    sale_price decimal(19,2),
    product_id int
); 

create table order_details(
    `id` int not null primary key auto_increment,
    clean_liness_status varchar(255),
    quantity int,
    temperature decimal(19,2),
    product_record_id int,
    purchase_order_id int
); 

alter table sellers
add locality_id INT;

/* DATA */

insert into buyers (id, card_number_id, first_name, last_name) values (1, '51442-543', 'Hercule', 'Gouldeby');
insert into buyers (id, card_number_id, first_name, last_name) values (2, '0228-2077', 'Kale', 'Worge');
insert into buyers (id, card_number_id, first_name, last_name) values (3, '31722-207', 'Winfield', 'Maxfield');
insert into buyers (id, card_number_id, first_name, last_name) values (4, '52164-1106', 'Delly', 'Yearns');
insert into buyers (id, card_number_id, first_name, last_name) values (5, '65437-035', 'Alyss', 'Van Brug');
insert into warehouses (id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature) values (1, '2985 Lunder Center', '(161) 7030736', '0338-0703', 1, 1);
insert into warehouses (id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature) values (2, '5 Myrtle Hill', '(601) 5899450', '11673-170', 2, 2);
insert into warehouses (id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature) values (3, '1 Morning Center', '(615) 9659486', '63777-223', 3, 3);
insert into warehouses (id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature) values (4, '40 Clyde Gallagher Plaza', '(862) 9958364', '67046-590', 4, 4);
insert into warehouses (id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature) values (5, '4002 Ridgeview Alley', '(699) 5099808', '54868-5345', 5, 5);
insert into sellers (id, cid, company_name, address, telephone) values (1, 1, 'Skyba', '3180 Roxbury Drive', '(618) 2011607');
insert into sellers (id, cid, company_name, address, telephone) values (2, 2, 'Latz', '6286 Moulton Parkway', '(387) 6865821');
insert into sellers (id, cid, company_name, address, telephone) values (3, 3, 'Wikizz', '9 Marquette Drive', '(129) 4018633');
insert into sellers (id, cid, company_name, address, telephone) values (4, 4, 'Meedoo', '91 Briar Crest Road', '(547) 2313975');
insert into sellers (id, cid, company_name, address, telephone) values (5, 5, 'Skidoo', '6924 Roxbury Park', '(558) 9207455');
insert into carries (id, cid, company_name, address, telephone, locality_id) values (1, 1, 'Trudoo', '94 Eastwood Way', '(981) 2938974', 1);
insert into carries (id, cid, company_name, address, telephone, locality_id) values (2, 2, 'Divape', '56 Red Cloud Terrace', '(331) 4585300', 2);
insert into carries (id, cid, company_name, address, telephone, locality_id) values (3, 3, 'Livepath', '101 Arrowood Place', '(890) 1282013', 3);
insert into carries (id, cid, company_name, address, telephone, locality_id) values (4, 4, 'Tavu', '6124 West Trail', '(550) 7194074', 4);
insert into carries (id, cid, company_name, address, telephone, locality_id) values (5, 5, 'Tekfly', '10 Commercial Park', '(445) 9818922', 5);
insert into localities (id, locality_name, province_id) values (1, 'Quigley, Bauch and Willms', 1);
insert into localities (id, locality_name, province_id) values (2, 'Von, Schmeler and Hyatt', 2);
insert into localities (id, locality_name, province_id) values (3, 'Johns-Abshire', 3);
insert into localities (id, locality_name, province_id) values (4, 'Bernhard Inc', 4);
insert into localities (id, locality_name, province_id) values (5, 'Gutkowski, Sipes and Rowe', 5);
insert into sections (id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, id_product_type) values (1, 1, 1, 1, 1, 1, 1, 1, 1);
insert into sections (id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, id_product_type) values (2, 2, 2, 2, 2, 2, 2, 2, 2);
insert into sections (id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, id_product_type) values (3, 3, 3, 3, 3, 3, 3, 3, 3);
insert into sections (id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, id_product_type) values (4, 4, 4, 4, 4, 4, 4, 4, 4);
insert into sections (id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, id_product_type) values (5, 5, 5, 5, 5, 5, 5, 5, 5);
insert into employees (id, card_number_id, first_name, last_name, warehouse_id) values (1, 1, 'Mattie', 'Smallpeice', 1);
insert into employees (id, card_number_id, first_name, last_name, warehouse_id) values (2, 2, 'Kary', 'Gavrielli', 2);
insert into employees (id, card_number_id, first_name, last_name, warehouse_id) values (3, 3, 'Kerwinn', 'Woller', 3);
insert into employees (id, card_number_id, first_name, last_name, warehouse_id) values (4, 4, 'Putnem', 'Pheazey', 4);
insert into employees (id, card_number_id, first_name, last_name, warehouse_id) values (5, 5, 'Tamas', 'Piletic', 5);
insert into products (id, description, expiration_rate, freezing_rate, height, length, netweight, product_code, recommended_freezing_temperature, width, id_product_type, id_seller) values (1, 'pretium iaculis diam erat', 22, 4, 88, 71, 80, '0536-3587', 77, 42, 1, 1);
insert into products (id, description, expiration_rate, freezing_rate, height, length, netweight, product_code, recommended_freezing_temperature, width, id_product_type, id_seller) values (2, 'pede morbi porttitor lorem id ligula', 25, 20, 69, 73, 77, '67046-089', 85, 82, 2, 2);
insert into products (id, description, expiration_rate, freezing_rate, height, length, netweight, product_code, recommended_freezing_temperature, width, id_product_type, id_seller) values (3, 'pede venenatis non sodales', 37, 88, 56, 70, 67, '63323-270', 41, 69, 3, 3);
insert into products (id, description, expiration_rate, freezing_rate, height, length, netweight, product_code, recommended_freezing_temperature, width, id_product_type, id_seller) values (4, 'turpis adipiscing lorem vitae mattis', 68, 25, 70, 51, 89, '0338-0552', 85, 60, 4, 4);
insert into products (id, description, expiration_rate, freezing_rate, height, length, netweight, product_code, recommended_freezing_temperature, width, id_product_type, id_seller) values (5, 'donec ut mauris eget', 12, 45, 97, 29, 64, '41268-029', 95, 19, 5, 5);
insert into products_types (id, description) values (1, 'consectetuer eget rutrum at lorem');
insert into products_types (id, description) values (2, 'vulputate ut ultrices');
insert into products_types (id, description) values (3, 'pellentesque eget nunc donec');
insert into products_types (id, description) values (4, 'vel lectus in quam');
insert into products_types (id, description) values (5, 'justo maecenas rhoncus aliquam lacus');
insert into order_status (id, description) values (1, 'mi nulla ac enim in tempor');
insert into order_status (id, description) values (2, 'blandit nam nulla integer pede');
insert into order_status (id, description) values (3, 'ipsum dolor sit');
insert into order_status (id, description) values (4, 'ligula in lacus curabitur at ipsum');
insert into order_status (id, description) values (5, 'sit amet justo morbi');
insert into countries (id, country_name) values (1, 'Greece');
insert into countries (id, country_name) values (2, 'Greece');
insert into countries (id, country_name) values (3, 'Burkina Faso');
insert into countries (id, country_name) values (4, 'China');
insert into countries (id, country_name) values (5, 'Venezuela');
insert into provinces (id, province_name, id_country) values (1, 'Frederiksberg', 1);
insert into provinces (id, province_name, id_country) values (2, 'Shuangta', 2);
insert into provinces (id, province_name, id_country) values (3, 'Quibdó', 3);
insert into provinces (id, province_name, id_country) values (4, 'Nantes', 4);
insert into provinces (id, province_name, id_country) values (5, 'Xiaosong', 5);
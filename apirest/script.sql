create database prueba_db;

create table persona(
    id serial primary key,
    nombre varchar(80) not null,
    apellido varchar(80) not null,
    direccion varchar(80),
    telefono varchar(80)
);
insert into persona(nombre,apellido,direccion,telefono)
values
('ESTEFANIA', 'AROCAS PASADAS','PADR╙ , 109','938205580'),
('QUERALT' , ' VISO GILABERT',' CASA CORDELLAS , ','936545115'),
('JOAN'   ,  'AYALA FERRERAS','DOCTOR FLEMING , 11','938202768'),
('JOAN'  ,   'BAEZ TEJADO','BERTRAND I SERRA , 11, 3R.','938727844'),
('MARC'  ,   'BASTARDES SOTO','CARRI╙ , 12, 5╚ A','938350521'),
('JOSEP'  ,  ' ANGUERA VILAFRANC','PIRINEUS , 10','938755645');
CREATE TABLE IF NOT EXISTS gallery
(
    id bigint not null,
    title text not null
);
alter table gallery
    add constraint gallery_pk
        primary key (id);
alter table gallery
    alter column id add generated by default as identity;
insert into gallery (id,title) values (default,'Догляд за садом');
insert into gallery (id,title) values (default,'Дерева');
insert into gallery (id,title) values (default,'Плоди');
insert into gallery (id,title) values (default,'Цвітіння');
insert into gallery (id,title) values (default,'Висадка');

CREATE TABLE IF NOT EXISTS gallery_images
(
    id bigint not null,
    gallery_id bigint not null,
    file text
);
alter table gallery_images
    add constraint gallery_images_pk
        primary key (id);
alter table gallery_images
    add constraint gallery_images_gallery_null_fk
        foreign key (gallery_id) references gallery (id);
alter table gallery_images
    alter column id add generated by default as identity;
insert into gallery_images (id,gallery_id,file) values (default,1,'https://live.staticflickr.com/65535/52119219176_ee0d3a25cd_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,1,'https://live.staticflickr.com/65535/52119467429_ddf7dca8d7_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,1,'https://live.staticflickr.com/65535/52119718225_fe6d0af38c_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,1,'https://live.staticflickr.com/65535/52119218561_f58509bd8e_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,1,'https://live.staticflickr.com/65535/52119218421_785674a142_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,1,'https://live.staticflickr.com/65535/52119466659_1f5264d390_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,1,'https://live.staticflickr.com/65535/52119244808_344926f356_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,1,'https://live.staticflickr.com/65535/52119466454_7b60c3a7f1_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,2,'https://live.staticflickr.com/65535/52119244478_f166933a6f_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,2,'https://live.staticflickr.com/65535/52119717210_cac01188d2_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,2,'https://live.staticflickr.com/65535/52119717010_6a3061a366_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,2,'https://live.staticflickr.com/65535/52119243968_b17a41cb32_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,2,'https://live.staticflickr.com/65535/52119465574_756826cb22_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,2,'https://live.staticflickr.com/65535/52119216981_cda40f8122_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,2,'https://live.staticflickr.com/65535/52119716240_d0277b0f6a_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,2,'https://live.staticflickr.com/65535/52119716120_019a8d0232_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,2,'https://live.staticflickr.com/65535/52119715970_314ed0e3fd_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,2,'https://live.staticflickr.com/65535/52119715870_ddbbf3ab03_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,2,'https://live.staticflickr.com/65535/52119216446_48d6ff9f48_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,2,'https://live.staticflickr.com/65535/52119242848_47a7b2bd6b_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,2,'https://live.staticflickr.com/65535/52118185497_69089db1f7_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,2,'https://live.staticflickr.com/65535/52119241188_5e984e8ea1_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,2,'https://live.staticflickr.com/65535/52119714015_7bb41ba174_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,2,'https://live.staticflickr.com/65535/52119462929_0d7328474e_z.jpg');
insert into gallery_images (id,gallery_id,file) values (default,2,'https://live.staticflickr.com/65535/52119241148_d1ba7a33a1_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,2,'https://live.staticflickr.com/65535/52119214571_8f95175cb7_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,2,'https://live.staticflickr.com/65535/52118183912_b0978b2c0b_z.jpg');
insert into gallery_images (id,gallery_id,file) values (default,2,'https://live.staticflickr.com/65535/52118183887_348d57a239_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,2,'https://live.staticflickr.com/65535/52119214506_91b9fa75fd_z.jpg');
insert into gallery_images (id,gallery_id,file) values (default,2,'https://live.staticflickr.com/65535/52119462839_3e9f43a10d_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,3,'https://live.staticflickr.com/65535/52119715450_1f06269028_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,3,'https://live.staticflickr.com/65535/52119715205_16e97cd87b_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,3,'https://live.staticflickr.com/65535/52118183817_af714d66d5_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,3,'https://live.staticflickr.com/65535/52119462634_836b04042e_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,3,'https://live.staticflickr.com/65535/52119240738_63b24fe16b_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,3,'https://live.staticflickr.com/65535/52119240653_e2157b45da_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,3,'https://live.staticflickr.com/65535/52119214096_f385634e7b_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,3,'https://live.staticflickr.com/65535/52119462419_c876577dba_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,3,'https://live.staticflickr.com/65535/52119713355_dc89a096d0_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,3,'https://live.staticflickr.com/65535/52119240273_44dbca7776_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,3,'https://live.staticflickr.com/65535/52119713175_1529990278_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,3,'https://live.staticflickr.com/65535/52119713115_cb09fcb9ca_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,3,'https://live.staticflickr.com/65535/52119239993_f29b47b9e3_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,3,'https://live.staticflickr.com/65535/52119712945_c131703abe_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,3,'https://live.staticflickr.com/65535/52119712905_060b8927c7_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,4,'https://live.staticflickr.com/65535/52119213256_d2cd1c4feb_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,4,'https://live.staticflickr.com/65535/52119213156_9aa70bc15f_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,4,'https://live.staticflickr.com/65535/52119712585_65a44c84e9_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,4,'https://live.staticflickr.com/65535/52119712585_65a44c84e9_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,4,'https://live.staticflickr.com/65535/52119712530_d3024f298f_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,4,'https://live.staticflickr.com/65535/52119461469_55d4361924_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,4,'https://live.staticflickr.com/65535/52119461379_1402ac0c84_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,4,'https://live.staticflickr.com/65535/52119461324_88a7e53ee5_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,5,'https://live.staticflickr.com/65535/52119463799_a6045356fa_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,5,'https://live.staticflickr.com/65535/52119463449_40def873c3_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,5,'https://live.staticflickr.com/65535/52119714450_3e12a36603_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,5,'https://live.staticflickr.com/65535/52119714325_e96a135b1f_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,5,'https://live.staticflickr.com/65535/52119714190_bd54bf90c5_c.jpg');
insert into gallery_images (id,gallery_id,file) values (default,5,'https://live.staticflickr.com/65535/52119239253_62d0705cd6_c.jpg');
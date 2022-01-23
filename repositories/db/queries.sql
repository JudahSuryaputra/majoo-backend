-- SQLINES LICENSE FOR EVALUATION USE ONLY
CREATE SEQUENCE Users_seq;

CREATE TABLE Users (
  id bigint NOT NULL DEFAULT NEXTVAL ('Users_seq'),
  name varchar(45) DEFAULT NULL,
  user_name varchar(45) DEFAULT NULL,
  password varchar(225) DEFAULT NULL,
  created_at timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_by bigint NOT NULL,
  updated_at timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_by bigint NOT NULL,
  PRIMARY KEY (id)
)  ;

ALTER SEQUENCE Users_seq RESTART WITH 1;



insert into Users values (1, 'Admin 1', 'admin1', MD5('admin1'), now(), 1, now(),1), (2, 'Admin 2', 'admin2', MD5('admin2'), now(), 2, now(),2);

-- SQLINES LICENSE FOR EVALUATION USE ONLY
CREATE SEQUENCE Merchants_seq;

CREATE TABLE Merchants (
	id bigint NOT NULL DEFAULT NEXTVAL ('Merchants_seq'),
	user_id int NOT NULL,
	merchant_name varchar(40) NOT NULL,
	created_at timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
	created_by bigint NOT NULL,
	updated_at timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_by bigint NOT NULL,
	PRIMARY KEY (id)
  )  ;

ALTER SEQUENCE Merchants_seq RESTART WITH 1;

  insert into Merchants values (1, 1, 'merchant 1', now(), 1, now(),1), (2, 2, 'Merchant 2', now(), 2, now(),2);


  -- SQLINES LICENSE FOR EVALUATION USE ONLY
  CREATE SEQUENCE Outlets_seq;

  CREATE TABLE Outlets (
	id bigint NOT NULL DEFAULT NEXTVAL ('Outlets_seq'),
	merchant_id bigint NOT NULL,
	outlet_name varchar(40) NOT NULL,
	created_at timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
	created_by bigint NOT NULL,
	updated_at timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_by bigint NOT NULL,
	PRIMARY KEY (id)
  )  ;

ALTER SEQUENCE Outlets_seq RESTART WITH 1;
  insert into Outlets values (1, 1, 'Outlet 1', now(), 1, now(),1), (2, 2, 'Outlet 1', now(), 2, now(),2), (3, 1, 'Outlet 2', now(), 1, now(),1);


  -- SQLINES LICENSE FOR EVALUATION USE ONLY
  CREATE SEQUENCE Transactions_seq;

  CREATE TABLE Transactions (
	id bigint NOT NULL DEFAULT NEXTVAL ('Transactions_seq'),
	merchant_id bigint NOT NULL,
	outlet_id bigint NOT NULL,
	bill_total double precision NOT NULL,
	created_at timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
	created_by bigint NOT NULL,
	updated_at timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_by bigint NOT NULL,
	PRIMARY KEY (id)
  )  ;

ALTER SEQUENCE Transactions_seq RESTART WITH 1;
  insert into Transactions values 
  (1, 1, 1, 2000, '2021-11-01 12:30:04', 1, '2021-11-01 12:30:04',1), 
  (2, 1, 1, 2500, '2021-11-01 17:20:14', 1, '2021-11-01 17:20:14',1),
  (3, 1, 1, 4000, '2021-11-02 12:30:04', 1, '2021-11-02 12:30:04',1),
  (4, 1, 1, 1000, '2021-11-04 12:30:04', 1, '2021-11-04 12:30:04',1),
  (5, 1, 1, 7000, '2021-11-05 16:59:30', 1, '2021-11-05 16:59:30',1),
  (6, 1, 3, 2000, '2021-11-02 18:30:04', 1, '2021-11-02 18:30:04',1), 
  (7, 1, 3, 2500, '2021-11-03 17:20:14', 1, '2021-11-03 17:20:14',1),
  (8, 1, 3, 4000, '2021-11-04 12:30:04', 1, '2021-11-04 12:30:04',1),
  (9, 1, 3, 1000, '2021-11-04 12:31:04', 1, '2021-11-04 12:31:04',1),
  (10, 1, 3, 7000, '2021-11-05 16:59:30', 1, '2021-11-05 16:59:30',1),
  (11, 2, 2, 2000, '2021-11-01 18:30:04', 2, '2021-11-01 18:30:04',2), 
  (12, 2, 2, 2500, '2021-11-02 17:20:14', 2, '2021-11-02 17:20:14',2),
  (13, 2, 2, 4000, '2021-11-03 12:30:04', 2, '2021-11-03 12:30:04',2),
  (14, 2, 2, 1000, '2021-11-04 12:31:04', 2, '2021-11-04 12:31:04',2),
  (15, 2, 2, 7000, '2021-11-05 16:59:30', 2, '2021-11-05 16:59:30',2),
  (16, 2, 2, 2000, '2021-11-05 18:30:04', 2, '2021-11-05 18:30:04',2), 
  (17, 2, 2, 2500, '2021-11-06 17:20:14', 2, '2021-11-06 17:20:14',2),
  (18, 2, 2, 4000, '2021-11-07 12:30:04', 2, '2021-11-07 12:30:04',2),
  (19, 2, 2, 1000, '2021-11-08 12:31:04', 2, '2021-11-08 12:31:04',2),
  (20, 2, 2, 7000, '2021-11-09 16:59:30', 2, '2021-11-09 16:59:30',2),
  (21, 2, 2, 1000, '2021-11-10 12:31:04', 2, '2021-11-10 12:31:04',2),
  (22, 2, 2, 7000, '2021-11-11 16:59:30', 2, '2021-11-11 16:59:30',2);
 
 CREATE SEQUENCE access_tokens_id_seq;
CREATE TABLE IF NOT EXISTS access_tokens
(
  id BIGINT NOT NULL default nextval('access_tokens_id_seq') PRIMARY KEY,
  user_name TEXT NOT NULL,
  token TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
create index access_token_email_idx on access_tokens(email);
create index access_token_access_token_idx on access_tokens(token);

select Date(t.created_at) as date, sum(t.bill_total) as omzet
from transactions t
where merchant_id = 1 and created_at >= '2021-11-01' and created_at < '2021-12-01'
group by date
order by date





CREATE TABLE orders (
  id varchar(60) NOT NULL,
  price decimal(10,2) NOT NULL,
  tax decimal(10,2) NOT NULL,
  final_price decimal(10,2) NOT NULL,
  PRIMARY KEY (id)
);
CREATE TABLE flats (
	id				integer CONSTRAINT firstkey PRIMARY KEY,
	name 			varchar(150) NOT NULL,
	price 			integer NOT NULL,
	rooms 			integer NOT NULL,
	size 			integer NOT NULL,
	store 			integer NOT NULL,
	elevator 		boolean NOT NULL,
	link 			varchar NOT NULL,
	area 			varchar(50) NOT NULL,
	date_published	date NOT NULL
);
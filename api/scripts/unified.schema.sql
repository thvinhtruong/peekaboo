create database if not exists pe;

use pe;

create table users (
	id INT NOT NULL AUTO_INCREMENT,
	full_name VARCHAR(255) NOT NULL,
	username VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	gender VARCHAR(255) NOT NULL,
	address VARCHAR(255),
	mail VARCHAR(255),
	phone VARCHAR(255),
	dob DATETIME,
	qualification VARCHAR(255),
	entity_code INT NOT NULL,
	active INT NOT NULL,
	datecreated DATETIME NOT NULL,
	dateupdated DATETIME,
	UNIQUE (username),
	PRIMARY KEY (id),
	CONSTRAINT valid_role_check CHECK (entity_code IN (1, 2, 3, 4))
);

create table classes (
	id INT NOT NULL AUTO_INCREMENT,
	class_name VARCHAR(255) NOT NULL,
	info TEXT NOT NULL,
	announcement TEXT NOT NULL,
	room_code VARCHAR(255) NOT NULL,
	level VARCHAR(255) NOT NULL,
	active TINYINT NOT NULL,
	datecreated DATETIME NOT NULL,
	dateupdated DATETIME,
	primary key (id),
	UNIQUE(class_name)
);

create table user_class (
	uid INT NOT NULL,
	cid INT NOT NULL,
	active TINYINT NOT NULL,
	FOREIGN KEY (cid) REFERENCES classes (id),
	FOREIGN KEY (uid) REFERENCES users (id),
	PRIMARY KEY (uid, cid)
);

create table tags (
	id INT NOT NULL AUTO_INCREMENT,
	tag VARCHAR(255) NOT NULL,
	info TEXT NOT NULL,
	active TINYINT NOT NULL,
	datecreated DATETIME NOT NULL,
	dateupdated DATETIME,
	PRIMARY KEY (id),
	UNIQUE(tag)
);

create table testbank (
	id INT NOT NULL AUTO_INCREMENT,
	tag_id INT NOT NULL,
	test_name VARCHAR(255) NOT NULL,
	created_user_id INT NOT NULL,
	target_entity_code INT NOT NULL,
	title VARCHAR(255) NOT NULL,
	info TEXT NOT NULL,
	duration INT,
	dateassigned DATETIME NOT NULL,
	deadline DATETIME,
	active TINYINT,
	datecreated DATETIME NOT NULL,
	dateupdated DATETIME,
	PRIMARY KEY (id),
	UNIQUE(test_name),
	FOREIGN KEY (tag_id) REFERENCES tags (id),
	FOREIGN KEY (created_user_id) REFERENCES users (id),
	CONSTRAINT valid_role_check_testbank CHECK (target_entity_code IN (1, 2, 3, 4))
);

create table test_class (
	id INT NOT NULL AUTO_INCREMENT,
	cid INT NOT NULL,
	tid INT NOT NULL,
	primary key (id),
	UNIQUE(cid, tid),
	FOREIGN KEY (cid) REFERENCES classes (id),
	FOREIGN KEY (tid) REFERENCES testbank (id)
);

create table testresults (
	id INT NOT NULL AUTO_INCREMENT,
	test_class_id INT NOT NULL,
	user_id INT NOT NULL,
	entity_code INT NOT NULL,
	datecreated DATETIME NOT NULL,
	score FLOAT NOT NULL,
	comment varchar(255),
	resultnote varchar(255),
	active TINYINT NOT NULL,
	dateupdated DATETIME,
	PRIMARY KEY (id),
	FOREIGN KEY (test_class_id) REFERENCES test_class (id),
	FOREIGN KEY (user_id) REFERENCES users (id),
	UNIQUE KEY (test_class_id, user_id, datecreated)
);

create table skilltests (
	id INT NOT NULL AUTO_INCREMENT,
	media_url VARCHAR(200),
	title VARCHAR(200) NOT NULL,
	content MEDIUMTEXT NOT NULL,
	description VARCHAR(200) NOT NULL,
	section MEDIUMTEXT NOT NULL,
	type VARCHAR(30) NOT NULL,
	datecreated DATETIME NOT NULL,
	dateupdated DATETIME,
	PRIMARY KEY (id),
	UNIQUE(title)
);

create table test_answer (
	id INT NOT NULL PRIMARY KEY,
	section_answer MEDIUMTEXT NOT NULL,
	UNIQUE KEY (id),
	FOREIGN KEY (id) REFERENCES testresults (id)
);

create table skilltest_test (
	tid INT NOT NULL AUTO_INCREMENT,
	stid INT NOT NULL,
	PRIMARY KEY (tid, stid),
	FOREIGN KEY (tid) REFERENCES testbank (id),
	FOREIGN KEY (stid) REFERENCES skilltests (id)
);

INSERT INTO
	users(
		full_name,
		username,
		password,
		gender,
		mail,
		entity_code,
		datecreated,
		active
	)
VALUES
	(
		"GiaNghi",
		"gianghi",
		"$2a$12$rSmHGKGlpiPrPA.oshhyB.k6seTEpl378iS/kEJFksqYWS1gRYSQG",
		"female",
		"gianghi@gmail.com",
		3,
		"2022-06-25",
		1
	);

INSERT INTO
	classes(
		class_name,
		info,
		announcement,
		room_code,
		level,
		datecreated,
		active
	)
VALUES
	(
		"Realtime",
		"sample class",
		"hello world",
		"1.2",
		"level max",
		"2022-06-25",
		1
	);

INSERT INTO
	user_class(uid, cid, active)
VALUES
	(1, 1, 1);

INSERT INTO
	tags(tag, info, active, datecreated)
VALUES
	("reading", "reading", 1, "2022-06-25");

INSERT INTO
	testbank(
		tag_id,
		test_name,
		created_user_id,
		target_entity_code,
		title,
		info,
		duration,
		dateassigned,
		deadline,
		active,
		datecreated
	)
VALUES
	(
		1,
		"test1",
		1,
		3,
		"test1",
		"test1",
		60,
		"2022-06-25",
		"2022-06-25",
		1,
		"2022-06-25"
	);

INSERT INTO
	test_class(cid, tid)
VALUES
	(1, 1);

INSERT INTO
	skilltests(
		media_url,
		title,
		content,
		description,
		section,
		type,
		datecreated
	)
VALUES
	(
		"https://i.imgur.com/4XOW6u7.jpeg",
		"Just not gonna ever!",
		"title: Hello world",
		"description: Hello world",
		"section: update this",
		"reading",
		"2022-06-25"
	);

INSERT INTO
	skilltest_test (tid, stid)
VALUES
	(1, 1);
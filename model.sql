CREATE DATABASE IF NOT EXISTS bookstore;
USE bookstore;

CREATE TABLE  IF NOT EXISTS accounts(
  id char(32),
  username varchar(30) NOT NULL,
  password varchar(32) NOT NULL,
  nickname varchar(30)  NOT NULL,
  email varchar(30) NOT NULL,
  role INT default 0,
  createat DATETIME,
  updatetime DATETIME,
  description varchar(50),
  phone varchar(15),
  header varchar(50),
  PRIMARY KEY(id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS document(
  id char(32),
  bookname varchar(30) NOT NULL,
  sectioncount INT default 0,
  owner char(32) NOT NULL ,
  readcount INT default 0,
  favoritescount INT default 0,
  description varchar(100),
  score INT default 4,
  url varchar(200),
  privately_owned INT,
  identify varchar(20),
  picture varchar(50),
  createat DATETIME,
  last_modify_text DATETIME,
  path varchar(50),
  PRIMARY KEY (id),
  FOREIGN KEY (owner) REFERENCES accounts(id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS label(
  id char(32),
  name char(32) NOT NULL,
  PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

#favlurites
CREATE TABLE IF NOT EXISTS label_document(
  labelid char(32),
  documentid char(32),
  FOREIGN KEY (labelid) REFERENCES label(id),
  FOREIGN KEY (documentid) REFERENCES document(id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

#comments
CREATE TABLE IF NOT EXISTS comments(
  id char(32),
  data varchar(100) NOT NULL,
  ownerid char(32),
  createat DATETIME NOT NULL ,
  PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS document_comments(
  documentid char(32),
  commentid char(32),
  FOREIGN KEY (commentid) REFERENCES comments(id),
  FOREIGN KEY (documentid) REFERENCES document(id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;








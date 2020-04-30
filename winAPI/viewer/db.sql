PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;

-- 注释测试

CREATE TABLE Type (
 ID integer not null,
 Name text not null unique primary key
);

CREATE TABLE Function (
 ID integer not null,
 Name text not null primary key,
 TypeID integer not null check(typeof(TypeID)='integer') references Type(ID) on update cascade on delete cascade,
 IsFunction integer not null check(IsFunction=1 or IsFunction=0) ,
 FullName text not null,
 Explain text,
 Return text,
 Parameter text,
 Remark text,
 Libraries text,
 Example text
);

CREATE TABLE Constant (
 ID integer not null,
 Name text not null primary key,
 Value text not null
);

CREATE TABLE Struct (
 ID integer not null primary key,
 Name text not null primary key,
 FullName text not null
);

COMMIT;
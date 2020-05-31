PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;

-- 注释测试

-- 函数分类（自己定义）
CREATE TABLE Type (
	ID integer not null,
	Name text not null primary key
);

-- 函数
CREATE TABLE Func (
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

-- Func参数
CREATE TABLE FuncParameter (
	ID integer not null primary key,
	FuncID integer not null check(typeof(FuncID)='integer') references Func(ID) on update cascade on delete cascade,
	Name text not null,
	-- 所有类型都放在Struct里，一些内置的也可以那么处理，真正使用的时候进行转换
	StructID integer not null check(typeof(StructID)='integer') references Struct(ID) on update cascade on delete cascade
);

-- 常量
CREATE TABLE Const (
	ID integer not null,
	Name text not null primary key,
	Value text not null
);

-- 结构体成员
CREATE TABLE StructField (
	ID integer not null,
	StructID integer not null check(typeof(StructID)='integer') references Struct(ID) on update cascade on delete cascade,
	SortNo integer not null check(typeof(SortNo)='integer'),
	Name text not null primary key,
	-- 这里要注意不要循环引用了！
	DataType integer not null check(typeof(DataType)='integer') references Struct(ID) on update cascade on delete cascade,
	-- 说明介绍
	Explain text,
	primary key(StructID, Name, DataType)
);

-- 结构体，普通的int、long等都可以认为是一种结构体
-- 具体的成员信息到StructField里面去查找
CREATE TABLE Struct (
	ID integer not null,
	Name text not null primary key,
	-- 结构体的说明介绍
	Explain text
);

-- 指明Table Rel与FuncID对应的是什么（函数、常量、结构体）
CREATE TABLE RelType (
	ID integer not null primary key,
	-- Func、Const、Struct
	Name text not null primary key
);

-- 记录函数与哪些函数、常量、结构体有关系，主要是使用某个函数的时候，可能会同时使用到的那些
CREATE TABLE Rel (
	ID integer not null primary key,
	-- 函数ID
	FuncID integer not null check(typeof(FuncID)='integer') references Func(ID) on update cascade on delete cascade,
	-- 指明与FuncID对应的是什么（函数、常量、结构体）
	RelTypeID integer not null check(typeof(RelTypeID)='integer') references RelType(ID) on update cascade on delete cascade,
	-- 在其他Table的ID
	pID integer not null check(typeof(pID)='integer')
);

COMMIT;
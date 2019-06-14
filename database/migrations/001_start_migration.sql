create table Person (
    ID        integer      not null primary key autoincrement,
    Username  varchar(255) not null unique,
    FirstName varchar(255),
    LastName  varchar(255)
);
create index idx_person on Person(ID,Username);

create table Room(
    ID   integer      not null primary key autoincrement,
    Name varchar(255) not null
);
create index idx_room on Room(ID,Name);

create table Message (
    ID      integer      not null primary key autoincrement,
    Message varchar(255) not null,
    IsRead  bit          default 0,
    RoomID  int,
    FromID  int,
    ToID    int,
    foreign key(RoomID) references Room(ID),
    foreign key(FromID) references Person(ID),
    foreign key(ToID)   references Person(ID)
);
create index idx_message on Message(ID,Message,RoomID,FromID,ToID);

create table RoomPerson (
    ID       integer not null primary key autoincrement,
    PersonID int,
    RoomID   int,
    foreign key(PersonID) references Person(ID),
    foreign key(RoomID) references Room(ID)
);
create index idx_room_person on RoomPerson(ID,PersonID,RoomID);

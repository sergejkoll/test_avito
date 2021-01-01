CREATE TABLE if not exists room (
    id serial primary key,
    description varchar(255),
    price integer,
    date_added integer
);

CREATE TABLE if not exists booking (
    booking_id serial primary key,
    room_id integer references room(id) on delete cascade,
    date_start date,
    date_end date
);
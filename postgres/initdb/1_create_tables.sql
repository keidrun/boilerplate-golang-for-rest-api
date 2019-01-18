create table users
(
  id serial primary key,
  email text not null UNIQUE,
  password text not null,
  name text not null,
  age integer,
  gender text,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp
);

create function set_update_time() returns opaque as '
  begin
    new.updated_at := ''now'';
    return new;
  end;
' language 'plpgsql';

create trigger update_trigger
  before update on users
  for each row execute procedure set_update_time();

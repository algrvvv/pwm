create table if not exists notes
(
  id integer primary key autoincrement,
  name text unique not null,
  value text not null,
  created_at timestamp default current_timestamp not null
)

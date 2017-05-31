create table matches (
    id serial primary key,
    home_team varchar(200) not null,
    away_team varchar(200) not null,
    home_user_id integer not null,
    away_user_id integer not null,
    home_points integer not null,
    away_points integer not null
);

create table users (
    id serial primary key,
    email varchar(200) not null unique,
    password char(64) not null,
    bio text default '',
    salt char(32) not null,
    first_name varchar(100) not null,
    last_name varchar(100) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

create table users_matches (
    user_id integer not null,
    match_id integer not null
);

alter table matches add constraint matches_home_user_id_fk foreign key (home_user_id) references users(id);
alter table matches add constraint matches_away_user_id_fk foreign key (away_user_id) references users(id);
alter table users_matches add constraint users_matches_user_id_fk foreign key (user_id) references users(id);
alter table users_matches add constraint users_matches_match_id_fk foreign key (match_id) references matches(id);

create index user_email_idx on users(email);
create index matches_away_user_id_idx on matches(away_user_id);
create index matches_home_user_id_idx on matches(home_user_id);
create index users_matches_user_id_idx on users_matches(user_id);
create index users_matches_match_id_idx on users_matches(match_id);
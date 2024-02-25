
-- drop  table cars
-- drop  table users
CREATE TABLE users (
  id SERIAL PRIMARY key NOT null,
  user_name varchar(100) NOT NULL,
  password varchar(100) DEFAULT '',
  email varchar(100) DEFAULT '',
  auth_token varchar(1000) DEFAULT '',
  status varchar(20) DEFAULT 'ACTIVE',
  created timestamp default CURRENT_TIMESTAMP,
  updated timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tigers (
   id SERIAL PRIMARY key NOT null,
   name varchar(100) NOT NULL,
   dob DATE NOT NULL,
   last_seen timestamp NOT NULL,
   last_seen_location POINT NOT NULL,
   status varchar(20) DEFAULT 'ACTIVE',
   created timestamp DEFAULT CURRENT_TIMESTAMP,
   updated timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tiger_sightings (
   id SERIAL PRIMARY key NOT null,
   user_id INT REFERENCES users(id),
   tiger_id INT REFERENCES tigers(id),
   loc POINT NOT NULL,
   sight_time timestamp NOT NULL,
   image_url varchar(100) NOT NULL,
   status varchar(20) DEFAULT 'ACTIVE',
   created timestamp DEFAULT CURRENT_TIMESTAMP,
   updated timestamp DEFAULT CURRENT_TIMESTAMP
);

select loc[0],loc[1]  from tiger_sightings  ORDER BY sight_time  desc  limit 1;

insert
	into
	users (user_name,
	password,
	email)
values('Pankaj',
'Banshiwal',
'pankaj92banshiwal')

select * from users user_name = 'Pankaj92Banshiwal'


select id,name,dob,last_seen ,status , last_seen_location[0],last_seen_location[1]  from tigers ORDER BY last_seen  desc  limit 10 offset 0;

select id,user_id,tiger_id,loc[0],loc[1],sight_time,image_url ,status  from tiger_sightings where tiger_id = 2 ORDER BY sight_time  desc  limit 10 offset 0;

select
	u.id,
	u.user_name,
	u.email
from
	tiger_sightings ts
inner join users u 
ON(u.id = ts.user_id)
where
	ts.tiger_id = 2
	and u.status = 'ACTIVE'
	and ts.status = 'ACTIVE'
group by u.id

select
	u.id,
	u.user_name,
	u.email,
	(select name from tigers where id = 2)
	
from
	tiger_sightings as ts
inner join users u on
	(u.id = ts.user_id)
where
	ts.tiger_id = 2
	and u.status = 'ACTIVE'
	and ts.status = 'ACTIVE'
group by
	u.id



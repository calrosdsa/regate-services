create table if not exists conversations(
    conversation_id serial primary key,
    establecimiento_id int not null,
    profile_id int not null
);


create table if not exists conversation_message(
  id serial primary key,
  conversation_id int not null,
  reply_to int,
  sender_id int,
  created_at TIMESTAMP DEFAULT current_timestamp,
  content TEXT NOT NULL,
  CONSTRAINT fk_conversation
  FOREIGN KEY(conversation_id) 
  REFERENCES conversations(conversation_id)  on delete cascade
);

insert into conversation_message(conversation_id,profile_id,content)values(1,1,'First Message');
insert into conversations(profile_id,establecimiento_id) values (1,1);
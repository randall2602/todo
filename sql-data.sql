use todo;

insert into users (id) values ("randallcsturgis@gmail.com");
insert into users (id) values ("jennifernsturgis@gmail.com");

insert into lists (name, owner) values ("shopping list", "jennifernsturgis@gmail.com");
insert into lists (name, owner) values ("work todo", "randallcsturgis@gmail.com");
insert into lists (name, owner) values ("books to read", "randallcsturgis@gmail.com");

insert into tasks (parent_list, description, completed) values ("1","milk",false);
insert into tasks (parent_list, description, completed) values ("1","eggs",false);
insert into tasks (parent_list, description, completed) values ("1","bread",false);

insert into tasks (parent_list, description, completed) values ("2","tps report",false);
insert into tasks (parent_list, description, completed) values ("2","timecard",false);
insert into tasks (parent_list, description, completed) values ("2","get coffee",false);

insert into tasks (parent_list, description, completed) values ("3","how to read",false);
insert into tasks (parent_list, description, completed) values ("3","illustrated dictionary",false);
insert into tasks (parent_list, description, completed) values ("3","how to program",false);
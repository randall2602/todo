create view tasks_summary as
select tasks.id,
       tasks.parent_list,
       tasks.description,
       tasks.completed,
       lists.id as list_id,
       lists.name,
       lists.owner
from tasks inner join lists on tasks.parent_list=lists.id;
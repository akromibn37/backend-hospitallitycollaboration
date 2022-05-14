alter table t_customer_service 
drop column hos_id

alter table t_customer_service
add start_date datetime

alter table t_customer_service
add end_date datetime
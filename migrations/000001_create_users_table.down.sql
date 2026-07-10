DROP TABLE IF EXISTS users;
--this is the down migration file we will add the sql statements to drop the table in the down migration file and the sql statements to create a table in the up migration file.
-- NOW TO APPLY THE MIGRATION  WE WILL RUN THE FOLLOWING COMMAND IN THE TERMINAL
-- migrate -path migrations -database sqlite3://socialNetwork.db up

package sqlite

const createTableQuery = `CREATE TABLE IF NOT EXISTS scheduler(
	id  integer primary key autoincrement,
	date char(8),
	title varchar,
	comment varchar,
	repeat varchar);
CREATE INDEX IF NOT EXISTS scheduler_date on scheduler(date);
`
const addTaskQuery = `INSERT INTO scheduler
(date, title, comment, repeat)
VALUES(?, ?, ?, ?);`
const getTasksQuery = `SELECT id, date, title, comment, repeat 
FROM scheduler ORDER 
BY date ASC;`

const getTaskByIdQuery = `SELECT id, date, title, comment, repeat 
FROM scheduler 
WHERE id = :id;`

const updateTaskQuery = `UPDATE scheduler
SET date=:date, 
title=:title, 
comment=:comment, 
repeat=:repeat
WHERE id=:id;`

const deleteTaskQuery = "DELETE FROM scheduler WHERE id=:id;"

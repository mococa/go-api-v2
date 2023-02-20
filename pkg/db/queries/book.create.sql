INSERT INTO books ("name", "author_id", "release_year")
VALUES (?, ?, ?)
RETURNING *
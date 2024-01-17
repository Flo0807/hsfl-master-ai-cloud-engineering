INSERT INTO users (username, password) VALUES('test', '\x243261243130247971356D7A35353263657173353066314869396478756D58464A3143664177796958725074706D6D786C78354F384D386353435879');

INSERT INTO posts (created_at, updated_at, title, content) VALUES
(current_timestamp, current_timestamp, 'First Post', 'This is the content of the first post.'),
(current_timestamp, current_timestamp, 'Second Post', 'This is the content of the second post.'),
(current_timestamp, current_timestamp, 'Third Post', 'This is the content of the third post.');
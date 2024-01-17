INSERT INTO users (id, email, password) VALUES(1, 'user@example.com', '\x2432612431302476304c7233415262754477677536596631796775464f6d71474278784c5a426d466c32714e537278334c746a642f6a4f743870522e');

INSERT INTO posts (created_at, updated_at, title, content) VALUES
(current_timestamp, current_timestamp, 'First Post', 'This is the content of the first post.'),
(current_timestamp, current_timestamp, 'Second Post', 'This is the content of the second post.'),
(current_timestamp, current_timestamp, 'Third Post', 'This is the content of the third post.');
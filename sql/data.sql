-- Active: 1662320489072@@127.0.0.1@3306@devbook
INSERT INTO users (name, nickname, email, password)
VALUES
("User 1", "user_1", "user1@email.com", "$2a$10$9TyzlPqUvGR/.A575VsOF.2uaSF/cHJrTvoDn6zt67gDaJmIwvYzq"),
("User 2", "user_2", "user2@email.com", "$2a$10$9TyzlPqUvGR/.A575VsOF.2uaSF/cHJrTvoDn6zt67gDaJmIwvYzq"),
("User 3", "user_3", "user3@email.com", "$2a$10$9TyzlPqUvGR/.A575VsOF.2uaSF/cHJrTvoDn6zt67gDaJmIwvYzq"),
("User 4", "user_4", "user4@email.com", "$2a$10$9TyzlPqUvGR/.A575VsOF.2uaSF/cHJrTvoDn6zt67gDaJmIwvYzq"),
("User 5", "user_5", "user5@email.com", "$2a$10$9TyzlPqUvGR/.A575VsOF.2uaSF/cHJrTvoDn6zt67gDaJmIwvYzq");

INSERT INTO followers (user_id, follower_id)
VALUES
(1, 2),
(1, 3),
(1, 4),
(1, 5),
(2, 5),
(3, 5);

-- TEST: insert non existent user
INSERT INTO followers (user_id, follower_id)
VALUES
(3, 6);



create database messages;

use messages;

create table messages (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,	
	message VARCHAR(250) NOT NULL
);

create table users (
	userid VARCHAR(30) NOT NULL PRIMARY KEY,
	name VARCHAR(50) NOT NULL
);

create table usermessages(
	userid VARCHAR(30) NOT NULL,
	messageid INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	message VARCHAR(250) NOT NULL,
	FOREIGN KEY (userid) REFERENCES users(userid)	
);

-- GetMessages

DELIMITER //

CREATE PROCEDURE GetMessages()
BEGIN
	SELECT * FROM messages;
END //

DELIMITER ;

-- GetUserMessages

DELIMITER //

CREATE PROCEDURE GetUserMessages(IN in_userid VARCHAR(30))
BEGIN
	SELECT * FROM usermessages WHERE userid=in_userid;
END //

DELIMITER ;

-- AddMessage

DELIMITER //

CREATE PROCEDURE AddMessage(IN in_message VARCHAR(250))
BEGIN
	insert into messages (message) values (in_message);
END //

DELIMITER ;

-- AddUserMessage
DELIMITER //

CREATE PROCEDURE AddUserMessage(IN in_userid VARCHAR(30), IN in_message VARCHAR(250))
BEGIN
	insert into usermessages (userid, message) values  (in_userid, in_message);
END //

DELIMITER ;

-- UpdateUserMessage

DELIMITER //

CREATE PROCEDURE UpdateUserMessage(IN in_userid VARCHAR(30), IN in_messageid INT(6), IN in_message VARCHAR(250))
BEGIN
	update usermessages set message=in_message where userid=in_userid and messageid=in_messageid;
END //

DELIMITER ;

-- DeleteUserMessage

DELIMITER //

CREATE PROCEDURE DeleteUserMessage(IN in_userid VARCHAR(30), IN in_messageid INT(6))
BEGIN
	delete from usermessages where userid=in_userid and messageid=in_messageid;
END //

DELIMITER ;



insert into users (userid, name) values ("abc", "AlvarDev");
CALL AddMessage("This is a simple new message");
CALL AddUserMessage("abc", "Hola Mundo!");

CALL GetMessages();
CALL GetUserMessages("abc");


SET DateStyle TO European;

DROP TABLE IF EXISTS cc_courses;

CREATE TABLE cc_courses (
   id INT PRIMARY KEY NOT NULL,
   coursename VARCHAR(100) NOT NULL,
   coursecode VARCHAR(100) UNIQUE NOT NULL,
   description VARCHAR(500) NOT NULL,
   instructor VARCHAR(100)  NOT NULL,
   backgroundurl VARCHAR(500)  NOT NULL
);


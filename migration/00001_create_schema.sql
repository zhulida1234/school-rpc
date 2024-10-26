
CREATE TABLE IF NOT EXISTS students (
                          id BIGSERIAL PRIMARY KEY,
                          name VARCHAR(100),
                          age INT,
                          gender INT,
                          mobile VARCHAR(100),
                          class_name VARCHAR(100),
                          grade INT
);

CREATE TABLE IF NOT EXISTS clazz (
                                     id BIGSERIAL PRIMARY KEY,
                                     name VARCHAR(100),
                                     grade INT
)

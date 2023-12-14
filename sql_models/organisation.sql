
CREATE TABLE "branch" (
  "id" UUID PRIMARY KEY,
  "branch_code" VARCHAR,
  "name" VARCHAR,
  "address" VARCHAR,
  "phone_number" VARCHAR,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP
);


-- Employee
-- Id          string `json:"id"`
-- 	Name        string `json:"name"`
-- 	LastName    string `json:"last_name"`
-- 	PhoneNumber string `json:"phone_number"`
-- 	Login       string `json:"login"`
-- 	Password    string `json:"password"`
-- 	Branch      string `json:"branch"`
-- 	SaleCenter  string `json:"sale_center"`
-- 	UserType    string `json:"user_type"`
-- 	CreatedAt   string `json:"created_at"`
-- 	UpdatedAt   string `json:"updated_at"`

CREATE TABLE "employee" (
  "id" UUID PRIMARY KEY,
  "name" VARCHAR ,
  "last_name" VARCHAR ,
  "phone_number" VARCHAR ,
  "login" VARCHAR ,
  "password" VARCHAR ,
  "branch" VARCHAR ,
  "sale_center" UUID REFERENCES "sale_center"("id"),
  "user_type" VARCHAR ,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP
);
-- Provider 
-- Id          string `json:"id"`
-- 	Name        string `json:"name"`
-- 	PhoneNumber string `json:"phone_number"`
-- 	Active      bool   `json:"active"`
-- 	CreatedAt   string `json:"created_at"`
-- 	UpdatedAt   string `json:"updated_at"`

CREATE TABLE "provider" (
  "id" UUID PRIMARY KEY ,
  "name" VARCHAR ,
  "phone_number" VARCHAR , 
  "active" BOOLEAN,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP
);

-- SaleCenter  
-- 	Id        string `json:"id"`
-- 	Name      string `json:"name"`
-- 	Branch    string `json:"branch"`
-- 	CreatedAt string `json:"created_at"`
-- 	UpdatedAt string `json:"updated_at"`
CREATE TABLE "sale_center" (
  "id" UUID PRIMARY KEY ,
  "name" VARCHAR ,
  "branch" UUID REFERENCES "branch"("id"),
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP
);
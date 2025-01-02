# API Documentation

## /**api/user**

---

### **POST** - Create a user

This endpoint creates a new user. It expects the request body to contain a JSON object with the user's email and password.

### **Request**

```json
{
  "email": "user@example.com",
  "password": "yourpassword"
}

```

```bash
curl --location 'http://localhost:8080/api/user' \
--header 'Content-Type: application/json' \
--data-raw '{
  "email": "user@example.com",
  "password": "yourpassword"
}'

```

### **Status Codes:**

- **201 Created**
    
    **Description:** User was created successfully.
    
    **Response Example:**
    
    ```json
    {
      "message": "User created successfully"
    }
    
    ```
    
- **400 Bad Request**
    
    **Description:** Error parsing the request body. This can happen if the JSON format is incorrect.
    
    **Response Example:**
    
    ```json
    {
      "message": "Error parsing request body",
      "error": "Detailed error message"
    }
    
    ```
    
- **422 Unprocessable Entity**
    
    **Description:** Missing required fields (email or password).
    
    **Response Example:**
    
    ```json
    {
      "message": "Email and password are required"
    }
    
    ```
    
- **409 Conflict**
    
    **Description:** The email is already taken (duplicate entry).
    
    **Response Example:**
    
    ```json
    {
      "message": "Email already taken"
    }
    
    ```
    
- **500 Internal Server Error**
    
    **Description:** Error occurred while hashing the password.
    
    **Response Example:**
    
    ```json
    {
      "message": "Error hashing password",
      "error": "Detailed error message"
    }
    
    ```
    
- **501 Not Implemented**
    
    **Description:** Error inserting the user into the database, possibly due to some unknown issue.
    
    **Response Example:**
    
    ```json
    {
      "message": "Error inserting into database",
      "error": "Detailed error message"
    }
    
    ```
    

---

# 10. User authentication

## 10.1. Routes setup

| **Route pattern** | **Handler** | **Action** |
| --- | --- | --- |
| GET /{$} | home | Display the home page |
| GET /snippet/view/{id} | snippetView | Display a specific snippet |
| GET /snippet/create | snippetCreate | Display a form for creating a new snippet |
| POST /snippet/create | snippetCreatePost | Create a new snippet |
| GET /user/signup | userSignup | Display a form for signing up a new user |
| POST /user/signup | userSignupPost | Create a new user |
| GET /user/login | userLogin | Display a form for logging in a user |
| POST /user/login | userLoginPost | Authenticate and login the user |
| POST /user/logout | userLogoutPost | Logout the user |
| GET /static/ | http.FileServer | Serve a specific static file |


## 10.2. Creating a users model
## 10.3. User signup and password encryption
## 10.4. User login
## 10.5. User logout
## 10.6. User authorization
## 10.7. CSRF protection

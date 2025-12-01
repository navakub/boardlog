@echo off
setlocal enabledelayedexpansion

REM -----------------------------
REM 1. Register
REM -----------------------------
echo Registering user...
FOR /F "tokens=*" %%i IN ('curl -s -w "%%{http_code}" -o temp_response.json -X POST http://localhost:5001/api/auth/register ^
-H "Content-Type: application/json" ^
-d "{\"email\":\"testuser@example.com\",\"password\":\"123456\"}"') DO set STATUS=%%i

if NOT !STATUS! == 201 (
    echo Failed to register user, status code: !STATUS!
    type temp_response.json
    del temp_response.json
    goto :EOF
)

set /p REGISTER_RESPONSE=<temp_response.json
del temp_response.json

echo User registered: !REGISTER_RESPONSE!
echo.

REM -----------------------------
REM 2. Login
REM -----------------------------
echo Logging in...
FOR /F "tokens=*" %%i IN ('curl -s -w "%%{http_code}" -o temp_response.json -X POST http://localhost:5001/api/auth/login ^
-H "Content-Type: application/json" ^
-d "{\"email\":\"testuser@example.com\",\"password\":\"123456\"}"') DO set STATUS=%%i

if NOT !STATUS! == 200 (
    echo Failed to login, status code: !STATUS!
    type temp_response.json
    del temp_response.json
    goto :EOF
)

set /p LOGIN_RESPONSE=<temp_response.json
del temp_response.json

echo Login response: !LOGIN_RESPONSE!
REM Extract token if needed here
echo.

REM -----------------------------
REM 3. Logout
REM -----------------------------
echo Logging out...
FOR /F "tokens=*" %%i IN ('curl -s -w "%%{http_code}" -o temp_response.json -X POST http://localhost:5001/api/auth/logout ^
-H "Content-Type: application/json"') DO set STATUS=%%i

if NOT !STATUS! == 200 (
    echo Failed to logout, status code: !STATUS!
    type temp_response.json
    del temp_response.json
    goto :EOF
)

del temp_response.json
echo Logged out.
echo.

REM -----------------------------
REM 4. Delete user
REM -----------------------------
echo Deleting user...
FOR /F "tokens=*" %%i IN ('curl -s -w "%%{http_code}" -o temp_response.json -X DELETE http://localhost:5001/api/user/1 ^
-H "Content-Type: application/json"') DO set STATUS=%%i

if NOT !STATUS! == 200 (
    echo Failed to delete user, status code: !STATUS!
    type temp_response.json
    del temp_response.json
    goto :EOF
)

del temp_response.json
echo User deleted.
pause

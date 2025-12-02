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

echo Login response: !LOGIN_RESPONSE!
REM Extract access token if needed here
for /f "delims=" %%a in ('powershell -NoLogo -NoProfile -Command ^
    "(Get-Content 'temp_response.json' | ConvertFrom-Json).access_token"') do (
    set access_token=%%a
)

echo access token extracted!

REM Extract refresh token if needed here
for /f "delims=" %%a in ('powershell -NoLogo -NoProfile -Command ^
    "(Get-Content 'temp_response.json' | ConvertFrom-Json).refresh_token"') do (
    set refresh_token=%%a
)

echo refresh token extracted!

set /p LOGIN_RESPONSE=<temp_response.json
del temp_response.json

echo.


REM -----------------------------
REM 3. Me
REM -----------------------------
echo Getting current authenticated user...
FOR /F "tokens=*" %%i IN (
    'curl -s -w "%%{http_code}" -o temp_response.json -X GET http://localhost:5001/api/auth/me ^
    -H "Content-Type: application/json" ^
    -H "Authorization: Bearer !access_token!"'
) DO set STATUS=%%i

if NOT !STATUS! == 200 (
    echo Failed to get current user, status code: !STATUS!
    type temp_response.json
    del temp_response.json
    goto :EOF
)

set /p ME_RESPONSE=<temp_response.json
del temp_response.json
echo Current user info: !ME_RESPONSE!
echo.


REM -----------------------------
REM 3. Refresh
REM -----------------------------
echo Generating new refresh token...
echo Refreshing access token...
FOR /F "tokens=*" %%i IN (
    'curl -s -w "%%{http_code}" -o temp_response.json -X POST http://localhost:5001/api/auth/refresh ^
    -H "Content-Type: application/json" ^
    -H "Authorization: Bearer !access_token!"' ^
    -d "{\"refresh_token\":\"!refresh_token!\"}"'
) DO set STATUS=%%i

if NOT !STATUS! == 200 (
    echo Failed to refresh token, status code: !STATUS!
    type temp_response.json
    del temp_response.json
    goto :EOF
)

REM Extract new access token using PowerShell
for /f "delims=" %%a in ('powershell -NoLogo -NoProfile -Command ^
    "(Get-Content 'temp_response.json' | ConvertFrom-Json).access_token"') do (
    set access_token=%%a
)

set /p REFRESH_RESPONSE=<temp_response.json
del temp_response.json

echo New access token created...
echo Refreshed...
echo.

REM -----------------------------
REM 3. Logout
REM -----------------------------
echo Logging out...
FOR /F "tokens=*" %%i IN (
    'curl -s -w "%%{http_code}" -o temp_response.json -X POST http://localhost:5001/api/auth/logout ^
    -H "Content-Type: application/json" ^
    -H "Authorization: Bearer !access_token!"'
) DO set STATUS=%%i

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

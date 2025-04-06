@ECHO OFF

IF "%~1" == "" GOTO END
IF "%~1" == "build" GOTO MAIN
IF "%~1" == "serve" GOTO SERVE

:MAIN
nodemon --ext go --exec go build -o main.exe
goto END

:SERVE
@REM nodemon --watch main.exe --signal SIGTERM --exec main.exe
nodemon --watch main.exe --ext none --exec main.exe
goto END

:END
ENDLOCAL
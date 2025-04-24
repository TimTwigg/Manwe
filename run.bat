@ECHO OFF

IF "%~1" == "" GOTO RUN
IF "%~1" == "build" GOTO MAIN
IF "%~1" == "serve" IF "%~2" == "exe" GOTO SERVEEXEC
IF "%~1" == "serve" GOTO SERVE
IF "%~1" == "v" GOTO VERIFY

:RUN
go run main.go
goto END

:MAIN
nodemon --ext go --exec go build -o main.exe
goto END

:SERVE
nodemon --ext go --exec go run main.go
goto END

:SERVEEXEC
nodemon --watch main.exe --ext none --exec main.exe
goto END

:VERIFY
py verify.py
goto END

:END
ENDLOCAL
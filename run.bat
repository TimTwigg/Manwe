@ECHO OFF

IF "%~1" == "" GOTO MAIN

:MAIN
go run main.go
goto END

:END
ENDLOCAL
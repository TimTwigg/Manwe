@ECHO OFF

IF "%~1" == "" GOTO MAIN
IF "%~1" == "v" GOTO VALIDATE
IF "%~1" == "validate" GOTO VALIDATE

:MAIN
go run main.go
goto END

:VALIDATE
cd scripts
python validate_schemas.py
cd ..
goto END

:END
ENDLOCAL
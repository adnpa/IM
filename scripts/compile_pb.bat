
@echo off
setlocal enabledelayedexpansion

:: 指定 proto 文件所在的目录
set PROTO_DIR=../proto

:: 指定输出目录
set OUTPUT_DIR=../pkg/pb

:: 检查 protoc 命令是否存在
where protoc >nul 2>&1
if %errorlevel% neq 0 (
    echo protoc could not be found, please install it first.
    exit /b
)

:: 循环遍历目录中的所有 .proto 文件
for /r "%PROTO_DIR%" %%f in (*.proto) do (
    echo Compiling "%%f"...
    protoc --go_out="%OUTPUT_DIR%" --go-grpc_out="%OUTPUT_DIR%" -I="%PROTO_DIR%" "%%~nf%%~xf"
)

echo Compilation complete.

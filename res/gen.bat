set WORKSPACE=.
set LUBAN_DLL=%WORKSPACE%\Luban4.1.0\Luban.dll

dotnet %LUBAN_DLL% ^
    -t all ^
    -d json ^
    -d bin ^
    -c go-bin ^
    --conf %WORKSPACE%\luban.conf ^
    -x json.outputDataDir=data/json ^
    -x bin.outputDataDir=data/bin ^
    -x outputCodeDir=../pkg/res ^
    -x lubanGoModule=KryptonGo/pkg/luban

pause
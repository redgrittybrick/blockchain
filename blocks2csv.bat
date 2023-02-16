@echo off
pushd %APPDATA%\Bitcoin\blocks
for %%f in (blk033*.dat blk034*.dat) do (
    echo %%f 1>&2
    blockchain -file %%f -blocks -transactions -format csv || exit /b
)
popd

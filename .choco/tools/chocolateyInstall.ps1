
$packageName = 'ferebook'
$version = "0.01"
$url = "https://github.com/idcpj/ferebook/releases/download/$version/ferebook_$version`_win64_x64.zip"

Install-ChocolateyZipPackage -PackageName $packageName `
 -Url $url `
 -UnzipLocation "$(Split-Path -Parent $MyInvocation.MyCommand.Definition)" `
 -Url64 $url64
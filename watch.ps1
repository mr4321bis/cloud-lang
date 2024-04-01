#! /snap/bin/pwsh

# watch a file changes in the current directory, 
# execute all tests when a file is changed or renamed

$watcher = New-Object System.IO.FileSystemWatcher
$watcher.Path = get-location
$watcher.IncludeSubdirectories = $true
$watcher.EnableRaisingEvents = $false
$watcher.NotifyFilter = [System.IO.NotifyFilters]::LastWrite `
	-bor [System.IO.NotifyFilters]::FileName


$gobin = '/snap/bin/go'
$app = 'cloud-lang'

Start-Process $gobin 'build'
$process = Start-Process $app -PassThru
while ($TRUE) {
	$result = $watcher.WaitForChanged([System.IO.WatcherChangeTypes]::Changed `
			-bor [System.IO.WatcherChangeTypes]::Renamed `
			-bOr [System.IO.WatcherChangeTypes]::Created,
		1000);
	if ($result.TimedOut) {
		continue;
	}
	if (($result.Name -notmatch '\.go$') -and ($result.Name -notmatch '\.tmpl$')) {
		Write-Host " no match: " $result.Name
		continue;
	}
	Write-Host "... change" $result.Name
	Stop-Process $process
	Start-Process $gobin 'build' 
	$process = Start-Process $app -PassThru
}

Write-Host "exiting..."
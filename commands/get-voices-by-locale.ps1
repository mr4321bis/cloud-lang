#! /usr/bin/pwsh

$headers = @{
  'Ocp-Apim-Subscription-Key' = ''
}

$uri = 'https://eastus.tts.speech.microsoft.com/cognitiveservices/voices/list'

Invoke-RestMethod -Method GET -Uri $uri -Headers $headers | ForEach-Object {
  $Output = @{}
  $_ | Select-Object -Property Locale, DisplayName | ForEach-Object {
    if (!$Output.ContainsKey($_.Locale)) {
      $Output.Add($_.Locale, $_.DisplayName)
    }
  }
  $Output | ConvertTo-Json
}
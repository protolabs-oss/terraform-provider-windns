package windns

import (
    "bytes"
    "text/template"
)

var createTemplate = `
try { 
    $newRecord = $record = Get-DnsServerResourceRecord -ZoneName '{{.ZoneName}}' -RRType '{{.RecordType}}' -Name '{{.RecordName}}' -ComputerName '{{.DomainController}}' -ErrorAction Stop 
} catch { $record = $null }; 
if ($record -and $record.RecordType -eq '{{.RecordType}}') { 
    Write-Host 'Existing Record Found, Modifying record.'
    Switch ('{{.RecordType}}')
    {
        'A'     { $newRecord.RecordData.IPv4Address = '{{.IPv4Address }}' }
        'CNAME' { $newRecord.RecordData.HostNameAlias = '{{.HostnameAlias}}' }
    }
    Set-DnsServerResourceRecord -ZoneName '{{.ZoneName}}' -OldInputObject $record -NewInputObject $newRecord -PassThru -ComputerName '{{.DomainController}}'
}
else {
    if ($record) {
        Remove-DnsServerResourceRecord -InputObject $record -ZoneName '{{.ZoneName}}' -ComputerName '{{.DomainController}}' -PassThru -Force
    }
    Write-Host 'Creating record.'
    Switch ('{{.RecordType}}')
    {
        'A'     { Add-DnsServerResourceRecord -ZoneName '{{.ZoneName}}' -{{.RecordType}} -Name '{{.RecordName}}' -ComputerName '{{.DomainController}}' -IPv4Address '{{.IPv4Address}}' }
        'CNAME' { Add-DnsServerResourceRecord -ZoneName '{{.ZoneName}}' -{{.RecordType}} -Name '{{.RecordName}}' -ComputerName '{{.DomainController}}' -HostNameAlias '{{.HostnameAlias}}' }
    }
}`

func getCreateCommand(record DNSRecord) (string, err) {
    t := template.New("CreateTemplate")
    t, err := t.Parse(createTemplate)
    if err != nil {
        return nil, err
    }

    var createComandBuffer bytes.Buffer
    if err := t.Execute(&createComandBuffer, record); err != nil {
        return nil, err
    }

    createCommand := createComandBuffer.String()
    return createCommand, nil
}
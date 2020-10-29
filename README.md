### Systems Design Practice:

The task is to come up with a prototype for a finding severity categorization system:

- **1)** A data consumer, some runtime that processes findings and has a constant need to access severity categorization mappings
- **2)** A data publishing mechanism

## Design

#### needs: 

- some **mock findings generator utility** that constantly (every ~2s or so) picks from a random list of findings / threat purposes
- some **data model** for the actual severity categorization mapping with O(1) access by finding / threat purpose
- some data **storage interface** with a file system implementation (read and write severity categorization mappings)

#### software:

For the data consumer we'll create a Go program that will process an infinite stream of incoming findings. Upon getting a new finding, the program will consult an in-memory cache of finding severity categorizations. 

The cache will be configurable in terms of how often to refresh the data.

For updating the, for starter's we'll have a simple JSON structure in the file system.

#### data model:

We will literally just need a big map of finding type / threat purpose to severity categorization. There aren't that many different finding severity types to get any fancier than that.

## Results

```
✔ ~/go/src/github.com/adrianosela/guardduty
17:32 $ go build
✔ ~/go/src/github.com/adrianosela/guardduty
17:32 $ ./guardduty
2020/10/28 17:32:29 Processed High severity finding UnauthorizedAccess:EC2/MetadataDNSRebind
2020/10/28 17:32:29 Processed Low severity finding Stealth:IAMUser/CloudTrailLoggingDisabled
2020/10/28 17:32:30 Processed Medium severity finding Recon:IAMUser/MaliciousIPCaller
2020/10/28 17:32:30 Processed Medium severity finding Behavior:EC2/NetworkPortUnusual
2020/10/28 17:32:31 Processed High severity finding Trojan:EC2/DNSDataExfiltration
2020/10/28 17:32:31 Processed Low severity finding UnauthorizedAccess:EC2/RDPBruteForce
2020/10/28 17:32:32 Processed Medium severity finding Recon:IAMUser/ResourcePermissions
2020/10/28 17:32:32 Processed High severity finding Trojan:EC2/DNSDataExfiltration
2020/10/28 17:32:33 Processed High severity finding Impact:EC2/PortSweep
2020/10/28 17:32:33 Processed High severity finding Policy:S3/BucketAnonymousAccessGranted
2020/10/28 17:32:34 Processed Medium severity finding Recon:EC2/Portscan
2020/10/28 17:32:34 Processed Medium severity finding Persistence:IAMUser/ResourcePermissions
2020/10/28 17:32:35 Processed High severity finding Trojan:EC2/PhishingDomainRequest!DNS
2020/10/28 17:32:35 Processed Low severity finding Stealth:S3/ServerAccessLoggingDisabled
2020/10/28 17:32:36 Processed Medium severity finding Recon:IAMUser/ResourcePermissions
2020/10/28 17:32:36 Processed High severity finding Discovery:S3/MaliciousIPCaller.Custom
2020/10/28 17:32:37 Processed High severity finding Trojan:EC2/DNSDataExfiltration
```

package generator

import (
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var findings = []string{
	"Backdoor:EC2/C&CActivity.B!DNS",
	"Backdoor:EC2/DenialOfService.Dns",
	"Backdoor:EC2/DenialOfService.Tcp",
	"Backdoor:EC2/DenialOfService.Udp",
	"Backdoor:EC2/DenialOfService.UdpOnTcpPorts",
	"Backdoor:EC2/DenialOfService.UnusualProtocol",
	"Backdoor:EC2/Spambot",
	"Behavior:EC2/NetworkPortUnusual",
	"Behavior:EC2/TrafficVolumeUnusual",
	"CryptoCurrency:EC2/BitcoinTool.B",
	"CryptoCurrency:EC2/BitcoinTool.B!DNS",
	"Impact:EC2/WinRMBruteForce",
	"Impact:EC2/PortSweep",
	"Recon:EC2/PortProbeEMRUnprotectedPort",
	"Recon:EC2/PortProbeUnprotectedPort",
	"Recon:EC2/Portscan",
	"Trojan:EC2/BlackholeTraffic",
	"Trojan:EC2/BlackholeTraffic!DNS",
	"Trojan:EC2/DGADomainRequest.B",
	"Trojan:EC2/DGADomainRequest.C!DNS",
	"Trojan:EC2/DNSDataExfiltration",
	"Trojan:EC2/DriveBySourceTraffic!DNS",
	"Trojan:EC2/DropPoint",
	"Trojan:EC2/DropPoint!DNS",
	"Trojan:EC2/PhishingDomainRequest!DNS",
	"UnauthorizedAccess:EC2/MaliciousIPCaller.Custom",
	"UnauthorizedAccess:EC2/MetadataDNSRebind",
	"UnauthorizedAccess:EC2/RDPBruteForce",
	"UnauthorizedAccess:EC2/SSHBruteForce",
	"UnauthorizedAccess:EC2/TorClient",
	"UnauthorizedAccess:EC2/TorRelay",
	"PenTest:IAMUser/KaliLinux",
	"PenTest:IAMUser/ParrotLinux",
	"PenTest:IAMUser/PentooLinux",
	"Persistence:IAMUser/NetworkPermissions",
	"Persistence:IAMUser/ResourcePermissions",
	"Persistence:IAMUser/UserPermissions",
	"Policy:IAMUser/RootCredentialUsage",
	"PrivilegeEscalation:IAMUser/AdministrativePermissions",
	"Recon:IAMUser/MaliciousIPCaller",
	"Recon:IAMUser/MaliciousIPCaller.Custom",
	"Recon:IAMUser/NetworkPermissions",
	"Recon:IAMUser/ResourcePermissions",
	"Recon:IAMUser/TorIPCaller",
	"Recon:IAMUser/UserPermissions",
	"ResourceConsumption:IAMUser/ComputeResources",
	"Stealth:IAMUser/CloudTrailLoggingDisabled",
	"Stealth:IAMUser/LoggingConfigurationModified",
	"Stealth:IAMUser/PasswordPolicyChange",
	"UnauthorizedAccess:IAMUser/ConsoleLogin",
	"UnauthorizedAccess:IAMUser/ConsoleLoginSuccess.B",
	"UnauthorizedAccess:IAMUser/InstanceCredentialExfiltration",
	"UnauthorizedAccess:IAMUser/MaliciousIPCaller",
	"UnauthorizedAccess:IAMUser/MaliciousIPCaller.Custom",
	"UnauthorizedAccess:IAMUser/TorIPCaller",
	"Discovery:S3/BucketEnumeration.Unusual",
	"Discovery:S3/MaliciousIPCaller.Custom",
	"Discovery:S3/TorIPCaller",
	"Exfiltration:S3/ObjectRead.Unusual",
	"Impact:S3/PermissionsModification.Unusual",
	"Impact:S3/ObjectDelete.Unusual",
	"PenTest:S3/KaliLinux",
	"PenTest:S3/ParrotLinux",
	"PenTest:S3/PentooLinux",
	"Policy:S3/AccountBlockPublicAccessDisabled",
	"Policy:S3/BucketBlockPublicAccessDisabled",
	"Policy:S3/BucketAnonymousAccessGranted",
	"Policy:S3/BucketPublicAccessGranted",
	"Stealth:S3/ServerAccessLoggingDisabled",
	"UnauthorizedAccess:S3/MaliciousIPCaller.Custom",
	"UnauthorizedAccess:S3/TorIPCaller",
}

// Stream returns a channel where generated findings
// will be written to periodically ever interval
func Stream(buffSize int, interval time.Duration) chan string {
	ticker := time.NewTicker(interval)

	ch := make(chan string, buffSize)

	// ending condition
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-ticker.C:
				ch <- findings[rand.Intn(len(findings))]
			case <-sig:
				close(sig)
				close(ch)
				return
			}
		}
	}()

	return ch
}

# Create single VSI in dal09. Hourly billed with private network connection only. 

resource "ibm_compute_vm_instance" "vm1" {
	hostname             = "vm1"
	domain               = "example.com"
	os_reference_code    = "CENTOS_7_64"
	datacenter           = "dal09"
	network_speed        = 100
	hourly_billing       = true
        private_network_only = true
	cores                = 1
	memory               = 1024
	disks                = [25]
	user_metadata = "{\"value\":\"newvalue\"}"
	dedicated_acct_host_only = true
	local_disk = false
	public_vlan_id = 1391277
	private_vlan_id = 7721931
	private_security_group_ids = ["576973"]
	#private_security_group_ids = ["1","2"]
	ipv6_enabled = true
	wait_time_minutes = 90
	public_bandwidth_unlimited = false
	ipv6_static_enabled = false
}


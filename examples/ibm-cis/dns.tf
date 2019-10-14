# Reference DNS registration
data "ibm_dns_domain_registration" "web_domain" {
  name = var.domain
}

# Set DNS name servers for CIS  
resource "ibm_dns_domain_registration_nameservers" "web_domain" {
  name_servers        = ibm_cis_domain.web_domain.name_servers
  dns_registration_id = data.ibm_dns_domain_registration.web_domain.id
}

# IBM Cloud Resource Group the CIS instance will be created under
data "ibm_resource_group" "wordpress_group" {
  name = var.resource_group
}

resource "ibm_cis" "web_domain" {
  name              = "web_domain"
  resource_group_id = data.ibm_resource_group.wordpress_group.id
  plan              = "standard"
  location          = "global"
}

resource "ibm_cis_domain_settings" "web_domain" {
  cis_id          = ibm_cis.web_domain.id
  domain_id       = ibm_cis_domain.web_domain.id
  waf             = "on"
  ssl             = "full"
  min_tls_version = "1.2"
}

resource "ibm_cis_domain" "web_domain" {
  cis_id = ibm_cis.web_domain.id
  domain = var.domain
}

resource "ibm_cis_healthcheck" "root" {
  cis_id         = ibm_cis.web_domain.id
  description    = "Websiteroot"
  expected_body  = ""
  expected_codes = "200"
  path           = "/"
}

resource "ibm_lbaas" "lbaas1" {
  name        = "terraformLB"
  description = "delete this"
  subnets     = [1511875]

  /*protocols = [{
    frontend_protocol     = "HTTPS"
    frontend_port         = 443
    backend_protocol      = "HTTP"
    backend_port          = 80
    load_balancing_method = "round_robin"
    tls_certificate_id    = 11670
  },
    {
      frontend_protocol     = "HTTP"
      frontend_port         = 80
      backend_protocol      = "HTTP"
      backend_port          = 80
      load_balancing_method = "round_robin"
    },
  ]*/
}
resource "ibm_lbaas" "lbaas2" {
  name        = "terraformLB"
  description = "delete this"
  subnets     = [1511875]

  /*protocols = [{
    frontend_protocol     = "HTTPS"
    frontend_port         = 443
    backend_protocol      = "HTTP"
    backend_port          = 80
    load_balancing_method = "round_robin"
    tls_certificate_id    = 11670
  },
    {
      frontend_protocol     = "HTTP"
      frontend_port         = 80
      backend_protocol      = "HTTP"
      backend_port          = 80
      load_balancing_method = "round_robin"
    },
  ]*/
}


resource "ibm_cis_origin_pool" "lon" {
  cis_id        = ibm_cis.web_domain.id
  name          = var.datacenter1
  check_regions = ["WEU"]

  monitor = ibm_cis_healthcheck.root.id

  origins {
    name    = var.datacenter1
    address = ibm_lbaas.lbaas1.vip
    enabled = true
  }

  description = "LON pool"
  enabled     = true
}

resource "ibm_cis_origin_pool" "ams" {
  cis_id        = ibm_cis.web_domain.id
  name          = var.datacenter2
  check_regions = ["WEU"]

  monitor = ibm_cis_healthcheck.root.id

  origins {
    name    = var.datacenter2
    address = ibm_lbaas.lbaas2.vip
    enabled = true
  }

  description = "AMS pool"
  enabled     = true
}

# GLB name - name advertised by DNS for the website: prefix + domain 
resource "ibm_cis_global_load_balancer" "web_domain" {
  cis_id           = ibm_cis.web_domain.id
  domain_id        = ibm_cis_domain.web_domain.id
  name             = "${var.dns_name}${var.domain}"
  fallback_pool_id = ibm_cis_origin_pool.lon.id
  default_pool_ids = [ibm_cis_origin_pool.lon.id, ibm_cis_origin_pool.ams.id]
  description      = "Load balancer"
  proxied          = true
  session_affinity = "cookie"
}

# Configuration replacing CIS service resource with a data source for reuse of existing CIS instance
# data "ibm_cis" "web_domain" {
#   name              = "web_domain"
#   resource_group_id = "${data.ibm_resource_group.wordpress_group.id}"
# }
# resource "ibm_cis_domain_settings" "web_domain" {
#   cis_id            = "${data.ibm_cis.web_domain.id}"
#   domain_id         = "${ibm_cis_domain.web_domain.id}"
#   "waf"             = "on"
#   "ssl"             = "full"
#   "min_tls_version" = "1.2"
# }
# resource "ibm_cis_domain" "web_domain" {
#   cis_id = "${data.ibm_cis.web_domain.id}"
#   domain = "${var.domain}"
# }
# resource "ibm_cis_healthcheck" "root" {
#   cis_id         = "${data.ibm_cis.web_domain.id}"
#   description    = "Websiteroot"
#   expected_body  = ""
#   expected_codes = "200"
#   path           = "/"
# }
# resource "ibm_cis_origin_pool" "lon" {
#   cis_id        = "${data.ibm_cis.web_domain.id}"
#   name          = "${var.datacenter1}"
#   check_regions = ["WEU"]
#   monitor = "${ibm_cis_healthcheck.root.id}"
#   origins = [{
#     name    = "${var.datacenter1}"
#     address = "${ibm_lbaas.lbaas1.vip}"
#     enabled = true
#   }]
#   description = "LON pool"
#   enabled     = true
# }
# resource "ibm_cis_origin_pool" "ams" {
#   cis_id        = "${data.ibm_cis.web_domain.id}"
#   name          = "${var.datacenter2}"
#   check_regions = ["WEU"]
#   monitor = "${ibm_cis_healthcheck.root.id}"
#   origins = [{
#     name    = "${var.datacenter2}"
#     address = "${ibm_lbaas.lbaas2.vip}"
#     enabled = true
#   }]
#   description = "AMS pool"
#   enabled     = true
# }
# # GLB name - name advertised by DNS for the website: prefix + domain 
# resource "ibm_cis_global_load_balancer" "web_domain" {
#   cis_id           = "${data.ibm_cis.web_domain.id}"
#   domain_id        = "${ibm_cis_domain.web_domain.id}"
#   name             = "${var.dns_name}${var.domain}"
#   fallback_pool_id = "${ibm_cis_origin_pool.lon.id}"
#   default_pool_ids = ["${ibm_cis_origin_pool.lon.id}", "${ibm_cis_origin_pool.ams.id}"]
#   description      = "Load balancer"
#   proxied          = true
#   session_affinity = "cookie"
# }

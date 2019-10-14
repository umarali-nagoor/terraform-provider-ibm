variable "iaas_classic_username" {
  description = "Enter the user name to access IBM Cloud classic infrastructure."
  default = "IBM1982912"
}

variable "iaas_classic_api_key" {
  description = "Enter the API key to access IBM Cloud classic infrastructure"
  default = "51391806799f2ff49cf32ab359060bc185fd47650059200071e70c7b34f2365d"
}

variable "ibmcloud_api_key" {
  description = "Enter your IBM Cloud API Key"
  default = "aB_8gGR8nyZozrGCerbvqYkacgnr1XH-dL1ARL4PF3hL"
}

provider "ibm" {
  iaas_classic_username = var.iaas_classic_username
  iaas_classic_api_key  = var.iaas_classic_api_key
  ibmcloud_api_key      = var.ibmcloud_api_key
}


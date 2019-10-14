variable "iaas_classic_username" {
  description = "Enter the user name to access IBM Cloud classic infrastructure. You can retrieve the user name by following the instructions for retrieving your classic infrastructure API key."

  type    = string
  default = "IBM1982912"
  //default = "IBM1683465"
}

variable "iaas_classic_api_key" {
  description = "Enter the API key to access IBM Cloud classic infrastructure. For more information for how to create an API key and retrieve it, see [Managing classic infrastructure API keys](https://cloud.ibm.com/docs/iam?topic=iam-classic_keys)."
  type=string
  default="51391806799f2ff49cf32ab359060bc185fd47650059200071e70c7b34f2365d"
  //default="222d21a45b4b8944b8711dff7e45a083e6ba9fa6cf868dd3bf8d98bf83441f52"
}

variable "ibmcloud_api_key" {
  description = "Enter your IBM Cloud API Key, you can get your IBM Cloud API key using: https://cloud.ibm.com/iam#/apikeys"
  type=string
  default="aB_8gGR8nyZozrGCerbvqYkacgnr1XH-dL1ARL4PF3hL"
  //default="jBABddXxmjKNL4Wil_BPh0djHTzymWKZCOArKoxY_-fq"
}

provider "ibm" {
  iaas_classic_username = var.iaas_classic_username
  iaas_classic_api_key  = var.iaas_classic_api_key
  ibmcloud_api_key      = var.ibmcloud_api_key
}


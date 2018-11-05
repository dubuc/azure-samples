# Getting the ProvisioninState and PowerState of a VM in a Scaleset

This little program queries the Azure Instance Metadata Service for the ProvisioningState and
PowerState of the Azure VM Instance running in a VM Scale Set. It then queries the Azure Resource 
Manager. API with the Azure SDK for Go to obtain the VM Instance View and print the PowerState and 
Provisioning State.

For this to work, your VM must have a MSI with the correct permissions to read the resource group
info.
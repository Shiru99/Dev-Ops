# Terraform (Infrastructure as Code)

Terraform is a tool for building, changing, and versioning infrastructure safely and efficiently. Terraform can manage existing and popular service providers as well as custom in-house solutions.


### Terraform Commands

- `terraform init` - Initialize a working directory containing Terraform configuration files

- `terraform plan` - Generate and show an execution plan

- `terraform apply` - Builds or changes infrastructure

- `terraform destroy` - Destroy Terraform-managed infrastructure

- `terraform show` - Inspect Terraform state or plan

- `terraform refresh` - Update local state file against real resources

- `terraform console` - Try Terraform expressions at an interactive command prompt

- `terraform graph` - Create a visual graph of Terraform resources


### Terraform Concepts

1. Providers
    * A provider is responsible for understanding API interactions and exposing resources. 
    * Providers are the mechanism by which Terraform interacts with any infrastructure type.
    * e.g. AWS, Azure, GCP, Kubernetes, etc.

2. Resources

    * A resource is a single object of infrastructure, such as a virtual machine, object storage bucket, or DNS zone.
    * e.g. aws_instance, aws_s3_bucket, etc.

3. Data Sources

    * A data source is a read-only/ existing data source that can be used to retrieve information from an external source.
    * e.g. existing_aws_vpc, aws_ami, aws_iam_policy_document, etc.

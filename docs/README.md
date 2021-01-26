## Compile multiverse provider

Check if GOPATH env variable is set, and set if it is not

```sh
echo $GOPATH
```

Clone multiverse's repository 

```sh
git clone git@github.com:mobfox/terraform-provider-multiverse.git $GOPATH/src/github.com/mobfox/terraform-provider-multiverse
```

Compile the provider

```sh
cd $GOPATH/src/github.com/mobfox/terraform-provider-multiverse
make build
```

## Move the provider to terraform's folder

```sh
mv $GOPATH/bin ~/.terraform.d/plugins/terraform-provider-multiverse_v0.0.1
```

Note: we are using terraform version 0.11 provider naming, which following the pattern 
`terraform-provider-<PROVIDER_NAME>_v<X.X.X>`.
Note: for terraform version 0.13 the pattern is described here: https://learn.hashicorp.com/tutorials/terraform/provider-use?in=terraform/providers#install-hashicups-provider

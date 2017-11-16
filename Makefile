
update-vendor:
	glide update

build: update-vendor
	revel build github.com/zeebox/terraform-server/server

run:
	revel run github.com/zeebox/terraform-server/server
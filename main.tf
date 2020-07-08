resource "sp_ipaddress" "my-ipaddress"{
          address = "1.2.3.6"
          description = "my ipaddress"
}

data "sp_myipaddress" "mypublicipaddress"{}

output "myip" {
          value = "${data.sp_myipaddress.mypublicipaddress.myip}"
}
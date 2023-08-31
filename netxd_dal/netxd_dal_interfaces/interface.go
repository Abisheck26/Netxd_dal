package netxd_dal_interfaces

import "github.com/Abisheck26/netxd_dal/netxd_dal/netxd_dal_models"

type ICustomer interface {
	CreateCustomer(customer *netxd_dal_models.Customer) (*netxd_dal_models.DBResponse, error)
}

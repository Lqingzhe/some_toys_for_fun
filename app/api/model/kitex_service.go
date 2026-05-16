package model

import kitexuserservice "aim/kitex_gen/kitexuserservice/userservice"

type ServiceClient struct {
	UserClient kitexuserservice.Client
}

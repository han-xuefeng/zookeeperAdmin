package admin

type AdminService struct {

}

func (a *AdminService)GetOneById(id int) (*Admin, error) {
	adminDao := GetAdminDao()
	admin := &Admin{}
	if err := adminDao.FindOneById(admin);err != nil {
		return nil, err
	}
	return admin,nil
}

func (a *AdminService) UpdateOne(admin *Admin) error{
	adminDao := GetAdminDao()
	return adminDao.UpdateOne(admin)
}

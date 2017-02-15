package easemob

type EaseMob struct {
	context *Context
	user    *User
}

func New(host, org, app, id, secret string) *EaseMob {
	em := &EaseMob{
		context: &Context{
			host:   host,
			org:    org,
			app:    app,
			id:     id,
			secret: secret,
			auth:   true,
			token:  "",
			uuid:   "",
			expire: 0,
		},
		user: &User{},
	}

	if err := em.context.init(); err != nil {
		return nil
	}

	return em
}

func (em *EaseMob) AddUser(user *User) error {
	_, err := em.user.add(em.context, user)
	return err
}

func (em *EaseMob) AddUsers(users []*User) error {
	_, err := em.user.add(em.context, users)
	return err
}

func (em *EaseMob) GetUser(username string) (resp []byte, err error) {
	em.user.Username = username
	return em.user.get(em.context)
}

func (em *EaseMob) GetUsers(limit int, cursor string) (resp []byte, err error) {
	return em.user.getUsers(em.context, limit, cursor)
}

func (em *EaseMob) DeleteUser(username string) error {
	em.user.Username = username
	_, err := em.user.delete(em.context)
	return err
}

func (em *EaseMob) DeleteUsers(limit int) error {
	_, err := em.user.deleteUsers(em.context, limit)
	return err
}

func (em *EaseMob) SetPassword(username, password string) error {
	em.user.Username = username
	em.user.Password = password
	_, err := em.user.setPassword(em.context)
	return err
}

func (em *EaseMob) SetNickname(username, nickname string) error {
	em.user.Username = username
	em.user.Nickname = nickname
	_, err := em.user.setNickname(em.context)
	return err
}

func (em *EaseMob) AddFriend(username, friendname string) error {
	em.user.Username = username
	_, err := em.user.addFriend(em.context, friendname)
	return err
}

func (em *EaseMob) DeleteFriend(username, friendname string) error {
	em.user.Username = username
	_, err := em.user.deleteFriend(em.context, friendname)
	return err
}

func (em *EaseMob) GetFriends(username string) (data []byte, err error) {
	em.user.Username = username
	return em.user.getFriends(em.context)
}

func (em *EaseMob) GetBlockUsers(username string) (data []byte, err error) {
	em.user.Username = username
	return em.user.getBlockUsers(em.context)
}

func (em *EaseMob) BlockUsers(username string, blockusers []string) error {
	em.user.Username = username
	_, err := em.user.blockUsers(em.context, blockusers)
	return err
}

func (em *EaseMob) DeleteBlockUser(username, blockuser string) error {
	em.user.Username = username
	_, err := em.user.deleteBlockUser(em.context, blockuser)
	return err
}

func (em *EaseMob) GetUserStatus(username string) (data []byte, err error) {
	em.user.Username = username
	return em.user.getUserStatus(em.context)
}

func (em *EaseMob) Activate(username string) error {
	em.user.Username = username
	_, err := em.user.activate(em.context)
	return err
}

func (em *EaseMob) Deactivate(username string) error {
	em.user.Username = username
	_, err := em.user.deactivate(em.context)
	return err
}

func (em *EaseMob) Offline(username string) error {
	em.user.Username = username
	_, err := em.user.offline(em.context)
	return err
}

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

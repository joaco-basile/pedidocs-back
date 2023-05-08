package user_service_test

// func TestRegister(t *testing.T) {

// 	oid := primitive.NewObjectID()
// 	userId = oid.Hex()

// 	user := m.User{
// 		ID:        oid,
// 		Name:      "Jesus",
// 		Email:     "jesus.matiz@micorreo.com",
// 		Password:  "123",
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}

// 	err := user_service.Register(user)

// 	if err != nil {
// 		t.Error("El usuario no se pudo registrar")
// 		t.Fail()
// 	} else {
// 		t.Log("la prueba de registrarse finalizo con exito")
// 	}

// }

// func TestLogin(t *testing.T) {

// 	usr, err := user_service.Login("Jesus", "jesus.matiz@micorreo.com", "123")

// 	if err != nil {
// 		t.Error("El login ha fallado")
// 		t.Fail()
// 	} else {
// 		t.Log("La prueba de login finalizo con exito!", usr)
// 	}
// }

// func TestUpdate(t *testing.T) {

// 	user := m.User{
// 		Name:  "Jesus Matiz",
// 		Email: "jesus.matiz.prg@gmail.com",
// 	}

// 	err := user_service.Update(user, userId)

// 	if err != nil {
// 		t.Error("Error al tratar de actualizar el usuario")
// 		t.Fail()
// 	} else {
// 		t.Log("La prueba de actualización finalizo con exito!")
// 	}
// }

// func TestDelete(t *testing.T) {

// 	err := user_service.Delete(userId)

// 	if err != nil {
// 		t.Error("Error al tratar de eliminar el usuario")
// 		t.Fail()
// 	} else {
// 		t.Log("La prueba de eliminación finalizo con exito!")
// 	}
// }

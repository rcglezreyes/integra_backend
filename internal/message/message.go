package message

const (
	//Starting Application
	MsgResponseStartError string = "ERROR READING CONFIGURATION FILE"
	//Responses JSON
	MsgResponseUserCreatedSuccess string = "User Created Successfully"
	MsgResponseUserListedSuccess  string = "Users Listed Successfully"
	MsgResponseUserUpdatedSuccess string = "Users Updated Successfully"
	MsgResponseUserDeletedSuccess string = "Users Deleted Successfully"
	//Controller and Model Operations
	MsgResponseBindDataError             string = "error bind data"
	MsgResponseInsertDataPgDatabaseError string = "error to insert data in Postgres DB"
	MsgResponseListDataPgDatabaseError   string = "error to list data in Postgres DB"
	MsgResponseUpdateDataPgDatabaseError string = "error to update data in Postgres DB"
	MsgResponseDeleteDataPgDatabaseError string = "error to delete data in Postgres DB"
	MsgComponentCreateUserController     string = "[UserController.CreateUser]"
	MsgComponentUpdateUserController     string = "[UserController.UpdateUser]"
	MsgComponentDeleteUserController     string = "[UserController.DeleteUser]"
	MsgComponentCreateUserModel          string = "[UserModel.CreateUser]"
	MsgComponentListUserModel            string = "[UserModel.ListUsers]"
	MsgComponentUpdateUserModel          string = "[UserModel.UpdateUser]"
	MsgComponentDeleteUserModel          string = "[UserModel.DeleteUser]"
	//Database
	MsgConnectingPgDatabaseError string = "error connecting to Postgres DB"
	//General
	MsgStatusOkSuccess string = "Posts Loaded"
)

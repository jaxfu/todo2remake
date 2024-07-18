export type T_FORMINFO_REGISTER = {
	username: string;
	firstPassword: string;
	secondPassword: string;
};

export type T_APIRESULT_REGISTER = {
	valid: boolean;
};

export type T_FORMINFO_LOGIN = {
	username: string;
	password: string;
};

export type T_APIRESULT_LOGIN = {
	valid: boolean;
};
